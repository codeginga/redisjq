package redisjq

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/backend/redis"
	"github.com/codeginga/redisjq/cfg"
	"github.com/codeginga/redisjq/cnst"
)

type ctxHandler func(*ctx)

type ctx struct {
	cont backend.Container
	key  string
	msg  *Message

	err error
}

func newCtx(cont backend.Container) *ctx {
	return &ctx{
		cont: cont,
		key:  "",
		msg:  nil,
		err:  nil,
	}
}

func pickFirst(h ctxHandler) ctxHandler {
	return func(c *ctx) {
		k, err := c.cont.Set.First()
		if err != nil {
			panic(err)
		}

		c.key = k
		h(c)
	}
}

func lockKey(h ctxHandler) ctxHandler {
	return func(c *ctx) {
		if err := c.cont.Locker.Lock(c.key); err != nil {
			panic(err)
		}

		h(c)
	}
}

func mkMessage(h ctxHandler) ctxHandler {
	return func(c *ctx) {
		str, err := c.cont.Task.Get(c.key)
		if err != nil {
			panic(err)
		}

		msg := Message{}
		err = msg.FrmString(str)
		if err != nil {
			panic(err)
		}

		c.msg = &msg

		h(c)
	}
}

func retryCheck(h ctxHandler) ctxHandler {
	return func(c *ctx) {
		if c.msg.Retry > -1 {
			h(c)
			return
		}

		if err := c.cont.Set.Remove(c.msg.Key()); err != nil {
			panic(err)
		}

		panic(cnst.ErrRetryExit)
	}
}

func decrementRetry(c *ctx) {
	c.msg.Retry--
	val, err := c.msg.String()
	if err != nil {
		panic(err)
	}

	if err = c.cont.Task.Save(c.msg.Key(), val); err != nil {
		panic(err)
	}
}

func recov(h ctxHandler) ctxHandler {
	return func(c *ctx) {
		defer func() {
			r := recover()
			if r == nil {
				return
			}

			err, ok := r.(error)
			if !ok {
				c.err = errors.New("unknown")
				return
			}

			c.err = err
		}()

		h(c)
	}
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

func (s *server) runTasks() (err error) {
	wg := sync.WaitGroup{}

	for i := 0; i < s.concurrentxWorker; {
		c := newCtx(s.backend)
		h := recov(pickFirst(lockKey(mkMessage(retryCheck(decrementRetry)))))
		h(c)

		if c.err == cnst.ErrEmptyQ {
			break
		}

		if c.err == cnst.ErrLocked {
			continue
		}

		if c.err == cnst.ErrEmptyTask {
			i++
			continue
		}

		if c.err == cnst.ErrRetryExit {
			i++
			continue
		}

		if c.err != nil {
			err = c.err
			break
		}

		w, ok := s.register[c.msg.Name]
		if !ok {
			i++
			continue
		}

		wg.Add(1)
		go runTask(&task{
			backend: s.backend,
			msg:     c.msg,
		}, w, &wg)
	}

	wg.Wait()

	return
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
