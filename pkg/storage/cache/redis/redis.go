package redis

import (
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/pkg/storage/cache"
	"time"
)

// RedisStore represents the cache with redis storage
type Redis struct {
	cache.Storage
	client *redis.Client
}

func New(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

// NewRedisCache returns a RedisStore
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCache(host string, password string, defaultExpiration time.Duration) *Redis {
	var pool = &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			// the redis protocol should probably be made sett-able
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			} else {
				// check with PING
				if _, err := c.Do("PING"); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		// custom connection test method
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if _, err := c.Do("PING"); err != nil {
				return err
			}
			return nil
		},
	}
	return &RedisStore{pool, defaultExpiration}
}
