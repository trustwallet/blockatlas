package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	rpcClient Client
	dexClient DexClient
}

func Init(api, dex string) *Platform {
	p := &Platform{
		rpcClient: Client{blockatlas.InitClient(api)},
		dexClient: DexClient{blockatlas.InitClient(dex)},
	}
	p.rpcClient.ErrorHandler = getHTTPError
	p.dexClient.ErrorHandler = getHTTPError
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
