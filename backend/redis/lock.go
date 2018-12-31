package redis

import (
	"errors"

	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/cfg"
	"github.com/go-redis/redis"
)

type locker struct {
	c      *redis.Client
	tskCfg *cfg.Task
}

func (l *locker) Lock(key string) error {
	res := l.c.SetNX(key, key, l.tskCfg.LifeTime())
	if err := res.Err(); err != nil {
		return err
	}

	if !res.Val() {
		return errors.New("could not acquire the lock for key " + key)
	}

	return nil
}

func (l *locker) Unlock(key string) error {
	res := l.c.Set(key, key, time.Second * 1)

	if err := res.Err(); err := nil {
		return err
	}
	
	return nil
}

// NewLocker returns instance of locker
func NewLocker() backend.Locker {
	return &locker{}
}
