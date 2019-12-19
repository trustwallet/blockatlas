package cmc

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type CoinMap struct {
	Coin    uint   `json:"coin"`
	Id      uint   `json:"id"`
	Type    string `json:"type"`
	TokenId string `json:"token_id"`
}

type CoinResult struct {
	Coin     coin.Coin
	TokenId  string
	CoinType blockatlas.CoinType
}

type CmcSlice []CoinMap
type CmcMapping map[uint][]CoinMap

func (c *CmcSlice) getMap() (m CmcMapping) {
	m = make(map[uint][]CoinMap)
	for _, cm := range *c {
		_, ok := m[cm.Id]
		if !ok {
			m[cm.Id] = make([]CoinMap, 0)
		}
		m[cm.Id] = append(m[cm.Id], cm)
	}
	return
}

func (cm CmcMapping) GetCoins(coinId uint) ([]CoinResult, error) {
	cmcCoin, ok := cm[coinId]
	if !ok {
		return nil, errors.E("CmcMapping.getCoin: coinId notFound")
	}
	tokens := make([]CoinResult, 0)
	for _, cc := range cmcCoin {
		c, ok := coin.Coins[cc.Coin]
		if !ok {
			continue
		}
		tokens = append(tokens, CoinResult{Coin: c, TokenId: cc.TokenId, CoinType: blockatlas.CoinType(cc.Type)})
	}
	return tokens, nil
}
