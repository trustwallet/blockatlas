package compound

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/market/clients/compound"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeTickers(t *testing.T) {
	type args struct {
		prices   compound.CoinPrices
		provider string
	}
	tests := []struct {
		name        string
		args        args
		wantTickers blockatlas.Tickers
	}{
		{
			"test normalize compound quote",
			args{prices: compound.CoinPrices{Data: []compound.CToken{
				{
					TokenAddress:    "0x39aa39c021dfbae8fac545936693ac917d5e7563",
					Symbol:          "cUSDC",
					UnderlyingPrice: compound.Amount{Value: 0.0021},
				},
				{
					TokenAddress:    "0x158079ee67fce2f58472a96584a73c7ab9ac95c1",
					Symbol:          "cREP",
					UnderlyingPrice: compound.Amount{Value: 0.02},
				},
			}}, provider: id},
			blockatlas.Tickers{
				&blockatlas.Ticker{CoinName: "ETH", TokenId: "0x39aa39c021dfbae8fac545936693ac917d5e7563", CoinType: blockatlas.TypeToken, LastUpdate: time.Unix(222, 0),
					Price: blockatlas.TickerPrice{
						Value:    0.0021,
						Currency: coin.Coins[coin.ETH].Symbol,
						Provider: id,
					},
				},
				&blockatlas.Ticker{CoinName: "ETH", TokenId: "0x158079ee67fce2f58472a96584a73c7ab9ac95c1", CoinType: blockatlas.TypeToken, LastUpdate: time.Unix(444, 0),
					Price: blockatlas.TickerPrice{
						Value:    0.02,
						Currency: coin.Coins[coin.ETH].Symbol,
						Provider: id,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTickers := normalizeTickers(tt.args.prices, tt.args.provider)
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
