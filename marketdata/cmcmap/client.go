package cmcmap

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type CoinMap struct {
	Coin    int    `json:"coin"`
	Type    string `json:"type"`
	TokenId string `json:"token_id"`
	Id      int    `json:"id"`
}

type CmcSlice []CoinMap
type CmcMapping map[int]CoinMap

func (c *CmcSlice) getMap() (m CmcMapping) {
	m = make(map[int]CoinMap)
	for _, cm := range *c {
		m[cm.Coin] = cm
	}
	return
}

func GetCmcMap() (CmcMapping, error) {
	var results CmcSlice
	request := blockatlas.Request{
		BaseUrl:      viper.GetString("market.cmc_map_url"),
		HttpClient:   blockatlas.DefaultClient,
		ErrorHandler: blockatlas.DefaultErrorHandler,
	}
	err := request.Get(&results, "", nil)
	if err != nil {
		return nil, errors.E(err).PushToSentry()
	}
	return results.getMap(), nil
}
