package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type Redis struct {
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

func (db *Redis) GetValue(key string, value interface{}) error {
	cmd := db.client.Get(key)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), util.ErrNotFound)
	}
	err := json.Unmarshal([]byte(cmd.Val()), value)
	if err != nil {
		return errors.E(err, util.ErrNotFound)
	}
	return nil
}

func (db *Redis) Add(key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"value": value})
	}
	cmd := db.client.Set(key, j, 0)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), util.ErrNotStored)
	}
	return nil
}

func (db *Redis) Delete(key string) error {
	cmd := db.client.Del(key)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), util.ErrNotDeleted)
	}
	return nil
}

func (db *Redis) IsReady() bool {
	if db.client == nil {
		return false
	}
	return db.client.Ping().Err() == nil
}
