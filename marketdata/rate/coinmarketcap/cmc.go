package cmc

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/cmcmap"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Cmc struct {
	rate.Rate
}

func InitRate() rate.Provider {
	cmc := &Cmc{
		Rate: rate.Rate{
			Id:         "cmc",
			Request:    blockatlas.InitClient(viper.GetString("market.cmc.api")),
			UpdateTime: viper.GetString("market.cmc.rate_update_time"),
		},
	}
	cmc.Headers["X-CMC_PRO_API_KEY"] = viper.GetString("market.cmc.api_key")
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
		currency := price.Symbol
		cmcCoin, _, err := cmap.GetCoin(price.Id)
		if err == nil {
			currency = cmcCoin.Symbol
		}
		rates = append(rates, blockatlas.Rate{
			Currency:  currency,
			Rate:      1.0 / price.Quote.USD.Price,
			Timestamp: price.LastUpdated.Unix(),
			Provider:  provider,
		})
	}
	return
}
