package storage

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/storage/redis"
)

type Storage struct {
	redis.Redis
	blockHeights BlockMap
}

func New() *Storage {
	s := new(Storage)
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
}

type Market interface {
	SaveTicker(coin *blockatlas.Ticker, pl ProviderList) error
	GetTicker(coin, token string) (*blockatlas.Ticker, error)
	SaveRates(rates blockatlas.Rates, pl ProviderList)
	GetRate(currency string) (*blockatlas.Rate, error)
}
