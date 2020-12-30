package zilliqa

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

func Init(api, apiKey, rpc string) *Platform {
	p := &Platform{
		client:    Client{client.InitClient(api)},
		rpcClient: RpcClient{client.InitClient(rpc)},
	}
	p.client.Headers["X-APIKEY"] = apiKey
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ZIL]
}
