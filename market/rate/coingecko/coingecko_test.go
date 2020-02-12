package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/market/clients/coingecko"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeRates(t *testing.T) {
	tests := []struct {
		name      string
		prices    coingecko.CoinPrices
		wantRates blockatlas.Rates
	}{
		{
			"test normalize coingecko rate 1",
			coingecko.CoinPrices{
				{
					Symbol:       "cUSDC",
					CurrentPrice: 0.0021,
				},
				{
					Symbol:       "cREP",
					CurrentPrice: 0.02,
				},
			},
			blockatlas.Rates{
				blockatlas.Rate{Currency: "CUSDC", Rate: 1 / 0.0021, Timestamp: 333, Provider: id},
				blockatlas.Rate{Currency: "CREP", Rate: 1 / 0.02, Timestamp: 333, Provider: id},
			},
		},
		{
			"test normalize coingecko rate 2",
			coingecko.CoinPrices{
				{
					Symbol:       "cUSDC",
					CurrentPrice: 110.0021,
				},
				{
					Symbol:       "cREP",
					CurrentPrice: 110.02,
				},
			},
			blockatlas.Rates{
				blockatlas.Rate{Currency: "CUSDC", Rate: 1 / 110.0021, Timestamp: 123, Provider: id},
				blockatlas.Rate{Currency: "CREP", Rate: 1 / 110.02, Timestamp: 123, Provider: id},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRates := normalizeRates(tt.prices, id)
			now := time.Now().Unix()
			sort.Slice(gotRates, func(i, j int) bool {
				gotRates[i].Timestamp = now
				gotRates[j].Timestamp = now
				return gotRates[i].Rate < gotRates[j].Rate
			})
			sort.Slice(tt.wantRates, func(i, j int) bool {
				tt.wantRates[i].Timestamp = now
				tt.wantRates[j].Timestamp = now
				return tt.wantRates[i].Rate < tt.wantRates[j].Rate
			})
			if !assert.ObjectsAreEqualValues(gotRates, tt.wantRates) {
				t.Errorf("normalizeRates() = %v, want %v", gotRates, tt.wantRates)
			}
		})
	}
}
