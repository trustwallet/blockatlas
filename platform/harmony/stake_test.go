package harmony

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/mock"
)

var (
	validatorSrc, _   = mock.JsonFromFilePathToString("mocks/" + "validator.json")
	delegationsSrc, _ = mock.JsonFromFilePathToString("mocks/" + "delegation.json")
	validatorMap      = blockatlas.ValidatorMap{
		"one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy": validator1,
	}
	validator1 = blockatlas.StakeValidator{
		ID:     "one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Harmony One",
			Description: "Stake and earn rewards with the most secure and stable validator. Operated by Harmony One Inc.",
			Image:       "https://assets.trustwalletapp.com/blockchains/harmony/validators/assets/one1pdv9lrdwl0rg5vglh4xtyrv3wjk3wsqket7zxy/logo.png",
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
)

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
		{
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
		{
			DelegatorAddress: "one1a0au0p33zrns49h3qw7prn02s4wphu0ggcqrhm",
			ValidatorAddress: "one1a0au0p33zrns49h3qw7prn02s4wphu0ggcqrhm",
			Amount:           100,
		},
	}

	var validators = []Validator{
		{
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
