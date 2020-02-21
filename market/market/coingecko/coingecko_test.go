package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/market/clients/coingecko"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeTickers(t *testing.T) {
	coins := coingecko.GeckoCoins{
		coingecko.GeckoCoin{
			Id:        "ethtereum",
			Symbol:    "eth",
			Name:      "eth",
			Platforms: nil,
		},
		coingecko.GeckoCoin{
			Id:        "bitcoin",
			Symbol:    "btc",
			Name:      "btc",
			Platforms: nil,
		},
		coingecko.GeckoCoin{
			Id:     "cREP",
			Symbol: "cREP",
			Name:   "cREP",
			Platforms: coingecko.Platforms{
				"ethtereum": "0x158079ee67fce2f58472a96584a73c7ab9ac95c1",
			},
		},
		coingecko.GeckoCoin{
			Id:     "cUSDC",
			Symbol: "cUSDC",
			Name:   "cUSDC",
			Platforms: coingecko.Platforms{
				"ethtereum": "0x39aa39c021dfbae8fac545936693ac917d5e7563",
			},
		},
	}

	m := Market{}
	m.cache = coingecko.NewCache(coins)
	type args struct {
		prices   coingecko.CoinPrices
		provider string
	}
	tests := []struct {
		name        string
		args        args
		wantTickers blockatlas.Tickers
	}{
		{
			"test normalize coingecko quote",
			args{prices: coingecko.CoinPrices{
				{
					Id:           "cUSDC",
					Symbol:       "cUSDC",
					CurrentPrice: 0.0021,
				},
				{
					Id:           "cREP",
					Symbol:       "cREP",
					CurrentPrice: 0.02,
				},
			}, provider: id},
			blockatlas.Tickers{
				&blockatlas.Ticker{CoinName: "ETH", TokenId: "0x39aa39c021dfbae8fac545936693ac917d5e7563", CoinType: blockatlas.TypeToken, LastUpdate: time.Unix(222, 0),
					Price: blockatlas.TickerPrice{
						Value:    0.0021,
						Currency: blockatlas.DefaultCurrency,
						Provider: id,
					},
				},
				&blockatlas.Ticker{CoinName: "ETH", TokenId: "0x158079ee67fce2f58472a96584a73c7ab9ac95c1", CoinType: blockatlas.TypeToken, LastUpdate: time.Unix(444, 0),
					Price: blockatlas.TickerPrice{
						Value:    0.02,
						Currency: blockatlas.DefaultCurrency,
						Provider: id,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTickers := m.normalizeTickers(tt.args.prices, tt.args.provider)
			now := time.Now()
			sort.Slice(gotTickers, func(i, j int) bool {
				gotTickers[i].LastUpdate = now
				gotTickers[j].LastUpdate = now
				return gotTickers[i].Coin > gotTickers[j].Coin
			})
			sort.Slice(tt.wantTickers, func(i, j int) bool {
				tt.wantTickers[i].LastUpdate = now
				tt.wantTickers[j].LastUpdate = now
				return tt.wantTickers[i].Coin > tt.wantTickers[j].Coin
			})
			assert.Equal(t, tt.wantTickers, gotTickers)
		})
	}
}
