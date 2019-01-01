package cfg

import "github.com/go-redis/redis"

// Redis holds config for redis connection
type Redis struct {
	Addr     string
	Password string
	DB       int
}

func (r *Redis) client(c *redis.Client, err error) {
	c = redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})

	_, err = c.Ping().Result()
	return
}
