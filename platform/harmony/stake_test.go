package harmony

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const validatorSrc = `
{
	"one-address": "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
	"bls-public-keys": [
		"65f55eb3052f9e9f632b2923be594ba77c55543f5c58ee1454b9cfd658d25e06373b0f7d42a19c84768139ea294f6204"
	],
	"min-self-delegation": 2000000000000000000,
	"max-total-delegation": 11553255926290448384,
	"active": true,
	"commission": {
		"rate": "0.100000000000000000",
		"max-rate": "0.900000000000000000",
		"max-change-rate": "0.050000000000000000"
	},
	"description": {
		"name": "John",
		"identity": "john",
		"website": "john@harmony.one",
		"security_contact": "Alex",
		"details": "John the validator"
	},
	"creation-height": 51
}`

const delegationsSrc = `
[
	{
		"validator_address": "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
		"delegator_address": "one1pf75h0t4am90z8uv3y0dgunfqp4lj8wr3t5rsp",
		"amount": 10000000000000000000,
		"reward": 15854399877248931866418,
		"Undelegations": []
	}
]`

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: v.Active,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 10},
			LockTime:      0,
			MinimumAmount: "0",
			Type:          blockatlas.DelegationTypeAuto,
		},
	}
	result, _ := normalizeValidator(v)
	assert.Equal(t, expected, result)
}

var validator1 = blockatlas.StakeValidator{
	ID:     "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "",
		Description: "",
		Image:       "",
		Website:     "",
	},
	Details: blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: 10,
		},
		LockTime:      0,
		MinimumAmount: "0",
	},
}

var validatorMap = blockatlas.ValidatorMap{
	"one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy": validator1,
}

func TestNormalizeDelegations(t *testing.T) {
	var delegations []Delegation
	err := json.Unmarshal([]byte(delegationsSrc), &delegations)
	assert.NoError(t, err)
	assert.NotNil(t, delegations)

	expected := []blockatlas.Delegation{
		{
			Delegator: validator1,
			Value:     "10000000000000000000",
			Status:    blockatlas.DelegationStatusActive,
		},
	}
	result := NormalizeDelegations(delegations, validatorMap)
	assert.Equal(t, expected, result)

}
