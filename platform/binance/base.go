package binance

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/client"
)

type Platform struct {
	client    Client
	dexClient DexClient
}

func (p *Platform) Init() error {
	p.client = Client{client.InitClient(viper.GetString("binance.api"))}
	p.client.ErrorHandler = getHTTPError

	p.dexClient = DexClient{client.InitClient(viper.GetString("binance.dex"))}
	p.dexClient.ErrorHandler = getHTTPError
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}
