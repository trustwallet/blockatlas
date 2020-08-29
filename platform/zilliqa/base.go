package zilliqa

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
	udClient  Client
}

func Init(api, apiKey, rpc, udClient string) *Platform {
	p := &Platform{
		client:    Client{blockatlas.InitClient(api)},
		rpcClient: RpcClient{blockatlas.InitClient(rpc)},
		udClient:  Client{blockatlas.InitClient(udClient)},
	}
	p.client.Headers["X-APIKEY"] = apiKey
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.ZILLIQA]
}
