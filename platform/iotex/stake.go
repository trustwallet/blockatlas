package iotex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
	"github.com/trustwallet/golibs/types"
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

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	return p.client.GetValidators()
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return p.client.GetDelegations(address)
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 0},
		MinimumAmount: types.Amount("100000000000000000000"),
		LockTime:      259200,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}

	return account.AccountMeta.Balance, nil
}
