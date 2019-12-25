package chart

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Provider interface {
	GetId() string
	GetChartData(coin uint, token string, currency string, timeStart int64) (blockatlas.ChartData, error)
	GetCoinData(coin uint, token string, currency string) (blockatlas.ChartCoinInfo, error)
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
