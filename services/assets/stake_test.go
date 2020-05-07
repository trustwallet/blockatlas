package assets

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var (
	cosmosCoin = coin.Coin{Handle: "cosmos"}
	tezosCoin  = coin.Coin{Handle: "tezos"}

	validators = []blockatlas.Validator{
		{
			ID:     "test1",
			Status: true,
		},
		{
			ID:     "test2",
			Status: true,
		},
	}
	assets1 = []AssetValidator{
		{
			ID:          "test1",
			Name:        "Spider",
			Description: "yo",
			Website:     "https://tw.com",
			Status:      ValidatorStatus{Disabled: false},
		},
		{
			ID:          "test2",
			Name:        "Man",
			Description: "lo",
			Website:     "https://tw.com",
			Status:      ValidatorStatus{Disabled: true},
		},
	}
	assets2 = []AssetValidator{
		{
			ID:          "test1",
			Name:        "Spider",
			Description: "yo",
			Website:     "https://tw.com",
			Status:      ValidatorStatus{Disabled: true},
		},
		{
			ID:          "test2",
			Name:        "Man",
			Description: "lo",
			Website:     "https://tw.com",
			Status:      ValidatorStatus{Disabled: true},
		},
	}

	tezosValidators = []blockatlas.Validator{
		{
			ID:     "test1",
			Status: true,
		},
	}

	tezosAssets1 = []AssetValidator{
		{
			ID:          "test1",
			Name:        "🐠stake.fish",
			Description: "Leading validator for Proof of Stake blockchains.",
			Website:     "https://stake.fish/",
			Status:      ValidatorStatus{Disabled: false},
			Staking:     StakingInfo{MinDelegation: 10},
		},
	}
	expectTezosVal1 = blockatlas.StakeValidator{
		ID: "test1", Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "🐠stake.fish",
			Description: "Leading validator for Proof of Stake blockchains.",
			Image:       getImage(tezosCoin, "test1"),
			Website:     "https://stake.fish/",
		},
		Details: blockatlas.StakingDetails{
			MinimumAmount: blockatlas.Amount("10"),
		},
	}
	expectedCosmosStakeValidator = blockatlas.StakeValidator{
		ID: "test1", Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Spider",
			Description: "yo",
			Image:       getImage(cosmosCoin, "test1"),
			Website:     "https://tw.com",
		},
		Details: blockatlas.StakingDetails{
			MinimumAmount: blockatlas.Amount("0"),
		},
	}
	expectedCosmosStakeValidatorDisabled1 = blockatlas.StakeValidator{
		ID: "test1", Status: false,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Spider",
			Description: "yo",
			Image:       getImage(cosmosCoin, "test1"),
			Website:     "https://tw.com",
		},
		Details: blockatlas.StakingDetails{
			MinimumAmount: blockatlas.Amount("0"),
		},
	}
	expectedCosmosStakeValidatorDisabled2 = blockatlas.StakeValidator{
		ID: "test2", Status: false,
		Info: blockatlas.StakeValidatorInfo{
			Name:        "Man",
			Description: "lo",
			Image:       getImage(cosmosCoin, "test2"),
			Website:     "https://tw.com",
		},
		Details: blockatlas.StakingDetails{
			MinimumAmount: blockatlas.Amount("0"),
		},
	}
)

func TestGetImage(t *testing.T) {
	image := getImage(cosmosCoin, "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp")
	expected := "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp/logo.png"
	assert.Equal(t, expected, image)
}

func TestCalcAnnual(t *testing.T) {
	type args struct {
		annual     float64
		commission float64
	}

	tests := []struct {
		name   string
		args   args
		wanted float64
	}{
		{
			name: "test TestCalcAnnual 1",
			args: args{
				annual:     10,
				commission: 10,
			},
			wanted: 9,
		},
		{
			name: "test TestCalcAnnual 2",
			args: args{
				annual:     100,
				commission: 10,
			},
			wanted: 90,
		},
		{
			name: "test TestCalcAnnual 3",
			args: args{
				annual:     20,
				commission: 10,
			},
			wanted: 18,
		},
		{
			name: "test TestCalcAnnual 1",
			args: args{
				annual:     30,
				commission: 10,
			},
			wanted: 27,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfo := calculateAnnual(tt.args.annual, tt.args.commission)
			assert.Equal(t, tt.wanted, gotInfo)
		})
	}
}

func TestNormalizeValidator(t *testing.T) {
	result := normalizeValidator(validators[0], assets1[0], cosmosCoin)
	assert.Equal(t, expectedCosmosStakeValidator, result)
}

func Test_normalizeValidators(t *testing.T) {
	type args struct {
		assets     AssetValidators
		validators []blockatlas.Validator
		coin       coin.Coin
	}
	tests := []struct {
		name string
		args args
		want blockatlas.StakeValidators
	}{
		{"normalize validator 1", args{assets1, validators, cosmosCoin}, blockatlas.StakeValidators{expectedCosmosStakeValidator, expectedCosmosStakeValidatorDisabled2}},
		{"normalize validator 2", args{assets2, validators, cosmosCoin}, blockatlas.StakeValidators{expectedCosmosStakeValidatorDisabled1, expectedCosmosStakeValidatorDisabled2}},
		{"normalize tezos validator", args{tezosAssets1, tezosValidators, tezosCoin}, blockatlas.StakeValidators{expectTezosVal1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeValidators(tt.args.assets, tt.args.validators, tt.args.coin)
			assert.Equal(t, tt.want, got)
		})
	}
}
