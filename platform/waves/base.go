package waves

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/client"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = Client{client.InitClient(viper.GetString("waves.api"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.WAVES]
}
