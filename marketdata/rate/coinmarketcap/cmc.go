package cmc

import (
	"github.com/trustwallet/blockatlas/marketdata/clients/cmc"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
	cmap, err := cmc.GetCmcMap(c.mapApi)
	if err != nil {
		return nil, err
	}
	prices, err := c.client.GetData()
	if err != nil {
		return
	}
	rates = normalizeRates(prices, cmap, c.GetId())
	return
}

func normalizeRates(prices cmc.CoinPrices, cmap cmc.CmcMapping, provider string) (rates blockatlas.Rates) {
	for _, price := range prices.Data {
		if price.Platform != nil {
			continue
		}
		rates = append(rates, blockatlas.Rate{
			Currency:  price.Symbol,
			Rate:      1.0 / price.Quote.USD.Price,
			Timestamp: price.LastUpdated.Unix(),
			Provider:  provider,
		})
	}
	return
}
