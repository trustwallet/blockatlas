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
			assert.Equal(t, tt.expected, coinIds(tt.coins))
		})
	}
}

func Test_prepareCache(t *testing.T) {
	tests := []struct {
		name     string
		coins    GeckoCoins
		expected map[string][]CoinResult
	}{
		{
			name: "test prepare cache map",
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
				GeckoCoin{
					Id:     "0x",
					Symbol: "btc",
					Name:   "0x",
					Platforms: Platforms{
						"ethtereum": "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
					},
				},
				GeckoCoin{
					Id:     "usdt",
					Symbol: "usdt",
					Name:   "usdt",
					Platforms: Platforms{
						"ethtereum": "0xdac17f958d2ee523a2206206994597c13d831ec7",
					},
				},
			},
			expected: map[string][]CoinResult{
				"0x": {
					CoinResult{
						Symbol:   "eth",
						TokenId:  "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
						CoinType: "token",
					},
				},
				"usdt": {
					CoinResult{
						Symbol:   "eth",
						TokenId:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
						CoinType: "token",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, prepareCache(tt.coins))
		})
	}
}

func TestClient_GetCoinsBySymbol(t *testing.T) {
	coins := GeckoCoins{
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
		GeckoCoin{
			Id:     "0x",
			Symbol: "btc",
			Name:   "0x",
			Platforms: Platforms{
				"ethtereum": "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
			},
		},
		GeckoCoin{
			Id:     "usdt",
			Symbol: "usdt",
			Name:   "usdt",
			Platforms: Platforms{
				"ethtereum": "0xdac17f958d2ee523a2206206994597c13d831ec7",
			},
		},
	}
	tests := []struct {
		name     string
		id       string
		expected []CoinResult
	}{
		{
			name: "test fetching 0x",
			id:   "0x",
			expected: []CoinResult{
				{
					Symbol:   "eth",
					TokenId:  "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
					CoinType: "token",
				},
			},
		},
		{
			name: "test fetching usdt",
			id:   "usdt",
			expected: []CoinResult{
				{
					Symbol:   "eth",
					TokenId:  "0xdac17f958d2ee523a2206206994597c13d831ec7",
					CoinType: "token",
				},
			},
		},
	}
	client := NewClient("")
	client.m = prepareCache(coins)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := client.GetCoinsBySymbol(tt.id)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}
