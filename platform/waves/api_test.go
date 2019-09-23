package waves

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
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
	Coin:   5741564,
	From:   "3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
	To:     "3PKWyVAmHom1sevggiXVfbGUc3kS85qT4Va",
	Fee:    "100000",
	Date:   1561048131,
	Block:  1580410,
	Status: blockatlas.StatusCompleted,
	Memo:   "",
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("9481600000"),
		Symbol:   "WAVES",
		Decimals: 8,
	},
}

var differentTxsObj = blockatlas.Tx{
	ID:     "52GG9U2e6foYRKp5vAzsTQ86aDAABfRJ7synz7ohBp19",
	Coin:   5741564,
	From:   "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	To:     "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
	Fee:    "100000",
	Date:   1479313236,
	Block:  7782,
	Memo:   "string",
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("100000"),
		Symbol:   "WAVES",
		Decimals: 8,
	},
}

type txParseTest struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

type txFilterTest struct {
	name        string
	apiResponse string
	expected    blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testParseTx(t, &txParseTest{
		name:        "transfer",
		apiResponse: transferV1,
		expected:    &transferV1Obj,
	})
	testFilterTxs(t, &txFilterTest{
		name:        "filter transfer transactions txParseTest",
		apiResponse: differentTxs,
		expected:    differentTxsObj,
	})
}

func testParseTx(t *testing.T, _test *txParseTest) {
	var tx Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &tx)
	if err != nil {
		t.Error(err)
		return
	}

	res, _ := NormalizeTx(&tx)

	resJSON, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&_test.expected)
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
	var res blockatlas.Tx
	for _, tx := range txs[0] {
		if tx.Type == 4 {
			n, ok := NormalizeTx(&tx)
			if ok {
				res = n
			}
		}
	}

	resJSON, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error(_test.name + ": txs don't equal")
	}
}
