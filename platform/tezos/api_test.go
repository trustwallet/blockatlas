package tezos

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"testing"
)

const transferSrc = `
{
	"hash": "oo3zTBHCkRkYDumt5t3rUyJ777wsr3dVMxYCU1FEV5xyftoih2Y",
	"block_hash": "BLczuBWhHKEwKCEPft9c7SfdZZ9oCxwqhcukKiwCNjfjfKeZPvU",
	"network_hash": "NetXdQprcVkpaWU",
	"type": {
		"kind": "manager",
		"source": {
			"tz": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q"
		},
		"operations": [
			{
				"kind": "transaction",
				"src": {
					"tz": "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q"
				},
				"amount": "2597000000",
				"destination": {
					"tz": "tz1gcsKzDRzEkN6HNyngmoiGEuxojYrAJeC6"
				},
				"failed": false,
				"internal": false,
				"burn": 0,
				"counter": 11080,
				"fee": 1420,
				"gas_limit": "10300",
				"storage_limit": "300",
				"op_level": 393070,
				"timestamp": "2019-04-12T20:55:04Z"
			}
		]
	}
}
`

var transferDst = models.Tx{
	Id:    "oo3zTBHCkRkYDumt5t3rUyJ777wsr3dVMxYCU1FEV5xyftoih2Y",
	Coin:  coin.XTZ,
	From:  "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
	To:    "tz1gcsKzDRzEkN6HNyngmoiGEuxojYrAJeC6",
	Fee:   "1420",
	Date:  1555102504,
	Block: 393070,
	Meta: models.Transfer{
		Value: "2597000000",
	},
}

func TestNormalize(t *testing.T) {
	var srcTx Tx
	err := json.Unmarshal([]byte(transferSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx, ok := Normalize(&srcTx)
	if !ok {
		t.Errorf("transfer: tx could not be normalized")
	}

	resJson, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		println(string(resJson))
		println(string(dstJson))
		t.Error("basic: tx don't equal")
	}
}
