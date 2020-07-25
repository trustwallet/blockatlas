package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/platform/tezos/bakingbad"
	"github.com/trustwallet/blockatlas/services/assets"
	"strconv"
)

const (
	lockTime           = 0
	minimumStakeAmount = "0"
)

func (p *Platform) GetActiveValidators() (blockatlas.StakeValidators, error) {
	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	result := make(blockatlas.StakeValidators, 0, len(validators))
	for _, v := range validators {
		if v.Status {
			result = append(result, v)
		}
	}
	return result, nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	account, err := p.rpcClient.GetAccount(address)
	if err != nil {
		return nil, err
	}
	if len(account.Delegate) == 0 {
		return make(blockatlas.DelegationsPage, 0), nil
	}

	validators, err := assets.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	return NormalizeDelegation(account, validators)
}

func NormalizeDelegation(account Account, validators blockatlas.ValidatorMap) (blockatlas.DelegationsPage, error) {
	validator, ok := validators[account.Delegate]
	if !ok {
		logger.Warn("Validator not found", logger.Params{"platform": "tezos", "delegation": account.Delegate})
		validator = getUnknownValidator(account.Delegate)
	}
	return blockatlas.DelegationsPage{
		{
			Delegator: validator,
			Value:     account.Balance,
			Status:    blockatlas.DelegationStatusActive,
		},
	}, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)

	validators, err := p.bakingbadClient.GetBakers()
	if err != nil {
		return results, err
	}
	for _, v := range *validators {
		results = append(results, normalizeValidator(v))
	}
	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: p.GetMaxAPR(),
		},
		MinimumAmount: minimumStakeAmount,
		LockTime:      lockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) GetMaxAPR() float64 {
	validators, err := p.GetActiveValidators()
	if err != nil {
		logger.Error("GetMaxAPR", logger.Params{"details": err, "platform": p.Coin().Symbol})
		return blockatlas.DefaultAnnualReward
	}

	var max = 0.0
	for _, e := range validators {
		v := e.Details.Reward.Annual
		if v > max {
			max = v
		}
	}

	return max
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.rpcClient.GetAccount(address)
	if err != nil {
		return "0", err
	}
	return account.Balance, nil
}

func getDetails(baker bakingbad.Baker) blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: baker.EstimatedRoi * 100},
		MinimumAmount: blockatlas.Amount(strconv.Itoa(baker.MinDelegation)),
		LockTime:      lockTime,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func normalizeValidator(v bakingbad.Baker) (validator blockatlas.Validator) {
	return blockatlas.Validator{
		Status:  true,
		ID:      v.Address,
		Details: getDetails(v),
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
			LockTime:      lockTime,
			MinimumAmount: minimumStakeAmount,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
