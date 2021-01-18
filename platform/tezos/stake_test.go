package tezos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	accountSrc, _          = mock.JsonStringFromFilePath("mocks/" + "account.json")
	validatorSrc, _        = mock.JsonStringFromFilePath("mocks/" + "validator.json")
	mockedTezosResponse, _ = mock.JsonStringFromFilePath("mocks/" + "delegation_response.json")

	validator = blockatlas.Validator{
		Status: true,
		ID:     "tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m",
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: Annual},
			MinimumAmount: types.Amount("0"),
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}

	stakeValidator = blockatlas.StakeValidator{
		ID:     "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "stake.fish",
			Description: "Leading validator for Proof of Stake blockchains. Stake your cryptocurrencies with us. We know validating.",
			Image:       "https://assets.trustwalletapp.com/blockchains/tezos/validators/assets/tz2fcnbrerxtattnx6iimr1uj5jsdxvdhm93/logo.png",
			Website:     "https://stake.fish/",
		},
		Details: getDetails(),
	}

	validatorMap = blockatlas.ValidatorMap{
		"tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93": stakeValidator,
	}

	delegationsBalance = "91237897"

	delegation = blockatlas.DelegationsPage{
		{
			Delegator: stakeValidator,
			Value:     delegationsBalance,
			Status:    blockatlas.DelegationStatusActive,
		},
	}
)

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

func TestPlatform_isValidatorActive(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()

	p := Init(server.URL, server.URL, server.URL)
	assert.True(t, p.isValidatorActive("tz1V3yg82mcrPJbegqVCPn6bC8w1CSTRp3f8"))
}

func createMockedAPI() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/chains/main/blocks/head/context/delegates/tz1V3yg82mcrPJbegqVCPn6bC8w1CSTRp3f8", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedTezosResponse); err != nil {
			panic(err)
		}
	})

	return r
}
