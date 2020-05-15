package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	rpcClient      Client
	explorerClient ExplorerClient
}

func Init(rpcApi, explorerApi string) *Platform {
	p := Platform{
		rpcClient:      Client{blockatlas.InitClient(rpcApi)},
		explorerClient: ExplorerClient{blockatlas.InitClient(explorerApi)},
	}
	p.rpcClient.ErrorHandler = handleHTTPError
	p.explorerClient.ErrorHandler = handleHTTPError
	return &p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
