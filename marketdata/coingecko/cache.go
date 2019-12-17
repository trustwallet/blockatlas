package coingecko

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type Cache map[string][]CoinResult

func (c Cache) GetCoinsBySymbol(id string) (coins []CoinResult, err error) {
	coins, ok := c[id]
	if !ok {
		err = errors.E("No coin found by id", errors.Params{"id": id})
	}
	return
}

func NewCache(coins GeckoCoins) *Cache {
	m := Cache{}
	coinsMap := make(map[string]GeckoCoin)
	for _, coin := range coins {
		coinsMap[coin.Id] = coin
	}

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
			m[coin.Id] = append(m[coin.Id], CoinResult{
				Symbol:   platformCoin.Symbol,
				TokenId:  address,
				CoinType: blockatlas.TypeToken,
			})
		}
	}
	return &m
}
