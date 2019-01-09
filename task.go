package redisjq

import (
	"log"
	"sync"

	"github.com/codeginga/redisjq/backend"
)

type task struct {
	backend backend.Container

	msg *Message
}

func (t *task) Message() Message {
	return *t.msg
}

func (t *task) Done() error {
	return t.backend.Set.Remove(t.msg.Key())
}

func (t *task) Retry() error {
	return t.backend.Set.Add(t.msg.popupTime(), t.msg.Key())
}

func runTask(tsk Task, w Worker, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		err := recover()
		if err != nil {
			log.Println("panic task run, ", err)
		}
	}()

	w(tsk)
}
