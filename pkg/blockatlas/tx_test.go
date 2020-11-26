package blockatlas

import (
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/tokentype"
	"reflect"
	"sort"
	"testing"
)

var transferDst1 = Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	Fee:    "125000",
	Date:   1555049867,
	Block:  7761368,
	Status: StatusCompleted,
	Memo:   "test",
	Meta: Transfer{
		Value:    "10000000000000",
		Decimals: 8,
		Symbol:   "BNB",
	},
}

var nativeTransferDst1 = Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	Coin:   coin.BNB,
	From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	Fee:    "125000",
	Date:   1555117625,
	Block:  7928667,
	Status: StatusCompleted,
	Memo:   "test",
	Meta: NativeTokenTransfer{
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	},
}

var utxoTransferDst1 = Tx{
	ID:   "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
	Coin: coin.BTC,
	Inputs: []TxOutput{
		{
			Address: "bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h",
			Value:   "1",
		},
		{
			Address: "bc1qc7ekqf2t0elfsmtgr2mgd7da2up4vgq8uqk2nh",
			Value:   "1",
		},
		{
			Address: "bc1qv454wacvnenr3hzzldjqn8cgfltdlxwe96h737",
			Value:   "1",
		},
	},
	Outputs: []TxOutput{
		{
			Address: "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6",
			Value:   "3",
		},
	},
	Fee:    "125000",
	Date:   1555117625,
	Block:  592400,
	Status: StatusCompleted,
	Memo:   "test",
}

var utxoTransferDst2 = Tx{
	ID:   "zpub6ruK9k6YGm8BRHWvTiQcrEPnFkuRDJhR7mPYzV2LDvjpLa5CuGgrhCYVZjMGcLcFqv9b2WvsFtY2Gb3xq8NVq8qhk9veozrA2W9QaWtihrC",
	Coin: coin.BTC,
	Inputs: []TxOutput{
		{
			Address: "bc1q6e8sdxlgc7ekqkqyevtrx8wshfv7sg66z3z6ce",
			Value:   "4",
		},
		{
			Address: "bc1q7nn4txus4g6fc5v7d2tha35ely8mfpd8qvv6eg",
			Value:   "2",
		},
	},
	Outputs: []TxOutput{
		{
			Address: "bc1q2fpry7zwqh575huc9urwfdvjtuvz508wez56ff",
			Value:   "3",
		},
		{
			Address: "bc1qk3yj6h79qw7tnsg4durc9sd5fpd3qt0p0m8u5p",
			Value:   "1",
		},
		{
			Address: "bc1qm8836plkzft2rhh23z6j8s9s8fxrzd4zag95z8",
			Value:   "2",
		},
	},
	Fee:    "125000",
	Date:   1555117625,
	Block:  592400,
	Status: StatusCompleted,
	Memo:   "test",
}

func TestTx_GetAddresses(t *testing.T) {
	assert.Equal(t, transferDst1.GetAddresses(), []string{"tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2", "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"})
	assert.Equal(t, nativeTransferDst1.GetAddresses(), []string{"tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a", "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"})
}

func TestTx_GetUtxoAddresses(t *testing.T) {
	assert.Equal(t, utxoTransferDst1.GetUtxoAddresses(), []string{"bc1qhn03cww757mnnlpkdvvfkaydxqygm86nvkm92h", "bc1qc7ekqf2t0elfsmtgr2mgd7da2up4vgq8uqk2nh", "bc1qv454wacvnenr3hzzldjqn8cgfltdlxwe96h737", "bc1qjcslq88cht8llqmh3aqshjx9we9msv386jvxl6"})
	assert.Equal(t, utxoTransferDst2.GetUtxoAddresses(), []string{"bc1q6e8sdxlgc7ekqkqyevtrx8wshfv7sg66z3z6ce", "bc1q7nn4txus4g6fc5v7d2tha35ely8mfpd8qvv6eg", "bc1q2fpry7zwqh575huc9urwfdvjtuvz508wez56ff", "bc1qk3yj6h79qw7tnsg4durc9sd5fpd3qt0p0m8u5p", "bc1qm8836plkzft2rhh23z6j8s9s8fxrzd4zag95z8"})
}

