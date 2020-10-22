package ripple

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	return &Platform{
		client: Client{blockatlas.InitClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XRP]
}
