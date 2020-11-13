package assets

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"time"
)

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func fetchValidatorsInfo(coin coin.Coin) (AssetValidators, error) {
	var results AssetValidators
	request := blockatlas.InitClient(AssetsURL + coin.Handle)
	err := request.GetWithCache(&results, "validators/list.json", nil, time.Hour*1)
	if err != nil {
		return nil, err
	}
	return results, nil
}
