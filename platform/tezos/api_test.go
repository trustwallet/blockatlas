package tezos

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
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

const validatorSrc = `
[
	{"pkh":"tz2TSvNTh2epDMhZHrw73nV9piBX7kLZ9K9m","rolls":3726}
]
`

var transferDst = blockatlas.Tx{
	ID:    "oo3zTBHCkRkYDumt5t3rUyJ777wsr3dVMxYCU1FEV5xyftoih2Y",
	Coin:  coin.XTZ,
	From:  "tz1WCd2jm4uSt4vntk4vSuUWoZQGhLcDuR9q",
	To:    "tz1gcsKzDRzEkN6HNyngmoiGEuxojYrAJeC6",
	Fee:   "1420",
	Date:  1555102504,
	Block: 393070,
	Meta: blockatlas.Transfer{
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

	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, _ := json.Marshal(&transferDst)

	assert.Equal(t, resJSON, dstJSON)
}

func TestNormalizeValidator(t *testing.T) {
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	expected := blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Reward: blockatlas.StakingReward{Annual: Annual},
	}

	result := normalizeValidator(v)

	assert.Equal(t, result, expected)
}
