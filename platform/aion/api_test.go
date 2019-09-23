package aion

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `
{
	"blockHash": "68364cfa1873c42f3c2ef659349ca101c4c691a0385fd1c1677f92a96f7332ca",
	"nrgPrice": 10000000000,
	"toAddr": "a09b8c4c40bd7a81e969b8f6f291074206196a99948b03c6a469892931a3c258",
	"contractAddr": "",
	"data": "",
	"year": 2019,
	"transactionIndex": 7,
	"nonce": "170c",
	"transactionHash": "af3c2f5087fc3332154dc9d11c27e312f30ff829dbc5436aec8cc4342c7dc384",
	"transactionTimestamp": 1554862205533375,
	"nrgConsumed": 21000,
	"month": 4,
	"blockNumber": 2880919,
	"blockTimestamp": 1554862228,
	"transactionLog": "[]",
	"fromAddr": "a07981da70ce919e1db5f051c3c386eb526e6ce8b9e2bfd56e3f3d754b0a17f3",
	"day": 10,
	"value": 11.903810405853733,
	"txError": ""
}`

var transferDst = blockatlas.Tx{
	ID:     "af3c2f5087fc3332154dc9d11c27e312f30ff829dbc5436aec8cc4342c7dc384",
	Coin:   coin.AION,
	From:   "0xa07981da70ce919e1db5f051c3c386eb526e6ce8b9e2bfd56e3f3d754b0a17f3",
	To:     "0xa09b8c4c40bd7a81e969b8f6f291074206196a99948b03c6a469892931a3c258",
	Fee:    "21000",
	Date:   1554862228,
	Block:  2880919,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "11903810405853733000",
		Symbol:   "AION",
		Decimals: 18,
	},
}

func TestNormalize(t *testing.T) {
	var srcTx Tx
	err := json.Unmarshal([]byte(transferSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	resTx, ok := NormalizeTx(&srcTx)
	if !ok {
		t.Fatal("Can't normalize transaction")
	}

	resJSON, err := json.Marshal(&resTx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("tx don't equal")
	}
}
