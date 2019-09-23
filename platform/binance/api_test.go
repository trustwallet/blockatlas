package binance

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
)

const nativeTransferTransaction = `
{
	"blockHeight": 7761368,
	"code": 0,
	"confirmBlocks": 2089441,
	"fromAddr": "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	"hasChildren": 0,
	"log": "Msg 0: ",
	"timeStamp": 1555049867552,
	"toAddr": "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	"txAge": 836729,
	"txAsset": "BNB",
	"txFee": 0.00125,
	"txHash": "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	"txType": "TRANSFER",
	"value": 100000,
	"memo": "test"
}`

const nativeTokenTransferTransaction = `
{
	"blockHeight": 7928667,
	"code": 0,
	"confirmBlocks": 1922024,
	"fromAddr": "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	"hasChildren": 0,
	"log": "Msg 0: ",
	"timeStamp": 1555117625829,
	"toAddr": "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	"txAge": 768924,
	"txAsset": "YLC-D8B",
	"txFee": 0.00125,
	"txHash": "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	"txType": "TRANSFER",
	"value": 2.10572645,
	"memo": "test"
}`

var transferDst = blockatlas.Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	Fee:    "125000",
	Date:   1555049867,
	Block:  7761368,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.Transfer{
		Value:    "10000000000000",
		Decimals: 8,
		Symbol:   "BNB",
	},
}

var nativeTransferDst = blockatlas.Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	Coin:   coin.BNB,
	From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	Fee:    "125000",
	Date:   1555117625,
	Block:  7928667,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.NativeTokenTransfer{
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
	token       string
}

func TestNormalizeTx(t *testing.T) {
	testNormalizeTx(t, &test{
		name:        "transfer",
		apiResponse: nativeTransferTransaction,
		expected:    &transferDst,
		token:       "",
	})
	testNormalizeTx(t, &test{
		name:        "native token transfer",
		apiResponse: nativeTokenTransferTransaction,
		expected:    &nativeTransferDst,
		token:       "YLC-D8B",
	})
}

func testNormalizeTx(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx, ok := NormalizeTx(&srcTx, txTypeTransfer)
	if !ok {
		t.Errorf("transfer: tx could not be normalized")
	}

	resJSON, err := json.Marshal(&tx)
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
		t.Error("transfer: tx don't equal")
	}
}

const myToken = `
{
	"free": "17199.38841739",
	"frozen": "0.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`

const tokenList = `
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

var tokenDst = blockatlas.Token{
	Name:     "Aeron",
	Symbol:   "ARN",
	Decimals: 8,
	TokenID:  "ARN-71B",
	Coin:     coin.BNB,
	Type:     blockatlas.TokenTypeBEP2,
}

type testToken struct {
	name        string
	apiResponse string
	expected    *blockatlas.Token
	tokens      string
}

func TestNormalizeToken(t *testing.T) {
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: myToken,
		tokens:      tokenList,
		expected:    &tokenDst,
	})
}

func testNormalizeToken(t *testing.T, _test *testToken) {
	var srcToken Balance
	err := json.Unmarshal([]byte(_test.apiResponse), &srcToken)
	if err != nil {
		t.Error(err)
		return
	}

	var srcTokens TokenPage
	err = json.Unmarshal([]byte(_test.tokens), &srcTokens)
	if err != nil {
		t.Error(err)
		return
	}

	tk, ok := NormalizeToken(&srcToken, &srcTokens)
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

func TestTokenSymbol(t *testing.T) {
	assert.Equal(t, "UGAS", TokenSymbol("UGAS"))
	assert.Equal(t, "UGAS", TokenSymbol("UGAS-B0C"))
}
