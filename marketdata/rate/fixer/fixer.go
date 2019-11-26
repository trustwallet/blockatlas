package fixer

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/marketdata/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

type Fixer struct {
	rate.Rate
	APIKey string
}

func InitRate() rate.Provider {
	return &Fixer{
		Rate: rate.Rate{
			Id:         "fixer",
			Request:    blockatlas.InitClient(viper.GetString("market.fixer.api")),
			UpdateTime: viper.GetString("market.fixer.rate_update_time"),
		},
		APIKey: viper.GetString("market.fixer.api_key"),
	}
}

func (f *Fixer) FetchLatestRates() (rates blockatlas.Rates, err error) {
	values := url.Values{
		"access_key": {f.APIKey},
		"base":       {blockatlas.DefaultCurrency}, // Base USD supported only in paid api
	}
	var latest Latest
	err = f.Get(&latest, "latest", values)
	if err != nil {
		return
	}
	rates = normalizeRates(latest, f.GetId())
	return
}

func normalizeRates(latest Latest, provider string) (rates blockatlas.Rates) {
	for currency, rate := range latest.Rates {
		rates = append(rates, blockatlas.Rate{Currency: currency, Rate: rate, Timestamp: latest.Timestamp, Provider: provider})
	}
	return
}
