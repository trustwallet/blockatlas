package theta

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api, key string) *Platform {
	request := client.InitClient(api, middleware.SentryErrorHandler)
	request.Headers = map[string]string{"x-api-token": key}
	return &Platform{
		client: Client{request},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.THETA]
}
