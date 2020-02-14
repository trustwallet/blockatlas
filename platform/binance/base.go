package binance

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	dexClient DexClient
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("binance.api"))}
	p.client.ErrorHandler = getHTTPError

	p.dexClient = DexClient{blockatlas.InitClient(viper.GetString("binance.dex"))}
	p.dexClient.ErrorHandler = getHTTPError
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
