package blockatlas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTicker_ApplyRate(t *testing.T) {
	tests := []struct {
		name  string
		price float64
		rate  float64
		want  float64
	}{
		{"apply rate 1", 443, 344, 152392},
		{"apply rate 2", 1111.22, 0.88, 977.8736},
		{"apply rate 3", 3.33, 22.3, 74.259},
		{"apply rate 4", 9.3332, 22, 205.3304},
		{"apply rate 5", 0.00000001, 0.2, 0.000000002},
		{"apply rate 6", 0.00000001, 10, 0.0000001},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			ticker := &Ticker{
				Price: TickerPrice{Value: tt.price},
			}
			ticker.ApplyRate(tt.rate, DefaultCurrency)
			assert.Equal(t, tt.want, ticker.Price.Value)
		})
	}
}
