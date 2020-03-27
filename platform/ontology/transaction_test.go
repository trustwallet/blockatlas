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
  "tx_hash": "ea0e5d8e389cb96760887094194ca359ac998b2f607be470a576861b91e2bf52",
  "tx_type": 209,
  "tx_time": 1578903628,
  "block_height": 7571132,
  "fee": "0.01",
  "block_index": 1,
  "confirm_flag": 1,
  "transfers": [
    {
      "amount": "5",
      "from_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "to_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "asset_name": "ont",
      "contract_hash": "0100000000000000000000000000000000000000"
    },
    {
      "amount": "0.01",
      "from_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
      "asset_name": "ong",
      "contract_hash": "0200000000000000000000000000000000000000"
    }
  ]
}`

	srcOngTransfer = `
{
  "tx_hash": "e5946ba02f56e17c3709db2bc91f43f76ee3a359006586024daa5c4ad8c54e78",
  "tx_type": 209,
  "tx_time": 1577631515,
  "block_height": 7457989,
  "fee": "0.01",
  "block_index": 1,
  "confirm_flag": 1,
  "transfers": [
    {
      "amount": "14.69",
      "from_address": "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
      "to_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "asset_name": "ong",
      "contract_hash": "0200000000000000000000000000000000000000"
    }
  ]
}`

	srcRewardTransfer = `
{
  "tx_hash": "d7554dcdf01f394b9107ff598df6d84e4c3b00ccf1e720b8c09abf085cbe4987",
  "tx_type": 209,
  "tx_time": 1579699532,
  "block_height": 7644328,
  "fee": "0.01",
  "block_index": 1,
  "confirm_flag": 1,
  "transfers": [
    {
      "amount": "0.03534404",
      "from_address": "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV",
      "to_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "asset_name": "ong",
      "contract_hash": "0200000000000000000000000000000000000000"
    },
    {
      "amount": "0.01",
      "from_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
      "asset_name": "ong",
      "contract_hash": "0200000000000000000000000000000000000000"
    }
  ]
}`
	srcFeeTransfer = `
{
  "tx_hash": "92a79aed526f31999e22d9e3912d4125d3f85ec3c63eede4b7dde4a041826095",
  "tx_type": 209,
  "tx_time": 1578902776,
  "block_height": 7571071,
  "fee": "0.01",
  "block_index": 1,
  "confirm_flag": 1,
  "transfers": [
    {
      "amount": "0.01",
      "from_address": "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
      "to_address": "AFmseVrdL9f9oyCzZefL9tG6UbviEH9ugK",
      "asset_name": "ong",
      "contract_hash": "0200000000000000000000000000000000000000"
    }
  ]
}`
)

var (
	dstOntTransfer = blockatlas.Tx{
		ID:     "ea0e5d8e389cb96760887094194ca359ac998b2f607be470a576861b91e2bf52",
		Coin:   coin.ONT,
		From:   "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		To:     "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		Fee:    "10000000",
		Date:   1578903628,
		Type:   "transfer",
		Status: blockatlas.StatusCompleted,
		Block:  7571132,
		Meta: blockatlas.Transfer{
			Value:    "5",
			Symbol:   "ONT",
			Decimals: 0,
		},
	}
	dstOngTransfer = blockatlas.Tx{
		ID:     "e5946ba02f56e17c3709db2bc91f43f76ee3a359006586024daa5c4ad8c54e78",
		Coin:   coin.ONT,
		From:   "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
		To:     "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		Fee:    "10000000",
		Date:   1577631515,
		Type:   blockatlas.TxNativeTokenTransfer,
		Status: blockatlas.StatusCompleted,
		Block:  7457989,
		Meta: blockatlas.NativeTokenTransfer{
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  "ong",
			Decimals: 9,
			Value:    "14690000000",
			From:     "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
			To:       "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		},
	}
	dstRewardTransfer = blockatlas.Tx{
		ID:     "d7554dcdf01f394b9107ff598df6d84e4c3b00ccf1e720b8c09abf085cbe4987",
		Coin:   coin.ONT,
		From:   "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV",
		To:     "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		Fee:    "10000000",
		Date:   1579699532,
		Type:   blockatlas.TxAnyAction,
		Status: blockatlas.StatusCompleted,
		Block:  7644328,
		Meta: blockatlas.AnyAction{
			Coin:     coin.Ontology().ID,
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  "ong",
			Decimals: 9,
			Value:    "35344040",
			Title:    blockatlas.AnyActionClaimRewards,
			Key:      blockatlas.KeyStakeClaimRewards,
		},
	}
)

func TestNormalize(t *testing.T) {
	tests := []struct {
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
	ontTxResp1 = TxResult{
		BaseResponse: BaseResponse{Code: 0, Msg: "SUCCESS"},
		Result: Tx{
			Hash:        "266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb",
			Time:        1580481541,
			Height:      7707834,
			Fee:         "0.01",
			Description: "transfer",
			BlockIndex:  2,
			ConfirmFlag: 1,
			EventType:   3,
			Details: Detail{
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

	ontTxResp2 = TxResult{
		BaseResponse: BaseResponse{Code: 0, Msg: "SUCCESS"},
		Result: Tx{
			Hash:        "2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd",
			Time:        1580481541,
			Height:      7707834,
			Fee:         "0.01",
			Description: "transfer",
			BlockIndex:  3,
			ConfirmFlag: 1,
			EventType:   3,
			Details: Detail{
				Transfers: Transfers{
					{
						Amount:      "113.2",
						AssetName:   "ont",
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

	ontTxResp3 = TxResult{
		BaseResponse: BaseResponse{Code: 0, Msg: "SUCCESS"},
		Result: Tx{
			Hash:        "40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3",
			Time:        1580481541,
			Height:      7707834,
			Fee:         "0.01",
			Description: "transfer",
			BlockIndex:  1,
			ConfirmFlag: 1,
			EventType:   3,
			Details: Detail{
				Transfers: Transfers{
					{
						Amount:      "10949",
						AssetName:   "ont",
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
	block := normalizeTxs([]Tx{ontTxResp1.Result, ontTxResp2.Result, ontTxResp3.Result}, AssetAll)
	got, err := json.Marshal(block)
	if err != nil {
		t.Fatal(err)
	}
	want := `[{"id":"266d9d7282a5601bf6cb8fc5368a76a2aa54f45731a063a699a692487bcbd0cb","coin":1024,"from":"AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"native_token_transfer","memo":"","metadata":{"name":"Ontology Gas","symbol":"ONG","token_id":"ong","decimals":9,"value":"51000000000000","from":"AbEeCHUWpzQaxUN7G1a83N3P2XtVLuMLaE","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx"}},{"id":"2935268c5715f1f2015ba828681c39399dedbe7a24ed628ef7b85d9aac8045fd","coin":1024,"from":"ANdrA47zDXUu8MCkMdD3FYPmpSNGYeAvKz","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"transfer","memo":"","metadata":{"value":"113.2","symbol":"ONT","decimals":0}},{"id":"40976edc1306b0e5f55b90c8d3ca248bb544e5ebbadb02be6146ba0a0de402c3","coin":1024,"from":"Abg2gs6pfpQu82jXbm8EYGiipRBvf9ktVS","to":"ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx","fee":"10000000","date":1580481541,"block":7707834,"status":"completed","type":"transfer","memo":"","metadata":{"value":"10949","symbol":"ONT","decimals":0}}]`
	assert.Equal(t, want, string(got))
}
