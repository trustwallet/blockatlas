package tron

import (
	"github.com/trustwallet/blockatlas/pkg/address"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	services "github.com/trustwallet/blockatlas/services/assets"
	"strconv"
	"time"
)

const Annual = 0.74

func (p *Platform) GetValidators() (blockatlas.ValidatorPage, error) {
	results := make(blockatlas.ValidatorPage, 0)
	validators, err := p.client.GetValidators()
	if err != nil {
		return results, err
	}

	for _, v := range validators.Witnesses {
		if val, ok := normalizeValidator(v); ok {
			results = append(results, val)
		}
	}
	return results, nil
}

func (p *Platform) GetDetails() blockatlas.StakingDetails {
	return getDetails()
}

func getDetails() blockatlas.StakingDetails {
	return blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: blockatlas.Amount("1000000"),
		LockTime:      259200,
		Type:          blockatlas.DelegationTypeDelegate,
	}
}

func (p *Platform) GetDelegations(address string) (blockatlas.DelegationsPage, error) {
	results := make(blockatlas.DelegationsPage, 0)
	votes, err := p.client.GetAccountVotes(address)
	if err != nil {
		return nil, err
	}
	if len(votes.Votes) == 0 {
		return results, nil
	}
	validators, err := services.GetValidatorsMap(p)
	if err != nil {
		return nil, err
	}
	results = append(results, NormalizeDelegations(votes, validators)...)
	return results, nil
}

func (p *Platform) UndelegatedBalance(address string) (string, error) {
	account, err := p.client.GetAccount(address)
	if err != nil {
		return "0", err
	}

	for _, data := range account.Data {
		return strconv.FormatUint(uint64(data.Balance), 10), nil
	}
	return "0", nil
}

func normalizeValidator(v Validator) (validator blockatlas.Validator, ok bool) {
	address, err := address.HexToAddress(v.Address)
	if err != nil {
		return validator, false
	}

	return blockatlas.Validator{
		Status:  true,
		ID:      address,
		Details: getDetails(),
	}, true
}

func NormalizeDelegations(data *AccountData, validators blockatlas.ValidatorMap) []blockatlas.Delegation {
	results := make([]blockatlas.Delegation, 0)
	for _, v := range data.Votes {
		validator, ok := validators[v.VoteAddress]
		if !ok {
			logger.Warn("Validator not found", logger.Params{"address": v.VoteAddress, "platform": "tron"})
			continue
		}
		delegation := blockatlas.Delegation{
			Delegator: validator,
			Value:     strconv.Itoa(v.VoteCount * 1000000),
			Status:    blockatlas.DelegationStatusActive,
		}
		for _, f := range data.Frozen {
			t2 := time.Now().UnixNano() / int64(time.Millisecond)
			if f.ExpireTime > t2 {
				delegation.Status = blockatlas.DelegationStatusPending
			}
		}
		results = append(results, delegation)
	}
	return results
}
