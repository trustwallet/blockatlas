package tezos

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

func Init(api, rpc string) *Platform {
	p := &Platform{
		client:    Client{blockatlas.InitClient(api)},
		rpcClient: RpcClient{blockatlas.InitClient(rpc)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TEZOS]
}
