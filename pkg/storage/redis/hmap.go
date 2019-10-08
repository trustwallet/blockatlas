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

func (db *HMap) Init(host string) error {
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

func (db *HMap) Into(value interface{}) error {
	if len(db.EntityName) == 0 {
		return errors.E(storage.ErrEmptyEntity)
	}
	q, ok := db.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	cmd := db.client.HMGet(db.EntityName, q)
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

func (db *HMap) Add(value interface{}) error {
	if len(db.EntityName) == 0 {
		return errors.E(storage.ErrEmptyEntity)
	}
	q, ok := db.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"value": value})
	}
	cmd := db.client.HMSet(db.EntityName, map[string]interface{}{q: j})
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotStored)
	}
	return nil
}

func (db *HMap) Delete() error {
	if len(db.EntityName) == 0 {
		return errors.E(storage.ErrEmptyEntity)
	}
	q, ok := db.QueryValue.(string)
	if !ok {
		return errors.E(storage.ErrEmptyQuery)
	}
	cmd := db.client.HDel(db.EntityName, q)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotDeleted)
	}
	return nil
}

func (db *HMap) Update(value interface{}) error {
	return db.Add(value)
}
