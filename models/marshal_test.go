package models

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"reflect"
	"testing"
)

func TestTx_UnmarshalJSON(t *testing.T) {
	source := `{
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
    }`
	expected := Tx{
		Id:   "14beb212aaefd06d7c6c0b25fc5ec242a2de2725af0a2827c105e743222cacd6",
		From: "NQ11 P00L 2HYP TUK8 VY6L 2N22 MMBU MHHR BSAA",
		To:   "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
		Fee:  "138",
		Date: 1548954343,
		Type: TxTransfer,
		Meta: &Transfer{
			Name:     "Nimiq",
			Symbol:   "NIM",
			Decimals: 5,
			Value:    "5004160",
		},
	}
	var got Tx
	err := json.Unmarshal([]byte(source), &got)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(expected, got) {
		spew.Println("Expected")
		spew.Dump(expected)
		spew.Println("Got")
		spew.Dump(got)
		t.Error("txs not equal")
	}
}
