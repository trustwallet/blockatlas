package tron

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client     Client
	gridClient GridClient
}

func Init(api, gridApi string) *Platform {
	return &Platform{
		client:     Client{internal.InitClient(api)},
		gridClient: GridClient{internal.InitClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}
