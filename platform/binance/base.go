package binance

import (
	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client      Client
}

func Init(rpcApi string) *Platform {
	p := Platform{
		client:      InitClient(rpcApi),
	}
	return &p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
