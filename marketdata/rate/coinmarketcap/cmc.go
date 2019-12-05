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
			UpdateTime: getUpdateTime(),
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

func getUpdateTime() string {
	updateTime := viper.GetString("market.cmc.rate_update_time")
	if len(updateTime) == 0 {
		return viper.GetString("market.rate_update_time")
	}
	return updateTime
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
