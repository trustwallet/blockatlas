package ethereum

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/ethereum/source"
)

const tokenTransferSrc = `
{
    "operations": [
        {
            "transactionId": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2-0",
            "contract": {
                "address": "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
                "symbol": "KBC",
                "decimals": 7,
                "totalSupply": "120000000000000000",
                "name": "KaratBank Coin"
            },
            "from": "0xd35f30d194684a391c63a6deced7d3dd5207c265",
            "to": "0xaa4d790076f1bf7511a0a0ac498c89e13e1efe17",
            "type": "token_transfer",
            "value": "4291000000",
            "coin": 60
        }
    ],
    "contract": null,
    "_id": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
    "blockNumber": 7491945,
    "timeStamp": "1554248437",
    "nonce": 88,
    "from": "0xd35f30d194684a391c63a6deced7d3dd5207c265",
    "to": "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
    "value": "0",
    "gas": "67497",
    "gasPrice": "6900000256",
    "gasUsed": "51921",
    "input": "0xa9059cbb000000000000000000000000aa4d790076f1bf7511a0a0ac498c89e13e1efe1700000000000000000000000000000000000000000000000000000000ffc376c0",
    "error": "",
    "id": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
    "coin": 60
}`

const contractCallSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	"blockNumber": 7522627,
	"timeStamp": "1554661737",
	"nonce": 534,
	"from": "0xc9a16a82c284efc3cb0fe8c891ab85d6eba0eefb",
	"to": "0xc67f9c909c4d185e4a5d21d642c27d05a145a76c",
	"value": "1800000000000000000",
	"gas": "1000000",
	"gasPrice": "2000000000",
	"gasUsed": "21340",
	"input": "0xfffdefefed",
	"error": "",
	"id": "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	"coin": 60
}
`

const transferSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	"blockNumber": 7522781,
	"timeStamp": "1554663642",
	"nonce": 88,
	"from": "0xf5aea47e57c058881b31ee8fce1002c409188f06",
	"to": "0x0ae933a89d9e249d0873cfc7ca022fcb3f1280ce",
	"value": "1999895000000000000",
	"gas": "21000",
	"gasPrice": "5000000000",
	"gasUsed": "21000",
	"input": "0x",
	"error": "",
	"id": "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	"coin": 60
}`

const failedSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	"blockNumber": 7522678,
	"timeStamp": "1554662399",
	"nonce": 1,
	"from": "0x4b55af7ce28a113d794f9a9940fe1506f37aa619",
	"to": "0xe65f787c8561a4b15771111bb427274dedfe85d7",
	"value": "59859820000000000",
	"gas": "21000",
	"gasPrice": "3000000000",
	"gasUsed": "21000",
	"input": "0x",
	"error": "Error",
	"id": "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	"coin": 60
}`

var tokenTransferDst = models.Tx{
	Id:   "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
	Coin: coin.ETH,
	From: "0xd35f30d194684a391c63a6deced7d3dd5207c265",
	To:   "0xaa4d790076f1bf7511a0a0ac498c89e13e1efe17",
	//To:   "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a", Contract
	Fee:    "67497",
	Date:   1554248437,
	Block:  7491945,
	Status: models.StatusCompleted,
	Meta: models.TokenTransfer{
		Name:     "KaratBank Coin",
		Symbol:   "KBC",
		TokenID:  "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
		Decimals: 7,
		Value:    "4291000000",
	},
}

var contractCallBaseDst = models.Tx{
	Id:     "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	Coin:   coin.ETH,
	From:   "0xc9a16a82c284efc3cb0fe8c891ab85d6eba0eefb",
	To:     "0xc67f9c909c4d185e4a5d21d642c27d05a145a76c",
	Fee:    "1000000",
	Date:   1554661737,
	Block:  7522627,
	Status: models.StatusCompleted,
}

var contractCallMeta1Dst = models.Transfer{
	Value: "1800000000000000000",
}

var contractCallMeta2Dst = models.ContractCall{
	Input: "0xfffdefefed",
	Value: "123",
}

var transferDst = models.Tx{
	Id:     "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	Coin:   coin.ETH,
	From:   "0xf5aea47e57c058881b31ee8fce1002c409188f06",
	To:     "0x0ae933a89d9e249d0873cfc7ca022fcb3f1280ce",
	Fee:    "21000",
	Date:   1554663642,
	Block:  7522781,
	Status: models.StatusCompleted,
	Meta: models.Transfer{
		Value: "1999895000000000000",
	},
}

var failedDst = models.Tx{
	Id:     "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	Coin:   coin.ETH,
	From:   "0x4b55af7ce28a113d794f9a9940fe1506f37aa619",
	To:     "0xe65f787c8561a4b15771111bb427274dedfe85d7",
	Fee:    "21000",
	Date:   1554662399,
	Block:  7522678,
	Status: models.StatusFailed,
	Error:  "Error",
	Meta: models.Transfer{
		Value: "59859820000000000",
	},
}

var contractCallDst []models.Tx

func init() {
	{
		// Transfer
		tx1 := contractCallBaseDst
		tx1.Meta = contractCallMeta1Dst
		// Contract Call
		tx2 := contractCallBaseDst
		tx2.Meta = contractCallMeta2Dst
		contractCallDst = []models.Tx{tx1, tx2}
	}
}

type test struct {
	name        string
	apiResponse string
	expected    []models.Tx
	token       bool
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    []models.Tx{transferDst},
	})
	testNormalize(t, &test{
		name:        "token transfer",
		apiResponse: tokenTransferSrc,
		expected:    []models.Tx{tokenTransferDst},
		token:       true,
	})
	testNormalize(t, &test{
		name:        "contract call",
		apiResponse: contractCallSrc,
		expected:    contractCallDst,
	})
	testNormalize(t, &test{
		name:        "failed transaction",
		apiResponse: failedSrc,
		expected:    []models.Tx{failedDst},
	})
}

func testNormalize(t *testing.T, _test *test) {
	var doc source.Doc
	err := json.Unmarshal([]byte(_test.apiResponse), &doc)
	if err != nil {
		t.Error(err)
		return
	}
	var res []models.Tx
	if _test.token {
		res = AppendTokenTxs(nil, &doc)
	} else {
		res = AppendTxs(nil, &doc)
	}

	resJson, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		t.Error(_test.name + ": tx don't equal")
	}
}
