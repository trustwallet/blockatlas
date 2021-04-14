package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/binance/staking"
	"github.com/trustwallet/blockatlas/services/assets"
)

const (
	lockTime      = 604800      // 7 days
	minimumAmount = "100000000" // 1 BNB
)

func (p *Platform) GetActiveValidators() (blockatlas.StakeValidators, error) {
	return blockatlas.StakeValidators{}, nil
}

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	validators, err := p.stakingClient.GetValidators()
	if err != nil {
		return nil, err
	}

	assetsValidators, err := assets.GetValidatorsInfo(p.Coin())
	if err != nil {
		return nil, err
	}

	assetsMap := assetsValidators.ToMap()
	result := make(blockatlas.ValidatorPage, 0)

	for _, validator := range validators.Validators {
		// filter trusted
		if _, ok := assetsMap[validator.Validator]; !ok {
			continue
		}
		// filter inactive
		if validator.Status != 0 {
			continue
		}
		result = append(result, normalizeValidator(validator))
	}
	return result, nil
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

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	return "0", nil
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	return blockatlas.DelegationsPage{}, nil
}

func normalizeValidator(v staking.Validator) blockatlas.Validator {
	return blockatlas.Validator{
		Status: v.Status == 0,
		ID:     v.Validator,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: v.APR},
			MinimumAmount: minimumAmount,
			LockTime:      lockTime,
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
}
