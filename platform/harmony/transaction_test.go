package harmony

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
		name    string
		args    args
		wantTx  types.Tx
		wantB   bool
		wantErr bool
	}{
		{
			name: "Test normalize transaction",
			args: args{
				filename: "transfer.json",
			},
			wantTx: types.Tx{
				ID:     "0x230798fe22abff459b004675bf827a4089326a296fa4165d0c2ad27688e03e0c",
				Coin:   coin.HARMONY,
				From:   "one103q7qe5t2505lypvltkqtddaef5tzfxwsse4z7",
				To:     "one129r9pj3sk0re76f7zs3qz92rggmdgjhtwge62k",
				Fee:    "21000000000000",
				Date:   1576346446,
				Block:  18,
				Type:   types.TxTransfer,
				Status: types.StatusCompleted,
				Meta: types.Transfer{
					Value:    "100000000000000000",
					Symbol:   "ONE",
					Decimals: 18,
				},
			},
			wantB:   true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Transaction
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx, gotB, err := NormalizeTx(&srcTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
			if gotB != tt.wantB {
				t.Errorf("NormalizeTx() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
