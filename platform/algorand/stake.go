package algorand

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	//TODO: Find a way to have a dynamic
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
	return make(blockatlas.ValidatorPage, 0), nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return make(blockatlas.DelegationsPage, 0), nil
}
