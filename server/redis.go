package server

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

var (
	REDIS_SERVER  = ""
	REDIS_DB      = 1
	REDIS_TIMEOUT = 5
)

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_SERVER)

			if err != nil {
				return nil, err
			}

			if _, err = c.Do("SELECT", REDIS_DB); err != nil {
				c.Close()
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			c.Close()
			return err
		},
	}
}

func RedisStorage(ra, rp, rd, rt string) (*redis.Pool, error) {

	if ra != "" && rp != "" {
		REDIS_SERVER = ra + ":" + rp
	} else {
		return nil, errors.New("invalid redis server configure")
	}

	if val, err := strconv.Atoi(rd); err == nil {
		REDIS_DB = val
	}

	if val, err := strconv.Atoi(rt); err == nil {
		REDIS_TIMEOUT = val
	}

	pool := newPool()

	if _, err := pool.Get().Do("PING"); err != nil {
		return nil, err
	}

	return pool, nil
}
