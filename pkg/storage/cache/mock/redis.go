package mock

import (
	"github.com/trustwallet/blockatlas/pkg/storage/util"
)

type Mock struct {
}

func (db *Mock) Init(host string) error {
	return nil
}

func (db *Mock) GetValue(key string, value interface{}) error {
	return util.ErrNotFound
}

func (db *Mock) Add(key string, value interface{}) error {
	return nil
}

func (db *Mock) Delete(key string) error {
	return nil
}

func (db *Mock) IsReady() bool {
	return true
}
