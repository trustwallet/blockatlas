package kava

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func Init(coin uint, api string) *Platform {
	return &Platform{
		CoinIndex: coin,
		client:    Client{client.InitClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
