package market

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/storage"
	"time"
)

type Provider interface {
	Init(storage.Market) error
	GetId() string
	GetUpdateTime() time.Duration
	GetData() (blockatlas.Tickers, error)
	GetLogType() string
}

type Providers map[int]Provider

func (ps Providers) GetPriority(providerId string) int {
	for priority, provider := range ps {
		if provider.GetId() == providerId {
			return priority
		}
	}
	return -1
}
