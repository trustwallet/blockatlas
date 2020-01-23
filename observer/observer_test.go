package observer

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

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
						{"DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", "72934112534"},
						{"DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", "500000000"},
					},
					Inputs: []blockatlas.TxOutput{
						{"DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", "73196112534"},
					},
				}, "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S",
			}, blockatlas.DirectionSelf,
		},
		{"Test UTXO Direction Outgoing",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{"3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", "4471835"},
						{"324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", "1600000"},
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1268998877"},
					},
				}, "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd",
			}, blockatlas.DirectionOutgoing,
		},
		{"Test UTXO Direction Incoming",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{"3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", "4471835"},
						{"324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", "1600000"},
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1268998877"},
					},
				}, "3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4",
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
						{"DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", "72934112534"},
						{"DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", "500000000"},
					},
					Inputs: []blockatlas.TxOutput{
						{"DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", "73196112534"},
					},
				}, "DAzruJfMBhd1vcQ13gVVyqb2g1vSEo2d5S", 3,
			}, blockatlas.Amount("72934112534"),
		},
		{"Test UTXO Direction Outgoing",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{"3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", "4471835"},
						{"324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", "1600000"},
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1268998877"},
					},
				}, "32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", 0,
			}, blockatlas.Amount("4471835"),
		},
		{"Test UTXO Direction Incoming",
			args{
				blockatlas.Tx{
					Outputs: []blockatlas.TxOutput{
						{"3BMEXVshYmWqc8qcQLyBQPgRgAPfogWdJ4", "4471835"},
						{"324Wmkkbr9uT9xnLUqFvCA3ntqqpqEZQDp", "1600000"},
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1262899630"},
					},
					Inputs: []blockatlas.TxOutput{
						{"32yRH5tNnFtAXE844wNrHN7Bf3SBcb3Uhd", "1268998877"},
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
