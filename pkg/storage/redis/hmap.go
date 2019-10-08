package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage"
)

type HMap struct {
	storage.Db
	client *redis.Client
}

func (r *HMap) Init(host string) error {
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

func (r *HMap) Into(value interface{}) error {
	q, ok := r.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	if len(r.EntityName) == 0 {
		return errors.E(storage.ErrEmptyEntity)
	}
	cmd := r.client.HMGet(r.EntityName, q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotFound)
	}
	val, ok := cmd.Val()[0].(string)
	if !ok {
		return errors.E(storage.ErrNotFound)
	}
	err := json.Unmarshal([]byte(val), value)
	if err != nil {
		return errors.E(err, storage.ErrNotFound)
	}
	return nil
}

func (r *HMap) Set(value interface{}) error {
	q, ok := r.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	if len(r.EntityName) == 0 {
		return errors.E(storage.ErrEmptyEntity)
	}
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"value": value})
	}
	cmd := r.client.HMSet(r.EntityName, map[string]interface{}{q: j})
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotStored)
	}
	return nil
}

func (r *HMap) Delete() error {
	q, ok := r.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	if len(r.EntityName) == 0 {
		return errors.E(storage.ErrEmptyEntity)
	}
	cmd := r.client.HDel(r.EntityName, q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotDeleted)
	}
	return nil
}
