package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/marketdata/clients/cmc"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeRates(t *testing.T) {
	provider := "cmc"
	tests := []struct {
		name      string
		prices    cmc.CoinPrices
		wantRates blockatlas.Rates
	}{
		{
			"test normalize cmc rate 1",
			cmc.CoinPrices{
				Data: []cmc.Data{
					{
						Coin: cmc.Coin{
							Symbol: "BTC",
						},
						Quote: cmc.Quote{
							USD: cmc.USD{
								Price: 223.5,
							},
						},
						LastUpdated: time.Unix(333, 0),
					},
					{
						Coin: cmc.Coin{
							Symbol: "ETH",
						},
						Quote: cmc.Quote{
							USD: cmc.USD{
								Price: 11.11,
							},
						},
						LastUpdated: time.Unix(333, 0),
					},
				},
			},
			blockatlas.Rates{
				blockatlas.Rate{Currency: "BTC", Rate: 1 / 223.5, Timestamp: 333, Provider: provider},
				blockatlas.Rate{Currency: "ETH", Rate: 1 / 11.11, Timestamp: 333, Provider: provider},
			},
		},
		{
			"test normalize cmc rate 2",
			cmc.CoinPrices{
				Data: []cmc.Data{
					{
						Coin: cmc.Coin{
							Symbol: "BNB",
						},
						Quote: cmc.Quote{
							USD: cmc.USD{
								Price: 30.333,
							},
						},
						LastUpdated: time.Unix(123, 0),
					},
					{
						Coin: cmc.Coin{
							Symbol: "XRP",
						},
						Quote: cmc.Quote{
							USD: cmc.USD{
								Price: 0.4687,
							},
						},
						LastUpdated: time.Unix(123, 0),
					},
				},
			},
			blockatlas.Rates{
				blockatlas.Rate{Currency: "BNB", Rate: 1 / 30.333, Timestamp: 123, Provider: provider},
				blockatlas.Rate{Currency: "XRP", Rate: 1 / 0.4687, Timestamp: 123, Provider: provider},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRates := normalizeRates(tt.prices, cmc.CmcMapping{}, provider)
			sort.SliceStable(gotRates, func(i, j int) bool {
				return gotRates[i].Rate < gotRates[j].Rate
			})
			if !assert.ObjectsAreEqualValues(gotRates, tt.wantRates) {
				t.Errorf("normalizeRates() = %v, want %v", gotRates, tt.wantRates)
			}
		})
	}
}
