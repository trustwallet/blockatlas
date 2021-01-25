package icon

import (
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
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
			name: "Test normalize transaction",
			args: args{
				filename: "transfer.json",
			},
			wantTx: types.Tx{
				ID:     "0x34b8b6ec3a52710c24074f5e298f4a9c67bb61a0a1dde20e695efaeb30ff3754",
				Coin:   coin.ICON,
				From:   "hx1b8959dd5c57d2c502e22ee0a887d33baec09091",
				To:     "cx334db6519871cb2bfd154cec0905ced4ea142de1",
				Fee:    "1747600000000000",
				Date:   1555396594,
				Block:  357832,
				Status: "completed",
				Type:   "transfer",
				Meta: types.Transfer{
					Value:    "3470000000000000",
					Symbol:   "ICX",
					Decimals: 18,
				},
			},
			ok: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Tx
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx, ok := Normalize(&srcTx)
			if ok != tt.ok {
				t.Errorf("Normalize() ok = %v, wantOk %v", ok, tt.ok)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("Normalize() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
