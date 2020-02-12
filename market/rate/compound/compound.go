package compound

import (
	c "github.com/trustwallet/blockatlas/market/clients/compound"
	"github.com/trustwallet/blockatlas/market/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
	"time"
)

const (
	compound = "compound"
)

type Compound struct {
	rate.Rate
	client *c.Client
}

func InitRate(api string, updateTime string) rate.Provider {
	return &Compound{
		Rate: rate.Rate{
			Id:         compound,
			UpdateTime: updateTime,
		},
		client: c.NewClient(api),
	}
}

func (c *Compound) FetchLatestRates() (rates blockatlas.Rates, err error) {
	coinPrices, err := c.client.GetData()
	if err != nil {
		return
	}
	rates = normalizeRates(coinPrices, c.GetId())
	return
}

func normalizeRates(coinPrices c.CoinPrices, provider string) (rates blockatlas.Rates) {
	for _, cToken := range coinPrices.Data {
		rates = append(rates, blockatlas.Rate{
			Currency:  strings.ToUpper(cToken.Symbol),
			Rate:      1.0 / cToken.UnderlyingPrice.Value,
			Timestamp: time.Now().Unix(),
			Provider:  provider,
		})
	}
	return
}
