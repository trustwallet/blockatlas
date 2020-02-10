package cmc

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/market/clients/cmc"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"reflect"
	"sort"
	"testing"
	"time"
)

func Test_normalizeInfo(t *testing.T) {
	type args struct {
		currency string
		cmcCoin  uint
		data     cmc.ChartInfo
	}
	tests := []struct {
		name     string
		args     args
		wantInfo blockatlas.ChartCoinInfo
	}{
		{
			"test normalize cmc chart info 1",
			args{
				currency: "USD",
				cmcCoin:  1,
				data: cmc.ChartInfo{
					Data: cmc.ChartInfoData{
						Rank:              1,
						CirculatingSupply: 111,
						TotalSupply:       222,
						Quotes: map[string]cmc.ChartInfoQuote{
							"USD": {Price: 333, Volume24: 444, MarketCap: 555},
						},
					},
				},
			},
			blockatlas.ChartCoinInfo{
				Vol24:             444,
				MarketCap:         555,
				CirculatingSupply: 111,
				TotalSupply:       222,
			},
		},
		{
			"test normalize cmc chart info 2",
			args{
				currency: "EUR",
				cmcCoin:  2,
				data: cmc.ChartInfo{
					Data: cmc.ChartInfoData{
						Rank:              2,
						CirculatingSupply: 111,
						TotalSupply:       222,
						Quotes: map[string]cmc.ChartInfoQuote{
							"EUR": {Price: 333, Volume24: 444, MarketCap: 555},
						},
					},
				},
			},
			blockatlas.ChartCoinInfo{
				Vol24:             444,
				MarketCap:         555,
				CirculatingSupply: 111,
				TotalSupply:       222,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo, err := normalizeInfo(tt.args.currency, tt.args.cmcCoin, tt.args.data)
			assert.Nil(t, err)
			assert.True(t, reflect.DeepEqual(tt.wantInfo, gotInfo))
		})
	}
}

func Test_normalizeCharts(t *testing.T) {
	type args struct {
		currency string
		symbol   string
		charts   cmc.Charts
	}

	timeStr1 := "2019-12-19T18:27:23.453Z"
	d1, _ := time.Parse(time.RFC3339, timeStr1)
	timeStr2 := "2019-11-19T18:27:23.453Z"
	d2, _ := time.Parse(time.RFC3339, timeStr2)
	tests := []struct {
		name     string
		args     args
		wantInfo blockatlas.ChartData
	}{
		{
			"test normalize cmc chart 1",
			args{
				currency: "USD",
				symbol:   "BTC",
				charts: cmc.Charts{
					Data: cmc.ChartQuotes{
						timeStr1: cmc.ChartQuoteValues{
							"USD": []float64{111, 222, 333},
						},
					},
				},
			},
			blockatlas.ChartData{
				Prices: []blockatlas.ChartPrice{
					{
						Price: 111,
						Date:  d1.Unix(),
					},
				},
			},
		},
		{
			"test normalize cmc chart 2",
			args{
				currency: "EUR",
				symbol:   "BTC",
				charts: cmc.Charts{
					Data: cmc.ChartQuotes{
						timeStr1: cmc.ChartQuoteValues{
							"EUR": []float64{333, 444, 555},
						},
						timeStr2: cmc.ChartQuoteValues{
							"EUR": []float64{555, 666, 777},
						},
					},
				},
			},
			blockatlas.ChartData{
				Prices: []blockatlas.ChartPrice{
					{
						Price: 333,
						Date:  d1.Unix(),
					},
					{
						Price: 555,
						Date:  d2.Unix(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo := normalizeCharts(tt.args.currency, tt.args.charts)
			sort.SliceStable(gotInfo.Prices, func(i, j int) bool {
				return gotInfo.Prices[i].Price < gotInfo.Prices[j].Price
			})
			if !assert.ObjectsAreEqualValues(tt.wantInfo, gotInfo) {
				t.Errorf("normalizeCharts() = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func Test_getInterval(t *testing.T) {
	tests := []struct {
		name     string
		days     int
		wantInfo string
	}{
		{
			"test getInterval 1",
			1,
			"5m",
		},
		{
			"test getInterval 2",
			5,
			"5m",
		},
		{
			"test getInterval 3",
			7,
			"15m",
		},
		{
			"test getInterval 4",
			8,
			"15m",
		},
		{
			"test getInterval 5",
			30,
			"1h",
		},
		{
			"test getInterval 6",
			40,
			"1h",
		},
		{
			"test getInterval 7",
			90,
			"2h",
		},
		{
			"test getInterval 8",
			120,
			"2h",
		},
		{
			"test getInterval 9",
			360,
			"1d",
		},
		{
			"test getInterval 10",
			800,
			"1d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo := getInterval(tt.days)
			assert.Equal(t, tt.wantInfo, gotInfo)
		})
	}
}
