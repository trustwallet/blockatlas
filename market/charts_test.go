package market

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"reflect"
	"testing"
)

func Test_normalizeInfo(t *testing.T) {
	type args struct {
		prices   []blockatlas.ChartPrice
		maxItems int
	}
	tests := []struct {
		args     args
		wantInfo []blockatlas.ChartPrice
	}{
		{
			args{
				prices: []blockatlas.ChartPrice{
					{
						Price: 1,
						Date:  1578741541,
					},
					{
						Price: 1,
						Date:  1578741542,
					},
					{
						Price: 1,
						Date:  1578741549,
					},
					{
						Price: 1,
						Date:  1578741545,
					},
					{
						Price: 1,
						Date:  1578741547,
					},
					{
						Price: 1,
						Date:  1578741546,
					},
				},
				maxItems: 3,
			},
			[]blockatlas.ChartPrice{
				{
					Price: 1,
					Date:  1578741541,
				},
				{
					Price: 1,
					Date:  1578741546,
				},
				{
					Price: 1,
					Date:  1578741549,
				},
			},
		},
		{
			args{
				prices: []blockatlas.ChartPrice{
					{
						Price: 1,
						Date:  1578741541,
					},
					{
						Price: 1,
						Date:  1578741542,
					},
					{
						Price: 1,
						Date:  1578741549,
					},
					{
						Price: 1,
						Date:  1578741545,
					},
					{
						Price: 1,
						Date:  1578741547,
					},
					{
						Price: 1,
						Date:  1578741546,
					},
				},
				maxItems: 20,
			},
			[]blockatlas.ChartPrice{
				{
					Price: 1,
					Date:  1578741541,
				},
				{
					Price: 1,
					Date:  1578741542,
				},
				{
					Price: 1,
					Date:  1578741545,
				},
				{
					Price: 1,
					Date:  1578741546,
				},
				{
					Price: 1,
					Date:  1578741547,
				},
				{
					Price: 1,
					Date:  1578741549,
				},
			},
		},
		{
			args{
				prices: []blockatlas.ChartPrice{
					{
						Price: 1,
						Date:  1578741541,
					},
					{
						Price: 1,
						Date:  1578741542,
					},
					{
						Price: 1,
						Date:  1578741545,
					},
					{
						Price: 1,
						Date:  1578741547,
					},
					{
						Price: 1,
						Date:  1578741546,
					},
				},
				maxItems: 3,
			},
			[]blockatlas.ChartPrice{
				{
					Price: 1,
					Date:  1578741541,
				},
				{
					Price: 1,
					Date:  1578741545,
				},
				{
					Price: 1,
					Date:  1578741547,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run("Test prices normalize", func(t *testing.T) {
			assert.True(t, reflect.DeepEqual(normalizePrices(tt.args.prices, tt.args.maxItems), tt.wantInfo))
		})
	}
}
