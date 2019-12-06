package cmc

import (
	"github.com/trustwallet/blockatlas/marketdata/cmcmap"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

const (
	id = "cmc"
)

type Cmc struct {
	rate.Rate
}

func InitRate(api string, apiKey string, updateTime string) rate.Provider {
	cmc := &Cmc{
		Rate: rate.Rate{
			Id:         id,
			Request:    blockatlas.InitClient(api),
			UpdateTime: updateTime,
		},
	}
	cmc.Headers["X-CMC_PRO_API_KEY"] = apiKey
	return cmc
}

func (c *Cmc) FetchLatestRates() (rates blockatlas.Rates, err error) {
	cmap, err := cmcmap.GetCmcMap()
	if err != nil {
		return nil, err
	}
	var prices CoinPrices
	err = c.Get(&prices, "v1/cryptocurrency/listings/latest",
		url.Values{"limit": {"5000"}, "convert": {blockatlas.DefaultCurrency}})
	if err != nil {
		return
	}
	rates = normalizeRates(prices, cmap, c.GetId())
	return
}

func normalizeRates(prices CoinPrices, cmap cmcmap.CmcMapping, provider string) (rates blockatlas.Rates) {
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
