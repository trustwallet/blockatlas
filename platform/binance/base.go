package binance

import (
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	p := Platform{
		client: InitClient(api),
	}
	return &p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Binance()
}
