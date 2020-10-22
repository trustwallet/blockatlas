package harmony

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"testing"
)

const transferSrc = `
{
	"blockHash": "0x0bde901cd3599aa082482777fd0a7fed3f02b7b5a9096b7ea7b2fcb8addaa05d",
	"blockNumber": "0x12",
	"from": "one103q7qe5t2505lypvltkqtddaef5tzfxwsse4z7",
	"gas": "0x5208",
	"gasPrice": "0x3b9aca00",
	"hash": "0x230798fe22abff459b004675bf827a4089326a296fa4165d0c2ad27688e03e0c",
	"input": "0x",
	"nonce": "0x0",
	"to": "one129r9pj3sk0re76f7zs3qz92rggmdgjhtwge62k",
	"transactionIndex": "0x1",
	"value": "0x16345785d8a0000",
	"shardID": 0,
	"toShardID": 0,
	"v": "0x27",
	"r": "0x57766aa1304e97f8b71a9fa54a61b61ce8ef9ad177fcb337dd81827aad184327",
	"s": "0x3b3e5767899e8af5e75d62243a725371f08705b91e2305459e6fd8e8d2646651",
	"timestamp": "0x5DF5234E"
}
`

var transferDst = blockatlas.Tx{
	ID:     "0x230798fe22abff459b004675bf827a4089326a296fa4165d0c2ad27688e03e0c",
	Coin:   coin.ONE,
	From:   "one103q7qe5t2505lypvltkqtddaef5tzfxwsse4z7",
	To:     "one129r9pj3sk0re76f7zs3qz92rggmdgjhtwge62k",
	Fee:    "21000000000000",
	Date:   1576346446,
	Block:  18,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "100000000000000000",
		Symbol:   "ONE",
		Decimals: 18,
	},
}

func TestNormalize(t *testing.T) {
	var srcTx Transaction
	err := json.Unmarshal([]byte(transferSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	resTx, isGood, err := NormalizeTx(&srcTx)

	if !isGood || err != nil {
		t.Fatal()
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
