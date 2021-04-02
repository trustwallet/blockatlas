package tron

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api, apiKey string) *Platform {
	request := client.InitClient(api, middleware.SentryErrorHandler)
	request.Headers = map[string]string{"TRON-PRO-API-KEY": apiKey}
	return &Platform{
		client: Client{request},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Tron()
}
