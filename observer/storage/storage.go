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
	s.blockHeights.heights = make(map[uint]Block)
	return s
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64)
	GetBlock(coin uint) (Block, bool)
}

type Addresses interface {
	Lookup(coin uint, addresses ...string) ([]Subscription, error)
	AddSubscriptions([]interface{}) error
	DeleteSubscriptions([]interface{}) error
	GetAddressFromXpub(coin uint, xpub string) ([]Xpub, error)
	GetXpubFromAddress(coin uint, address string) (string, error)
	SaveXpubAddresses(coin uint, addresses []string, xpub string) error
}
