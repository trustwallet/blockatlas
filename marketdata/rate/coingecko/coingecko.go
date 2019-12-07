package coingecko

import (
	"github.com/trustwallet/blockatlas/marketdata/coingecko"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

const (
	id = "coingecko"
)

type Coingecko struct {
	client *coingecko.Client
	rate.Rate
}

func InitRate(api string, updateTime string) rate.Provider {
	return &Coingecko{
		client: coingecko.NewClient(api),
		Rate: rate.Rate{
			Id:         id,
			UpdateTime: updateTime,
		},
	}
}

func (c *Coingecko) FetchLatestRates() (rates blockatlas.Rates, err error) {
	coins, err := c.client.FetchLatestRates()
	if err != nil {
		return
	}

	rates = normalizeRates(coins, c.GetId())
	return
}

func normalizeRates(coinPrices coingecko.CoinPrices, provider string) (rates blockatlas.Rates) {
	for _, price := range coinPrices {
		rates = append(rates, blockatlas.Rate{
			Currency:  strings.ToUpper(price.Symbol),
			Rate:      1.0 / price.CurrentPrice,
			Timestamp: price.LastUpdated.Unix(),
			Provider:  provider,
		})
	}
	return
}
