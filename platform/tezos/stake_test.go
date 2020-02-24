package tezos

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const accountSrc = `
{
  "delegate": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
  "balance": "91237897"
}`

const validatorSrc = `
[
	{"pkh":"tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m","rolls":3726}
]
`

var validator = blockatlas.Validator{
	Status: true,
	ID:     "tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m",
	Details: blockatlas.StakingDetails{
		Reward:        blockatlas.StakingReward{Annual: Annual},
		MinimumAmount: blockatlas.Amount("0"),
		Type:          blockatlas.DelegationTypeDelegate,
	},
}

var stakeValidator = blockatlas.StakeValidator{
	ID:     "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "stake.fish",
		Description: "Leading validator for Proof of Stake blockchains. Stake your cryptocurrencies with us. We know validating.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/tezos/validators/assets/tz2fcnbrerxtattnx6iimr1uj5jsdxvdhm93/logo.png",
		Website:     "https://stake.fish/",
	},
	Details: getDetails(),
}

var validatorMap = blockatlas.ValidatorMap{
	"tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93": stakeValidator,
}

var delegationsBalance = "91237897"

var delegation = blockatlas.DelegationsPage{
	{
		Delegator: stakeValidator,
		Value:     delegationsBalance,
		Status:    blockatlas.DelegationStatusActive,
	},
}

func TestNormalizeValidator(t *testing.T) {
	var v []Validator
	err := json.Unmarshal([]byte(validatorSrc), &v)
	assert.Nil(t, err)
	result := normalizeValidator(v[0])
	assert.Equal(t, validator, result)
}

func TestNormalizeDelegations(t *testing.T) {
	var account Account
	err := json.Unmarshal([]byte(accountSrc), &account)
	assert.NoError(t, err)
	assert.NotNil(t, account)
	result, err := NormalizeDelegation(account, validatorMap)
	assert.NoError(t, err)
	assert.Equal(t, delegation, result)
}
