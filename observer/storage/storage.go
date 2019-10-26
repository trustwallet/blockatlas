package storage

import (
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

type Storage struct {
	sql.PgSql
	blockHeights BlockMap
	subsMap      SubsMap
}

func New() *Storage {
	s := new(Storage)
	s.blockHeights.heights = make(map[interface{}]*Block)
	s.subsMap.subs = make(map[string][]Subscription)
	s.subsMap.xpubSubs = make(map[string][]Subscription)
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
