package coingecko

import "time"

type CoinPrices []CoinPrice

type CoinPrice struct {
	CoinDetails
	CoingeckoPrice
}

type CoingeckoPrices []CoingeckoPrice

type CoingeckoPrice struct {
	Id                           string    `json:"id"`
	Symbol                       string    `json:"symbol"`
	Name                         string    `json:"name"`
	CurrentPrice                 float64   `json:"current_price"`
	PriceChange24h               float64   `json:"price_change_24h"`
	PriceChangePercentage24h     float64   `json:"price_change_percentage_24h"`
	MarketCapChange24h           float64   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64   `json:"market_cap_change_percentage_24h"`
	LastUpdated                  time.Time `json:"last_updated"`
}

type Coins []Coin

type Coin struct {
	CoinDetails
	CoingeckoCoin
}

type CoingeckoCoins []CoingeckoCoin

type CoingeckoCoin struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CoinDetails struct {
	AssetPlatformId string `json:"asset_platform_id"`
	ContractAddress string `json:"contract_address"`
}
