package terra

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/stretchr/testify/assert"
)

const validatorSrc = `
{
	"operatorAddress":"terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9",
	"consensusPubkey":"terravalconspub1zcjduepqqhddjnl8uyudw6qkdtqachp9eexgxuexk0g7y86swg6mhxx8fklqmruxaj",
	"description":{
		"moniker":"hashed",
		"identity":"",
		"website":"",
		"details":"",
		"profileIcon":""
	},
	"tokens":"38068697723657",
	"delegatorShares":"38072504974063.539668711908260929",
	"votingPower":{
		"amount":"38068697000000",
		"weight":"0.17918231449278281226"
	},
	"commissionInfo":{
		"rate":"0.100000000000000000",
		"maxRate":"0.200000000000000000",
		"maxChangeRate":"0.010000000000000000",
		"updateTime":"2019-04-24T06:00:00Z"
	},
	"upTime":1,
	"status":"active",
	"rewardsPool":{
		"total":"96307562684.9429864235676171098",
		"denoms":[
			{
				"denom":"uluna",
				"amount":"31304951486.993503510079423492",
				"adjustedAmount":"31304951486.993503510079423492"
			},
			{
				"denom":"ukrw",
				"amount":"23196862000540.698448130374238539",
				"adjustedAmount":"64977204483.3072785661915244777"
			},
			{
				"denom":"usdr",
				"amount":"5115482.983843863860806657",
				"adjustedAmount":"22931885.07627057320072834121"
			},
			{
				"denom":"uusd",
				"amount":"375.706218008274534676",
				"adjustedAmount":"1217.53030792722960875746"
			},
			{
				"denom":"umnt",
				"amount":"2089380667.833843813037780727",
				"adjustedAmount":"2473612.03562584686633204143"
			}
		]
	},
	"stakingReturn":"0.08275367696699465695"
	}`

const delegationsSrc = `
[
  {
    "delegator_address": "terra135qla4294zxarqhhgxsx0sw56yssa3z006ape0",
    "validator_address": "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5",
    "shares": "109999.000001746056062372",
    "balance": "0"
  }
]`

const unbondingDelegationsSrc = `
[
  {
    "delegator_address": "terra135qla4294zxarqhhgxsx0sw56yssa3z006ape0",
    "validator_address": "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5",
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

var stakingPool = Pool{"1222", "200"}

var terraValidator = Validator{Commission: TerraCommission{Rate: "0.4"}}

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 8.275367696699467},
			LockTime:      1814400,
			MinimumAmount: "0",
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
	result := normalizeValidator(v)
	assert.Equal(t, expected, result)
}

var validator1 = blockatlas.StakeValidator{
	ID:     "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Gazua",
		Description: "to the moon",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/terra/validators/assets/terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5/logo.png",
		Website:     "https://terra.money",
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
	"terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5": validator1,
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
