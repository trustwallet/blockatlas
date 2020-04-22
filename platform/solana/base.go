package solana

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
)

type Platform struct {
	client Client
	assets assets.AssetsServiceIface
}

// Requires assets
func Init(api string) *Platform {
	p := &Platform{
		client: Client{blockatlas.InitJSONClient(api)},
		assets: assets.GetService(),
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.SOL]
}
