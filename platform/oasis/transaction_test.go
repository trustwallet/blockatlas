package oasis

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
	"reflect"
	"testing"
)

func TestNormalizeTx(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name   string
		args   args
		wantTx types.Tx
		ok     bool
	}{
		{
			name: "Test normalize transaction 1",
			args: args{
				filename: "tx_1.json",
			},
			wantTx: types.Tx{
				ID:     "7DYN6sCpXcHuOqomRoRp63mzOinTnSJEqIX+dMtrJWQ=",
				Coin:   coin.ROSE,
				From:   "oasis1qz9re9hc0k9qxrhvww7x9zrfv8x8jpr4kcr2twr2",
				To:     "oasis1qp29h8ykmxet46eqzw0wennrmmy4al3xzv37m3ca",
				Fee:    "0",
				Date:   1605718103,
				Block:  702480,
				Memo:   "",
				// TODO: check if add status field
				//Status: types.StatusCompleted,
				// TODO: check if hardcode
				Meta: types.Transfer{
					Value:    "0",
					Symbol:   "ROSE",
					Decimals: 9,
				},
			},
			ok: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: add error cases on Normalize fn
			var srcTx Transaction
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx := NormalizeTx(srcTx)
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, wantTx %v", gotTx, tt.wantTx)
			}
			t.Log(tt.wantTx)
			t.Log(gotTx)
		})
	}
}
