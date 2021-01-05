package tezos

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

func Init(api, rpc string) *Platform {
	p := &Platform{
		client:    Client{internal.InitClient(api)},
		rpcClient: RpcClient{internal.InitClient(rpc)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}
