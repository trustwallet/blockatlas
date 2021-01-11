package aion

import (
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

func TestNormalizeTx(t *testing.T) {
	type args struct {
		srcTx string
	}
	tests := []struct {
		name   string
		args   args
		wantTx types.Tx
		wantOk bool
	}{
		{
			name: "Test normalize transfer",
			args: args{
				srcTx: "transfer.json",
			},
			wantTx: types.Tx{
				ID:     "0xaf3c2f5087fc3332154dc9d11c27e312f30ff829dbc5436aec8cc4342c7dc384",
				Coin:   coin.AION,
				From:   "0xa07981da70ce919e1db5f051c3c386eb526e6ce8b9e2bfd56e3f3d754b0a17f3",
				To:     "0xa09b8c4c40bd7a81e969b8f6f291074206196a99948b03c6a469892931a3c258",
				Fee:    "21000",
				Date:   1554862228,
				Block:  2880919,
				Status: types.StatusCompleted,
				Meta: types.Transfer{
					Value:    "11903810405853733000",
					Symbol:   "AION",
					Decimals: 18,
				},
			},
			wantOk: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Tx
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.srcTx, &srcTx)
			gotTx, gotOk := NormalizeTx(&srcTx)
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
			if gotOk != tt.wantOk {
				t.Errorf("NormalizeTx() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
