package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	rpcClient      Client
	explorerClient ExplorerClient
}

func Init(api, explorer string) *Platform {
	p := &Platform{
		rpcClient:      Client{blockatlas.InitClient(api)},
		explorerClient: ExplorerClient{blockatlas.InitClient(explorer)},
	}
	p.rpcClient.ErrorHandler = getHTTPError
	p.explorerClient.ErrorHandler = getHTTPError
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
