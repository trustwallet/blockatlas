package tezos

import (
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/client"
)

type Platform struct {
	client    Client
	rpcClient RpcClient
}

func (p *Platform) Init() error {
	p.client = Client{client.InitClient(viper.GetString("tezos.api"))}
	p.client.SetTimeout(35)
	p.rpcClient = RpcClient{client.InitClient(viper.GetString("tezos.rpc"))}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}
