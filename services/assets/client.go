package assets

import (
	"time"

	"github.com/trustwallet/golibs/network/middleware"

	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

const (
	URL = "https://assets.trustwalletapp.com/blockchains/"
)

func GetValidatorsInfo(coin coin.Coin) (AssetValidators, error) {
	var results AssetValidators
	request := client.InitClient(URL+coin.Handle, middleware.SentryErrorHandler)
	err := request.GetWithCache(&results, "validators/list.json", nil, time.Hour*1)
	if err != nil {
		return nil, err
	}
	return results, nil
}
