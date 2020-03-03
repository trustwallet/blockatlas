package ethereum

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const tokenSrc = `
{
	"balance": "0",
	"contract": {
		"contract": "0xa14839c9837657efcde754ebeaf5cbecdd801b2a",
		"address": "0xa14839c9837657efcde754ebeaf5cbecdd801b2a",
		"name": "FusChain",
		"decimals": 18,
		"symbol": "FUS"
	}
}`

var tokenDst = blockatlas.Token{
	Name:     "FusChain",
	Symbol:   "FUS",
	Decimals: 18,
	TokenID:  "0xa14839c9837657efcde754ebeaf5cbecdd801b2a",
	Coin:     coin.ETH,
	Type:     blockatlas.TokenTypeERC20,
}

type testToken struct {
	name        string
	apiResponse string
	expected    *blockatlas.Token
}

func TestNormalizeToken(t *testing.T) {
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: tokenSrc,
		expected:    &tokenDst,
	})
}

func testNormalizeToken(t *testing.T, _test *testToken) {
	var token Token
	err := json.Unmarshal([]byte(_test.apiResponse), &token)
	if err != nil {
		t.Error(err)
		return
	}
	tk, ok := NormalizeToken(&token, coin.ETH)
	if !ok {
		t.Errorf("token: token could not be normalized")
	}

	resJSON, err := json.Marshal(&tk)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("token: token don't equal")
	}
}
