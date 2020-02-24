package vechain

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
)

const (
	// TODO: Find a way to have a dynamic APR
	// The current value comes from https://www.stakingrewards.com/asset/vechain
	Annual = 1.35
)

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
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
	balance, err := numbers.HexToDecimal(acc.Balance)
	if err != nil {
		return "0", errors.E("Invalid asset balance", errors.Params{"balance": acc.Balance})
	}
	return balance, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	return blockatlas.ValidatorPage{}, nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return blockatlas.DelegationsPage{}, nil
}
