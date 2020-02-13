package coingecko

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strings"
)

type Cache map[string][]CoinResult
type SymbolsCache map[string]GeckoCoin

const (
	PlatformEthereum = "ethereum"
	PlatformClassic  = "ethereum-classic"
)

func (c Cache) GetCoinsById(id string) (coins []CoinResult, err error) {
	coins, ok := c[id]
	if !ok {
		err = errors.E("No coin found by id", errors.Params{"id": id})
	}
	return
}

func (c SymbolsCache) generateId(symbol, token string) string {
	if len(token) > 0 {
		return fmt.Sprintf("%s:%s", strings.ToUpper(symbol), address.EIP55Checksum(token))
	}
	return strings.ToUpper(symbol)
}

func (c SymbolsCache) GetCoinsBySymbol(symbol, token string) (coin GeckoCoin, err error) {
	coin, ok := c[c.generateId(symbol, token)]
	if !ok {
		err = errors.E("No coin found by symbol", errors.Params{"symbol": symbol, "token": token})
	}
	return
}

func NewSymbolsCache(coins GeckoCoins) *SymbolsCache {
	m := SymbolsCache{}
	coinsMap := getCoinsMap(coins)

	for _, coin := range coins {
		if len(coin.Platforms) == 0 {
			m[m.generateId(coin.Symbol, "")] = coin
		}
		for platform, address := range coin.Platforms {
			if len(platform) == 0 || len(address) == 0 {
				continue
			}
			platformCoin, ok := coinsMap[platform]
			if !ok {
				continue
			}
			m[m.generateId(platformCoin.Symbol, address)] = coin
		}
	}

	return &m
}

func NewCache(coins GeckoCoins) *Cache {
	m := Cache{}
	coinsMap := getCoinsMap(coins)

	for _, coin := range coins {
		for platform, address := range coin.Platforms {
			if len(platform) == 0 || len(address) == 0 {
				continue
			}
			platformCoin, ok := coinsMap[platform]
			if !ok {
				continue
			}

			_, ok = m[coin.Id]
			if !ok {
				m[coin.Id] = make([]CoinResult, 0)
			}

			tokenId := normalizeTokenId(platform, address)
			if tokenId == "" {
				continue
			}

			m[coin.Id] = append(m[coin.Id], CoinResult{
				Symbol:   platformCoin.Symbol,
				TokenId:  tokenId,
				CoinType: blockatlas.TypeToken,
			})
		}
	}
	return &m
}

func getCoinsMap(coins GeckoCoins) map[string]GeckoCoin {
	coinsMap := make(map[string]GeckoCoin)
	for _, coin := range coins {
		coinsMap[coin.Id] = coin
	}
	return coinsMap
}

func normalizeTokenId(platform, addr string) string {
	if platform == "" || addr == "" {
		return ""
	}
	switch platform {
	case PlatformEthereum, PlatformClassic:
		if len(addr) == 42 && strings.HasPrefix(addr, "0x") {
			return address.EIP55Checksum(addr)
		}
		return ""
	default:
		return addr
	}
}
