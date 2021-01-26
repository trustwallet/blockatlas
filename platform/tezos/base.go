package tezos

import (
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/network/middleware"
)

type Platform struct {
	client      Client
	rpcClient   RpcClient
	bakerClient BakerClient
}

func Init(api, rpc, baker string) *Platform {
	p := &Platform{
		client:      Client{client.InitClient(api, middleware.SentryErrorHandler)},
		rpcClient:   RpcClient{client.InitClient(rpc, middleware.SentryErrorHandler)},
		bakerClient: BakerClient{client.InitClient(baker, middleware.SentryErrorHandler)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Tezos()
}
