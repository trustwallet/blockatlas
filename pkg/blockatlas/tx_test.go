package blockatlas

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
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

func TestTxSet_Add(t *testing.T) {
	set := TxSet{}
	set.Add(&transferDst1)
	var txs = set.Txs()
	assert.Equal(t, txs[0].ID, transferDst1.ID)
	set.Add(&transferDst1)
	assert.Equal(t, set.Size(), 1)
	set.Add(&nativeTransferDst1)
	assert.Equal(t, set.Size(), 2)
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

var (
	transferDstOne = Tx{
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

	transferDst2 = Tx{
		ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44556",
		Coin:   coin.BNB,
		From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
		To:     "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
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

	nativeTransferDstOne = Tx{
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

	nativeTransferDst2 = Tx{
		ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4D0",
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
)

func TestGetTxsTx(t *testing.T) {
	tx := Txs{
		transferDstOne,
		transferDst2,
		nativeTransferDstOne,
		nativeTransferDst2,
	}
	txs := tx.GetTransactionsMap()
	assert.Equal(t, len(txs.Map), 4)
	assert.Equal(t, txs.Map["tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"].Size(), 2)
	assert.Equal(t, txs.Map["tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"].Size(), 1)
	assert.Equal(t, txs.Map["tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a"].Size(), 2)
	assert.Equal(t, txs.Map["tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"].Size(), 2)
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
