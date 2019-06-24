package waves

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const transferV1 = `
{
	"type":4,
	"id":"7QoQc9qMUBCfY4QV35mgBsT8eTXybvGkM2HTumtAvBUL",
	"sender":"3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
	"senderPublicKey":"Ao159h5j1piHBhoEbCAYyaiKNd6uoKvcdwzRZF9za3Vv",
	"fee":100000,
	"timestamp":1561048131740,
	"signature":"4WjDwn5t34PLHzgH1NfA4DYdt4PdTbGQDjDdxwKrp82QTQSHFRrgSJXWU2FTYe82afvgUDhnipSKxaiGzMWWo2HW",
	"proofs":["4WjDwn5t34PLHzgH1NfA4DYdt4PdTbGQDjDdxwKrp82QTQSHFRrgSJXWU2FTYe82afvgUDhnipSKxaiGzMWWo2HW"],
	"version":1,
	"recipient":"3PKWyVAmHom1sevggiXVfbGUc3kS85qT4Va",
	"assetId":null,
	"feeAssetId":null,
	"feeAsset":null,
	"amount":9481600000,
	"attachment":"",
	"height":1580410
}`

// btc asset field was added, signature is invalid but this's ok for us
const transferV2 = `
{
	"type":4,
	"id":"EZxdiLRKvjouSGy4NW8sqfbiEYXgcVvGiAsa3q4pdFcj",
	"sender":"3P4w1A96SojL9VnJMAaCenNySEL11NDvXK8",
	"senderPublicKey":"6yHwbWHmGQEeBXSvEhpyxhJNXWFXroQDQFqCxVzeKzk1",
	"fee":100000,
    "timestamp":1560414208156,
	"proofs":["2drH4pV4LsoNESm6Ec3X9eqJYWetSaRr4mE7P4vmVMETKRJYwN5KidhGGWddfjiRkj1Pxb7nZVBAXkiu7GwZi5mS"],
	"version":2,
	"recipient":"3P6cRBnTfnvNWsYKdznDPCTRMZhjJHJxJxs",
	"assetId":"8LQW8f7P5d5PZM7GtZEBgaqRPGSzS3DfPuiXrURJ4AJS",
	"feeAssetId":null,
	"feeAsset":null,
	"amount":675133100,
	"attachment":"",
	"height":1569685
}
`

const differentTxs = `
[[
	{
	 "type": 10,
	 "timestamp": 1516171819000,
	 "sender": "3MtrNP7AkTRuBhX4CBti6iT21pQpEnmHtyw",
	 "fee": 100000,
	 "alias": "ALIAS"
	},
	{
	 "type":10,
	 "id":"9q7X84wFuVvKqRdDQeWbtBmpsHt9SXFbvPPtUuKBVxxr",
	 "sender":"3MtrNP7AkTRuBhX4CBti6iT21pQpEnmHtyw",
	 "senderPublicKey":"G6h72icCSjdW2A89QWDb37hyXJoYKq3XuCUJY2joS3EU",
	 "fee":100000000,
	 "timestamp":46305781705234713,
	 "signature":"4gQyPXzJFEzMbsCd9u5n3B2WauEc4172ssyrXCL882oNa8NfNihnpKianHXrHWnZs1RzDLbQ9rcRYnSqxKWfEPJG",
	 "alias":"dajzmj6gfuzmbfnhamsbuxivc"
	},
	{
	  "type": 4,
	  "id": "52GG9U2e6foYRKp5vAzsTQ86aDAABfRJ7synz7ohBp19",
	  "sender": "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	  "senderPublicKey": "CRxqEuxhdZBEHX42MU4FfyJxuHmbDBTaHMhM3Uki7pLw",
	  "recipient": "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	  "assetId": null,
	  "amount": 100000,
	  "feeAsset": null,
	  "fee": 100000,
	  "timestamp": 1479313236091,
	  "attachment": "string",
	  "signature": "GknccUA79dBcwWgKjqB7vYHcnsj7caYETfncJhRkkaetbQon7DxbpMmvK9LYqUkirJp17geBJCRTNkHEoAjtsUm",
	  "height": 7782
	},
	{
	  "type": 2,
	  "id": "4XE4M9eSoVWVdHwDYXqZsXhEc4q8PH9mDMUBegCSBBVHJyP2Yb1ZoGi59c1Qzq2TowLmymLNkFQjWp95CdddnyBW",
	  "fee": 100000,
	  "timestamp": 1479313097422,
	  "signature": "4XE4M9eSoVWVdHwDYXqZsXhEc4q8PH9mDMUBegCSBBVHJyP2Yb1ZoGi59c1Qzq2TowLmymLNkFQjWp95CdddnyBW",
	  "sender": "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	  "senderPublicKey": "CRxqEuxhdZBEHX42MU4FfyJxuHmbDBTaHMhM3Uki7pLw",
	  "recipient": "3N9iRMou3pgmyPbFZn5QZQvBTQBkL2fR6R1",
	  "amount": 1000000000
	}
]]`

