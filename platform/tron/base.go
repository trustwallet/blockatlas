package tron

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/servicerepo"
	"github.com/trustwallet/blockatlas/services/assets"
)

type Platform struct {
	client Client
	assets assets.AssetsServiceIface
}

// Requires assetsService from serviceRepo.
func Init(serviceRepo *servicerepo.ServiceRepo, api string) *Platform {
	return &Platform{
		client: Client{blockatlas.InitClient(api)},
		assets: assets.GetService(serviceRepo),
	}
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.TRX]
}
