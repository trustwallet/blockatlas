package tezos

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestNormalizeStakeValidators(t *testing.T) {
	tests := []struct {
		name      string
		baker     Baker
		validator blockatlas.StakeValidator
	}{
		{
			name: "Test normalize negative free space",
			baker: Baker{

				Address:           "tz1fJHFn6sWEd3NnBPngACuw2dggTv6nQZ7g",
				Name:              "Baking Team",
				Logo:              "https://services.tzkt.io/v1/avatars/tz1fJHFn6sWEd3NnBPngACuw2dggTv6nQZ7g",
				FreeSpace:         -54723.23208699998,
				Fee:               0,
				MinDelegation:     1000,
				OpenForDelegation: true,
				EstimatedRoi:      0.060643,
				ServiceHealth:     "active",
			},
			validator: blockatlas.StakeValidator{
				ID:     "tz1fJHFn6sWEd3NnBPngACuw2dggTv6nQZ7g",
				Status: false,
				Info: blockatlas.StakeValidatorInfo{
					Name:  "Baking Team",
					Image: "https://services.tzkt.io/v1/avatars/tz1fJHFn6sWEd3NnBPngACuw2dggTv6nQZ7g",
				},
				Details: blockatlas.StakingDetails{
					Reward: blockatlas.StakingReward{
						Annual: 6.06,
					},
					MinimumAmount: "1000",
					Type:          "delegate",
				},
			},
		},
		{
			name: "Test normalize negative free space",
			baker: Baker{
				Address:           "tz1gcna2xxZj2eNp1LaMyAhVJ49mEFj4FH26",
				Name:              "Exaion Baker",
				Logo:              "https://services.tzkt.io/v1/avatars/tz1gcna2xxZj2eNp1LaMyAhVJ49mEFj4FH26",
				FreeSpace:         7947.711756999997,
				Fee:               0.04,
				MinDelegation:     0,
				OpenForDelegation: true,
				EstimatedRoi:      0.058896,
				ServiceHealth:     "active",
			},
			validator: blockatlas.StakeValidator{
				ID:     "tz1gcna2xxZj2eNp1LaMyAhVJ49mEFj4FH26",
				Status: true,
				Info: blockatlas.StakeValidatorInfo{
					Name:  "Exaion Baker",
					Image: "https://services.tzkt.io/v1/avatars/tz1gcna2xxZj2eNp1LaMyAhVJ49mEFj4FH26",
				},
				Details: blockatlas.StakingDetails{
					Reward: blockatlas.StakingReward{
						Annual: 5.89,
					},
					MinimumAmount: "0",
					Type:          "delegate",
				},
			},
		},
		{
			name: "Test",
			baker: Baker{
				Address:           "tz1dbfppLAAxXZNtf2SDps7rch3qfUznKSoK",
				Name:              "Coinhouse",
				Logo:              "https://services.tzkt.io/v1/avatars/tz1dbfppLAAxXZNtf2SDps7rch3qfUznKSoK",
				FreeSpace:         91005.65305700002,
				Fee:               0.08,
				MinDelegation:     0,
				OpenForDelegation: false,
				EstimatedRoi:      0.056598,
				ServiceHealth:     "active",
			},
			validator: blockatlas.StakeValidator{
				ID:     "tz1dbfppLAAxXZNtf2SDps7rch3qfUznKSoK",
				Status: false,
				Info: blockatlas.StakeValidatorInfo{
					Name:  "Coinhouse",
					Image: "https://services.tzkt.io/v1/avatars/tz1dbfppLAAxXZNtf2SDps7rch3qfUznKSoK",
				},
				Details: blockatlas.StakingDetails{
					Reward: blockatlas.StakingReward{
						Annual: 5.66,
					},
					MinimumAmount: "0",
					Type:          "delegate",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValidator := NormalizeStakeValidator(tt.baker); !reflect.DeepEqual(gotValidator, tt.validator) {
				t.Errorf("NormalizeStakeValidators() = %v, want %v", gotValidator, tt.validator)
			}
		})
	}
}
