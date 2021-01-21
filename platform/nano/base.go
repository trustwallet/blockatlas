package nano

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	p := &Platform{
		client: Client{client.InitJSONClient(api, middleware.SentryErrorHandler)},
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.NANO]
}
