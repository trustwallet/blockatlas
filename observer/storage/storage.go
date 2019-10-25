package storage

import (
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

type Storage struct {
	sql.PgSql
	blockHeights BlockMap
}

func New() *Storage {
	s := new(Storage)
	s.blockHeights.heights = make(map[interface{}]*Block)
	return s
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64)
}

type Addresses interface {
	Lookup(coin uint, addresses ...string) (observers []Subscription, err error)
	AddSubscriptions(subscriptions []interface{}) error
	DeleteSubscriptions(subscriptions []interface{}) error
}
