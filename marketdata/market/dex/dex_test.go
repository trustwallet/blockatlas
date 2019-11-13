package dex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
	"reflect"
	"testing"
	"time"
)

func Test_normalizeTickers(t *testing.T) {
	type args struct {
		prices   []CoinPrice
		provider string
	}
	tests := []struct {
		name        string
		args        args
		wantTickers blockatlas.Tickers
	}{
		{
			"test normalize dex quote",
			args{prices: []CoinPrice{
				{
					BaseAssetName:      "RAVEN-F66",
					QuoteAssetName:     "BNB",
					LastPrice:          "0.00001082",
					PriceChangePercent: "-2.2500",
				},
				{
					BaseAssetName:      "SLV-986",
					QuoteAssetName:     "BNB",
					LastPrice:          "0.04494510",
					PriceChangePercent: "-5.3700",
				},
				{
					BaseAssetName:      "CBIX-3C9",
					QuoteAssetName:     "TAUD-888",
					LastPrice:          "0.00100235",
					PriceChangePercent: "5.2700",
				},
			},
				provider: "dex"},
			blockatlas.Tickers{
				blockatlas.Ticker{CoinName: "BNB", TokenId: "RAVEN-F66", CoinType: blockatlas.TypeToken, LastUpdate: time.Now(),
					Price: blockatlas.TickerPrice{
						Value:     big.NewFloat(0.00001082),
						Change24h: big.NewFloat(-2.2500),
						Currency:  "BNB",
						Provider:  "dex",
					},
				},
				blockatlas.Ticker{CoinName: "BNB", TokenId: "SLV-986", CoinType: blockatlas.TypeToken, LastUpdate: time.Now(),
					Price: blockatlas.TickerPrice{
						Value:     big.NewFloat(0.0449451),
						Change24h: big.NewFloat(-5.3700),
						Currency:  "BNB",
						Provider:  "dex",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTickers := normalizeTickers(tt.args.prices, tt.args.provider)
			for i, got := range gotTickers {
				want := tt.wantTickers[i]
				want.LastUpdate = got.LastUpdate
				if !reflect.DeepEqual(got, want) {
					t.Errorf("normalizeTickers() = %v, want %v", got, want)
				}
			}
		})
	}
}
