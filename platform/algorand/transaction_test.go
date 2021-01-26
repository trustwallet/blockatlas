package algorand

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
				ID:     "C2LK3CGBPIGERLPFUXE6INSBJGHOXU7YZMEGELWMVSBASFJYOOQQ",
				Coin:   coin.ALGORAND,
				From:   "5TSQNIL54GB545B3WLC6OVH653SHAELMHU6MSVNGTUNMOEHAMWG7EC3AA4",
				To:     "4EZFQABCVQTHQCK3HQBIYGC4NV2VM42FZHEFTVH77ROG4ZGREC6Y7V5T2U",
				Fee:    types.Amount("1000"),
				Date:   1569123058,
				Block:  2031351,
				Status: types.StatusCompleted,
				Type:   types.TxTransfer,
				Meta: types.Transfer{
					Value:    types.Amount("1"),
					Symbol:   "ALGO",
					Decimals: 6,
				},
			},
			ok: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response TransactionsResponse
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &response)
			gotTx, ok := Normalize(response.Transactions[0])
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
