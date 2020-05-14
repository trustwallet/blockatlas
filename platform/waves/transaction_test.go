package waves

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferV1 = `
{
	"type":4,
	"id":"7QoQc9qMUBCfY4QV35mgBsT8eTXybvGkM2HTumtAvBUL",
	"sender":"3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
	"senderPublicKey":"Ao159h5j1piHBhoEbCAYyaiKNd6uoKvcdwzRZF9za3Vv",
	"fee":100000,
	"timestamp":1561048131740,
	"signature":"4WjDwn5t34PLHzgH1NfA4DYdt4PdTbGQDjDdxwKrp82QTQSHFRrgSJXWU2FTYe82afvgUDhnipSKxaiGzMWWo2HW",
	"proofs":["4WjDwn5t34PLHzgH1NfA4DYdt4PdTbGQDjDdxwKrp82QTQSHFRrgSJXWU2FTYe82afvgUDhnipSKxaiGzMWWo2HW"],
	"version":1,
	"recipient":"3PKWyVAmHom1sevggiXVfbGUc3kS85qT4Va",
	"assetId":null,
	"feeAssetId":null,
	"feeAsset":null,
	"amount":9481600000,
	"attachment":"",
	"height":1580410
}`

var transferV1Obj = blockatlas.Tx{
	ID:     "7QoQc9qMUBCfY4QV35mgBsT8eTXybvGkM2HTumtAvBUL",
	Coin:   5741564,
	From:   "3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
	To:     "3PKWyVAmHom1sevggiXVfbGUc3kS85qT4Va",
	Fee:    "100000",
	Date:   1561048131,
	Block:  1580410,
	Status: blockatlas.StatusCompleted,
	Memo:   "",
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("9481600000"),
		Symbol:   "WAVES",
		Decimals: 8,
	},
}

func TestNormalizeTxs(t *testing.T) {
	var (
		tests = []struct {
			name        string
			apiResponse string
			expected    []blockatlas.Tx
		}{
			{"transfer", transferV1, []blockatlas.Tx{transferV1Obj}},
		}
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tx Transaction
			err := json.Unmarshal([]byte(tt.apiResponse), &tx)
			if err != nil {
				t.Error(err)
				return
			}
			res := NormalizeTxs([]Transaction{tx})

			resJSON, err := json.Marshal(&res)
			if err != nil {
				t.Fatal(err)
			}

			dstJSON, err := json.Marshal(&tt.expected)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(resJSON, dstJSON) {
				println(string(resJSON))
				println(string(dstJSON))
				t.Error(tt.name + ": tx don't equal")
			}
		})
	}

}

