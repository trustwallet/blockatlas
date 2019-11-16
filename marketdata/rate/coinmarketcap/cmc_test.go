package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/marketdata/cmcmap"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
	"sort"
	"testing"
	"time"
)

func Test_normalizeRates(t *testing.T) {
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
								Price: 223.55,
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
				blockatlas.Rate{Currency: "BTC", Rate: new(big.Float).Quo(big.NewFloat(1), big.NewFloat(223.55)), Timestamp: 333},
				blockatlas.Rate{Currency: "ETH", Rate: new(big.Float).Quo(big.NewFloat(1), big.NewFloat(11.11)), Timestamp: 333},
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
				blockatlas.Rate{Currency: "BNB", Rate: new(big.Float).Quo(big.NewFloat(1), big.NewFloat(30.333)), Timestamp: 123},
				blockatlas.Rate{Currency: "XRP", Rate: new(big.Float).Quo(big.NewFloat(1), big.NewFloat(0.4687)), Timestamp: 123},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRates := normalizeRates(tt.prices, cmcmap.CmcMapping{})
			sort.SliceStable(gotRates, func(i, j int) bool {
				y := gotRates[i].Rate.Cmp(gotRates[j].Rate)
				return y <= 0
			})
			if !assert.ObjectsAreEqualValues(gotRates, tt.wantRates) {
				t.Errorf("normalizeRates() = %v, want %v", gotRates, tt.wantRates)
			}
		})
	}
}
