package tron

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/mock"
)

var (
	delegationsSrc1, _ = mock.JsonStringFromFilePath("mocks/" + "delegation.json")
	delegationsSrc2, _ = mock.JsonStringFromFilePath("mocks/" + "delegation_2.json")

	validator1 = blockatlas.StakeValidator{
		ID:     "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Sesameseed",
			Description: "Sesameseed is a blockchain community providing fair and transparent representation in delegated governance by rewarding Voters for their participation on Tron and Ontology.",
			Image:       "https://assets.trustwalletapp.com/blockchains/tron/validators/assets/tgzz8gjyiyrqpfmdwnlxfgpulvnmpcswvp/logo.png",
			Website:     "https://www.sesameseed.org",
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: 4.32,
			},
			LockTime:      259200,
			MinimumAmount: "1000000",
		},
	}

	validator2 = blockatlas.StakeValidator{
		ID:     "TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw",
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "InfStones",
			Description: "World's leading cloud infrastructure and staking as a service provicer for blockchains. Supernodes on EOS, TRON, VeChain, Ontology, LOOM, IOST and many other chains.",
			Image:       "https://assets.trustwalletapp.com/blockchains/tron/validators/assets/tpmgfspxlqgom8skutrbhcdkthjrhfbgkw/logo.png",
			Website:     "https://infstones.io/",
		},
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: 4.32,
			},
			LockTime:      259200,
			MinimumAmount: "1000000",
		},
	}

	delegation1 = blockatlas.Delegation{
		Delegator: validator1,
		Value:     "21000000",
		Status:    blockatlas.DelegationStatusPending,
	}
	delegation2 = blockatlas.Delegation{
		Delegator: validator2,
		Value:     "5000000",
		Status:    blockatlas.DelegationStatusPending,
	}
	delegation3 = blockatlas.Delegation{
		Delegator: validator2,
		Value:     "5000000",
		Status:    blockatlas.DelegationStatusPending,
	}
	delegation4 = blockatlas.Delegation{
		Delegator: validator2,
		Value:     "5000000",
		Status:    blockatlas.DelegationStatusActive,
	}
	validatorMap = blockatlas.ValidatorMap{
		"TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp": validator1,
		"TPMGfspxLQGom8sKutrbHcDKtHjRHFbGKw": validator2,
	}
)

func TestNormalizeValidator(t *testing.T) {
	validator := Validator{Address: "414d1ef8673f916debb7e2515a8f3ecaf2611034aa"}

	actual, _ := normalizeValidator(validator)
	expected := blockatlas.Validator{
		ID:     "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp",
		Status: true,
		Details: blockatlas.StakingDetails{
			Reward: blockatlas.StakingReward{
				Annual: Annual,
			},
			LockTime:      259200,
			MinimumAmount: "1000000",
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}
	assert.Equal(t, expected, actual)
}

func TestNormalizeDelegations(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  []blockatlas.Delegation
	}{
		{"Status Pending", delegationsSrc1, []blockatlas.Delegation{delegation1, delegation2, delegation3}},
		{"Status Active", delegationsSrc2, []blockatlas.Delegation{delegation4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testNormalizeDelegations(t, tt.value, tt.want)
		})
	}
}

func testNormalizeDelegations(t *testing.T, value string, want []blockatlas.Delegation) {
	var accountData *AccountData
	err := json.Unmarshal([]byte(value), &accountData)
	assert.NoError(t, err)
	assert.NotNil(t, accountData)
	result := NormalizeDelegations(accountData, validatorMap)
	assert.Equal(t, result, want)
}
