package oasis

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
	"github.com/trustwallet/golibs/types"
	"time"
)

func (p *Platform) GetActiveValidators() (blockatlas.StakeValidators, error) {
	validators, err := p.client.GetValidators()
	if err != nil {
		return nil, err
	}
	consensusParams, err := p.client.GetConsensusParams()
	if err != nil {
		return nil, err
	}

	result := make(blockatlas.StakeValidators, 0)
	for _, validator := range validators.Validators {
		result = append(result, normalizeActiveValidator(validator, *consensusParams))
	}

	return result, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)

	validators, err := p.client.GetValidators()
	if err != nil {
		return nil, err
	}
	consensusParams, err := p.client.GetConsensusParams()
	if err != nil {
		return nil, err
	}

	for _, validator := range validators.Validators {
		results = append(results, normalizeValidator(validator, *consensusParams))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	apr := blockatlas.DefaultAnnualReward
	minimumAmount := types.Amount("-")
	lockTime := 0

	validators, err := p.GetValidators()
	consensusParams, err := p.client.GetConsensusParams()
	if err == nil {
		apr = blockatlas.FindHightestAPR(validators)
		minimumAmount = types.Amount(consensusParams.MinDelegationAmount)
		lockTime = int(consensusParams.DebondingInterval)
	}

	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: apr},
		MinimumAmount: minimumAmount,
		LockTime:      lockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	var amount int64 = 0
	delegations, err := p.client.GetDelegationsFor(address)
	if err != nil {
		return "0", err
	}

	for _, v := range delegations.List {
		amount += v.Shares.Int64()
	}

	return fmt.Sprintf("%d", amount), nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	delegations, err := p.client.GetDelegationsFor(address)
	if err != nil {
		return nil, err
	}
	debondingDelegations, err := p.client.GetUnbondingDelegationsFor(address)
	if err != nil {
		return nil, err
	}

	if delegations.List == nil && debondingDelegations.List == nil {
		return results, nil
	}
	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}

	results = append(results, NormalizeDelegations(delegations.List, validators)...)
	results = append(results, NormalizeUnbondingDelegations(debondingDelegations.List, validators)...)

	return results, nil
}

func normalizeValidator(v Validator, consensusParams ConsensusParams) (validator blockatlas.Validator) {
	return blockatlas.Validator{
		Status: true,
		ID:     v.ID,

		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: v.EffectiveAnnualReward},
			MinimumAmount: types.Amount(consensusParams.MinDelegationAmount),
			LockTime:      int(consensusParams.DebondingInterval),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}

func normalizeActiveValidator(v Validator, consensusParams ConsensusParams) (validator blockatlas.StakeValidator) {
	return blockatlas.StakeValidator{
		Status: true,
		ID:     v.ID,
		Info: blockatlas.StakeValidatorInfo{
			Name:    v.Name,
			Website: v.URL,
		},
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: v.EffectiveAnnualReward},
			MinimumAmount: types.Amount(consensusParams.MinDelegationAmount),
			LockTime:      int(consensusParams.DebondingInterval),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}

func NormalizeDelegations(delegations map[string]Delegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for k, v := range delegations {
		validator, ok := validators[k]
		if !ok {
			log.WithFields(
				log.Fields{"address": k, "platform": "cosmos", "delegation": k}, // FIXME check where to find the required addresses (ValidatorAddress and DelegationAddress)
			).Warn("Validator not found")
			validator = getUnknownValidator(k)

		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     v.Shares.String(),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}

func NormalizeUnbondingDelegations(delegations map[string][]DebondingDelegation, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for k, v := range delegations {
		for _, entry := range v {
			validator, ok := validators[k]
			if !ok {
				log.WithFields(
					log.Fields{"address": k, "platform": "cosmos", "delegation": k}, // FIXME check where to find the required addresses (ValidatorAddress and DelegationAddress)
				).Warn("Validator not found")
				validator = getUnknownValidator(k)
			}
			t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%d", entry.DebondEndTime)) // FIXME check if we convert the date from epoch to RFC3339 correctly
			delegation := blockatlas.Delegation{
				Delegator: validator,
				Value:     entry.Shares.String(),
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
			LockTime:      0,                // FIXME
			MinimumAmount: types.Amount(""), // FIXME
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
