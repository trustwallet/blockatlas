package algorand

import (
	"github.com/trustwallet/blockatlas/services/assets"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetActiveValidators() (blockatlas.StakeValidators, error) {
	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	result := make(blockatlas.StakeValidators, 0, len(validators))
	for _, v := range validators {
		result = append(result, v)
	}
	return result, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 6.1},
		MinimumAmount: "0",
		LockTime:      0,
		Type:          blockatlas.DelegationTypeAuto,
	}
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	acc, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}
	return strconv.FormatUint(acc.Amount, 10), nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	return blockatlas.ValidatorPage{}, nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return blockatlas.DelegationsPage{}, nil
}