var transferV1Obj = blockatlas.Tx{
	ID:     "7QoQc9qMUBCfY4QV35mgBsT8eTXybvGkM2HTumtAvBUL",
	Coin:   coin.WAVES,
	From:   "3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
	To:     "3PKWyVAmHom1sevggiXVfbGUc3kS85qT4Va",
	Fee:    "100000",
	Date:   1561048131740,
	Block:  1580410,
	Status: blockatlas.StatusCompleted,
	Memo:   "",
	Meta: blockatlas.Transfer{
		Value: blockatlas.Amount("9481600000"),
	},
}

var transferV2Obj = blockatlas.Tx{
	ID:     "EZxdiLRKvjouSGy4NW8sqfbiEYXgcVvGiAsa3q4pdFcj",
	Coin:   coin.WAVES,
	From:   "3P4w1A96SojL9VnJMAaCenNySEL11NDvXK8",
	To:     "3P6cRBnTfnvNWsYKdznDPCTRMZhjJHJxJxs",
	Fee:    "100000",
	Date:   1560414208156,
	Block:  1569685,
	Status: blockatlas.StatusCompleted,
	Memo:   "",
	Meta: blockatlas.NativeTokenTransfer{
		Name:     "Bitcoin Token",
		Symbol:   "WBTC",
		TokenID:  "8LQW8f7P5d5PZM7GtZEBgaqRPGSzS3DfPuiXrURJ4AJS",
		Decimals: 8,
		Value:    blockatlas.Amount("675133100"),
		From:     "3P4w1A96SojL9VnJMAaCenNySEL11NDvXK8",
		To:       "3P6cRBnTfnvNWsYKdznDPCTRMZhjJHJxJxs",
	},
}

var differentTxsObjs = []blockatlas.Tx{{
	ID:     "52GG9U2e6foYRKp5vAzsTQ86aDAABfRJ7synz7ohBp19",
	Coin:   coin.WAVES,
	From:   "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	To:     "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	Fee:    "100000",
	Date:   1479313236091,
	Block:  7782,
	Memo:   "string",
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value: blockatlas.Amount("100000"),
	},
}}

type txParseTest struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
	tokenInfo   *TokenInfo
}

type txFilterTest struct {
	name        string
	apiResponse string
	expected    []blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testParseTx(t, &txParseTest{
		name:        "transfer",
		apiResponse: transferV1,
		expected:    &transferV1Obj,
	})
	testParseTx(t, &txParseTest{
		name:        "token transfer",
		apiResponse: transferV2,
		expected:    &transferV2Obj,
		tokenInfo: &TokenInfo{
			Name:        "WBTC",
			Description: "Bitcoin Token",
			Decimals:    8,
		},
	})
	testFilterTxs(t, &txFilterTest{
		name:        "filter transfer transactions txParseTest",
		apiResponse: differentTxs,
		expected:    differentTxsObjs,
	})
}

func testParseTx(t *testing.T, _test *txParseTest) {
	var tx Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &tx)
	if err != nil {
		t.Error(err)
		return
	}
	tx.Asset = _test.tokenInfo
	res := AppendTxs(nil, &tx, coin.WAVES)

	resJSON, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal([]blockatlas.Tx{*_test.expected})
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error(_test.name + ": tx don't equal")
	}
}

func testFilterTxs(t *testing.T, _test *txFilterTest) {
	var txs [][]Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &txs)
	if err != nil {
		t.Error(err)
		return
	}
	var res []blockatlas.Tx
	for _, tx := range txs[0] {
		if tx.Type == 4 {
			res = AppendTxs(nil, &tx, coin.WAVES)
		}
	}

	resJSON, err := json.Marshal(res)
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
		t.Error(_test.name + ": txs don't equal")
	}
}
