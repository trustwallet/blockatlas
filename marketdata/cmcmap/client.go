package cmcmap

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"time"
)

type CoinMap struct {
	Coin    uint   `json:"coin"`
	Id      uint   `json:"id"`
	Type    string `json:"type"`
	TokenId string `json:"token_id"`
}

type CmcSlice []CoinMap
type CmcMapping map[uint]CoinMap

func (c *CmcSlice) getMap() (m CmcMapping) {
	m = make(map[uint]CoinMap)
	for _, cm := range *c {
		m[cm.Id] = cm
	}
	return
}

func (cm CmcMapping) GetCoin(coinId uint) (coin.Coin, string, error) {
	cmcCoin, ok := cm[coinId]
	if !ok {
		return coin.Coin{}, "", errors.E("CmcMapping.getCoin: coinId notFound")
	}
	c, ok := coin.Coins[cmcCoin.Coin]
	if !ok {
		return coin.Coin{}, "", errors.E("CmcMapping.getCoin: Invalid cmcCoin.CoinId")
	}
	return c, cmcCoin.TokenId, nil
}

func GetCmcMap() (CmcMapping, error) {
	var results CmcSlice
	request := blockatlas.Request{
		BaseUrl:      viper.GetString("market.cmc.map_url"),
		HttpClient:   blockatlas.DefaultClient,
		ErrorHandler: blockatlas.DefaultErrorHandler,
	}
	err := request.GetWithCache(&results, "mapping.json", nil, time.Hour*1)
	if err != nil {
		return nil, errors.E(err).PushToSentry()
	}
	return results.getMap(), nil
}
