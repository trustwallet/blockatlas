package tron

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client         Client
	explorerClient ExplorerClient
}

func Init(api, explorerApi string) *Platform {
	return &Platform{
		client:         Client{client.InitClient(api)},
		explorerClient: ExplorerClient{client.InitClient(explorerApi)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}
