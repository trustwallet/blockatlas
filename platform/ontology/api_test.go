package ontology

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const (
	srcOntTransfer = `
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
}`

	srcOngTransfer = `
{
	"TxnType": 209,
	"ConfirmFlag": 1,
	"Fee": "0.010000000",
	"BlockIndex": 2,
	"TransferList": [
		{
			"FromAddress": "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
			"Amount": "2.1455",
			"ToAddress": "ASwdosWb2wH8Y3HYCYJpUJWWD3joyvvYGN",
			"AssetName": "ong"
		}
	],
	"TxnTime": 1555341286,
	"TxnHash": "a483d1d854e47a20692f472d72ff45b9a2bfc542f84dceb3171a48f68ba322cb",
	"Height": 2863855
}`
	srcRewardTransfer = `
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
}`
	srcFeeTransfer = `
{
	"TxnType": 209,
	"ConfirmFlag": 1,
	"Fee": "0.010000000",
	"BlockIndex": 1,
	"TransferList": [
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
}`
)

var (
	dstOntTransfer = blockatlas.Tx{
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
	dstOngTransfer = blockatlas.Tx{
		ID:     "a483d1d854e47a20692f472d72ff45b9a2bfc542f84dceb3171a48f68ba322cb",
		Coin:   coin.ONT,
		From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
		To:     "ASwdosWb2wH8Y3HYCYJpUJWWD3joyvvYGN",
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
			Value:    "2145500000",
			From:     "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
			To:       "ASwdosWb2wH8Y3HYCYJpUJWWD3joyvvYGN",
		},
	}
	dstRewardTransfer = blockatlas.Tx{
		ID:     "eccbfd040925a22884d87e73f818f30ab42d06046460b86e9a042a1e9cba7561",
		Coin:   coin.ONT,
		From:   "AUyL4TZ1zFEcSKDJrjFnD7vsq5iFZMZqT7",
		To:     "AQ9kzzHNLCcyrPwJuVMrSPgGzqmuQNVwMF",
		Fee:    "10000000",
		Date:   1556952520,
		Type:   blockatlas.TxNativeTokenTransfer,
		Status: blockatlas.StatusCompleted,
		Block:  3411141,
		Meta: blockatlas.AnyAction{
			Coin:     coin.Ontology().ID,
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  "ong",
			Decimals: 9,
			Value:    "10000000000",
			Title:    blockatlas.AnyActionClaimRewards,
			Key:      blockatlas.KeyStakeClaimRewards,
		},
	}
)

func TestNormalize(t *testing.T) {
	var tests = []struct {
		name        string
		Transaction string
		AssetName   AssetType
		Expected    blockatlas.Tx
		wantErr     bool
	}{
		{"normalize ont", srcOntTransfer, AssetONT, dstOntTransfer, false},
		{"normalize ong", srcOngTransfer, AssetONG, dstOngTransfer, false},
		{"normalize claim reward", srcRewardTransfer, AssetONG, dstRewardTransfer, false},
		{"normalize fee", srcFeeTransfer, AssetONG, blockatlas.Tx{}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var sourceTx Tx
			err := json.Unmarshal([]byte(test.Transaction), &sourceTx)
			assert.Nil(t, err)
			tx, ok := Normalize(&sourceTx, test.AssetName)
			if test.wantErr {
				assert.False(t, ok)
				return
			}
			assert.True(t, ok)
			assert.Equal(t, test.Expected, tx)
		})
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
				Transfers: Transfers{
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
				Transfers: Transfers{
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
				Transfers: Transfers{
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
	want := `{"number":7707834,"id":"a5f3ee1a102df7196bb1e262a05435f260392fae6be676ae2c0a6147f8ecf94c","txs":[{"id":"266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb","coin":1024,"from":"AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"any_action","memo":"","metadata":{"coin":1024,"title":"Claim Rewards","key":"stake_claim_rewards","token_id":"ong","name":"Ontology Gas","symbol":"ONG","decimals":9,"value":"51000000000000"}},{"id":"2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd","coin":1024,"from":"ANdrA47zDXUu8MCkMdD3FYPmpSNGYeAvKz","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"any_action","memo":"","metadata":{"coin":1024,"title":"Claim Rewards","key":"stake_claim_rewards","token_id":"ong","name":"Ontology Gas","symbol":"ONG","decimals":9,"value":"113200000000"}},{"id":"40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3","coin":1024,"from":"Abg2gs6pfpQu82jXbm8EYGiipRBvf9ktVS","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"any_action","memo":"","metadata":{"coin":1024,"title":"Claim Rewards","key":"stake_claim_rewards","token_id":"ong","name":"Ontology Gas","symbol":"ONG","decimals":9,"value":"10949000000000"}}]}`
	assert.Equal(t, want, string(got))
}
