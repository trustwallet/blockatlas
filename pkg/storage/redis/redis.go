package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage"
)

type Redis struct {
	storage.Db
	client *redis.Client
}

func (r *Redis) Init(host string) error {
	options, err := redis.ParseURL(host)
	if err != nil {
		return errors.E(err, "Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		return errors.E(err, "Redis connection test failed")
	}
	r.client = client
	return nil
}

func (r *Redis) Into(value interface{}) error {
	q, ok := r.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	cmd := r.client.Get(q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotFound)
	}
	err := json.Unmarshal([]byte(cmd.Val()), value)
	if err != nil {
		return errors.E(err, storage.ErrNotFound)
	}
	return nil
}

func (r *Redis) Set(value interface{}) error {
	q, ok := r.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"value": value})
	}
	cmd := r.client.Set(q, j, 0)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotStored)
	}
	return nil
}

func (r *Redis) Delete() error {
	q, ok := r.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	cmd := r.client.Del(q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotDeleted)
	}
	return nil
}
