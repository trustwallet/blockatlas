package coingecko

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/market/clients/coingecko"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"reflect"
	"testing"
	"time"
)

func Test_normalizeInfo(t *testing.T) {
	type args struct {
		data coingecko.CoinPrice
	}
	tests := []struct {
		name     string
		args     args
		wantInfo blockatlas.ChartCoinInfo
	}{
		{
			"test normalize coingecko chart info 1",
			args{
				data: coingecko.CoinPrice{
					MarketCap:         555,
					TotalVolume:       444,
					CirculatingSupply: 111,
					TotalSupply:       222,
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
			gotInfo := normalizeInfo(tt.args.data)
			assert.True(t, reflect.DeepEqual(tt.wantInfo, gotInfo))
		})
	}
}

func Test_normalizeCharts(t *testing.T) {
	type args struct {
		charts coingecko.Charts
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
			"test normalize coingecko chart 1",
			args{
				charts: coingecko.Charts{
					Prices: []coingecko.ChartVolume{
						[]float64{float64(d1.UnixNano() / int64(time.Millisecond)), 222},
						[]float64{float64(d2.UnixNano() / int64(time.Millisecond)), 333},
					},
				},
			},
			blockatlas.ChartData{
				Prices: []blockatlas.ChartPrice{
					{
						Price: 222,
						Date:  d1.Unix(),
					},
					{
						Price: 333,
						Date:  d2.Unix(),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo := normalizeCharts(tt.args.charts)
			assert.True(t, reflect.DeepEqual(tt.wantInfo, gotInfo))
		})
	}
}
