package cosmos

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
)

type Platform struct {
	client    Client
	CoinIndex uint
	assets    assets.AssetsServiceIface
}

// Requires assets
func Init(coin uint, api string) *Platform {
	return &Platform{
		CoinIndex: coin,
		client:    Client{blockatlas.InitClient(api)},
		assets:    assets.GetService(),
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
