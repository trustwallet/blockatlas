package blockatlas

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"reflect"
	"sort"
	"testing"
)

var txJSON = []byte(`{
	"id": "14beb212aaefd06d7c6c0b25fc5ec242a2de2725af0a2827c105e743222cacd6",
	"coin": 242,
	"from": "NQ11 P00L 2HYP TUK8 VY6L 2N22 MMBU MHHR BSAA",
	"to": "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	"fee": "138",
	"date": 1548954343,
	"block": 419040,
	"status": "completed",
	"type": "transfer",
	"metadata": {
		"value": "5004160"
	}
}`)

var txModel = Tx{
	ID:     "14beb212aaefd06d7c6c0b25fc5ec242a2de2725af0a2827c105e743222cacd6",
	Coin:   coin.NIM,
	From:   "NQ11 P00L 2HYP TUK8 VY6L 2N22 MMBU MHHR BSAA",
	To:     "NQ86 2H8F YGU5 RM77 QSN9 LYLH C56A CYYR 0MLA",
	Fee:    "138",
	Date:   1548954343,
	Block:  419040,
	Status: StatusCompleted,
	Meta: &Transfer{
		Value: "5004160",
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

func TestSortTxPage(t *testing.T) {
	tests := []struct {
		name string
		page TxPage
		want TxPage
	}{
		{"test sort 1", TxPage{{Date: 5}, {Date: 2}, {Date: 1}, {Date: 4}, {Date: 3}}, TxPage{{Date: 5}, {Date: 4}, {Date: 3}, {Date: 2}, {Date: 1}}},
		{"test sort 2", TxPage{{Date: 100}, {Date: 2}, {Date: 33}, {Date: 409}}, TxPage{{Date: 409}, {Date: 100}, {Date: 33}, {Date: 2}}},
		{"test sort 3", TxPage{{Date: 100}, {Date: 2}, {Date: 100}}, TxPage{{Date: 100}, {Date: 100}, {Date: 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Sort(tt.page)
			assert.Equal(t, tt.want, tt.page)
		})
	}
}
