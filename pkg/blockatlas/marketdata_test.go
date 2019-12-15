package blockatlas

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestTicker_ApplyRate(t *testing.T) {
	type args struct {
		price            float64
		percentChange24h float64
		rate             float64
		rateChange24h    float64
		currency         string
	}
	type want struct {
		price            float64
		percentChange24h float64
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			"apply rate 1",
			args{443, -0.44, 344, -0.33, "BTC"},
			want{152392, -0.10999999999999999},
		}, {
			"apply rate 2",
			args{1111.22, 3.04, 0.88, -2.12, "BTC"},
			want{977.8736, 5.16},
		}, {
			"apply rate 3",
			args{3.33, -1.02, 22.3, 1.02, "BTC"},
			want{74.259, -2.04},
		}, {
			"apply rate 4",
			args{9.3332, -2, 22, -0.555, "BTC"},
			want{205.3304, -1.4449999999999998},
		}, {
			"apply rate 5",
			args{0.00000001, 0, 0.2, 1.333, "BTC"},
			want{0.000000002, -1.333},
		}, {
			"apply rate 6",
			args{0.00000001, -0.0003, 10, -0.0003, "BTC"},
			want{0.0000001, 0},
		}, {
			"apply for same currency",
			args{0.33333, -0.0003, 10, -0.0003, "USD"},
			want{0.33333, -0.0003},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			ticker := &Ticker{
				Price: TickerPrice{
					Value:     tt.args.price,
					Change24h: tt.args.percentChange24h,
					Currency:  DefaultCurrency,
				},
			}
			ticker.ApplyRate(tt.args.currency, tt.args.rate, big.NewFloat(tt.args.rateChange24h))
			assert.Equal(t, tt.want.price, ticker.Price.Value)
			assert.Equal(t, tt.want.percentChange24h, ticker.Price.Change24h)
		})
	}
}
