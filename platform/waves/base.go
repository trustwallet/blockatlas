package waves

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	return &Platform{
		client: Client{internal.InitClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.WAVES]
}
