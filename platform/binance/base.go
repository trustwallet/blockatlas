package binance

import (
	"github.com/trustwallet/blockatlas/platform/binance/staking"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client        Client
	stakingClient staking.Client
}

func Init(api, apiKey, stakingApi string) *Platform {
	p := Platform{
		client:        InitClient(api, apiKey),
		stakingClient: staking.InitClient(stakingApi),
	}
	return &p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Binance()
}
