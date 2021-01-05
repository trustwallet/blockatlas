package assets

import (
	"time"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func fetchValidatorsInfo(coin coin.Coin) (AssetValidators, error) {
	var results AssetValidators
	request := client.InitClient(AssetsURL + coin.Handle)
	err := request.GetWithCache(&results, "validators/list.json", nil, time.Hour*1)
	if err != nil {
		return nil, err
	}
	return results, nil
}
