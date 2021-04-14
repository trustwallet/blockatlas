package aeternity

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	return &Platform{
		client: Client{client.InitClient(api, middleware.SentryErrorHandler)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Aeternity()
}
