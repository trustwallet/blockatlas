package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/storage/cache"
	"github.com/trustwallet/blockatlas/pkg/storage/cache/mock"
	"github.com/trustwallet/blockatlas/pkg/storage/cache/redis"
)

type Storage struct {
	cache.Backend
	blockHeights BlockMap
}

func New(useMock bool) *Storage {
	s := new(Storage)
	if useMock {
		s.Backend = &mock.Mock{}
	} else {
		s.Backend = &redis.Redis{}
	}

	s.blockHeights.heights = make(map[uint]int64)
	return s
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64) error
}

type Addresses interface {
	Lookup(coin uint, addresses []string) ([]blockatlas.Subscription, error)
	AddSubscriptions(subscriptions []blockatlas.Subscription)
	DeleteSubscriptions(subscriptions []blockatlas.Subscription)
	GetXpubFromAddress(coin uint, address string) (xpub string, addresses []string, err error)
	GetXpub(coin uint, xpub string) ([]string, error)
	CacheXpubs(subscriptions map[string][]string)
}

type Market interface {
	SaveTicker(coin *blockatlas.Ticker, pl ProviderList) error
	GetTicker(coin, token string) (*blockatlas.Ticker, error)
	SaveRates(rates blockatlas.Rates, pl ProviderList)
	GetRate(currency string) (*blockatlas.Rate, error)
}
