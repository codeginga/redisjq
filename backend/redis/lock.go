package redis

import (
	"time"

	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/cfg"
	"github.com/codeginga/redisjq/cnst"
	"github.com/go-redis/redis"
)

type locker struct {
	c      *redis.Client
	tskCfg cfg.Task
}

func (l *locker) Lock(key string) error {
	res := l.c.SetNX(key, key, l.tskCfg.RunTimeDuration())
	if err := res.Err(); err != nil {
		return err
	}

	if !res.Val() {
		return cnst.ErrLocked
	}

	return nil
}

func (l *locker) Unlock(key string) error {
	res := l.c.Set(key, key, time.Second*1)

	if err := res.Err(); err != nil {
		return err
	}

	return nil
}

// NewLocker returns instance of locker
func NewLocker(redis *redis.Client, tskCfg cfg.Task) backend.Locker {
	return &locker{
		c:      redis,
		tskCfg: tskCfg,
	}
}
