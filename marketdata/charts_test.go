package marketdata

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"reflect"
	"testing"
)

func Test_normalizeInfo(t *testing.T) {
	tests := []struct {
		args     []blockatlas.ChartPrice
		wantInfo []blockatlas.ChartPrice
	}{
		{
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
	}
	for _, tt := range tests {
		t.Run("Test prices normalize", func(t *testing.T) {
			normalizePrices(tt.args)
			assert.True(t, reflect.DeepEqual(tt.args, tt.wantInfo))
		})
	}
}
