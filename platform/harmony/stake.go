package harmony

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
	"strconv"
)

const (
	lockTime      = 604800 // 7 days
	minimumAmount = "1000"
)

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	balance, err := p.client.GetBalance(address)
	if err != nil {
		return "0", err
	}
	return balance, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	apr := blockatlas.DefaultAnnualReward
	validators, err := p.GetValidators()
	if err == nil {
		apr = blockatlas.FindHightestAPR(validators)
	}
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: apr * 100},
		MinimumAmount: minimumAmount,
		LockTime:      lockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.fetchValidators()
	if err != nil {
		return results, err
	}
	for _, v := range validators.Validators {
		var apr float64
		if apr, err = strconv.ParseFloat(v.Lifetime.Apr, 64); err != nil {
			apr = blockatlas.DefaultAnnualReward
		}
		results = append(results, normalizeValidator(v, apr))
	}
	return results, nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	delegations, err := p.client.fetchDelegations(address)
	if err != nil {
		return nil, err
	}
	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	return normalizeDelegations(delegations.List, validators), nil
}

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

func normalizeValidator(v Validator, apr float64) (validator blockatlas.Validator) {
	return blockatlas.Validator{
		Status: v.Active,
		ID:     v.Info.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: apr * 100},
			MinimumAmount: minimumAmount,
			LockTime:      lockTime,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}

func normalizeDelegations(delegations []Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		if validator, ok := validators[v.ValidatorAddress]; ok {
			delegation := blockatlas.Delegation{
				Delegator: validator,
				Value:     v.Amount.String(),
				Status:    blockatlas.DelegationStatusActive,
			}
			results = append(results, delegation)
		}
	}
	return results
}
