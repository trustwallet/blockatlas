package nebulas

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
	}{
		{
			name: "Test normalize transaction",
			args: args{
				filename: "transfer.json",
			},
			wantTx: types.Tx{
				ID:       "96bd280d60447b7dbcdb3fa76a99856e0422a76304e9d01d0c87e1dfceb6d952",
				Coin:     coin.NEBULAS,
				From:     "n1Yv9xJJcH4UjoJPVDGdUCL2CxK29asFuyV",
				To:       "n1TFrmLUDTe5ggQaWJiXHSqNSRzKYdaV6hQ",
				Fee:      "400000000000000",
				Sequence: 7,
				Date:     1565213205,
				Block:    2848548,
				Status:   types.StatusCompleted,
				Meta: types.Transfer{
					Value:    "500000000000000000",
					Symbol:   "NAS",
					Decimals: 18,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Transaction
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx := NormalizeTx(srcTx)
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
