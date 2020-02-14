package nimiq

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("nimiq.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.NIM]
}
