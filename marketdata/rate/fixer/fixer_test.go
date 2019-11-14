package fixer

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
	"sort"
	"testing"
	"time"
)

func Test_normalizeRates(t *testing.T) {
	tests := []struct {
		name      string
		latest    Latest
		wantRates blockatlas.Rates
	}{
		{
			"test normalize fixer rate 1",
			Latest{
				Timestamp: 123,
				Rates:     map[string]float64{"USD": 22.111, "BRL": 33.2, "BTC": 44.99},
				UpdatedAt: time.Now(),
			},
			blockatlas.Rates{
				blockatlas.Rate{Currency: "USD", Rate: big.NewFloat(22.111), Timestamp: 123},
				blockatlas.Rate{Currency: "BRL", Rate: big.NewFloat(33.2), Timestamp: 123},
				blockatlas.Rate{Currency: "BTC", Rate: big.NewFloat(44.99), Timestamp: 123},
			},
		},
		{
			"test normalize fixer rate 2",
			Latest{
				Timestamp: 333,
				Rates:     map[string]float64{"LSK": 123.321, "IFC": 34.973, "DUO": 998.3},
				UpdatedAt: time.Now(),
			},
			blockatlas.Rates{
				blockatlas.Rate{Currency: "IFC", Rate: big.NewFloat(34.973), Timestamp: 333},
				blockatlas.Rate{Currency: "LSK", Rate: big.NewFloat(123.321), Timestamp: 333},
				blockatlas.Rate{Currency: "DUO", Rate: big.NewFloat(998.3), Timestamp: 333},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRates := normalizeRates(tt.latest)
			sort.SliceStable(gotRates, func(i, j int) bool {
				y := gotRates[i].Rate.Cmp(gotRates[j].Rate)
				return y <= 0
			})
			if !assert.ObjectsAreEqualValues(gotRates, tt.wantRates) {
				t.Errorf("normalizeRates() = %v, want %v", gotRates, tt.wantRates)
			}
		})
	}
}
