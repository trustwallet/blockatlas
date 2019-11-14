package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
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
				{Coin: Coin{Symbol: "BTC"}, LastUpdated: time.Unix(333, 0), Quote: Quote{
					USD: USD{Price: 223.55, PercentChange24h: 10}}},
				{Coin: Coin{Symbol: "ETH"}, LastUpdated: time.Unix(333, 0), Quote: Quote{
					USD: USD{Price: 11.11, PercentChange24h: 20}}}}},
				provider: "cmc"},
			blockatlas.Tickers{
				blockatlas.Ticker{CoinName: "BTC", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(333, 0),
					Price: blockatlas.TickerPrice{
						Value:     big.NewFloat(223.55),
						Change24h: big.NewFloat(10),
						Currency:  "USD",
						Provider:  "cmc",
					},
				},
				blockatlas.Ticker{CoinName: "ETH", CoinType: blockatlas.TypeCoin, LastUpdate: time.Unix(333, 0),
					Price: blockatlas.TickerPrice{
						Value:     big.NewFloat(11.11),
						Change24h: big.NewFloat(20),
						Currency:  "USD",
						Provider:  "cmc",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTickers := normalizeTickers(tt.args.prices, tt.args.provider)
			if !assert.ObjectsAreEqualValues(gotTickers, tt.wantTickers) {
				t.Errorf("normalizeTickers() = %v, want %v", gotTickers, tt.wantTickers)
			}
		})
	}
}

func Test_percentageChange(t *testing.T) {
	type args struct {
		value   float64
		percent float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test 10% for value 10", args{100, 10}, 10},
		{"test 100% for value 30", args{100, 30}, 30},
		{"test 300% for value 100", args{300, 100}, 300},
		{"test 55% for value 4", args{55, 4}, 2.2},
		{"test 10% for value 15", args{10, 15}, 1.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percentageChange(tt.args.value, tt.args.percent); got != tt.want {
				t.Errorf("percentageChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
