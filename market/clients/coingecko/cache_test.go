package coingecko

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewCache(t *testing.T) {
	tests := []struct {
		name     string
		coins    GeckoCoins
		expected *Cache
	}{
		{
			name: "test prepare cache map",
			coins: GeckoCoins{
				GeckoCoin{
					Id:        "ethereum",
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
					Symbol: "0x",
					Name:   "0x",
					Platforms: Platforms{
						"ethereum": "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
					},
				},
				GeckoCoin{
					Id:     "usdt",
					Symbol: "usdt",
					Name:   "usdt",
					Platforms: Platforms{
						"ethereum": "0xdac17f958d2ee523a2206206994597c13d831ec7",
					},
				},
			},
			expected: &Cache{
				"0x": {
					CoinResult{
						Symbol:   "eth",
						TokenId:  "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
						CoinType: "token",
					},
				},
				"usdt": {
					CoinResult{
						Symbol:   "eth",
						TokenId:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
						CoinType: "token",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, NewCache(tt.coins))
		})
	}
}

func TestClient_GetCoinsById(t *testing.T) {
	coins := GeckoCoins{
		GeckoCoin{
			Id:        "ethereum",
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
			Symbol: "0x",
			Name:   "0x",
			Platforms: Platforms{
				"ethereum": "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
			},
		},
		GeckoCoin{
			Id:     "usdt",
			Symbol: "usdt",
			Name:   "usdt",
			Platforms: Platforms{
				"ethereum": "0xdac17f958d2ee523a2206206994597c13d831ec7",
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
					TokenId:  "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
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
					TokenId:  "0xdAC17F958D2ee523a2206206994597C13D831ec7",
					CoinType: "token",
				},
			},
		},
	}
	cache := NewCache(coins)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := cache.GetCoinsById(tt.id)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func Test_NewSymbolsCache(t *testing.T) {
	tests := []struct {
		name     string
		coins    GeckoCoins
		expected *SymbolsCache
	}{
		{
			name: "test prepare cache map",
			coins: GeckoCoins{
				GeckoCoin{
					Id:        "ethereum",
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
					Symbol: "0x",
					Name:   "0x",
					Platforms: Platforms{
						"ethereum": "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
					},
				},
				GeckoCoin{
					Id:     "usdt",
					Symbol: "usdt",
					Name:   "usdt",
					Platforms: Platforms{
						"ethereum": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
					},
				},
			},
			expected: &SymbolsCache{
				"ETH:0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6": GeckoCoin{
					Id:     "0x",
					Symbol: "0x",
					Name:   "0x",
					Platforms: Platforms{
						"ethereum": "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
					},
				},
				"ETH:0xdAC17F958D2ee523a2206206994597C13D831ec7": GeckoCoin{
					Id:     "usdt",
					Symbol: "usdt",
					Name:   "usdt",
					Platforms: Platforms{
						"ethereum": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
					},
				},
				"BTC": GeckoCoin{Id: "bitcoin", Symbol: "btc", Name: "btc", Platforms: nil},
				"ETH": GeckoCoin{Id: "ethereum", Symbol: "eth", Name: "eth", Platforms: nil},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, NewSymbolsCache(tt.coins))
		})
	}
}

func TestClient_GetCoinsBySymbol(t *testing.T) {
	coins := GeckoCoins{
		GeckoCoin{
			Id:        "ethereum",
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
			Symbol: "0x",
			Name:   "0x",
			Platforms: Platforms{
				"ethereum": "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
			},
		},
		GeckoCoin{
			Id:     "usdt",
			Symbol: "usdt",
			Name:   "usdt",
			Platforms: Platforms{
				"ethereum": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			},
		},
	}
	tests := []struct {
		name     string
		symbol   string
		address  string
		expected GeckoCoin
	}{
		{
			name:    "test fetching 0x",
			symbol:  "eth",
			address: "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
			expected: GeckoCoin{
				Id:     "0x",
				Symbol: "0x",
				Name:   "0x",
				Platforms: Platforms{
					"ethereum": "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
				},
			},
		}, {
			name:    "test fetching usdt",
			symbol:  "eth",
			address: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			expected: GeckoCoin{
				Id:     "usdt",
				Symbol: "usdt",
				Name:   "usdt",
				Platforms: Platforms{
					"ethereum": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
				},
			},
		},
		{
			name:   "test fetching btc",
			symbol: "btc",
			expected: GeckoCoin{
				Id:        "bitcoin",
				Symbol:    "btc",
				Name:      "btc",
				Platforms: nil,
			},
		},
		{
			name:   "test fetching eth",
			symbol: "eth",
			expected: GeckoCoin{
				Id:        "ethereum",
				Symbol:    "eth",
				Name:      "eth",
				Platforms: nil,
			},
		},
	}
	cache := NewSymbolsCache(coins)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := cache.GetCoinsBySymbol(tt.symbol, tt.address)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res)
		})
	}
}

func Test_NormalizeTokenIDs(t *testing.T) {
	tests := []struct {
		name     string
		platform string
		address  string
		expected string
	}{
		{
			name:     "Should checksum Ethereum lowercase address",
			platform: PlatformEthereum,
			address:  "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
			expected: "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
		},
		{
			name:     "Should checksum Classic lowercase address",
			platform: PlatformClassic,
			address:  "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
			expected: "0x812f35b66Ec9EEe26CD7Fdf07Fbc1c9c0ac3C4D6",
		},
		{
			name:     "Check if one of the input empty - 1",
			platform: PlatformEthereum,
			address:  "",
			expected: "",
		},
		{
			name:     "Check if one of the input empty - 2",
			platform: PlatformClassic,
			address:  "",
			expected: "",
		},
		{
			name:     "Check if one of the input empty - 3",
			platform: "",
			address:  "0x812f35b66ec9eee26cd7fdf07fbc1c9c0ac3c4d6",
			expected: "",
		},
		{
			name:     "Should not process if address malformed",
			platform: PlatformEthereum,
			address:  "https://etherscan.io/address/0x8Ddc86DbA7ad728012eFc460b8A168Aba60B403B",
			expected: "",
		},
		{
			name:     "Return empty string if input empty",
			platform: "",
			address:  "",
			expected: "",
		},
		{
			name:     "Should return same Stellar address",
			platform: "stellar",
			address:  "SIX-GDMS6EECOH6MBMCP3FYRYEVRBIV3TQGLOFQIPVAITBRJUMTI6V7A2X6Z",
			expected: "SIX-GDMS6EECOH6MBMCP3FYRYEVRBIV3TQGLOFQIPVAITBRJUMTI6V7A2X6Z",
		},
		{
			name:     "Should return same Neo address",
			platform: "neo",
			address:  "ab38352559b8b203bde5fddfa0b07d8b2525e132",
			expected: "ab38352559b8b203bde5fddfa0b07d8b2525e132",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, normalizeTokenId(tt.platform, tt.address))
		})
	}
}
