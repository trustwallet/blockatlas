package terra

import (
	"strconv"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
)

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)

	var validators []Validator
	if res, err := p.client.GetValidators(); err != nil {
		return nil, err
	} else {
		validators = res.Validators
	}

	for _, validator := range validators {
		results = append(results, normalizeValidator(validator))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {

	// Dynamically retreive annual return
	annualReturn := 11.0
	stakingReturns, err := p.client.GetStakingReturns()
	if err == nil && len(stakingReturns) > 0 {
		latest, err := strconv.ParseFloat(stakingReturns[0].AnnualizedReturn, 64)
		if err == nil {
			annualReturn = latest
		}
	}

	lockTime, err := p.client.GetLockTime()
	if err != nil {
		lockTime = 1814400
	}

	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: annualReturn},
		MinimumAmount: blockatlas.Amount("0"),
		LockTime:      int(lockTime),
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	delegations, err := p.client.GetDelegations(address)
	if err != nil {
		return nil, err
	}
	unbondingDelegations, err := p.client.GetUnbondingDelegations(address)
	if err != nil {
		return nil, err
	}
	if delegations.List == nil && unbondingDelegations.List == nil {
		return results, nil
	}
	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(delegations.List, validators)...)
	results = append(results, NormalizeUnbondingDelegations(unbondingDelegations.List, validators)...)

	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}
	for _, coin := range account.Account.Value.Coins {
		if coin.Denom == DenomLuna {
			return coin.Amount, nil
		}
	}
	return "0", nil
}

func NormalizeDelegations(delegations []Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		validator, ok := validators[v.ValidatorAddress]
		if !ok {
			logger.Error(errors.E("Validator not found", errors.Params{"address": v.ValidatorAddress, "platform": "cosmos", "delegation": v.DelegatorAddress}))
			continue
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     v.Value(),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}

func NormalizeUnbondingDelegations(delegations []UnbondingDelegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		for _, entry := range v.Entries {
			validator, ok := validators[v.ValidatorAddress]
			if !ok {
				logger.Error(errors.E("Validator not found", errors.Params{"address": v.ValidatorAddress, "platform": "cosmos", "delegation": v.DelegatorAddress}))
				continue
			}
			t, _ := time.Parse(time.RFC3339, entry.CompletionTime)
			delegation := blockatlas.Delegation{
				Delegator: validator,
				Value:     entry.Balance,
				Status:    blockatlas.DelegationStatusPending,
				Metadata: blockatlas.DelegationMetaDataPending{
					AvailableDate: uint(t.Unix()),
				},
			}
			results = append(results, delegation)
		}
	}
	return results
}

func normalizeValidator(v Validator) (validator blockatlas.Validator) {
	reward, _ := strconv.ParseFloat(v.StakingReturn, 64)
	reward = reward * 100
	return blockatlas.Validator{
		Status: v.Status == "active",
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: reward},
			MinimumAmount: "0",
			LockTime:      1814400,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
