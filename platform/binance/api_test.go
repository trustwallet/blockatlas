package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferTransaction = `
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

const tokenTransferTransaction = `
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

const newOrderTransaction = `
{
      "txHash": "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
      "blockHeight": 10815565,
      "txType": "NEW_ORDER",
      "timeStamp": 1559689901929,
      "fromAddr": "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
      "value": 0.00649878,
      "txAsset": "BNB",
      "txQuoteAsset": "AWC-986",
      "txFee": 0,
      "txAge": 14346340,
      "orderId": "D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30",
      "data": "{\"orderData\":{\"symbol\":\"BNB_AWC-986\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00324939,\"quantity\":2.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30\"}}",
      "code": 0,
      "log": "Msg 0: ",
      "confirmBlocks": 0,
      "memo": "",
      "source": 0,
      "hasChildren": 0
}`

const cancelOrderTransaction = `
{
      "txHash": "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
      "blockHeight": 10815539,
      "txType": "CANCEL_ORDER",
      "timeStamp": 1559689892180,
      "fromAddr": "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
      "txFee": 0,
      "txAge": 14346349,
      "data": "{\"orderData\":{\"symbol\":\"BNB_GTO-908\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00104716,\"quantity\":1.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-28\"}}",
      "code": 0,
      "log": "Msg 0: ",
      "confirmBlocks": 0,
      "memo": "",
      "source": 0,
      "hasChildren": 0
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

var tokenTransferDst = blockatlas.Tx{
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

var newOrderTransferDst = blockatlas.Tx{
	ID:     "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
	Coin:   coin.BNB,
	From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
	Fee:    "0",
	Date:   1559689901,
	Block:  10815565,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.BNB,
		Title:    blockatlas.KeyTitlePlaceOrder,
		Key:      blockatlas.KeyPlaceOrder,
		TokenID:  "BNB_AWC-986",
		Name:     "AWC-986",
		Symbol:   "AWC",
		Value:    "649878",
		Decimals: 8,
	},
}

//TODO: temp test dst
var metaFreeNewOrderTransferDst = blockatlas.Tx{
	ID:     "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
	Coin:   coin.BNB,
	From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
	Fee:    "0",
	Date:   1559689901,
	Block:  10815565,
	Status: blockatlas.StatusCompleted,
	Meta:   nil,
}

var cancelOrdeTransferDst = blockatlas.Tx{
	ID:     "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
	Coin:   coin.BNB,
	From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
	Fee:    "0",
	Date:   1559689892,
	Block:  10815539,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.BNB,
		Title:    blockatlas.KeyTitleCancelOrder,
		Key:      blockatlas.KeyCancelOrder,
		TokenID:  "BNB_GTO-908",
		Name:     "GTO-908",
		Symbol:   "GTO",
		Value:    "104716",
		Decimals: 8,
	},
}

//TODO: temp test dst
var metaFreeCancelOrdeTransferDst = blockatlas.Tx{
	ID:     "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
	Coin:   coin.BNB,
	From:   "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
	Fee:    "0",
	Date:   1559689892,
	Block:  10815539,
	Status: blockatlas.StatusCompleted,
	Meta:   nil,
}

type testTx struct {
	name        string
	apiResponse string
	expected    blockatlas.Tx
	token       string
	wantError   bool
}

func TestNormalizeTx(t *testing.T) {
	testNormalizeTx(t, &testTx{
		name:        "bnb transfer",
		apiResponse: transferTransaction,
		expected:    transferDst,
		token:       "BNB",
		wantError:   false,
	})
	testNormalizeTx(t, &testTx{
		name:        "native token transfer",
		apiResponse: tokenTransferTransaction,
		expected:    tokenTransferDst,
		token:       "YLC-D8B",
		wantError:   false,
	})
	testNormalizeTx(t, &testTx{
		name:        "new order transfer",
		apiResponse: newOrderTransaction,
		expected:    newOrderTransferDst,
		token:       "AWC-986",
		wantError:   true,
	})
	testNormalizeTx(t, &testTx{
		name:        "cancel order transfer",
		apiResponse: cancelOrderTransaction,
		expected:    cancelOrdeTransferDst,
		token:       "GTO-908",
		wantError:   true,
	})
	testNormalizeTx(t, &testTx{
		name:        "new order transfer",
		apiResponse: newOrderTransaction,
		expected:    metaFreeNewOrderTransferDst,
		token:       "AWC-986",
		wantError:   true,
	})
	testNormalizeTx(t, &testTx{
		name:        "cancel order transfer",
		apiResponse: cancelOrderTransaction,
		expected:    metaFreeCancelOrdeTransferDst,
		token:       "GTO-908",
		wantError:   true,
	})
	testNormalizeTx(t, &testTx{
		name:        "normalize error transfer",
		apiResponse: tokenTransferTransaction,
		token:       "GTO-908",
		wantError:   true,
	})
}

func testNormalizeTx(t *testing.T, _test *testTx) {
	t.Run(_test.name, func(t *testing.T) {
		var srcTx Tx
		err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
		assert.Nil(t, err)
		tx, ok := NormalizeTx(&srcTx, _test.token)
		if _test.wantError {
			assert.False(t, ok, "transfer: tx could be normalized")
			return
		}
		assert.True(t, ok, "transfer: tx could not be normalized")
		assert.Equal(t, _test.expected, tx, "transfer: tx don't equal")
	})
}

const AllTransfersType = `[
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
  },
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
  },
  {
    "txHash": "B0677F3436C1B1661E94D192B84B98AA42AC2485D9808357796EE501CBF794F7",
    "blockHeight": 10815565,
    "txType": "NEW_ORDER",
    "timeStamp": 1559689901929,
    "fromAddr": "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
    "value": 0.00649878,
    "txAsset": "BNB",
    "txQuoteAsset": "AWC-986",
    "txFee": 0,
    "txAge": 14346340,
    "orderId": "D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30",
    "data": "{\"orderData\":{\"symbol\":\"BNB_AWC-986\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00324939,\"quantity\":2.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-30\"}}",
    "code": 0,
    "log": "Msg 0: ",
    "confirmBlocks": 0,
    "memo": "",
    "source": 0,
    "hasChildren": 0
  },
  {
    "txHash": "F48DE755170C10F4A4C0E6836A708C33EEF9A7144800F25187D5F2349FD15A34",
    "blockHeight": 10815539,
    "txType": "CANCEL_ORDER",
    "timeStamp": 1559689892180,
    "fromAddr": "bnb16ya67j7kvw8682kka09qujlw5u7lf4geqef0ku",
    "txFee": 0,
    "txAge": 14346349,
    "data": "{\"orderData\":{\"symbol\":\"BNB_GTO-908\",\"orderType\":\"limit\",\"side\":\"buy\",\"price\":0.00104716,\"quantity\":1.00000000,\"timeInForce\":\"GTE\",\"orderId\":\"D13BAF4BD6638FA3AAD6EBCA0E4BEEA73DF4D519-28\"}}",
    "code": 0,
    "log": "Msg 0: ",
    "confirmBlocks": 0,
    "memo": "",
    "source": 0,
    "hasChildren": 0
  }
]`

type testTxs struct {
	name        string
	apiResponse string
	expected    []blockatlas.Tx
	token       string
}

func TestNormalizeTxs(t *testing.T) {
	//testNormalizeTxs(t, &testTxs{
	//	name:        "all transfers",
	//	apiResponse: AllTransfersType,
	//	expected:    []blockatlas.Tx{transferDst, tokenTransferDst, newOrderTransferDst, cancelOrdeTransferDst},
	//	token:       "",
	//})
	testNormalizeTxs(t, &testTxs{
		name:        "bnb transfer",
		apiResponse: convertJsonToArray(transferTransaction),
		expected:    []blockatlas.Tx{transferDst},
		token:       "BNB",
	})
	testNormalizeTxs(t, &testTxs{
		name:        "native token transfer",
		apiResponse: convertJsonToArray(tokenTransferTransaction),
		expected:    []blockatlas.Tx{tokenTransferDst},
		token:       "YLC-D8B",
	})
	//testNormalizeTxs(t, &testTxs{
	//	name:        "new order transfer",
	//	apiResponse: convertJsonToArray(newOrderTransaction),
	//	expected:    []blockatlas.Tx{newOrderTransferDst},
	//	token:       "AWC-986",
	//})
	//testNormalizeTxs(t, &testTxs{
	//	name:        "cancel order transfer",
	//	apiResponse: convertJsonToArray(cancelOrderTransaction),
	//	expected:    []blockatlas.Tx{cancelOrdeTransferDst},
	//	token:       "GTO-908",
	//})
}

func convertJsonToArray(jsonString string) string {
	return "[" + jsonString + "]"
}

func testNormalizeTxs(t *testing.T, _test *testTxs) {
	t.Run(_test.name, func(t *testing.T) {
		var srcTxs []Tx
		err := json.Unmarshal([]byte(_test.apiResponse), &srcTxs)
		assert.Nil(t, err)
		txs := NormalizeTxs(srcTxs, _test.token)
		assert.Equal(t, _test.expected, txs, "transfer: tx don't equal")
	})
}

const myToken = `
{
	"free": "17199.38841739",
	"frozen": "0.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`

const myTokenAllZero = `
{
	"free": "0.00000000",
	"frozen": "0.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`

const myTokenFreeZero = `
{
	"free": "0.00000000",
	"frozen": "1.00000000",
	"locked": "0.00000000",
	"symbol": "ARN-71B"
}
`

const myTokenFrozenAndFreeZero = `
{
	"free": "0.00000000",
	"frozen": "0.00000000",
	"locked": "0.00000001",
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

var emptyTokenDst = blockatlas.Token{}

type testToken struct {
	name        string
	apiResponse string
	expected    blockatlas.Token
	tokens      string
	ok          bool
}

func TestNormalizeToken(t *testing.T) {
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: myToken,
		tokens:      tokenList,
		expected:    tokenDst,
		ok:          true,
	})
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: myTokenAllZero,
		tokens:      tokenList,
		expected:    emptyTokenDst,
		ok:          false,
	})
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: myTokenFreeZero,
		tokens:      tokenList,
		expected:    tokenDst,
		ok:          true,
	})
	testNormalizeToken(t, &testToken{
		name:        "token",
		apiResponse: myTokenFrozenAndFreeZero,
		tokens:      tokenList,
		expected:    tokenDst,
		ok:          true,
	})
}

func testNormalizeToken(t *testing.T, _test *testToken) {
	t.Run(_test.name, func(t *testing.T) {
		var srcToken Balance
		err := json.Unmarshal([]byte(_test.apiResponse), &srcToken)
		assert.Nil(t, err)

		var srcTokens TokenPage
		err = json.Unmarshal([]byte(_test.tokens), &srcTokens)
		assert.Nil(t, err)

		tk, ok := NormalizeToken(&srcToken, &srcTokens)
		assert.Equal(t, _test.ok, ok, "token: token could not be normalized")
		assert.Equal(t, _test.expected, tk, "token: token don't equal")
	})
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
