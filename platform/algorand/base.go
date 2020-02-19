package algorand

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/pkg/client"

	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{client.InitClient(viper.GetString("algorand.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ALGO]
}
