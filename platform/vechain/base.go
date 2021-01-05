package vechain

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	return &Platform{
		client: Client{internal.InitJSONClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.VET]
}
