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

func (db *Redis) Init(host string) error {
	options, err := redis.ParseURL(host)
	if err != nil {
		return errors.E(err, "Cannot connect to Redis")
	}
	client := redis.NewClient(options)
	if err := client.Ping().Err(); err != nil {
		return errors.E(err, "Redis connection test failed")
	}
	db.client = client
	return nil
}

func (db *Redis) Into(value interface{}) error {
	q, ok := db.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	cmd := db.client.Get(q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotFound)
	}
	err := json.Unmarshal([]byte(cmd.Val()), value)
	if err != nil {
		return errors.E(err, storage.ErrNotFound)
	}
	return nil
}

func (db *Redis) Add(value interface{}) error {
	q, ok := db.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"value": value})
	}
	cmd := db.client.Set(q, j, 0)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotStored)
	}
	return nil
}

func (db *Redis) Delete() error {
	q, ok := db.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	cmd := db.client.Del(q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotDeleted)
	}
	return nil
}

func (db *Redis) Update(value interface{}) error {
	return db.Add(value)
}
