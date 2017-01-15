package server

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	REDIS_SERVER  = ""
	REDIS_DB      = 1
	REDIS_TIMEOUT = 5
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     4,
		MaxActive:   16,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.Port))

			if err != nil {
				return nil, err
			}

			if _, err = c.Do("SELECT", cfg.DB); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		}, /*
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				c.Close()
				return err
			},*/
	}
}

func RedisStorage() (*redis.Pool, error) {
	pool := newPool()

	if _, err := pool.Get().Do("PING"); err != nil {
		return nil, err
	}

	return pool, nil
}
