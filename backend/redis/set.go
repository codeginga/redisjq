package redis

import (
	"strconv"
	"time"

	"github.com/codeginga/redisjq/cnst"

	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/cfg"
	"github.com/go-redis/redis"
)

type set struct {
	c      *redis.Client
	tskCfg cfg.Task
}

func (s *set) zkey() string {
	return cnst.App + "-" + s.tskCfg.QName
}

func (s *set) First() (key string, err error) {
	mx := strconv.FormatInt(time.Now().Unix(), 10)
	res := s.c.ZRangeByScore(s.zkey(), redis.ZRangeBy{
		Min:    "0",
		Max:    mx,
		Count:  1,
		Offset: 0,
	})

	if err = res.Err(); err != nil {
		return
	}

	vals, err := res.Result()
	if err != nil {
		return
	}

	if len(vals) == 0 {
		err = cnst.ErrEmptyQ
		return
	}

	key = vals[0]
	return
}

func (s *set) Add(tim time.Time, key string) (err error) {
	unix := tim.Unix()
	res := s.c.ZAdd(s.zkey(), redis.Z{
		Score:  float64(unix),
		Member: key,
	})

	err = res.Err()
	return
}

func (s *set) Remove(key string) (err error) {
	err = s.c.ZRem(s.zkey(), key).Err()
	return
}

// NewSet returns instance of backend.Set
func NewSet(redis *redis.Client, tskCfg cfg.Task) backend.Set {
	return &set{
		c:      redis,
		tskCfg: tskCfg,
	}
}
