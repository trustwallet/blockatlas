package fixer

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"time"
)

type Fixer struct {
	rate.Rate
	APIKey string
}

func InitRate() rate.Provider {
	return &Fixer{
		Rate: rate.Rate{
			Id:         "Fixer",
			Request:    blockatlas.InitClient(viper.GetString("market.fixer_api")),
			UpdateTime: time.Second * 30,
		},
		APIKey: viper.GetString("market.fixer_key"),
	}
}

func (f *Fixer) FetchLatestRates() (rates blockatlas.Rates, err error) {
	values := url.Values{
		"access_key": {f.APIKey},
		"base":       {"USD"}, // Base USD supported only in paid api
	}
	var latest Latest
	err = f.Get(&latest, "latest", values)
	if err != nil {
		return
	}
	rates = normalizeRates(latest)
	return
}

func normalizeRates(latest Latest) (rates blockatlas.Rates) {
	for currency, rate := range latest.Rates {
		rates = append(rates, blockatlas.Rate{Currency: currency, Rate: rate, Timestamp: latest.Timestamp})
	}
	return
}
