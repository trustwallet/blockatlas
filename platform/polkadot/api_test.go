package polkadot

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestNormalize(t *testing.T) {
	platform := Platform{}
	type args struct {
		srcTx *Transfer
	}
	tests := []struct {
		name   string
		args   args
		wantTx blockatlas.Tx
	}{
		{
			name: "Receive 1",
			args: args{
				srcTx: &Transfer{
					From:           "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
					To:             "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
					Amount:         "0.01",
					Module:         "balances",
					Hash:           "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
					BlockTimestamp: 1577176992,
					BlockNum:       360298,
					Success:        true,
				},
			},
			wantTx: blockatlas.Tx{
				ID:     "0x20cfbba19817e4b7a61e718d269de47e7067a24860fa978c2a8ead4c96a827c4",
				Coin:   459,
				Date:   1577176992,
				From:   "HKtMPUSoTC8Hts2uqcQVzPAuPRpecBt4XJ5Q1AT1GM3tp2r",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  360298,
				Status: "completed",
				Fee:    "100000000",
				Meta: blockatlas.Transfer{
					Value:    blockatlas.Amount("10000000000"),
					Symbol:   "KSM",
					Decimals: 12,
				},
			},
		},
		{
			name: "Receive 2",
			args: args{
				srcTx: &Transfer{
					From:           "DbCNECPna3k6MXFWWNZa5jGsuWycqEE6zcUxZYkxhVofrFk",
					To:             "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
					Amount:         "210",
					Module:         "balances",
					Hash:           "0xe0be47f6ce0e62a218f197dd68599989376ee7e951d54c3c6146e2c5c5eacd1f",
					BlockTimestamp: 1575525432,
					BlockNum:       90672,
					Success:        true,
				},
			},
			wantTx: blockatlas.Tx{
				ID:     "0xe0be47f6ce0e62a218f197dd68599989376ee7e951d54c3c6146e2c5c5eacd1f",
				Coin:   90672,
				Date:   1575525432,
				From:   "DbCNECPna3k6MXFWWNZa5jGsuWycqEE6zcUxZYkxhVofrFk",
				To:     "CtwdfrhECFs3FpvCGoiE4hwRC4UsSiM8WL899HjRdQbfYZY",
				Block:  90672,
				Status: "completed",
				Fee:    "100000000",
				Meta: blockatlas.Transfer{
					Value:    blockatlas.Amount("210000000000000"),
					Symbol:   "KSM",
					Decimals: 12,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTx := platform.NormalizeTransfer(tt.args.srcTx); !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("Normalize() = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
