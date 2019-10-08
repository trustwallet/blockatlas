package redis

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage"
)

func (db *Redis) GetHMValue(entity, key string, value interface{}) error {
	cmd := db.client.HMGet(entity, key)
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

func (db *Redis) AddHM(entity, key string, value interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		return errors.E(err, errors.Params{"value": value})
	}
	cmd := db.client.HMSet(entity, map[string]interface{}{key: j})
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotStored)
	}
	return nil
}

func (db *Redis) DeleteHM(entity, key string) error {
	cmd := db.client.HDel(entity, key)
	if cmd.Err() != nil {
		return errors.E(cmd.Err(), storage.ErrNotDeleted)
	}
	return nil
}
