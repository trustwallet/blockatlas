package cmc

import (
	"fmt"
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
	Id       uint
	Coin     coin.Coin
	TokenId  string
	CoinType blockatlas.CoinType
}

type CmcSlice []CoinMap
type CoinMapping map[string]CoinMap
type CmcMapping map[uint][]CoinMap

func (c *CmcSlice) coinToCmcMap() (m CoinMapping) {
	m = make(map[string]CoinMap)
	for _, cm := range *c {
		m[generateId(cm.Coin, cm.TokenId)] = cm
	}
	return
}

func (c *CmcSlice) cmcToCoinMap() (m CmcMapping) {
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
		tokens = append(tokens, CoinResult{Coin: c, Id: cc.Id, TokenId: cc.TokenId, CoinType: blockatlas.CoinType(cc.Type)})
	}
	return tokens, nil
}

func (cm CoinMapping) GetCoinByContract(coinId uint, contract string) (c CoinMap, err error) {
	c, ok := cm[generateId(coinId, contract)]
	if !ok {
		err = errors.E("No coin found", errors.Params{"coin": coinId, "token": contract})
	}

	return
}

func generateId(id uint, token string) string {
	return fmt.Sprintf("%d:%s", id, token)
}
