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

var tokenTransferDst []models.Tx

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
}

func TestExtractTxs(t *testing.T) {
	testExtractTokenTransfer(t)
}

func testExtractTokenTransfer(t *testing.T) {
	var doc source.Doc
	err := json.Unmarshal([]byte(tokenTransferSrc), &doc)
	if err != nil {
		t.Error(err)
		return
	}
	res := ExtractTxs(nil, &doc)

	resJson, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(tokenTransferDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		t.Error("tx don't equal")
		spew.Dump(res)
	}
}
