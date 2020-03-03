package observer

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
	"time"
)

var transferDst1 = blockatlas.Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	Fee:    "125000",
	Date:   1555049867,
	Block:  7761368,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.Transfer{
		Value:    "10000000000000",
		Decimals: 8,
		Symbol:   "BNB",
	},
}

var transferDst2 = blockatlas.Tx{
	ID:     "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44556",
	Coin:   coin.BNB,
	From:   "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To:     "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	Fee:    "125000",
	Date:   1555049867,
	Block:  7761368,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.Transfer{
		Value:    "10000000000000",
		Decimals: 8,
		Symbol:   "BNB",
	},
}

var nativeTransferDst1 = blockatlas.Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4C9",
	Coin:   coin.BNB,
	From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	Fee:    "125000",
	Date:   1555117625,
	Block:  7928667,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.NativeTokenTransfer{
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	},
}

var nativeTransferDst2 = blockatlas.Tx{
	ID:     "95CF63FAA27579A9B6AF84EF8B2DFEAC29627479E9C98E7F5AE4535E213FA4D0",
	Coin:   coin.BNB,
	From:   "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
	To:     "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	Fee:    "125000",
	Date:   1555117625,
	Block:  7928667,
	Status: blockatlas.StatusCompleted,
	Memo:   "test",
	Meta: blockatlas.NativeTokenTransfer{
		TokenID:  "YLC-D8B",
		Symbol:   "YLC",
		Value:    "210572645",
		Decimals: 8,
		From:     "tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a",
		To:       "tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex",
	},
}

var txsBlock = blockatlas.Block{
	Number: 12345,
	ID:     "12345",
	Txs: []blockatlas.Tx{
		transferDst1,
		transferDst2,
		nativeTransferDst1,
		nativeTransferDst2,
	},
}

func TestGetTxs(t *testing.T) {
	txs := GetTxs(&txsBlock)
	assert.Equal(t, len(txs), 4)
	assert.Equal(t, txs["tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2"].Size(), 2)
	assert.Equal(t, txs["tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5"].Size(), 1)
	assert.Equal(t, txs["tbnb1ttyn4csghfgyxreu7lmdu3lcplhqhxtzced45a"].Size(), 2)
	assert.Equal(t, txs["tbnb12hlquylu78cjylk5zshxpdj6hf3t0tahwjt3ex"].Size(), 2)
}

func Test_getDirection(t *testing.T) {
	type args struct {
		tx      blockatlas.Tx
		address string
	}
	tests := []struct {
		name string
		args args
		want blockatlas.Direction
	}{
		{"Test Direction Self",
			args{
				blockatlas.Tx{
					From: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1", To: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
				"0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"}, blockatlas.DirectionSelf,
		},
		{"Test Direction Outgoing",
			args{
				blockatlas.Tx{
					From: "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB", To: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7"},
				"0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB"}, blockatlas.DirectionOutgoing,
		},
		{"Test Direction Incoming",
			args{
				blockatlas.Tx{
					From: "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7", To: "0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"},
				"0xfc10cab6a50a1ab10c56983c80cc82afc6559cf1"}, blockatlas.DirectionIncoming,
		},
		{"Test UTXO Direction Self",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "72934112534"},
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "500000000"},
					},
					Inputs: []blockatlas.TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "73196112534"},
					},
				}, "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S",
			}, blockatlas.DirectionSelf,
		},
		{"Test UTXO Direction Outgoing",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd",
			}, blockatlas.DirectionOutgoing,
		},
		{"Test UTXO Direction Incoming",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4",
			}, blockatlas.DirectionIncoming,
		},
		{"Test NativeTokenTransfer Direction Self",
			args{
				blockatlas.Tx{
					Meta: blockatlas.NativeTokenTransfer{
						From:     "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:       "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, blockatlas.DirectionSelf,
		},
		{"Test NativeTokenTransfer Direction Outgoing",
			args{
				blockatlas.Tx{
					Meta: blockatlas.NativeTokenTransfer{
						From:     "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:       "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, blockatlas.DirectionOutgoing,
		},
		{"Test NativeTokenTransfer Direction Incoming",
			args{
				blockatlas.Tx{
					Meta: blockatlas.NativeTokenTransfer{
						From:     "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
						To:       "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, blockatlas.DirectionIncoming,
		},
		{"Test TokenTransfer Direction Self",
			args{
				blockatlas.Tx{
					Meta: blockatlas.TokenTransfer{
						From:     "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:       "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, blockatlas.DirectionSelf,
		},
		{"Test TokenTransfer Direction Outgoing",
			args{
				blockatlas.Tx{
					Meta: blockatlas.TokenTransfer{
						From:     "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
						To:       "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, blockatlas.DirectionOutgoing,
		},
		{"Test TokenTransfer Direction Incoming",
			args{
				blockatlas.Tx{
					Meta: blockatlas.TokenTransfer{
						From:     "0x74c8199372c584DAB8b14c519bc8BC8C622F37b7",
						To:       "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
					},
				}, "0x5574Cd97432cEd0D7Caf58ac3c4fEDB2061C98fB",
			}, blockatlas.DirectionIncoming,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirection(tt.args.tx, tt.args.address); got != tt.want {
				t.Errorf("getDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inferUtxoValue(t *testing.T) {
	type args struct {
		tx        blockatlas.Tx
		address   string
		coinIndex uint
	}
	tests := []struct {
		name       string
		args       args
		wantAmount blockatlas.Amount
	}{
		{"Test UTXO Direction Self",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "72934112534"},
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "500000000"},
					},
					Inputs: []blockatlas.TxOutput{
						{Address: "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", Value: "73196112534"},
					},
				}, "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", 3,
			}, blockatlas.Amount("72934112534"),
		},
		{"Test UTXO Direction Outgoing",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", 0,
			}, blockatlas.Amount("4471835"),
		},
		{"Test UTXO Direction Incoming",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{Address: "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", Value: "4471835"},
						{Address: "324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", Value: "1600000"},
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{Address: "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", Value: "1268998877"},
					},
				}, "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", 0,
			}, blockatlas.Amount("4471835"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expect := blockatlas.Transfer{
				Value:    tt.wantAmount,
				Symbol:   coin.Coins[tt.args.coinIndex].Symbol,
				Decimals: coin.Coins[tt.args.coinIndex].Decimals,
			}
			tt.args.tx.Direction = getDirection(tt.args.tx, tt.args.address)
			if inferUtxoValue(&tt.args.tx, tt.args.address, tt.args.coinIndex); tt.args.tx.Meta != expect {
				t.Errorf("inferUtxoValue() = %v, want %v", tt.args.tx.Meta, expect)
			}
		})
	}
}

func TestGetInterval(t *testing.T) {
	min, _ := time.ParseDuration("2s")
	max, _ := time.ParseDuration("30s")
	type args struct {
		blockTime   int
		minInterval time.Duration
		maxInterval time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			"test minimum",
			args{
				blockTime:   100,
				minInterval: min,
				maxInterval: max,
			},
			min,
		}, {
			"test maximum",
			args{
				blockTime:   600000,
				minInterval: min,
				maxInterval: max,
			},
			max,
		}, {
			"test right blocktime",
			args{
				blockTime:   5000,
				minInterval: min,
				maxInterval: max,
			},
			5000 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetInterval(tt.args.blockTime, tt.args.minInterval, tt.args.maxInterval)
			assert.EqualValues(t, tt.want, got)
		})
	}
}
