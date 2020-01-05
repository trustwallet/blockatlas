package mock

import (
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

func (db *Mock) GetAllHM(entity string) (map[string]string, error) {
	return nil, util.ErrNotFound
}

func (db *Mock) GetHMValue(entity, key string, value interface{}) error {
	return util.ErrNotFound
}

func (db *Mock) AddHM(entity, key string, value interface{}) error {
	return nil
}

func (db *Mock) DeleteHM(entity, key string) error {
	return nil
}
