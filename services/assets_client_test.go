package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

func TestGetImage(t *testing.T) {
	coin := coin.Coin{Handle: "cosmos"}
	image := GetImage(coin, "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys")

	expected := "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys/logo.png"

	assert.Equal(t, expected, image)
}
