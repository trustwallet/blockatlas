package assets

import (
	"time"

	"github.com/trustwallet/blockatlas/internal"

	"github.com/trustwallet/golibs/coin"
)

const (
	URL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func GetchValidatorsInfo(coin coin.Coin) (AssetValidators, error) {
	var results AssetValidators
	request := internal.InitClient(URL + coin.Handle)
	err := request.GetWithCache(&results, "validators/list.json", nil, time.Hour*1)
	if err != nil {
		return nil, err
	}
	return results, nil
}
