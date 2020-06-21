package tron

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client         Client
	explorerClient Client
}

func Init(api, explorerApi string) *Platform {
	return &Platform{
		client:         Client{blockatlas.InitClient(api)},
		explorerClient: Client{blockatlas.InitClient(explorerApi)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}
