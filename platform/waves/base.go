package waves

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client Client
}

func Init(api string) *Platform {
	return &Platform{
		client: Client{client.InitClient(api)},
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.WAVES]
}
