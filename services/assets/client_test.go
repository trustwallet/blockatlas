package assets

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

var c = coin.Coin{Handle: "cosmos"}
var validators = []blockatlas.Validator{
	{
		Coin:   c,
		ID:     "test",
		Status: true,
	},
	{
		Coin:   c,
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
	coin := coin.Coin{Handle: "cosmos"}
	image := GetImage(coin, "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys")

	expected := "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys/logo.png"

	assert.Equal(t, expected, image)
}

func TestNormalizeValidator(t *testing.T) {

	result := NormalizeValidator(validators[0], assets[0])

	assert.Equal(t, expectedStakeValidator, result)
}

func TestNormalizeValidators(t *testing.T) {

	result := NormalizeValidators(validators, assets)

	expected := []blockatlas.StakeValidator{expectedStakeValidator}

	assert.Equal(t, expected, result)
}
