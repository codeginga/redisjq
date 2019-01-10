package redis

import (
	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/cfg"
	"github.com/go-redis/redis"
)

type task struct {
	cfg cfg.Task
	c   *redis.Client
}

func (t *task) Save(key, value string) (err error) {
	res := t.c.Set(key, value, t.cfg.MessageTTLDuration())
	err = res.Err()

	return
}

func (t *task) Get(key string) (value string, err error) {
	res := t.c.Get(key)
	value, err = res.Result()
	if err != nil && errEmpty(err) {
		err = backend.ErrEmptyTask
	}
	return
}

// NewTask returns instance of backend.Task
func NewTask(c *redis.Client, cfg cfg.Task) backend.Task {
	return &task{
		c:   c,
		cfg: cfg,
	}
}
