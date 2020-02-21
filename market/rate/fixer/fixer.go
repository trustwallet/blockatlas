package fixer

import (
	"github.com/trustwallet/blockatlas/market/rate"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
)

const (
	id = "fixer"
)

type Fixer struct {
	rate.Rate
	APIKey string
	blockatlas.Request
}

func InitRate(api string, apiKey string, updateTime string) rate.Provider {
	return &Fixer{
		Rate: rate.Rate{
			Id:         id,
			UpdateTime: updateTime,
		},
		Request: blockatlas.InitClient(api),
		APIKey:  apiKey,
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
