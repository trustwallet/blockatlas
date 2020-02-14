package assets

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var cosmosCoin = coin.Coin{Handle: "cosmos"}
var validators = []blockatlas.Validator{
	{
		ID:     "test",
		Status: true,
	},
	{
		ID:     "test2",
		Status: true,
	},
}
var assets = []AssetValidator{
	{
		ID:          "test",
		Name:        "Spider",
		Description: "yo",
		Website:     "https://tw.com",
	},
}

var expectedStakeValidator = blockatlas.StakeValidator{
	ID: "test", Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Spider",
		Description: "yo",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/test/logo.png",
		Website:     "https://tw.com",
	},
}

func TestGetImage(t *testing.T) {
	image := getImage(cosmosCoin, "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp")

	expected := "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp/logo.png"

	assert.Equal(t, expected, image)
}

func TestNormalizeValidator(t *testing.T) {

	result := normalizeValidator(validators[0], assets[0], cosmosCoin)

	assert.Equal(t, expectedStakeValidator, result)
}

func TestNormalizeValidators(t *testing.T) {

	result := normalizeValidators(validators, assets, cosmosCoin)

	expected := []blockatlas.StakeValidator{expectedStakeValidator}

	assert.Equal(t, expected, result)
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
