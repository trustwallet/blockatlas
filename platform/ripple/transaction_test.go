package ripple

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
			name: "Test normalize payment",
			args: args{
				filename: "payment.json",
			},
			wantTx: types.Tx{
				ID:     "40279A3DE51148BD41409DADF29DE8DCCD50F5AEE30840827B2C4C81C4E36505",
				Coin:   coin.RIPPLE,
				From:   "rGSxFjoqmWz54PycrgQBQ5dB6e7TUpMxzq",
				To:     "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
				Fee:    "3115",
				Date:   1512168330,
				Block:  34698103,
				Memo:   "2500",
				Status: types.StatusCompleted,
				Meta: types.Transfer{
					Value:    "100000000",
					Symbol:   "XRP",
					Decimals: 6,
				},
			},
			ok: true,
		},
		{
			name: "Test normalize payment 2",
			args: args{
				filename: "payment_2.json",
			},
			wantTx: types.Tx{
				ID:     "3D8512E02414EF5A6BC00281D945735E85DED9EF739B1DCA9EABE04D9EEC72C1",
				Coin:   coin.RIPPLE,
				From:   "raz97dHvnyBcnYTbXGYxhV8bGyr1aPrE5w",
				To:     "rna8qC8Y9uLd2vzYtSEa1AJcdD3896zQ9S",
				Fee:    "120",
				Date:   1565114281,
				Block:  49163909,
				Memo:   "",
				Status: types.StatusCompleted,
				Meta: types.Transfer{
					Value:    "3100",
					Symbol:   "XRP",
					Decimals: 6,
				},
			},
			ok: true,
		},
		{
			name: "Test normalize failed payment",
			args: args{
				filename: "payment_failed.json",
			},
			wantTx: types.Tx{
				ID:     "B9086F7EB895E943C4DDA9F1B582E6E7DE35F4FB91AD13C50AB74F854DC0EBE0",
				Coin:   coin.RIPPLE,
				From:   "rJb5KsHsDHF1YS5B5DU6QCkH5NsPaKQTcy",
				To:     "rfHj5CuhajwdrzW2C8Y7EDXbx1QMiD5SXP",
				Fee:    "100000",
				Date:   1581590872,
				Block:  53401154,
				Memo:   "",
				Status: types.StatusError,
				Meta: types.Transfer{
					Value:    "24999750000",
					Symbol:   "XRP",
					Decimals: 6,
				},
			},
			ok: true,
		},
		{
			name: "Test normalize SetRegularKey transfer",
			args: args{
				filename: "payment_3.json",
			},
			wantTx: types.Tx{},
			ok:     false,
		},
		{
			name: "Test normalize token transfer",
			args: args{
				filename: "payment_4.json",
			},
			wantTx: types.Tx{},
			ok:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Tx
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx, err := NormalizeTx(&srcTx)
			if err != tt.ok {
				t.Errorf("NormalizeTx() error = %v, ok %v", err, tt.ok)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
