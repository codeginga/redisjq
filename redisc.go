package redisjq

import (
	"github.com/codeginga/redisjq/cfg"
	"github.com/go-redis/redis"
)

func redisClient(cfg cfg.Redis) (c *redis.Client, err error) {
	c = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err = c.Ping().Result()
	return

}
