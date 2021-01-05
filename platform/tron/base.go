package tron

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client         Client
	explorerClient ExplorerClient
}

func Init(api, explorerApi string) *Platform {
	return &Platform{
		client:         Client{internal.InitClient(api)},
		explorerClient: ExplorerClient{internal.InitClient(explorerApi)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}
