package zilliqa

import (
	"github.com/trustwallet/blockatlas/platform/zilliqa/rpc"
	"github.com/trustwallet/blockatlas/platform/zilliqa/viewblock"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client    viewblock.Client
	rpcClient rpc.Client
}

func Init(api, apiKey, rpcUrl string) *Platform {
	p := &Platform{
		client:    viewblock.InitClient(api, apiKey),
		rpcClient: rpc.InitClient(rpcUrl),
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Zilliqa()
}
