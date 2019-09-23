package assets

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var c = coin.Coin{Handle: "cosmos"}
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
	image := GetImage(c, "TGzz8gjYiYRqpfmDwnLxfgPuLVNmpCswVp")

	expected := "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/tgzz8gjyiyrqpfmdwnlxfgpulvnmpcswvp/logo.png"

	assert.Equal(t, expected, image)
}

func TestNormalizeValidator(t *testing.T) {

	result := NormalizeValidator(validators[0], assets[0], c)

	assert.Equal(t, expectedStakeValidator, result)
}

func TestNormalizeValidators(t *testing.T) {

	result := NormalizeValidators(validators, assets, c)

	expected := []blockatlas.StakeValidator{expectedStakeValidator}

	assert.Equal(t, expected, result)
}
