package ontology

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	srcOntTransfer, _    = mock.JsonStringFromFilePath("mocks/transfer_ont.json")
	srcOngTransfer, _    = mock.JsonStringFromFilePath("mocks/transfer_ong.json")
	srcRewardTransfer, _ = mock.JsonStringFromFilePath("mocks/transfer_rewards.json")
	srcFeeTransfer, _    = mock.JsonStringFromFilePath("mocks/transfer_fee.json")

	dstOntTransfer = types.Tx{
		ID:     "ea0e5d8e389cb96760887094194ca359ac998b2f607be470a576861b91e2bf52",
		Coin:   coin.ONTOLOGY,
		From:   "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		To:     "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		Fee:    "10000000",
		Date:   1578903628,
		Type:   "transfer",
		Status: types.StatusCompleted,
		Block:  7571132,
		Meta: types.Transfer{
			Value:    "5",
			Symbol:   "ONT",
			Decimals: 0,
		},
	}
	dstOngTransfer = types.Tx{
		ID:     "e5946ba02f56e17c3709db2bc91f43f76ee3a359006586024daa5c4ad8c54e78",
		Coin:   coin.ONTOLOGY,
		From:   "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
		To:     "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		Fee:    "10000000",
		Date:   1577631515,
		Type:   types.TxNativeTokenTransfer,
		Status: types.StatusCompleted,
		Block:  7457989,
		Meta: types.NativeTokenTransfer{
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  "ong",
			Decimals: 9,
			Value:    "14690000000",
			From:     "ASLbwuar3ZTbUbLPnCgjGUw2WHhMfvJJtx",
			To:       "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		},
	}
	dstRewardTransfer = types.Tx{
		ID:     "d7554dcdf01f394b9107ff598df6d84e4c3b00ccf1e720b8c09abf085cbe4987",
		Coin:   coin.ONTOLOGY,
		From:   "AFmseVrdL9f9oyCzZefL9tG6UbvhUMqNMV",
		To:     "ARFXGXSmgFT2h9EiS4D5fen127Lzi48Eij",
		Fee:    "10000000",
		Date:   1579699532,
		Type:   types.TxAnyAction,
		Status: types.StatusCompleted,
		Block:  7644328,
		Meta: types.AnyAction{
			Coin:     coin.Ontology().ID,
			Name:     "Ontology Gas",
			Symbol:   "ONG",
			TokenID:  "ong",
			Decimals: 9,
			Value:    "35344040",
			Title:    types.AnyActionClaimRewards,
			Key:      types.KeyStakeClaimRewards,
		},
	}
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		name        string
		Transaction string
		AssetName   AssetType
		Expected    types.Tx
		wantErr     bool
	}{
		{"normalize ont", srcOntTransfer, AssetONT, dstOntTransfer, false},
		{"normalize ong", srcOngTransfer, AssetONG, dstOngTransfer, false},
		{"normalize claim reward", srcRewardTransfer, AssetONG, dstRewardTransfer, false},
		{"normalize fee", srcFeeTransfer, AssetONG, types.Tx{}, true},
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
	var want types.Txs
	_ = mock.JsonModelFromFilePath("mocks/block_response.json", &want)
	lhs, _ := json.Marshal(block)
	rhs, _ := json.Marshal(want)
	assert.JSONEq(t, string(lhs), string(rhs))
}
