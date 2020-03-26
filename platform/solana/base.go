package solana

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	p := &Platform{
		client: Client{blockatlas.InitJSONClient(api)},
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.SOL]
}
