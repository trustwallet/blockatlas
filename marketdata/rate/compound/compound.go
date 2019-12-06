package compound

import (
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
	"time"
)

const (
	compound = "compound"
)

type Compound struct {
	rate.Rate
}

func InitRate(api string, updateTime string) rate.Provider {
	return &Compound{
		Rate: rate.Rate{
			Id:         compound,
			Request:    blockatlas.InitClient(api),
			UpdateTime: updateTime,
		},
	}
}

func (c *Compound) FetchLatestRates() (rates blockatlas.Rates, err error) {
	var coinPrices CoinPrices
	err = c.Get(&coinPrices, "v2/ctoken", nil)
	if err != nil {
		return
	}
	rates = normalizeRates(coinPrices, c.GetId())
	return
}

func normalizeRates(coinPrices CoinPrices, provider string) (rates blockatlas.Rates) {
	for _, cToken := range coinPrices.Data {
		rates = append(rates, blockatlas.Rate{
			Currency:  strings.ToUpper(cToken.Symbol),
			Rate:      1.0 / cToken.ExchangeRate.Value,
			Timestamp: time.Now().Unix(),
			Provider:  provider,
		})
	}
	return
}
