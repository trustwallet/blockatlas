package harmony

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
)

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()
	if err != nil {
		return results, err
	}

	apr, err := p.client.GetAPR()
	if err != nil {
		apr = Annual
	}

	for _, v := range validators.Validators {
		if val, ok := normalizeValidator(v, apr); ok {
			results = append(results, val)
		}
	}
	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	apr, err := p.client.GetAPR()
	if err != nil {
		apr = Annual
	}
	return getDetails(apr)
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	delegations, err := p.client.GetDelegations(address)
	if err != nil {
		return nil, err
	}

	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(delegations.List, validators)...)
	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	balance, err := p.client.GetBalance(address)
	if err != nil {
		return "0", err
	}
	return balance, nil
}

func NormalizeDelegations(delegations []Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		validator, ok := validators[v.ValidatorAddress]
		if !ok {
			logger.Error(errors.E("Validator not found", errors.Params{"address": v.ValidatorAddress, "platform": "harmony", "delegation": v.DelegatorAddress}))
			continue
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     v.Amount.String(),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}

func getDetails(apr float64) blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: apr},
		MinimumAmount: blockatlas.Amount("0"),
		LockTime:      0,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v Validator, apr float64) (validator blockatlas.Validator, ok bool) {
	return blockatlas.Validator{
		Status:  true,
		ID:      v.Address,
		Details: getDetails(apr),
	}, true
}
