package coingecko

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"time"
)

type Charts struct {
	Prices     []ChartVolume `json:"prices"`
	MarketCaps []ChartVolume `json:"market_caps"`
	Volumes    []ChartVolume `json:"total_volumes"`
}

type ChartVolume []float64

type CoinResult struct {
	Symbol   string
	TokenId  string
	CoinType blockatlas.CoinType
}

type CoinPrices []CoinPrice

type CoinPrice struct {
	Id                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	CurrentPrice                 float64   `json:"current_price"`
	PriceChange24h               float64   `json:"price_change_24h"`
	PriceChangePercentage24h     float64   `json:"price_change_percentage_24h"`
	MarketCapChange24h           float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64   `json:"market_cap_change_percentage_24h"`
	MarketCap                    float64   `json:"market_cap"`
	TotalVolume                  float64   `json:"total_volume"`
	CirculatingSupply            float64   `json:"circulating_supply"`
	TotalSupply                  float64   `json:"total_supply"`
	LastUpdated                  time.Time `json:"last_updated"`
}

type GeckoCoins []GeckoCoin

type GeckoCoin struct {
	Id        string    `json:"id"`
	Symbol    string    `json:"symbol"`
	Name      string    `json:"name"`
	Platforms Platforms `json:"platforms"`
}

type Platforms map[string]string
