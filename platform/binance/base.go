package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	dexClient DexClient
}

func Init(api, dex string) *Platform {
	p := &Platform{
		client:    Client{blockatlas.InitClient(api)},
		dexClient: DexClient{blockatlas.InitClient(dex)},
	}
	p.client.ErrorHandler = getHTTPError
	p.dexClient.ErrorHandler = getHTTPError
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
