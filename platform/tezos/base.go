package tezos

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/tezos/bakingbad"
)

type Platform struct {
	client          Client
	rpcClient       RpcClient
	bakingbadClient bakingbad.Client
}

func Init(api, rpc, bakingbadUrl string) *Platform {
	p := &Platform{
		client:          Client{blockatlas.InitClient(api)},
		rpcClient:       RpcClient{blockatlas.InitClient(rpc)},
		bakingbadClient: bakingbad.Client{blockatlas.InitClient(bakingbadUrl)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}
