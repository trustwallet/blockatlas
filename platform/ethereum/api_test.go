package ethereum

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/ethereum/source"
	"testing"
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

var tokenTransferBaseDst = models.Tx{
	Id:     "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
	Coin:   coin.IndexETH,
	From:   "0xd35f30d194684a391c63a6deced7d3dd5207c265",
	To:     "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
	Fee:    "67497",
	Date:   1554248437,
	Block:  7491945,
	Status: models.StatusCompleted,
}

var tokenTransferMeta1Dst = models.TokenTransfer{
	Name:     "KaratBank Coin",
	Symbol:   "KBC",
	Contract: "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
	Decimals: 7,
	Value:    "4291000000",
}

var tokenTransferMeta2Dst = models.ContractCall{
	Input: "0xa9059cbb000000000000000000000000aa4d790076f1bf7511a0a0ac498c89e13e1efe1700000000000000000000000000000000000000000000000000000000ffc376c0",
}

var contractCallBaseDst = models.Tx{
	Id:     "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	Coin:   coin.IndexETH,
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
}

var failedBaseDst = models.Tx{
	Id: "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	Coin: coin.IndexETH,
	From: "0x4b55af7ce28a113d794f9a9940fe1506f37aa619",
	To: "0xe65f787c8561a4b15771111bb427274dedfe85d7",
	Fee: "21000",
	Date: 1554662399,
	Block: 7522678,
	Status: models.StatusFailed,
	Error: "Error",
}

var failedMeta1Dst = models.Transfer{
	Value: "59859820000000000",
}

var tokenTransferDst []models.Tx
var contractCallDst []models.Tx
var transferDst []models.Tx
var failedDst []models.Tx

func init() {
	{
		// ERC-20
		tx1 := tokenTransferBaseDst
		tx1.Meta = tokenTransferMeta1Dst
		tx1.To = "0xaa4d790076f1bf7511a0a0ac498c89e13e1efe17"
		// Contract Call
		tx2 := tokenTransferBaseDst
		tx2.Meta = tokenTransferMeta2Dst
		tokenTransferDst = []models.Tx{ tx2, tx1 }
	}
	{
		// Transfer
		tx1 := contractCallBaseDst
		tx1.Meta = contractCallMeta1Dst
		// Contract Call
		tx2 := contractCallBaseDst
		tx2.Meta = contractCallMeta2Dst
		contractCallDst = []models.Tx{ tx1, tx2 }
	}
}

func TestExtractTxs(t *testing.T) {
	testExtract(t, "token transfer", tokenTransferSrc, tokenTransferDst)
	testExtract(t, "contract call", contractCallSrc, contractCallDst)
}

func testExtract(t *testing.T, test string, src string, dst []models.Tx) {
	var doc source.Doc
	err := json.Unmarshal([]byte(src), &doc)
	if err != nil {
		t.Error(err)
		return
	}
	res := ExtractTxs(nil, &doc)

	resJson, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(dst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		t.Error(test + ": tx don't equal")
		println("---- EXPECTED ----")
		spew.Dump(dst)
		println("---- GOT ----")
		spew.Dump(res)
	}
}
