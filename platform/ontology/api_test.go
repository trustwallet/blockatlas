package ontology

import (
	"bytes"
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var srcOntTransfer = `
{
	"TxnType": 209,
	"ConfirmFlag": 1,
	"Fee": "0.010000000",
	"BlockIndex": 2,
	"TransferList": [
		{
			"FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
			"Amount": "2.000000000",
			"ToAddress": "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
			"AssetName": "ont"
		}
	],
	"TxnTime": 1556952450,
	"TxnHash": "4804e1be63ebe1715d6b4a039cc9d84b86cde74c8a8c8411578e6dcadc1e5405",
	"Height": 3411115
}
`

var dstOntTransfer = blockatlas.Tx{
	ID:     "4804e1be63ebe1715d6b4a039cc9d84b86cde74c8a8c8411578e6dcadc1e5405",
	Coin:   coin.ONT,
	From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	To:     "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
	Fee:    "10000000",
	Date:   1556952450,
	Type:   "transfer",
	Status: blockatlas.StatusCompleted,
	Block:  3411115,
	Meta: blockatlas.Transfer{
		Value:    "2",
		Symbol:   "ONT",
		Decimals: 0,
	},
}

var srcOngTransfer1 = `
{
	"TxnType": 209,
	"ConfirmFlag": 1,
	"Fee": "0.010000000",
	"BlockIndex": 2,
	"TransferList": [
		{
			"FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
			"Amount": "0.010000000",
			"ToAddress": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
			"AssetName": "ong"
		}
	],
	"TxnTime": 1555341286,
	"TxnHash": "f494d7aeae2b88d313465d1f5f588b213795b38988a0bef182d9d3c2012f6e6e",
	"Height": 2863855
}
`

var dstOngTransfer1 = blockatlas.Tx{
	ID:     "f494d7aeae2b88d313465d1f5f588b213795b38988a0bef182d9d3c2012f6e6e",
	Coin:   coin.ONT,
	From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	To:     "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
	Fee:    "10000000",
	Date:   1555341286,
	Type:   blockatlas.TxNativeTokenTransfer,
	Status: blockatlas.StatusCompleted,
	Block:  2863855,
	Meta: blockatlas.NativeTokenTransfer{
		Name:     "Ontology Gas",
		Symbol:   "ONG",
		TokenID:  "ong",
		Decimals: 9,
		Value:    "0",
		From:     "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
		To:       "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
	},
}

var srcOngTransfer2 = `
{
	"TxnType": 209,
	"ConfirmFlag": 1,
	"Fee": "0.010000000",
	"BlockIndex": 1,
	"TransferList": [
		{
			"FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
			"Amount": "10.000000000",
			"ToAddress": "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
			"AssetName": "ong"
		},
		{
			"FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
			"Amount": "0.010000000",
			"ToAddress": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
			"AssetName": "ong"
		}
	],
	"TxnTime": 1556952520,
	"TxnHash": "eccbfd040925a22884d87e73f818f30ab42d06046460b86e9a042a1e9cba7561",
	"Height": 3411141
}
`
var dstOngTransfer = blockatlas.Tx{
	ID:     "eccbfd040925a22884d87e73f818f30ab42d06046460b86e9a042a1e9cba7561",
	Coin:   coin.ONT,
	From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
	To:     "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
	Fee:    "10000000",
	Date:   1556952520,
	Type:   blockatlas.TxNativeTokenTransfer,
	Status: blockatlas.StatusCompleted,
	Block:  3411141,
	Meta: blockatlas.NativeTokenTransfer{
		Name:     "Ontology Gas",
		Symbol:   "ONG",
		TokenID:  "ong",
		Decimals: 9,
		Value:    "10000000000",
		From:     "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
		To:       "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
	},
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		Transaction string
		AssetName   string
		Expected    blockatlas.Tx
	}{
		{srcOntTransfer, ONTAssetName, dstOntTransfer},
		{srcOngTransfer1, ONGAssetName, dstOngTransfer1},
		{srcOngTransfer2, ONGAssetName, dstOngTransfer},
	}

	for _, test := range tests {
		var (
			sourceTx Tx
			ok       bool
		)

		tErr := json.Unmarshal([]byte(test.Transaction), &sourceTx)
		if tErr != nil {
			t.Fatal("Ontology: Can't unmarshal transaction", tErr)
		}

		tx, ok := Normalize(sourceTx, test.AssetName)

		if !ok {
			t.Fatal("Ontology: Can't normalize transaction")
		}

		actual, err := json.Marshal(tx)
		if err != nil {
			t.Fatal(err)
		}

		expected, err := json.Marshal(&test.Expected)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}

var (
	ontBlockResult = BlockResult{
		Error: 0,
		Result: Block{
			Height: 7707834,
			TxnList: []Tx{
				{
					TxnHash:     "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb",
					ConfirmFlag: 1,
					TxnTime:     1580481541,
					Height:      7707834,
				},
				{
					TxnHash:     "2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd",
					ConfirmFlag: 1,
					TxnTime:     1580481541,
					Height:      7707834,
				},
				{
					TxnHash:     "40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3",
					ConfirmFlag: 1,
					TxnTime:     1580481541,
					Height:      7707834,
				},
			},
			Hash: "a5f3ee1a102df7196bb1e262a05435f260392fae6be676ae2c0a6147f8ecf94c",
		},
	}
)

var (
	ontTxResp1 = TxResponse{
		Code: 0,
		Msg:  "SUCCESS",
		Result: TxV2{
			Hash:        "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb",
			Type:        209,
			Time:        1580481541,
			BlockHeight: 7707834,
			Fee:         "0.01",
			Description: "transfer",
			BlockIndex:  2,
			ConfirmFlag: 1,
			EventType:   3,
			Details: TransactionDetails{
				Transfers: []TransferDetails{
					{
						Amount:      "51000",
						AssetName:   "ong",
						FromAddress: "AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE",
						ToAddress:   "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
					},
					{
						Amount:      "0.01",
						AssetName:   "ong",
						FromAddress: "ANKXUWXy6XrhQvqbjhKJPH9AnLa2CuEMRK",
						ToAddress:   "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
					},
				},
			},
		},
	}

	ontTxResp2 = TxResponse{
		Code: 0,
		Msg:  "SUCCESS",
		Result: TxV2{
			Hash:        "2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd",
			Type:        209,
			Time:        1580481541,
			BlockHeight: 7707834,
			Fee:         "0.01",
			Description: "transfer",
			BlockIndex:  3,
			ConfirmFlag: 1,
			EventType:   3,
			Details: TransactionDetails{
				Transfers: []TransferDetails{
					{
						Amount:      "113.2",
						AssetName:   "ong",
						FromAddress: "ANdrA47zDXUu8MCkMdD3FYPmpSNGYeAvKz",
						ToAddress:   "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
					},
					{
						Amount:      "0.01",
						AssetName:   "ong",
						FromAddress: "ANKXUWXy6XrhQvqbjhKJPH9AnLa2CuEMRK",
						ToAddress:   "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
					},
				},
			},
		},
	}

	ontTxResp3 = TxResponse{
		Code: 0,
		Msg:  "SUCCESS",
		Result: TxV2{
			Hash:        "40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3",
			Type:        209,
			Time:        1580481541,
			BlockHeight: 7707834,
			Fee:         "0.01",
			Description: "transfer",
			BlockIndex:  1,
			ConfirmFlag: 1,
			EventType:   3,
			Details: TransactionDetails{
				Transfers: []TransferDetails{
					{
						Amount:      "10949",
						AssetName:   "ong",
						FromAddress: "Abg2gs6pfpQu82jXbm8EYGiipRBvf9ktVS",
						ToAddress:   "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
					}, {
						Amount:      "0.01",
						AssetName:   "ong",
						FromAddress: "ANKXUWXy6XrhQvqbjhKJPH9AnLa2CuEMRK",
						ToAddress:   "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
					},
				},
			},
		},
	}
)

func TestNormalizeBlock(t *testing.T) {
	block := normalizeBlock(ontBlockResult, []TxV2{ontTxResp1.Result, ontTxResp2.Result, ontTxResp3.Result})
	got, err := json.Marshal(block)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"number":7707834,"id":"a5f3ee1a102df7196bb1e262a05435f260392fae6be676ae2c0a6147f8ecf94c","txs":[{"id":"266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb","coin":1024,"from":"AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"native_token_transfer","memo":"","metadata":{"name":"Ontology Gas","symbol":"ONG","token_id":"ong","decimals":9,"value":"51000000000000","from":"AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx"}},{"id":"2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd","coin":1024,"from":"ANdrA47zDXUu8MCkMdD3FYPmpSNGYeAvKz","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"native_token_transfer","memo":"","metadata":{"name":"Ontology Gas","symbol":"ONG","token_id":"ong","decimals":9,"value":"113200000000","from":"ANdrA47zDXUu8MCkMdD3FYPmpSNGYeAvKz","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx"}},{"id":"40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3","coin":1024,"from":"Abg2gs6pfpQu82jXbm8EYGiipRBvf9ktVS","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"native_token_transfer","memo":"","metadata":{"name":"Ontology Gas","symbol":"ONG","token_id":"ong","decimals":9,"value":"10949000000000","from":"Abg2gs6pfpQu82jXbm8EYGiipRBvf9ktVS","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx"}}]}`

	assert.Equal(t, string(got), want)
}
