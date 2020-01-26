package tezos

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const delegationsSrc = `
{
  "op": {
    "signature": "sigvHd2YBByFXU8nL4CZKSTYXNdMapMsJw1f239YRRjgz9NvrTyA6iGnpBDhi9kCB4zMHysrg9H4jxcpPH975WiQtEmkMjb5",
    "blockUuid": "4b292c55-41ba-4383-a1d6-03fb71b88f41",
    "opHash": "opGphHGNEZZN5rF78yxwe9BJydxYA2yqxECnZR6s6HcxXtCg8Sj",
    "uuid": "e4ec0e07-1601-4da3-bd92-090a820ed369",
    "chainId": "NetXdQprcVkpaWU",
    "blockHash": "BLkscXpE63gajVzmgBS7fQx63hERKQRCZFGtMXdYY6WPThHyji7",
    "protocol": "PsBabyM1eUXZseaJdmXFApDSBqj8YBfwELoxZHHW77EMcAbbwAS",
    "branch": "BKqtLegZfdPR3USyYYcMpedB59W5eUBuFZAVpVMPpFgEvMcZjr1",
    "blockLevel": 791778,
    "blockTimestamp": "2020-01-22T22:13:38Z",
    "insertedTimestamp": "2020-01-22 22:14:05.937406 UTC"
  },
  "delegation": {
    "storageLimit": "257",
    "delegate": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
    "opUuid": "e4ec0e07-1601-4da3-bd92-090a820ed369",
    "uuid": "6459fcd9-5eee-4999-ac4d-92330b9eaab3",
    "gasLimit": "10600",
    "kind": "delegation",
    "operationResultStatus": "applied",
    "fee": "1500",
    "operationResultUuid": "791f6ec7-ecec-43d5-82ca-a1497be0188c",
    "operationResultConsumedGas": "10000",
    "counter": "2409130",
    "operationResultErrors": null,
    "source": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
    "insertedTimestamp": "2020-01-22 22:15:12.038586 UTC",
    "metadataUuid": "85b42f50-1a89-421c-bcb0-a06926941bc4"
  }
}`

const accountSrc = `
{
  "address": "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
  "delegate": "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
  "total_balance": 68995.611927,
  "is_delegated": true
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

var delegationsBalance = "68995611927"

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

func Test_removeDecimals(t *testing.T) {
	tests := []struct {
		name   string
		volume float64
		want   string
	}{
		{"one float precision", 9.5, "9500000"},
		{"zero float precision", 9, "9000000"},
		{"five float precision", 9.00005, "9000050"},
		{"six float precision", 9.000005, "9000005"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDecimals(tt.volume); got != tt.want {
				t.Errorf("removeDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}
