package coingecko

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinIds(t *testing.T) {
	tests := []struct {
		name     string
		coins    GeckoCoins
		expected []string
	}{
		{
			name: "test construct coin Ids",
			coins: GeckoCoins{
				GeckoCoin{
					Id:        "ethtereum",
					Symbol:    "eth",
					Name:      "eth",
					Platforms: nil,
				},
				GeckoCoin{
					Id:        "bitcoin",
					Symbol:    "btc",
					Name:      "btc",
					Platforms: nil,
				},
			},
			expected: []string{
				"ethtereum",
				"bitcoin",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.coins.coinIds())
		})
	}
}