func Test_getDirection(t *testing.T) {
	type args struct {
		tx      Tx
		address string
	}
	tests := []struct {
		name string
		args args
		want Direction
	}{
		{"Test Direction Self",
			args{
				Tx{
					From: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1", To: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
				"0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"}, DirectionSelf,
		},
		{"Test Direction Outgoing",
			args{
				Tx{
					From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", To: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7"},
				"0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB"}, DirectionOutgoing,
		},
		{"Test Direction Incoming",
			args{
				Tx{
					From: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7", To: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
				"0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"}, DirectionIncoming,
		},
		{"Test UTXO Direction Self",
			args{
				Tx{
					Outputs: []TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "72934112534"},
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "500000000"},
					},
					Inputs: []TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "73196112534"},
					},
				}, "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S",
			}, DirectionSelf,
		},
		{"Test UTXO Direction Outgoing",
			args{
				Tx{
					Outputs: []TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd",
			}, DirectionOutgoing,
		},
		{"Test UTXO Direction Incoming",
			args{
				Tx{
					Outputs: []TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4",
			}, DirectionIncoming,
		},
		{"Test NativeTokenTransfer Direction Self",
			args{
				Tx{
					Meta: &NativeTokenTransfer{
						From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:   "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, DirectionSelf,
		},
		{"Test NativeTokenTransfer Direction Outgoing",
			args{
				Tx{
					Meta: &NativeTokenTransfer{
						From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:   "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, DirectionOutgoing,
		},
		{"Test NativeTokenTransfer Direction Incoming",
			args{
				Tx{
					Meta: &NativeTokenTransfer{
						From: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
						To:   "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, DirectionIncoming,
		},
		{"Test TokenTransfer Direction Self",
			args{
				Tx{
					Meta: &TokenTransfer{
						From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:   "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, DirectionSelf,
		},
		{"Test TokenTransfer Direction Outgoing",
			args{
				Tx{
					Meta: &TokenTransfer{
						From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:   "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, DirectionOutgoing,
		},
		{"Test TokenTransfer Direction Incoming",
			args{
				Tx{
					Meta: &TokenTransfer{
						From: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
						To:   "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, DirectionIncoming,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.tx.GetTransactionDirection(tt.args.address); got != tt.want {
				t.Errorf("getTransactionDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inferUtxoValue(t *testing.T) {
	type args struct {
		tx        Tx
		address   string
		coinIndex uint
	}
	tests := []struct {
		name       string
		args       args
		wantAmount Amount
	}{
		{"Test UTXO Direction Self",
			args{
				Tx{
					Outputs: []TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "72934112534"},
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "500000000"},
					},
					Inputs: []TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "73196112534"},
					},
				}, "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", 3,
			}, Amount("72934112534"),
		},
		{"Test UTXO Direction Outgoing",
			args{
				Tx{
					Outputs: []TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", 0,
			}, Amount("4471835"),
		},
		{"Test UTXO Direction Incoming",
			args{
				Tx{
					Outputs: []TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", 0,
			}, Amount("4471835"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect := Transfer{
				Value:    tt.wantAmount,
				Symbol:   coin.Coins[tt.args.coinIndex].Symbol,
				Decimals: coin.Coins[tt.args.coinIndex].Decimals,
			}
			tt.args.tx.Direction = tt.args.tx.GetTransactionDirection(tt.args.address)
			if tt.args.tx.InferUtxoValue(tt.args.address, tt.args.coinIndex); tt.args.tx.Meta != expect {
				t.Errorf("inferUtxoValue() = %v, want %v", tt.args.tx.Meta, expect)
			}
		})
	}
}

// zpub: zpub6r9CEhEkruYbEcu2yQCaRKQ1qufTa4zLrx6ezs31P627UpAepVNBE2td3d3mHnSaXyRbwksRwDJGzLBWQeZPFMut8N3BvXpcwRwEWGEwAnq
var (
	btcSet = mapset.NewSet("bc1qfrrncxmf7skye2glyef95xlpmrlmf2e8qlav2l", "bc1qxm90n0rxkadhdkvglev56k60qths73luzlnn7a",
		"bc1q2sykr9c342mjpm9mwnps8ksk6e35lz75rpdlfe", "bc1qs86ucvr3unce2grvfp77433npy66nzha9w0e3c")
	btcInputs1  = []TxOutput{{Address: "bc1q2sykr9c342mjpm9mwnps8ksk6e35lz75rpdlfe"}}
	btcOutputs1 = []TxOutput{{Address: "bc1q6wf7tj62f0uwr6almah3666th2ejefdg72ek6t"}}
	btcInputs2  = []TxOutput{{
		Address: "3CgvDkzcJ7yMZe75jNBem6Bj6nkMAWwMEf"},
		{Address: "3LyzYcB54pm9EAMmzXpFfb1kzEDAFvqBgT"},
		{Address: "3Q6DYour5q5WdMhyXsyPgBeAqPCXchzCsF"},
		{Address: "3JZZM1rwst7G5izxbFL7KNvy7ZiZ47SVqG"}}
	btcOutputs2 = []TxOutput{
		{Address: "139f1CrnLWvVajGzs3ZtpQhbGWxM599sho"},
		{Address: "3LyzYcB54pm9EAMmzXpFfb1kzEDAFvqBgT"},
		{Address: "bc1q9mx5tm66zs7epa4skvyuf2vfuwmtnlttj74cnl"},
		{Address: "3JZZM1rwst7G5izxbFL7KNvy7ZiZ47SVqG"}}

	dogeSet     = mapset.NewSet("DB49sNjVdxyREXEBEzUV54TrQYYpvi3Be7")
	dogeInputs  = []TxOutput{{Address: "DAukM5pPtGdbPxMX1u2LYHoyhbDhEFHbnH"}}
	dogeOutputs = []TxOutput{{Address: "DB49sNjVdxyREXEBEzUV54TrQYYpvi3Be7"}, {Address: "DAukM5pPtGdbPxMX1u2LYHoyhbDhEFHbnH"}}
)

func TestInferDirection(t *testing.T) {
	tests := []struct {
		AddressSet mapset.Set
		Inputs     []TxOutput
		Outputs    []TxOutput
		Expected   Direction
		Coin       uint
	}{
		{
			btcSet,
			btcInputs1,
			btcOutputs1,
			DirectionOutgoing,
			coin.BTC,
		},
		{
			btcSet,
			btcInputs2,
			btcOutputs2,
			DirectionIncoming,
			coin.BTC,
		},
		{
			dogeSet,
			dogeInputs,
			dogeOutputs,
			DirectionIncoming,
			coin.DOGE,
		},
	}

	for _, test := range tests {
		tx := Tx{
			Inputs:  test.Inputs,
			Outputs: test.Outputs,
		}

		direction := InferDirection(&tx, test.AddressSet)
		if direction != test.Expected {
			t.Errorf("direction is not %s", test.Expected)
		}
	}
}

func TestTx_GetTransactionDirection(t *testing.T) {
	txMeta := TokenTransfer{
		Name:     "Kyber Network Crystal",
		Symbol:   "KNC",
		TokenID:  "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Decimals: 18,
		Value:    "100000000000000",
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0x38d45371993eEc84f38FEDf93C646aA2D2267CEA",
	}

	tx := Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c4",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169424,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
		Meta:     txMeta,
	}

	tx.Direction = tx.GetTransactionDirection("0x38d45371993eEc84f38FEDf93C646aA2D2267CEA")
	assert.Equal(t, Direction("incoming"), tx.Direction)

	tx.Meta = &txMeta

	tx.Direction = tx.GetTransactionDirection("0x38d45371993eEc84f38FEDf93C646aA2D2267CEA")
	assert.Equal(t, Direction("incoming"), tx.Direction)

	tx.Direction = DirectionSelf
	tx.Direction = tx.GetTransactionDirection("0x38d45371993eEc84f38FEDf93C646aA2D2267CEA")
	assert.Equal(t, Direction("yourself"), tx.Direction)
}

func TestTxs_FilterUniqueID(t *testing.T) {
	tx := Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c4",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169424,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
	}
	tx2 := Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c4",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169424,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
	}

	txs := make([]Tx, 0)
	txs = append(txs, tx)
	txs = append(txs, tx2)

	entry := Txs(txs)

	result := entry.FilterUniqueID()

	assert.Equal(t, entry[:1], result)
}

func TestTxs_SortByDate(t *testing.T) {
	tx := Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c4",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169423,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
	}
	tx2 := Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c5",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169424,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
	}
	tx3 := Tx{
		ID:       "0xbcd1a43e796de4035e5e2991d8db332958e36031d54cb1d3a08d2cb790e338c6",
		Coin:     60,
		From:     "0x08777CB1e80F45642752662B04886Df2d271E049",
		To:       "0xdd974D5C2e2928deA5F71b9825b8b646686BD200",
		Fee:      "52473000000000",
		Date:     1585169425,
		Block:    9742705,
		Status:   "completed",
		Sequence: 149,
		Type:     "token_transfer",
	}

	txs := make([]Tx, 0)
	txs = append(txs, tx)
	txs = append(txs, tx2)
	txs = append(txs, tx3)

	entry := Txs(txs)
	isNotSorted := sort.SliceIsSorted(entry, func(i, j int) bool {
		return entry[i].Date > entry[j].Date
	})
	assert.True(t, !isNotSorted)
	result := entry.SortByDate()
	isSorted := sort.SliceIsSorted(result, func(i, j int) bool {
		return result[i].Date > result[j].Date
	})
	assert.True(t, isSorted)
}

func TestTx_TokenID(t *testing.T) {
	tx1 := Tx{
		Coin: 60,
		From: "A",
		To:   "B",
		Meta: NativeTokenTransfer{
			TokenID: "ABC",
			From:    "A",
			To:      "C",
		},
	}

	tx2 := Tx{
		Coin: 60,
		From: "D",
		To:   "V",
		Meta: TokenTransfer{
			TokenID: "EFG",
			From:    "D",
			To:      "F",
		},
	}

	tx3 := Tx{
		Coin: 60,
		From: "Q",
		To:   "L",
		Meta: AnyAction{
			TokenID: "HIJ",
		},
	}

	token1, ok1 := tx1.TokenID()
	assert.True(t, ok1)
	assert.Equal(t, token1, "ABC")
	token2, ok2 := tx2.TokenID()
	assert.True(t, ok2)
	assert.Equal(t, token2, "EFG")
	token3, ok3 := tx3.TokenID()
	assert.Equal(t, token3, "HIJ")
	assert.True(t, ok3)

}

func TestTokenType(t *testing.T) {
	type testStruct struct {
		Name       string
		ID         uint
		TokenID    string
		WantedType string
		WantedOk   bool
	}
	tests := []testStruct{
		{
			Name:       "Tron TRC20",
			ID:         coin.Tron().ID,
			TokenID:    "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t",
			WantedType: string(tokentype.TRC20),
			WantedOk:   true,
		},
		{
			Name:       "Tron TRC10",
			ID:         coin.Tron().ID,
			TokenID:    "1002000",
			WantedType: string(tokentype.TRC10),
			WantedOk:   true,
		},
		{
			Name:       "Ethereum ERC20",
			ID:         coin.Ethereum().ID,
			TokenID:    "dai",
			WantedType: string(tokentype.ERC20),
			WantedOk:   true,
		},
		{
			Name:       "Binance BEP20",
			ID:         coin.Smartchain().ID,
			TokenID:    "busd",
			WantedType: string(tokentype.BEP20),
			WantedOk:   true,
		},
		{
			Name:       "Binance BEP10",
			ID:         coin.Binance().ID,
			TokenID:    "busd",
			WantedType: string(tokentype.BEP2),
			WantedOk:   true,
		},
		{
			Name:       "Wrong",
			ID:         coin.Bitcoin().ID,
			TokenID:    "busd",
			WantedType: "",
			WantedOk:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			expectedType, expectedOk := GetTokenType(tt.ID, tt.TokenID)
			assert.Equal(t, tt.WantedType, expectedType)
			assert.Equal(t, tt.WantedOk, expectedOk)
		})
	}
}

var (
	beforeTransactionsToken = `[{"id":"59FE2D56F0E2320023D5D579497FF4E5FD8B226AEF14AADCE089A241BE0051DA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592484005,"block":95113723,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103643577","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089263465","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"DA10ACADF3BF23072CBAAEA02E51C48331DDE588EC16BA6AA0FF79F9F37CD08E","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592442063,"block":95008674,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"1B7696018A0BEB07EB41F1E8D7A20B64AA9DB6AFDC61CD30EFA082D4E8650BD3","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592429461,"block":94977026,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2511060524","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"B98EF6224FD385830CFE88B431FE4F098D1ABAFB349A138DD11235FD02B619B7","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592348463,"block":94777896,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"10998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"C6E37E9966BD52EB612B23DAD30BFD68BE05D4670B7FDC64DA2C109305247330","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592253063,"block":94538295,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4146300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"A9F58764D7F9555000FD0C1BD0B403953F05AE2DD9E094127315BE8FBF0FA19E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592250904,"block":94532906,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"63266290000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"91B7F8348917289F08A9C2B5BAF8F6673AB6808C9F2CAD4930BC5BB62AAA469B","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592246708,"block":94522375,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"104150629","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4512670728","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"62A224A212083CC53142230A0598107305C7659B687F4B17FDBDA1D03008679E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592245803,"block":94520016,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"58689300000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5E2D321690FFD6BBCD0247DDCE07A1D41AE659F6C11F0662DDD6C7D8242E3D00","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592244606,"block":94517032,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"108137171","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1053055990","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"E7A814D1E23F269914FB4CDECF61D6A3D8D900684CFC8280BBB7989A3EF6C071","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592237402,"block":94498680,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"8479394653","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"8B71C1AA5C4C636127E8A08D67AB7DF4BC5E4EFA25F99F12A34FFF3B8928BB14","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592223002,"block":94462381,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"38451401547","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"4C96992F1D5D971C1C9D787C0D96EE342875147D083E79759BF569CE9E2F63B4","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592220661,"block":94456416,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"99998300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"2001461D2CCF84864B999235B6A0B08A2A75649367A7F809651E73C2D49E94F3","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592201404,"block":94407633,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103229158","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089751117","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"0AE83A51E05BB34C8FE0AA0AC376B9C7AE8960A9A77EB9C687B29994A53D1159","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592194201,"block":94389286,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1098704296","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"68C4BA1864A52EA78FD99610FB9FE64DDB86EDDA2F020D7E250AB8E3E8E5CF32","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592179801,"block":94352682,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2461605848","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5D6939D3AE2AFDFF6A3ECA5DD48CB5D8D2F0389A35C414806CF284B453DA5133","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592177461,"block":94346701,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2324800000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"5714D3F5C442E2491EEF8BCA9D3CBDF18005C5B1D67BA1A7E873ADFC46F6F4B8","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592165402,"block":94316303,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"156671627595","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"1823E07FBB966050060D127B266F9FB64F30DF1F549C3906BDE234D66A060766","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592163611,"block":94311779,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"106835633","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2042060818","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"A0A6DA82F4F4610145635FFB1F52CF8EC5F8F51BDF0256B5B0F77D3D5692E2CA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592158202,"block":94298059,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"3101371","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"B380CF3BA7831EBD82240CFCC369572FCC022ECA3A91A7F11743D07C1599F636","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592154601,"block":94288998,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"367696585025","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"7ABEF721E2A3D2B70583F4641FE426D0CAE55AD11C7021809289E504917E4569","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592073191,"block":94084342,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2503750000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"C422B298A2EABD7EA503CC1F525C20BEB606FB83D01FE760A69F0FA57DB65B03","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592066528,"block":94067506,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"12959997250","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"FD656B1D90C660F4F6D74E4C0584AC11ED0626F6FBF5E8C33D0E9131F5DB04C2","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060464,"block":94052345,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2781326500","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"0447FA8D100F715DCE70513AB9ECF5BCF16D38A6C450B33DE0257BB654FBD50E","coin":714,"from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060463,"block":94052341,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"39644614000","from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"19583E4A440801AA458F06D655D9745FE0A75B76FF776CEA65455A9C9D9F59D9","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592053265,"block":94034212,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"40798300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}}]`
	//beforeTransactionsMemo  = `[{"id":"59FE2D56F0E2320023D5D579497FF4E5FD8B226AEF14AADCE089A241BE0051DA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592484005,"block":95113723,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"trust.com","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089263465","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"DA10ACADF3BF23072CBAAEA02E51C48331DDE588EC16BA6AA0FF79F9F37CD08E","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592442063,"block":95008674,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"1B7696018A0BEB07EB41F1E8D7A20B64AA9DB6AFDC61CD30EFA082D4E8650BD3","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592429461,"block":94977026,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2511060524","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"B98EF6224FD385830CFE88B431FE4F098D1ABAFB349A138DD11235FD02B619B7","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592348463,"block":94777896,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"10998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"C6E37E9966BD52EB612B23DAD30BFD68BE05D4670B7FDC64DA2C109305247330","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592253063,"block":94538295,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4146300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"A9F58764D7F9555000FD0C1BD0B403953F05AE2DD9E094127315BE8FBF0FA19E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592250904,"block":94532906,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"63266290000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"91B7F8348917289F08A9C2B5BAF8F6673AB6808C9F2CAD4930BC5BB62AAA469B","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592246708,"block":94522375,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"104150629","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4512670728","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"62A224A212083CC53142230A0598107305C7659B687F4B17FDBDA1D03008679E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592245803,"block":94520016,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"58689300000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5E2D321690FFD6BBCD0247DDCE07A1D41AE659F6C11F0662DDD6C7D8242E3D00","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592244606,"block":94517032,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"108137171","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1053055990","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"E7A814D1E23F269914FB4CDECF61D6A3D8D900684CFC8280BBB7989A3EF6C071","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592237402,"block":94498680,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"8479394653","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"8B71C1AA5C4C636127E8A08D67AB7DF4BC5E4EFA25F99F12A34FFF3B8928BB14","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592223002,"block":94462381,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"38451401547","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"4C96992F1D5D971C1C9D787C0D96EE342875147D083E79759BF569CE9E2F63B4","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592220661,"block":94456416,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"99998300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"2001461D2CCF84864B999235B6A0B08A2A75649367A7F809651E73C2D49E94F3","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592201404,"block":94407633,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103229158","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089751117","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"0AE83A51E05BB34C8FE0AA0AC376B9C7AE8960A9A77EB9C687B29994A53D1159","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592194201,"block":94389286,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1098704296","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"68C4BA1864A52EA78FD99610FB9FE64DDB86EDDA2F020D7E250AB8E3E8E5CF32","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592179801,"block":94352682,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2461605848","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5D6939D3AE2AFDFF6A3ECA5DD48CB5D8D2F0389A35C414806CF284B453DA5133","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592177461,"block":94346701,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2324800000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"5714D3F5C442E2491EEF8BCA9D3CBDF18005C5B1D67BA1A7E873ADFC46F6F4B8","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592165402,"block":94316303,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"156671627595","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"1823E07FBB966050060D127B266F9FB64F30DF1F549C3906BDE234D66A060766","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592163611,"block":94311779,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"106835633","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2042060818","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"A0A6DA82F4F4610145635FFB1F52CF8EC5F8F51BDF0256B5B0F77D3D5692E2CA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592158202,"block":94298059,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"3101371","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"B380CF3BA7831EBD82240CFCC369572FCC022ECA3A91A7F11743D07C1599F636","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592154601,"block":94288998,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"367696585025","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"7ABEF721E2A3D2B70583F4641FE426D0CAE55AD11C7021809289E504917E4569","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592073191,"block":94084342,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2503750000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"C422B298A2EABD7EA503CC1F525C20BEB606FB83D01FE760A69F0FA57DB65B03","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592066528,"block":94067506,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"12959997250","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"FD656B1D90C660F4F6D74E4C0584AC11ED0626F6FBF5E8C33D0E9131F5DB04C2","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060464,"block":94052345,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2781326500","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"0447FA8D100F715DCE70513AB9ECF5BCF16D38A6C450B33DE0257BB654FBD50E","coin":714,"from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060463,"block":94052341,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"39644614000","from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"19583E4A440801AA458F06D655D9745FE0A75B76FF776CEA65455A9C9D9F59D9","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592053265,"block":94034212,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"40798300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}}]`
	wantedTransactionsToken = `[{"id":"59FE2D56F0E2320023D5D579497FF4E5FD8B226AEF14AADCE089A241BE0051DA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592484005,"block":95113723,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103643577","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089263465","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"DA10ACADF3BF23072CBAAEA02E51C48331DDE588EC16BA6AA0FF79F9F37CD08E","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592442063,"block":95008674,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"1B7696018A0BEB07EB41F1E8D7A20B64AA9DB6AFDC61CD30EFA082D4E8650BD3","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592429461,"block":94977026,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2511060524","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"B98EF6224FD385830CFE88B431FE4F098D1ABAFB349A138DD11235FD02B619B7","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592348463,"block":94777896,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"10998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"C6E37E9966BD52EB612B23DAD30BFD68BE05D4670B7FDC64DA2C109305247330","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592253063,"block":94538295,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4146300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"A9F58764D7F9555000FD0C1BD0B403953F05AE2DD9E094127315BE8FBF0FA19E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592250904,"block":94532906,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"63266290000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"91B7F8348917289F08A9C2B5BAF8F6673AB6808C9F2CAD4930BC5BB62AAA469B","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592246708,"block":94522375,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"104150629","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4512670728","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"62A224A212083CC53142230A0598107305C7659B687F4B17FDBDA1D03008679E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592245803,"block":94520016,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"58689300000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5E2D321690FFD6BBCD0247DDCE07A1D41AE659F6C11F0662DDD6C7D8242E3D00","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592244606,"block":94517032,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"108137171","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1053055990","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"E7A814D1E23F269914FB4CDECF61D6A3D8D900684CFC8280BBB7989A3EF6C071","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592237402,"block":94498680,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"8479394653","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"8B71C1AA5C4C636127E8A08D67AB7DF4BC5E4EFA25F99F12A34FFF3B8928BB14","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592223002,"block":94462381,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"38451401547","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"4C96992F1D5D971C1C9D787C0D96EE342875147D083E79759BF569CE9E2F63B4","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592220661,"block":94456416,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"99998300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"2001461D2CCF84864B999235B6A0B08A2A75649367A7F809651E73C2D49E94F3","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592201404,"block":94407633,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103229158","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089751117","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"0AE83A51E05BB34C8FE0AA0AC376B9C7AE8960A9A77EB9C687B29994A53D1159","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592194201,"block":94389286,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1098704296","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"68C4BA1864A52EA78FD99610FB9FE64DDB86EDDA2F020D7E250AB8E3E8E5CF32","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592179801,"block":94352682,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2461605848","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5D6939D3AE2AFDFF6A3ECA5DD48CB5D8D2F0389A35C414806CF284B453DA5133","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592177461,"block":94346701,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2324800000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"5714D3F5C442E2491EEF8BCA9D3CBDF18005C5B1D67BA1A7E873ADFC46F6F4B8","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592165402,"block":94316303,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"156671627595","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"1823E07FBB966050060D127B266F9FB64F30DF1F549C3906BDE234D66A060766","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592163611,"block":94311779,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"106835633","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2042060818","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"A0A6DA82F4F4610145635FFB1F52CF8EC5F8F51BDF0256B5B0F77D3D5692E2CA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592158202,"block":94298059,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"3101371","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"B380CF3BA7831EBD82240CFCC369572FCC022ECA3A91A7F11743D07C1599F636","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592154601,"block":94288998,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"367696585025","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"7ABEF721E2A3D2B70583F4641FE426D0CAE55AD11C7021809289E504917E4569","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592073191,"block":94084342,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2503750000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"C422B298A2EABD7EA503CC1F525C20BEB606FB83D01FE760A69F0FA57DB65B03","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592066528,"block":94067506,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"12959997250","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"FD656B1D90C660F4F6D74E4C0584AC11ED0626F6FBF5E8C33D0E9131F5DB04C2","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060464,"block":94052345,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2781326500","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"0447FA8D100F715DCE70513AB9ECF5BCF16D38A6C450B33DE0257BB654FBD50E","coin":714,"from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060463,"block":94052341,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"39644614000","from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"19583E4A440801AA458F06D655D9745FE0A75B76FF776CEA65455A9C9D9F59D9","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592053265,"block":94034212,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"40798300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}}]`
	//wantedTransactionsMemo  = `[{"id":"59FE2D56F0E2320023D5D579497FF4E5FD8B226AEF14AADCE089A241BE0051DA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592484005,"block":95113723,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089263465","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"DA10ACADF3BF23072CBAAEA02E51C48331DDE588EC16BA6AA0FF79F9F37CD08E","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592442063,"block":95008674,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"1B7696018A0BEB07EB41F1E8D7A20B64AA9DB6AFDC61CD30EFA082D4E8650BD3","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592429461,"block":94977026,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2511060524","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"B98EF6224FD385830CFE88B431FE4F098D1ABAFB349A138DD11235FD02B619B7","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592348463,"block":94777896,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"10998400000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"C6E37E9966BD52EB612B23DAD30BFD68BE05D4670B7FDC64DA2C109305247330","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592253063,"block":94538295,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4146300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"A9F58764D7F9555000FD0C1BD0B403953F05AE2DD9E094127315BE8FBF0FA19E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592250904,"block":94532906,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"63266290000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"91B7F8348917289F08A9C2B5BAF8F6673AB6808C9F2CAD4930BC5BB62AAA469B","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592246708,"block":94522375,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"104150629","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"4512670728","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"62A224A212083CC53142230A0598107305C7659B687F4B17FDBDA1D03008679E","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592245803,"block":94520016,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"58689300000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5E2D321690FFD6BBCD0247DDCE07A1D41AE659F6C11F0662DDD6C7D8242E3D00","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592244606,"block":94517032,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"108137171","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1053055990","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"E7A814D1E23F269914FB4CDECF61D6A3D8D900684CFC8280BBB7989A3EF6C071","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592237402,"block":94498680,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"8479394653","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"8B71C1AA5C4C636127E8A08D67AB7DF4BC5E4EFA25F99F12A34FFF3B8928BB14","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592223002,"block":94462381,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"38451401547","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"4C96992F1D5D971C1C9D787C0D96EE342875147D083E79759BF569CE9E2F63B4","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592220661,"block":94456416,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"99998300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"2001461D2CCF84864B999235B6A0B08A2A75649367A7F809651E73C2D49E94F3","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592201404,"block":94407633,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103229158","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1089751117","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"0AE83A51E05BB34C8FE0AA0AC376B9C7AE8960A9A77EB9C687B29994A53D1159","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592194201,"block":94389286,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"1098704296","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"68C4BA1864A52EA78FD99610FB9FE64DDB86EDDA2F020D7E250AB8E3E8E5CF32","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592179801,"block":94352682,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2461605848","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"5D6939D3AE2AFDFF6A3ECA5DD48CB5D8D2F0389A35C414806CF284B453DA5133","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592177461,"block":94346701,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2324800000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"5714D3F5C442E2491EEF8BCA9D3CBDF18005C5B1D67BA1A7E873ADFC46F6F4B8","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592165402,"block":94316303,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"156671627595","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"1823E07FBB966050060D127B266F9FB64F30DF1F549C3906BDE234D66A060766","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592163611,"block":94311779,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"106835633","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2042060818","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"A0A6DA82F4F4610145635FFB1F52CF8EC5F8F51BDF0256B5B0F77D3D5692E2CA","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592158202,"block":94298059,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"3101371","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"B380CF3BA7831EBD82240CFCC369572FCC022ECA3A91A7F11743D07C1599F636","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592154601,"block":94288998,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103268674","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"367696585025","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"7ABEF721E2A3D2B70583F4641FE426D0CAE55AD11C7021809289E504917E4569","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592073191,"block":94084342,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2503750000","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"C422B298A2EABD7EA503CC1F525C20BEB606FB83D01FE760A69F0FA57DB65B03","coin":714,"from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23","fee":"37500","date":1592066528,"block":94067506,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"outgoing","memo":"103335769","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"12959997250","from":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","to":"bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23"}},{"id":"FD656B1D90C660F4F6D74E4C0584AC11ED0626F6FBF5E8C33D0E9131F5DB04C2","coin":714,"from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060464,"block":94052345,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"2781326500","from":"bnb1f87wdfjeqx5yju5vgjrz62u32lhrqxztscxrxj","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"0447FA8D100F715DCE70513AB9ECF5BCF16D38A6C450B33DE0257BB654FBD50E","coin":714,"from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592060463,"block":94052341,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"39644614000","from":"bnb1xyq7ttkzq4ekmn26kwn2ul2f6ludplnhlqcf6k","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}},{"id":"19583E4A440801AA458F06D655D9745FE0A75B76FF776CEA65455A9C9D9F59D9","coin":714,"from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp","fee":"37500","date":1592053265,"block":94034212,"status":"completed","sequence":0,"type":"native_token_transfer","direction":"incoming","memo":"","metadata":{"name":"","symbol":"BUSD","token_id":"BUSD-BD1","decimals":8,"value":"40798300000","from":"bnb1yjlk7f47qf0z97ph6pxc00s2s0yw048edmtjtw","to":"bnb1t7hpl286qgvsg08lvx6ac9ul0msy4k2ud8dukp"}}]`
)

func Test_filterTransactionsByToken(t *testing.T) {
	var p TxPage
	assert.Nil(t, json.Unmarshal([]byte(beforeTransactionsToken), &p))
	result := p.FilterTransactionsByToken("BUSD-BD1")
	rawResult, err := json.Marshal(result)
	assert.Nil(t, err)
	assert.Equal(t, wantedTransactionsToken, string(rawResult))
}

func Test_AllowMemo(t *testing.T) {
	type args struct {
		memo string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Numeric memo",
			args{memo: "123"},
			true,
		},
		{
			"Numeric memo",
			args{memo: "12356172321321"},
			true,
		},
		{
			"Numeric memo",
			args{memo: "test"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllowMemo(tt.args.memo); got != tt.want {
				t.Errorf("isMemoAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTxPage_FilterTransactionsByMemo(t *testing.T) {
	tests := []struct {
		name string
		txs  TxPage
		want TxPage
	}{
		{
			name: "Allow memo",
			txs: TxPage{
				{
					Memo: "123",
				},
			},
			want: TxPage{
				{
					Memo: "123",
				},
			},
		},
		{
			name: "Disallow memo",
			txs: TxPage{
				{
					Memo: "test",
				},
			},
			want: TxPage{
				{
					Memo: "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.txs.FilterTransactionsByMemo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterTransactionsByMemo() = %v, want %v", got, tt.want)
			}
		})
	}
}
