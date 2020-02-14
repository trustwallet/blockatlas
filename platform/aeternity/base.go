package aeternity

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{blockatlas.InitClient(viper.GetString("aeternity.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.AE]
}
