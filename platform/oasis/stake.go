package oasis

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
	"time"
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

	validators, err := p.client.GetValidators()
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

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	var amount int64 = 0;
	delegations, err := p.client.GetDelegationsFor(address)
	if err != nil {
		return "0", err
	}

	for _, v := range delegations.List {
		amount += v.Shares.Int64()
	}

	return string(amount), nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	delegations, err := p.client.GetDelegationsFor(address)
	if err != nil {
		return nil, err
	}
	unbondingDelegations, err := p.client.GetUnbondingDelegationsFor(address)
	if err != nil {
		return nil, err
	}

	if delegations.List == nil && unbondingDelegations.List == nil {
		return results, nil
	}
	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}

	results = append(results, NormalizeDelegations(delegations.List, validators)...)
	results = append(results, NormalizeUnbondingDelegations(unbondingDelegations.List, validators)...)

	return results, nil
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
			t, _ := time.Parse(time.RFC3339, string(entry.DebondEndTime)) // FIXME check if we convert the date from epoch to RFC3339 correctly
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
			LockTime:      lockTime,
			MinimumAmount: minimumAmount,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
