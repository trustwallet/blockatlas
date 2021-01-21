package elrond

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client    Client
	CoinIndex uint
}

func Init(coin uint, api string) *Platform {
	return &Platform{
		CoinIndex: coin,
		client:    Client{client.InitJSONClient(api, middleware.SentryErrorHandler)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[p.CoinIndex]
}
