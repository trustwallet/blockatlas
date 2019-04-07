package models

import (
	"bytes"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

var txJSON = []byte(`{
	"id": "14beb212aaefd06d7c6c0b25fc5ec242a2de2725af0a2827c105e743222cacd6",
	"from": "NQ11 P00L 2HYP TUK8 VY6L 2N22 MMBU MHHR BSAA",
	"to": "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	"fee": "138",
	"date": 1548954343,
	"type": "transfer",
	"metadata": {
		"name": "Nimiq",
		"symbol": "NIM",
		"decimals": 5,
		"value": "5004160"
	}
}`)

var txModel = Tx{
	Id:   "14beb212aaefd06d7c6c0b25fc5ec242a2de2725af0a2827c105e743222cacd6",
	From: "NQ11 P00L 2HYP TUK8 VY6L 2N22 MMBU MHHR BSAA",
	To:   "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	Fee:  "138",
	Date: 1548954343,
	Meta: &Transfer{
		Value:    "5004160",
	},
}

func TestTx_UnmarshalJSON(t *testing.T) {
	// Expect to get txModel, but with type set
	expected := txModel
	expected.Type = TxTransfer

	// Unmarshal source
	var got Tx
	err := json.Unmarshal(txJSON, &got)
	if err != nil {
		t.Fatal(err)
	}

	// Compare got and expected
	if !reflect.DeepEqual(expected, got) {
		spew.Println("Expected")
		spew.Dump(expected)
		spew.Println("Got")
		spew.Dump(got)
		t.Error("txs not equal")
	}
}

func TestTx_MarshalJSON(t *testing.T) {
	// Input is txModel without type
	input := txModel

	// Marshal transaction
	got, err := json.MarshalIndent(&input, "", "\t")
	if err != nil {
		t.Fatal(err)
	}

	// After marshal, the type should be set
	if input.Type == "" {
		t.Fatal("type was not set")
	} else if input.Type != TxTransfer {
		t.Fatal("wrong type set")
	}

	// Compare expected and output JSON
	bytes.Equal(got, txJSON)
}
