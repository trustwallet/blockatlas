package oasis

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
)

// FIXME We have to check these values are correct
const (
	lockTime      = 1814400 // in seconds (21 days)
	minimumAmount = "1"
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
	results := make(blockatlas.ValidatorPage, 0)

	// FIXME  Get the correct height to make the GetValidators request
	validators, err := p.client.GetValidators(1000)
	if err != nil {
		return nil, err
	}

	for _, validator := range *validators {
		results = append(results, normalizeValidator(validator))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	apr := blockatlas.DefaultAnnualReward
	validators, err := p.GetValidators()
	if err == nil {
		apr = blockatlas.FindHightestAPR(validators)
	}
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: apr},
		MinimumAmount: minimumAmount,
		LockTime:      lockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v Validator) (validator blockatlas.Validator) {
	reward := 123545.2 // FIXME Get the correct reward value
	return blockatlas.Validator{
		Status: true, // FIXME Check where to find the status
		ID:     v.ID, // FIXME Check if the public address we rcv is the address we need to pass
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: reward},
			MinimumAmount: minimumAmount,
			LockTime:      lockTime,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
