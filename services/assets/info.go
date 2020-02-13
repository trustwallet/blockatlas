package assets

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"time"
)

func GetCoinInfo(coinId int, token string) (info *blockatlas.CoinInfo, err error) {
	assetsURL := viper.GetString("assets.blockchain") + "/"

	c, ok := coin.Coins[uint(coinId)]
	if !ok {
		return info, errors.E("coin not found")
	}

	url := getCoinInfoUrl(c, token, assetsURL)
	request := blockatlas.InitClient(url)
	err = request.GetWithCache(&info, "info/info.json", nil, time.Hour*1)
	return
}

func getCoinInfoUrl(c coin.Coin, token, assetsURL string) string {
	if len(token) == 0 {
		return assetsURL + c.Handle
	}
	return assetsURL + c.Handle + "/assets/" + token
}
