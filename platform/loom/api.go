package cosmos

import (
	"github.com/trustwallet/blockatlas"

	"github.com/loomnetwork/blockatlas/coin"
	"github.com/spf13/viper"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = InitClient(viper.GetString("loom.rpc"))
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.LOOM]
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()
	if err != nil {
		return results, nil
	}

	rate, err := p.client.GetRate()
	if err != nil {
		return results, nil
	}

	for _, validator := range validators {
		resp := blockatlas.Validator{
			Status: bool(validator.Status == 2),
			ID:     validator.Address,
			Reward: blockatlas.StakingReward{Annual: rate},
		}
		results = append(results, resp)
	}

	return results, nil
}
