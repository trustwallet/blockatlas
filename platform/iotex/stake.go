package iotex

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	return p.client.GetValidators()
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return p.client.GetDelegations(address)
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 0},
		MinimumAmount: blockatlas.Amount("100000000000000000000"),
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
