package zilliqa

import (
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/platform/zilliqa/viewblock"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

func TestNormalizeTx(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantTx  types.Tx
		wantErr bool
	}{
		{
			name: "Test normalize transaction",
			args: args{
				filename: "transfer.json",
			},
			wantTx: types.Tx{
				ID:       "0xd44413c79e7518152f3b05ef1edff8ef59afd06119b16d09c8bc72e94fed7843",
				Coin:     coin.ZILLIQA,
				From:     "0x88af5ba10796d9091d6893eed4db23ef0bbbca37",
				To:       "0x7fccacf066a5f26ee3affc2ed1fa9810deaa632c",
				Fee:      "1000000000",
				Date:     1557889788,
				Block:    104282,
				Status:   types.StatusCompleted,
				Sequence: 3,
				Memo:     "",
				Meta: types.Transfer{
					Value:    "7997000000000",
					Symbol:   "ZIL",
					Decimals: 12,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx viewblock.Tx
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx := Normalize(&srcTx)
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
