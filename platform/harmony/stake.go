package harmony

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/services/assets"
	"math/big"
	"strconv"
)

const (
	lockTime = 604800 // in seconds (7 epochs or 7 days)
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
		return results, err
	}

	for _, v := range validators.Validators {
		var apr float64
		if apr, err = strconv.ParseFloat(v.Lifetime.Apr, 64); err != nil {
			apr = 0
		}
		results = append(results, normalizeValidator(v, apr))
	}

	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	apr := p.GetMaxAPR()
	return getDetails(apr)
}

func (p *Platform) GetMaxAPR() float64 {
	validators, err := p.client.GetValidators()
	if err != nil {
		log.WithFields(log.Fields{"details": err, "platform": p.Coin().Symbol}).Error("GetMaxAPR")
		return Annual
	}

	var max = 0.0
	for _, e := range validators.Validators {
		var apr float64
		if apr, err = strconv.ParseFloat(e.Lifetime.Apr, 64); err != nil {
			apr = 0.0
		}

		if apr > max {
			max = apr
		}
	}

	return max
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	delegations, err := p.client.GetDelegations(address)
	if err != nil {
		return nil, err
	}

	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}

	return NormalizeDelegations(delegations.List, validators), nil
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
			log.WithFields(
				log.Fields{"address": v.ValidatorAddress, "platform": "harmony", "delegation": v.DelegatorAddress},
			).Error("Validator not found")
			continue
		}

		bigval := new(big.Float)
		bigval.SetFloat64(v.Amount)

		result := new(big.Int)
		bigval.Int(result) // store converted number in result

		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     result.String(), // v.Amount.String(),
			Status:    blockatlas.DelegationStatusActive,
		}
		results = append(results, delegation)
	}
	return results
}

func getDetails(apr float64) blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: apr},
		MinimumAmount: blockatlas.Amount("1000"),
		LockTime:      lockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v Validator, apr float64) (validator blockatlas.Validator) {
	return blockatlas.Validator{
		Status:  v.Active,
		ID:      v.Info.Address,
		Details: getDetails(apr),
	}
}
