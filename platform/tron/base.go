package tron

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client     Client
	gridClient GridClient
}

func Init(api, gridApi string) *Platform {
	return &Platform{
		client:     Client{client.InitClient(api, middleware.SentryErrorHandler)},
		gridClient: GridClient{client.InitClient(api, middleware.SentryErrorHandler)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Tron()
}
