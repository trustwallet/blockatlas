package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/marketdata/cmcmap"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeRates(t *testing.T) {
	provider := "cmc"
	tests := []struct {
		name      string
		prices    CoinPrices
		wantRates blockatlas.Rates
	}{
		{
			"test normalize cmc rate 1",
			CoinPrices{
				Data: []Data{
					{
						Coin: Coin{
							Symbol: "BTC",
						},
						Quote: Quote{
							USD: USD{
								Price: 223.5,
							},
						},
						LastUpdated: time.Unix(333, 0),
					},
					{
						Coin: Coin{
							Symbol: "ETH",
						},
						Quote: Quote{
							USD: USD{
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
			CoinPrices{
				Data: []Data{
					{
						Coin: Coin{
							Symbol: "BNB",
						},
						Quote: Quote{
							USD: USD{
								Price: 30.333,
							},
						},
						LastUpdated: time.Unix(123, 0),
					},
					{
						Coin: Coin{
							Symbol: "XRP",
						},
						Quote: Quote{
							USD: USD{
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
			gotRates := normalizeRates(tt.prices, cmcmap.CmcMapping{}, provider)
			sort.SliceStable(gotRates, func(i, j int) bool {
				return gotRates[i].Rate < gotRates[j].Rate
			})
			if !assert.ObjectsAreEqualValues(gotRates, tt.wantRates) {
				t.Errorf("normalizeRates() = %v, want %v", gotRates, tt.wantRates)
			}
		})
	}
}
