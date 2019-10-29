package storage

import (
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/storage/sql"
)

type Storage struct {
	sql.PgSql
	blockHeights BlockMap
	xpubMap      XpubMap
}

func New() *Storage {
	s := new(Storage)
	s.blockHeights.heights = make(map[interface{}]*Block)
	s.xpubMap.xpub = make(map[string][]string)
	return s
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64)
}

type Addresses interface {
	Lookup(addresses ...string) (observers []Subscription, err error)
	AddSubscriptions(subscriptions []interface{}) error
	DeleteSubscriptions(subscriptions []interface{}) error
	GetXpubFromAddress(address string) (string, bool)
	GetXpub(xpub string) ([]string, bool)
}

func (s *Storage) LoadCacheData() {
	err := s.LoadXpubs()
	if err != nil {
		logger.Fatal(err)
	}
}

func (s *Storage) SaveCacheData() {
	err := s.SaveAllBlocks()
	if err != nil {
		logger.Error(err)
	}
	err = s.SaveAllXpubs()
	if err != nil {
		logger.Error(err)
	}
}

func (s *Storage) CloseStorageSafety() {
	s.SaveCacheData()
	err := s.Close()
	if err != nil {
		logger.Error(err)
	}
}
