package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/market/clients/cmc"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"sort"
	"testing"
	"time"
)

func Test_normalizeTickers(t *testing.T) {
	mapping := cmc.CmcMapping{
		666: {{
			Coin: 1023,
			Id:   666,
			Type: "coin",
		}},
	}
	type args struct {
		prices   cmc.CoinPrices
		provider string
	}
	tests := []struct {
		name        string
		args        args
		wantTickers blockatlas.Tickers
	}{
		{
			"test normalize cmc quote",
			args{prices: cmc.CoinPrices{Data: []cmc.Data{
				{Coin: cmc.Coin{Symbol: "BTC", Id: 0}, LastUpdated: time.Unix(111, 0), Quote: cmc.Quote{
					USD: cmc.USD{Price: 223.55, PercentChange24h: 10}}},
				{Coin: cmc.Coin{Symbol: "ETH", Id: 60}, LastUpdated: time.Unix(333, 0), Quote: cmc.Quote{
					USD: cmc.USD{Price: 11.11, PercentChange24h: 20}}},
				{Coin: cmc.Coin{Symbol: "SWP", Id: 6969}, LastUpdated: time.Unix(444, 0), Quote: cmc.Quote{
					USD: cmc.USD{Price: 463.22, PercentChange24h: -3}},
					Platform: &cmc.Platform{Coin: cmc.Coin{Symbol: "ETH"}, TokenAddress: "0x8ce9137d39326ad0cd6491fb5cc0cba0e089b6a9"}},
				{Coin: cmc.Coin{Symbol: "ONE", Id: 666}, LastUpdated: time.Unix(555, 0), Quote: cmc.Quote{
					USD: cmc.USD{Price: 123.09, PercentChange24h: -1.4}},
					Platform: &cmc.Platform{Coin: cmc.Coin{Symbol: "BNB"}, TokenAddress: "0x8ce9137d39326ad0cd6491fb5cc0cba0e089b6a9"}}}},
				provider: "cmc"},
			blockatlas.Tickers{
				&blockatlas.Ticker{CoinName: "BTC", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(111, 0),
					Price: blockatlas.TickerPrice{
						Value:     223.55,
						Change24h: 10,
						Currency:  blockatlas.DefaultCurrency,
						Provider:  "cmc",
					},
				},
				&blockatlas.Ticker{CoinName: "ETH", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(333, 0),
					Price: blockatlas.TickerPrice{
						Value:     11.11,
						Change24h: 20,
						Currency:  blockatlas.DefaultCurrency,
						Provider:  "cmc",
					},
				},
				&blockatlas.Ticker{CoinName: "ETH", TokenId: "0x8ce9137d39326ad0cd6491fb5cc0cba0e089b6a9", CoinType: blockatlas.TypeToken, LastUpdate: time.Unix(444, 0),
					Price: blockatlas.TickerPrice{
						Value:     463.22,
						Change24h: -3,
						Currency:  blockatlas.DefaultCurrency,
						Provider:  "cmc",
					},
				},
				&blockatlas.Ticker{CoinName: "ONE", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(555, 0),
					Price: blockatlas.TickerPrice{
						Value:     123.09,
						Change24h: -1.4,
						Currency:  blockatlas.DefaultCurrency,
						Provider:  "cmc",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTickers := normalizeTickers(tt.args.prices, tt.args.provider, mapping)
			sort.SliceStable(gotTickers, func(i, j int) bool {
				return gotTickers[i].LastUpdate.Unix() < gotTickers[j].LastUpdate.Unix()
			})
			if !assert.Equal(t, len(tt.wantTickers), len(gotTickers)) {
				t.Fatal("invalid tickers length")
			}
			for i, obj := range tt.wantTickers {
				assert.Equal(t, obj, gotTickers[i])
			}
		})
	}
}
