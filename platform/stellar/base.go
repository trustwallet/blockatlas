package stellar

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/client"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func (p *Platform) Init() error {
	handle := coin.Coins[p.CoinIndex].Handle
	api := fmt.Sprintf("%s.api", handle)
	p.client = Client{client.InitClient(viper.GetString(api))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
