package binance

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestPlatform_GetDetails(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)
	result := p.GetDetails()
	assert.Equal(t, 0.2, result.Reward.Annual)
	assert.Equal(t, blockatlas.Amount("0.0001"), result.MinimumAmount)
	assert.Equal(t, 86400, result.LockTime)
	assert.Equal(t, blockatlas.DelegationTypeDelegate, result.Type)
}

func TestPlatform_NormalizeDelegations(t *testing.T) {
	delegations := []Delegation{
		Delegation{
			Value:            99.9,
			DelegatorAddress: "delegatorAddress",
			ValidatorAddress: "validatorAddress",
		},
		Delegation{
			Value:            20,
			DelegatorAddress: "delegatorAddress",
			ValidatorAddress: "NoSuchValidator",
		},
	}
	validator := blockatlas.StakeValidator{
		ID:     "validatorAddress",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Unit Test Validator",
			Description: "this is a validator in this unit test",
			Image:       "path/logo.png",
			Website:     "https://site/validator",
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: 0.25,
			},
			LockTime:      123456000,
			MinimumAmount: blockatlas.Amount("0.01"),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
	validators := make(blockatlas.ValidatorMap, 0)
	validators[validator.ID] = validator
	result := NormalizeDelegations(delegations, validators)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
	assert.NotNil(t, result[0].Delegator)
	assert.Equal(t, validator, result[0].Delegator)
	assert.Equal(t, "99.90000000", result[0].Value)
	assert.Equal(t, blockatlas.DelegationStatusActive, result[0].Status)
	assert.Equal(t, "NoSuchValidator", result[1].Delegator.ID)
	assert.Equal(t, false, result[1].Delegator.Status)
	assert.Equal(t, "Decommissioned", result[1].Delegator.Info.Name)
	assert.Equal(t, "20.00000000", result[1].Value)
	assert.Equal(t, blockatlas.DelegationStatusActive, result[1].Status)
}
