package tron

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var tokenDst = blockatlas.Token{
	Name:     "Test",
	Symbol:   "TST",
	Decimals: 8,
	TokenID:  "1",
	Coin:     195,
	Type:     blockatlas.TokenTypeTRC10,
}

func TestNormalizeToken(t *testing.T) {
	asset := AssetInfo{Name: "Test", Symbol: "TST", ID: "1", Decimals: 8}
	actual := NormalizeToken(asset)
	assert.Equal(t, tokenDst, actual)
}
