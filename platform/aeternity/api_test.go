package aeternity

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const transferTransaction = `
{
	"block_hash": "mh_sJqfsWuuhA7vXDJLYFVtpagCSTmfmhzdqKWFR4pU5LK4D8W8T",
    "block_height": 113579,
    "hash": "th_oJfBC6KZKaKsL4WXTq1ZtFiSE8Wp2PQYEnwyZqtudyHcU3Qg6",
    "signatures": [
      "sg_F3Ecfu5g6FcPyHrgZue96hVHthnXW7CbuDUEoKwWqWvbE84xb3ifB57AGTaH1WzDr4x1cnv4biLqTorjq9ZqhzCFVJC5c"
    ],
    "time": 1563848658206,
    "tx": {
      "amount": 252550000000000000000,
      "fee": 20500000000000,
      "nonce": 251291,
      "payload": "ba_SGVsbG8sIE1pbmVyISAvWW91cnMgQmVlcG9vbC4vKXcQag==",
      "recipient_id": "ak_ZWrS6xGhzxBasKmMbVSACfRioWqPyM5jNqMpBQ5ngP75RS6pS",
      "sender_id": "ak_nv5B93FPzRHrGNmMdTDfGdd5xGZvep3MVSpJqzcQmMp59bBCv",
      "type": "SpendTx",
      "version": 1
    }
  }
`

var transferDst = blockatlas.Tx{
	ID:     "th_oJfBC6KZKaKsL4WXTq1ZtFiSE8Wp2PQYEnwyZqtudyHcU3Qg6",
	Coin:   coin.AE,
	From:   "ak_nv5B93FPzRHrGNmMdTDfGdd5xGZvep3MVSpJqzcQmMp59bBCv",
	To:     "ak_ZWrS6xGhzxBasKmMbVSACfRioWqPyM5jNqMpBQ5ngP75RS6pS",
	Fee:    "20500000000000",
	Date:   1563848658,
	Block:  113579,
	Status: blockatlas.StatusCompleted,
	Memo:   "",
	Meta: blockatlas.Transfer{
		Value: "252550000000000000000",
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
	token       string
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferTransaction,
		expected:    &transferDst,
		token:       "",
	})
}

func testNormalize(t *testing.T, _test *test) {
	var srcTx Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx := NormalizeTx(&srcTx)

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
