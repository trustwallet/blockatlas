package cmc

import (
	"github.com/trustwallet/blockatlas/market/clients/cmc"
	"github.com/trustwallet/blockatlas/market/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
)

const (
	id = "cmc"
)

type Cmc struct {
	rate.Rate
	mapApi string
	client *cmc.Client
}

func InitRate(api string, apiKey string, mapApi string, updateTime string) rate.Provider {
	cmc := &Cmc{
		Rate: rate.Rate{
			Id:         id,
			UpdateTime: updateTime,
		},
		mapApi: mapApi,
		client: cmc.NewClient(api, apiKey),
	}
	return cmc
}

func (c *Cmc) FetchLatestRates() (rates blockatlas.Rates, err error) {
	prices, err := c.client.GetData()
	if err != nil {
		return
	}
	rates = normalizeRates(prices, c.GetId())
	return
}

func normalizeRates(prices cmc.CoinPrices, provider string) (rates blockatlas.Rates) {
	for _, price := range prices.Data {
		if price.Platform != nil {
			continue
		}
		rates = append(rates, blockatlas.Rate{
			Currency:         price.Symbol,
			Rate:             1.0 / price.Quote.USD.Price,
			Timestamp:        price.LastUpdated.Unix(),
			PercentChange24h: big.NewFloat(price.Quote.USD.PercentChange24h),
			Provider:         provider,
		})
	}
	return
}
