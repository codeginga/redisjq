package redisjq

import (
	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/backend/redis"
	"github.com/codeginga/redisjq/cfg"
)

type publisher struct {
	task backend.Task
	set  backend.Set
}

func (p *publisher) Publish(msg Message) (err error) {
	strMsg, err := msg.String()
	if err != nil {
		return
	}

	key := msg.Key()

	if err = p.task.Save(key, strMsg); err != nil {
		return
	}

	err = p.set.Add(msg.popupTime(), key)
	return
}

// NewPublisher returns new instance of Publisher
func NewPublisher(cfg cfg.Publisher) (p Publisher, err error) {
	rc, err := redisClient(cfg.Redis)
	if err != nil {
		return
	}

	set := redis.NewSet(rc, cfg.Task)
	task := redis.NewTask(rc, cfg.Task)
	p = &publisher{
		task: task,
		set:  set,
	}

	return
}
