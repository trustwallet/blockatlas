package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const (
	myToken = `
{
	"free": "17199.38841739",
	"frozen": "0.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`
	myTokenAllZero = `
{
	"free": "0.00000000",
	"frozen": "0.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`
	myTokenFreeZero = `
{
	"free": "0.00000000",
	"frozen": "1.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`
	myTokenFrozenAndFreeZero = `
{
	"free": "0.00000000",
	"frozen": "0.00000000",
	"locked": "0.00000001",
	"symbol": "ARN-71B"
}
`
	tokenList = `
[
  {
    "mintable": false,
    "name": "Aeron",
    "original_symbol": "ARN",
    "owner": "bnb1dq8ae0ayztqp99peggq5sygzf3n7u2ze4t0jne",
    "symbol": "ARN-71B",
    "total_supply": "20000000.00000000"
  },
  {
    "mintable": false,
    "name": "BOLT Token",
    "original_symbol": "BOLT",
    "owner": "bnb177ujwmshxu8r9za4vy9ztqn65tmr54ddw958rt",
    "symbol": "BOLT-4C6",
    "total_supply": "995000000.00000000"
  }
]
`
)

var (
	tokenDst = blockatlas.Token{
		Name:     "Aeron",
		Symbol:   "ARN",
		Decimals: 8,
		TokenID:  "ARN-71B",
		Coin:     coin.BNB,
		Type:     blockatlas.TokenTypeBEP2,
	}
	emptyTokenDst = blockatlas.Token{}
)

type testToken struct {
	name        string
	apiResponse string
	expected    blockatlas.Token
	tokens      string
	ok          bool
}

func TestNormalizeToken(t *testing.T) {
	testingTokens := []testToken{
		{
			name:        "Test with not zero balance",
			apiResponse: myToken,
			tokens:      tokenList,
			expected:    tokenDst,
			ok:          true,
		},
		{
			name:        "Test with all zero balance",
			apiResponse: myTokenAllZero,
			tokens:      tokenList,
			expected:    emptyTokenDst,
			ok:          false,
		},
		{
			name:        "Test with only free zero balance",
			apiResponse: myTokenFreeZero,
			tokens:      tokenList,
			expected:    tokenDst,
			ok:          true,
		},
		{
			name:        "Test with free and frozen zero balances",
			apiResponse: myTokenFrozenAndFreeZero,
			tokens:      tokenList,
			expected:    tokenDst,
			ok:          true,
		},
	}
	for _, testToken := range testingTokens {
		t.Run(testToken.name, func(t *testing.T) {
			var srcToken Balance
			err := json.Unmarshal([]byte(testToken.apiResponse), &srcToken)
			assert.Nil(t, err)

			var srcTokens TokenList
			err = json.Unmarshal([]byte(testToken.tokens), &srcTokens)
			assert.Nil(t, err)

			tk, ok := NormalizeToken(&srcToken, &srcTokens)
			assert.Equal(t, testToken.ok, ok, "token: token could not be normalized")
			assert.Equal(t, testToken.expected, tk, "token: token don't equal")
		})
	}
}

func TestDecimalPlaces(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  int
	}{
		{"Test text value with dot", "decimal.places", 6},
		{"Test float value", "1234.543212222", 9},
		{"Test float value", "5.33333333", 8},
		{"Test text value", "decimal", 0},
		{"Test integer value", "4", 0},
		{"Test empty value", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decimalPlaces(tt.value); got != tt.want {
				t.Errorf("decimalPlaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
