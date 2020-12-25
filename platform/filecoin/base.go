package filecoin

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
	"github.com/trustwallet/blockatlas/platform/filecoin/rpc"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client   rpc.Client
	explorer explorer.Client
}

func Init(api, explorerApi string) *Platform {
	p := &Platform{
		client:   rpc.Client{Request: blockatlas.InitClient(api)},
		explorer: explorer.Client{Request: blockatlas.InitClient(explorerApi)},
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.FIL]
}
