package stellar

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func Init(coin uint, api string) *Platform {
	return &Platform{
		CoinIndex: coin,
		client:    Client{blockatlas.InitClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
