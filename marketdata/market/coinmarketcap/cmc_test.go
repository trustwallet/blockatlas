package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/marketdata/cmcmap"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeTickers(t *testing.T) {
	type args struct {
		prices   CoinPrices
		provider string
	}
	tests := []struct {
		name        string
		args        args
		wantTickers blockatlas.Tickers
	}{
		{
			"test normalize cmc quote",
			args{prices: CoinPrices{Data: []Data{
				{Coin: Coin{Symbol: "BTC"}, LastUpdated: time.Unix(111, 0), Quote: Quote{
					USD: USD{Price: 223.55, PercentChange24h: 10}}},
				{Coin: Coin{Symbol: "ETH"}, LastUpdated: time.Unix(333, 0), Quote: Quote{
					USD: USD{Price: 11.11, PercentChange24h: 20}}},
				{Coin: Coin{Symbol: "SWP"}, LastUpdated: time.Unix(444, 0), Quote: Quote{
					USD: USD{Price: 463.22, PercentChange24h: -3}},
					Platform: &Platform{Coin: Coin{Symbol: "ETH"}, TokenAddress: "0x8ce9137d39326ad0cd6491fb5cc0cba0e089b6a9"}}}},
				provider: "cmc"},
			blockatlas.Tickers{
				blockatlas.Ticker{CoinName: "BTC", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(111, 0),
					Price: blockatlas.TickerPrice{
						Value:     223.55,
						Change24h: 10,
						Currency:  "USD",
						Provider:  "cmc",
					},
				},
				blockatlas.Ticker{CoinName: "ETH", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(333, 0),
					Price: blockatlas.TickerPrice{
						Value:     11.11,
						Change24h: 20,
						Currency:  "USD",
						Provider:  "cmc",
					},
				},
				blockatlas.Ticker{CoinName: "ETH", TokenId: "0x8ce9137d39326ad0cd6491fb5cc0cba0e089b6a9", CoinType: blockatlas.TypeToken, LastUpdate: time.Unix(444, 0),
					Price: blockatlas.TickerPrice{
						Value:     463.22,
						Change24h: -3,
						Currency:  "USD",
						Provider:  "cmc",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTickers := normalizeTickers(tt.args.prices, tt.args.provider, cmcmap.CmcMapping{})
			sort.SliceStable(gotTickers, func(i, j int) bool {
				return gotTickers[i].LastUpdate.Unix() < gotTickers[j].LastUpdate.Unix()
			})
			if !assert.ObjectsAreEqualValues(gotTickers, tt.wantTickers) {
				t.Errorf("normalizeTickers() = %v, want %v", gotTickers, tt.wantTickers)
			}
		})
	}
}
