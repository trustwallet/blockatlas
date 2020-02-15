package assets

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"time"
)

func GetCoinInfo(coinId int, token string) (info *blockatlas.CoinInfo, err error) {
	c, ok := coin.Coins[uint(coinId)]
	if !ok {
		return info, errors.E("coin not found")
	}
	url := getCoinInfoUrl(c, token)
	request := blockatlas.InitClient(url)
	err = request.GetWithCache(&info, "/info.json", nil, time.Hour*1)
	return
}

func getCoinInfoUrl(c coin.Coin, token string) string {
	if len(token) == 0 {
		return AssetsURL + c.Handle + "/info"
	}
	return AssetsURL + c.Handle + "/assets/" + token
}
