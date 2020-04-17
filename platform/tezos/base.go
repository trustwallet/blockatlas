package tezos

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"github.com/trustwallet/blockatlas/services/assets"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
	assets    assets.AssetsServiceI
}

func Init(serviceRepo *servicerepo.ServiceRepo, api, rpc string) *Platform {
	p := &Platform{
		client:    Client{blockatlas.InitClient(api)},
		rpcClient: RpcClient{blockatlas.InitClient(rpc)},
		assets:    assets.GetService(serviceRepo),
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}
