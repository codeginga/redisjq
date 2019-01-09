package redisjq

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/backend/redis"
	"github.com/codeginga/redisjq/cfg"
	"github.com/codeginga/redisjq/cnst"
)

type errcont struct {
	err error
}

func (e *errcont) recov() {
	r := recover()
	if r == nil {
		return
	}

	err, ok := r.(error)
	if !ok {
		e.err = fmt.Errorf("%v", r)
		return
	}

	e.err = err
}

type mkmsg struct {
	errcont

	backend backend.Container
	key     string

	msg *Message
}

func (m *mkmsg) fetch() {
	val, err := m.backend.Task.Get(m.key)
	if err != nil {
		panic(err)
	}

	msg := Message{}
	err = msg.FrmString(val)
	if err != nil {
		panic(err)
	}

	m.msg = &msg
}

func (m *mkmsg) valid() {
	if m.msg.Retry < 0 {
		m.msg = nil
	}
}

func (m *mkmsg) update() {
	m.msg.Retry--
	val, err := m.msg.String()
	if err != nil {
		panic(err)
	}

	if err = m.backend.Task.Save(m.msg.Key(), val); err != nil {
		panic(err)
	}
}

func (m *mkmsg) do() {
	defer m.recov()

	m.fetch()
	m.valid()
	m.update()
}

func (m *mkmsg) empty() bool {
	if m.err == backend.ErrEmpty {
		return true
	}

	if m.msg == nil {
		return true
	}

	return false
}

func (m *mkmsg) Do() (*Message, error) {
	m.do()

	return m.msg, m.err
}

type pick struct {
	errcont

	backend backend.Container

	key string
}

func (p *pick) empty() bool {
	if p.err == backend.ErrEmpty {
		return true
	}
	return false
}

func (p *pick) locked() bool {
	if p.err == backend.ErrLocked {
		return true
	}

	return false
}

func (p *pick) first() {
	k, err := p.backend.Set.First()
	if err != nil {
		panic(err)
	}

	p.key = k
}

func (p *pick) lock() {
	if err := p.backend.Locker.Lock(p.key); err != nil {
		panic(err)
	}
}

func (p *pick) do() {
	defer p.recov()

	p.first()
	p.lock()
}

func (p *pick) Do() (key string, err error) {
	p.do()
	return p.key, p.err
}

type server struct {
	backend backend.Container

	concurrentxWorker int
	ctx               context.Context

	register map[string]Worker
}

func (s *server) RegisterTask(name string, worker Worker) (err error) {
	s.register[name] = worker
	return
}

func (s *server) runTasks() error {
	wg := sync.WaitGroup{}

	for i := 0; i < s.concurrentxWorker; {
		p := pick{backend: s.backend}
		key, err := p.Do()
		if p.empty() {
			break
		}

		if p.locked() {
			continue
		}

		if err != nil {
			return err
		}

		m := mkmsg{backend: s.backend, key: key}
		msg, err := m.Do()
		if m.empty() {
			i++
			continue
		}

		if err != nil {
			return err
		}

		w, ok := s.register[msg.Name]
		if !ok {
			i++
			continue
		}

		wg.Add(1)
		go runTask(&task{
			backend: s.backend,
			msg:     msg,
		}, w, &wg)
	}

	wg.Wait()

	return nil
}

func (s *server) Start(ctx context.Context) (err error) {
	tic := time.NewTicker(cnst.SleepDuration)
	defer tic.Stop()

	done := ctx.Done()
	for {
		select {
		case <-tic.C:
			err = s.runTasks()
			if err != nil {
				return
			}

		case <-done:
			return
		}
	}

	return
}

// NewServer returns new instance of Server
func NewServer(ctx context.Context, cfg cfg.Server) (Server, error) {
	rc, err := redisClient(cfg.Redis)
	if err != nil {
		return nil, err
	}

	cont := backend.Container{
		Task:   redis.NewTask(rc, cfg.Task),
		Set:    redis.NewSet(rc, cfg.Task),
		Locker: redis.NewLocker(rc, cfg.Task),
	}

	return &server{
		ctx:               ctx,
		backend:           cont,
		register:          map[string]Worker{},
		concurrentxWorker: cfg.ConcurrentWorker,
	}, nil
}
