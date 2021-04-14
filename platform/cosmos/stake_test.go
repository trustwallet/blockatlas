package cosmos

import (
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
)

var (
	validatorSrc, _            = mock.JsonStringFromFilePath("mocks/" + "validator.json")
	delegationsSrc, _          = mock.JsonStringFromFilePath("mocks/" + "delegation.json")
	unbondingDelegationsSrc, _ = mock.JsonStringFromFilePath("mocks/" + "unbonding.json")
	stakingPool                = Pool{"1222", "200"}
	cosmosValidator            = Validator{Commission: CosmosCommission{CosmosCommissionRates{Rate: "0.4"}}}
	inflation                  = 0.7
	validatorMap               = blockatlas.ValidatorMap{
		"cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys": validator1,
	}
	validator1 = blockatlas.StakeValidator{
		ID:     "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Certus One",
			Description: "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by Certus One Inc. By delegating, you confirm that you are aware of the risk of slashing and that Certus One Inc is not liable for any potential damages to your investment.",
			Image:       "https://assets.trustwalletapp.com/blockchains/cosmos/validators/assets/cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys/logo.png",
			Website:     "https://certus.one",
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: 9.259735525366604,
			},
			LockTime:      lockTime,
			MinimumAmount: minimumAmount,
		},
	}
)

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 462.6619201898575},
			LockTime:      lockTime,
			MinimumAmount: minimumAmount,
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

func TestNormalizeDelegations(t *testing.T) {
	var delegations []DelegationValue
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
