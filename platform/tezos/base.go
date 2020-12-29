package tezos

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client      Client
	rpcClient   RpcClient
	bakerClient BakerClient
}

func Init(api, rpc, baker string) *Platform {
	p := &Platform{
		client:      Client{client.InitClient(api)},
		rpcClient:   RpcClient{client.InitClient(rpc)},
		bakerClient: BakerClient{client.InitClient(baker)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}
