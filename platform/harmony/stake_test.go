package harmony

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
	"testing"
)

const validatorSrc = `
 {
	"validator": {
		"bls-public-keys": [
			"29fb5d202e2f6f7955b425dc706fc0b29047067e51ba583fbb017c0f51186d5e1eaf6dd4059848be311ab5a49d625309"
		],
		"last-epoch-in-committee": 18,
		"min-self-delegation": 10999000000000000000000,
		"max-total-delegation": 100000000000000000000000000,
		"rate": "0.100000000000000000",
		"max-rate": "0.100000000000000000",
		"max-change-rate": "0.100000000000000000",
		"update-height": 88,
		"name": "sieemma node",
		"identity": "sieemma node by ankr",
		"website": "www.ankr.com",
		"security-contact": "info@ankr.com",
		"details": "This validator is launched from app.ankr.com",
		"creation-height": 88,
		"address": "one1v8pukmelacy3xdap773rpg5pax3tmu40wmwr2j",
		"delegations": [
			{
				"delegator-address": "one1v8pukmelacy3xdap773rpg5pax3tmu40wmwr2j",
				"amount": 10999000000000000000000,
				"reward": 2328233463148225437028,
				"undelegations": []
			}
		]
	},
	"current-epoch-performance": {
		"current-epoch-signing-percent": {
			"current-epoch-signed": 3,
			"current-epoch-to-sign": 3,
			"num-beacon-blocks-until-next-epoch": 37,
			"current-epoch-signing-percentage": "1.000000000000000000"
		}
	},
	"metrics": {
		"by-bls-key": [
			{
				"key": {
					"bls-public-key": "29fb5d202e2f6f7955b425dc706fc0b29047067e51ba583fbb017c0f51186d5e1eaf6dd4059848be311ab5a49d625309",
					"group-percent": "0.056856187290969900",
					"effective-stake": "85000000000000000000000.000000000000000000",
					"earning-account": "one1v8pukmelacy3xdap773rpg5pax3tmu40wmwr2j",
					"overall-percent": "0.018193979933110368",
					"shard-id": 1
				},
				"earned-reward": 4478494623655913952
			}
		]
	},
	"total-delegation": 10999000000000000000000,
	"currently-in-committee": true,
	"epos-status": "currently elected",
	"epos-winning-stake": "85000000000000000000000.000000000000000000",
	"booted-status": null,
	"lifetime": {
		"reward-accumulated": 2328233463148225437028,
		"blocks": {
			"to-sign": 525,
			"signed": 504
		},
		"apr": "12.345"
	}
}`

const delegationsSrc = `
[
	{
		"validator_address": "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
		"delegator_address": "one1pf75h0t4am90z8uv3y0dgunfqp4lj8wr3t5rsp",
		"amount": 12345678900000000000,
		"reward": 15854399877248931866418,
		"Undelegations": []
	}
]`

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: v.Active,
		ID:     v.Info.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 12.345},
			LockTime:      lockTime,
			MinimumAmount: "1000",
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}

	var apr float64
	var err error
	if apr, err = strconv.ParseFloat(v.Lifetime.Apr, 64); err != nil {
		apr = 0
	}

	result := normalizeValidator(v, apr)
	assert.Equal(t, expected, result)
}

var validator1 = blockatlas.StakeValidator{
	ID:     "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Harmony One",
		Description: "Stake and earn rewards with the most secure and stable validator. Operated by Harmony One Inc.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/harmony/validators/assets/one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy/logo.png",
		Website:     "https://harmony.one",
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
			Value:     "12345678900000000000",
			Status:    blockatlas.DelegationStatusActive,
		},
	}
	result := NormalizeDelegations(delegations, validatorMap)
	assert.Equal(t, expected, result)
}

func TestHexToInt(t *testing.T) {
	result, _ := hexToInt("0x604800")

	assert.Equal(t, uint64(6309888), result)
}

func TestGetDetails(t *testing.T) {
	var expected = blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: 10},
		LockTime:      lockTime,
		MinimumAmount: "1000",
		Type:          blockatlas.DelegationTypeDelegate,
	}

	result := getDetails(10)
	assert.Equal(t, expected, result)
}

func TestGetValidators(t *testing.T) {
	var c Client

	p := Platform{
		client: c,
	}

	var validators = []Validator{
		Validator{
			Info: ValidatorInfo{
				Address: "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
			},
			Active:   true,
			Lifetime: LifetimeInfo{Apr: "10"},
		},
	}

	rpcCallStub = func(c *Client, result interface{}, method string, params interface{}) error {
		jsonData, _ := json.Marshal(validators)
		_ = json.Unmarshal(jsonData, result)
		return nil
	}

	result, _ := p.GetValidators()
	assert.Equal(t, lockTime, result[0].Details.LockTime)
	assert.Equal(t, float64(10), result[0].Details.Reward.Annual)
}

func TestGetDelegation(t *testing.T) {
	var c Client

	p := Platform{
		client: c,
	}

	var delegations = []Delegation{
		Delegation{
			DelegatorAddress: "one1a0au0p33zrns49h3qw7prn02s4wphu0ggcqrhm",
			ValidatorAddress: "one1a0au0p33zrns49h3qw7prn02s4wphu0ggcqrhm",
			Amount:           100,
		},
	}

	var validators = []Validator{
		Validator{
			Info: ValidatorInfo{
				Address: "one1a0au0p33zrns49h3qw7prn02s4wphu0ggcqrhm",
			},
			Active:   true,
			Lifetime: LifetimeInfo{Apr: "10"},
		},
	}

	rpcCallStub = func(c *Client, result interface{}, method string, params interface{}) error {
		if method == "hmy_getAllValidatorInformation" {
			jsonData, _ := json.Marshal(validators)
			_ = json.Unmarshal(jsonData, result)
		} else {
			jsonData, _ := json.Marshal(delegations)
			_ = json.Unmarshal(jsonData, result)
		}
		return nil
	}

	result, _ := p.GetDelegations("one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy")
	assert.Equal(t, delegations[0].DelegatorAddress, result[0].Delegator.ID)
}

func TestGeBalance(t *testing.T) {
	var c Client

	p := Platform{
		client: c,
	}

	var balance = "0x100"

	rpcCallStub = func(c *Client, result interface{}, method string, params interface{}) error {
		jsonData, _ := json.Marshal(balance)
		_ = json.Unmarshal(jsonData, result)
		return nil
	}

	result, _ := p.UndelegatedBalance("one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy")
	assert.Equal(t, "256", result)
}
