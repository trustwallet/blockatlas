package cosmos

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
)

const validatorSrc = `
{
  "operator_address": "cosmosvaloper1lktjhnzkpkz3ehrg8psvmwhafg56kfss3q3t8m",
  "consensus_pubkey": "cosmosvalconspub1zcjduepqelcwpat987h9yq0ck6g9fsc8t0mththk547gwvk0w4wnkpl0stnspr3hdc",
  "jailed": false,
  "status": 2,
  "tokens": "1557750969185",
  "delegator_shares": "1557750969185.000000000000000000",
  "description": {
    "moniker": "Umbrella ☔",
    "identity": "A530AC4D75991FE2",
    "website": "https://umbrellavalidator.com",
    "details": "One of the winners of Cosmos Game of Stakes, and HackAtom3."
  },
  "unbonding_height": "0",
  "unbonding_time": "1970-01-01T00:00:00Z",
  "commission": {
    "commission_rates": {
      "rate": "0.070400000000000000",
      "max_rate": "1.000000000000000000",
      "max_change_rate": "0.100000000000000000",
      "update_time": "2019-08-05T07:10:23.689753607Z"
    }
  },
  "min_self_delegation": "1"
}`

const delegationsSrc = `
[
  {
    "delegator_address": "cosmos135qla4294zxarqhhgxsx0sw56yssa3z0f78pm0",
    "validator_address": "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
    "shares": "109999.000001746056062372",
    "balance": "0"
  }
]`

const unbondingDelegationsSrc = `
[
  {
    "delegator_address": "cosmos135qla4294zxarqhhgxsx0sw56yssa3z0f78pm0",
    "validator_address": "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
    "entries": [
      {
        "creation_height": "0",
        "completion_time": "2020-01-01T06:54:18.441436491Z",
        "initial_balance": "109999",
        "balance": "109999"
      }
    ]
  }
]`

var delegateDst = blockatlas.Tx{
	ID:     "11078091D1D5BD84F4275B6CEE02170428944DB0E8EEC37E980551435F6D04C7",
	Coin:   coin.ATOM,
	From:   "cosmos1237l0vauhw78qtwq045jd24ay4urpec6r3xfn3",
	To:     "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
	Fee:    "5000",
	Date:   1564632616,
	Block:  1258202,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    blockatlas.AnyActionDelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     "ATOM",
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    "49920",
	},
}

var unDelegateDst = blockatlas.Tx{
	ID:     "A1EC36741FEF681F4A77B8F6032AD081100EE5ECB4CC76AEAC2174BC6B871CFE",
	Coin:   coin.ATOM,
	From:   "cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94",
	To:     "cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
	Fee:    "5000",
	Date:   1564624521,
	Block:  1257037,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    blockatlas.AnyActionUndelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     "ATOM",
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    "5100000000",
	},
}

var stakingPool = Pool{"1222", "200"}

var cosmosValidator = Validator{Commission: CosmosCommission{CosmosCommissionRates{Rate: "0.4"}}}

var inflation = 0.7

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 462.6619201898575},
			LockTime:      1814400,
			MinimumAmount: "0",
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
	result := normalizeValidator(v, stakingPool, inflation)
	assert.Equal(t, expected, result)
}

func TestCalculateAnnualReward(t *testing.T) {
	result := CalculateAnnualReward(Pool{"1222", "200"}, inflation, cosmosValidator)
	assert.Equal(t, 298.61999703347686, result)
}

var validator1 = blockatlas.StakeValidator{
	ID:     "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Certus One",
		Description: "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by Certus One Inc. By delegating, you confirm that you are aware of the risk of slashing and that Certus One Inc is not liable for any potential damages to your investment.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys/logo.png",
		Website:     "https://certus.one",
	},
	Details: blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: 9.259735525366604,
		},
		LockTime:      1814400,
		MinimumAmount: "0",
	},
}

var validatorMap = blockatlas.ValidatorMap{
	"cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys": validator1,
}

func TestNormalizeDelegations(t *testing.T) {
	var delegations []Delegation
	err := json.Unmarshal([]byte(delegationsSrc), &delegations)
	assert.NoError(t, err)
	assert.NotNil(t, delegations)

	expected := []blockatlas.Delegation{
		{
			Delegator: validator1,
			Value:     "109999",
			Status:    blockatlas.DelegationStatusActive,
		},
	}
	result := NormalizeDelegations(delegations, validatorMap)
	assert.Equal(t, expected, result)
}

func TestNormalizeUnbondingDelegations(t *testing.T) {
	var delegations []UnbondingDelegation
	err := json.Unmarshal([]byte(unbondingDelegationsSrc), &delegations)
	assert.NoError(t, err)
	assert.NotNil(t, delegations)

	expected := []blockatlas.Delegation{
		{
			Delegator: validator1,
			Value:     "109999",
			Status:    blockatlas.DelegationStatusPending,
			Metadata: blockatlas.DelegationMetaDataPending{
				AvailableDate: 1577861658,
			},
		},
	}
	result := NormalizeUnbondingDelegations(delegations, validatorMap)
	assert.Equal(t, expected, result)
}
