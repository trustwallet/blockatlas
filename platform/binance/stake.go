package binance

import (
	"strconv"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/services/assets"
)

const (
	chainID            = "bsc"
	dummyLockTime      = 1814400 // in seconds (21 days)
	dummyMinimumAmount = "1"
	dummyMaxAPR        = 0.2
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
	validators, err := p.client.GetValidators()
	if err != nil {
		return nil, err
	}
	for _, validator := range validators.Validators {
		results = append(results, normalizeValidator(validator))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: dummyMaxAPR,
		},
		MinimumAmount: dummyMinimumAmount,
		LockTime:      dummyLockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	delegations, err := p.client.GetDelegations(chainID, address)
	if err != nil {
		return nil, err
	}
	unbondingDelegations, err := p.client.GetUnbondingDelegations(chainID, address)
	if err != nil {
		return nil, err
	}
	if len(delegations) == 0 && len(unbondingDelegations) == 0 {
		return results, nil
	}
	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(delegations, validators)...)
	results = append(results, NormalizeUnbondingDelegations(unbondingDelegations, validators)...)

	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	accountMeta, err := p.client.FetchAccountMeta(address)
	if err != nil {
		return "0", err
	}
	for _, coin := range accountMeta.Balances {
		if coin.Symbol == "BNB" {
			return coin.Free, nil
		}
	}
	return "0", nil
}

func NormalizeDelegations(delegations []Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		validator, ok := validators[v.ValidatorAddress]
		if !ok {
			logger.Warn("Validator not found", logger.Params{"address": v.ValidatorAddress, "platform": "binance", "delegation": v.DelegatorAddress})
			validator = getUnknownValidator(v.ValidatorAddress)

		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     strconv.FormatFloat(v.Value, 'f', 8, 64),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}

func NormalizeUnbondingDelegations(delegations []UnbondingDelegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range delegations {
		validator, ok := validators[v.Validator]
		if !ok {
			logger.Warn("Validator not found", logger.Params{"address": v.Validator, "platform": "binance", "delegation": v.Delegator})
			validator = getUnknownValidator(v.Validator)
		}
		t, _ := time.Parse(time.RFC3339, v.CompleteTime)
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     strconv.FormatFloat(v.Balance, 'f', 8, 64),
			Status:    blockatlas.DelegationStatusPending,
			Metadata: blockatlas.DelegationMetaDataPending{
				AvailableDate: uint(t.Unix()),
			},
		}
		results = append(results, delegation)
	}
	return results
}

//func normalizeValidator(v Validator, p Pool, inflation float64) (validator blockatlas.Validator) {
func normalizeValidator(v Validator) (validator blockatlas.Validator) {
	reward := dummyMaxAPR // TODO
	return blockatlas.Validator{
		Status: v.Status == 2,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: reward},
			MinimumAmount: dummyMinimumAmount,
			LockTime:      dummyLockTime,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}

func getUnknownValidator(address string) blockatlas.StakeValidator {
	return blockatlas.StakeValidator{
		ID:     address,
		Status: false,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Decommissioned",
			Description: "Decommissioned",
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: 0,
			},
			LockTime:      dummyLockTime,
			MinimumAmount: dummyMinimumAmount,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
