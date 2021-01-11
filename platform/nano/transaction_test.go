package nano

import (
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/types"
)

func TestNormalize(t *testing.T) {
	platform := Platform{}
	type args struct {
		srcTx   *Transaction
		account string
	}
	tests := []struct {
		name   string
		args   args
		wantTx types.Tx
	}{
		{
			name: "Send",
			args: args{
				srcTx: &Transaction{
					Type:           "send",
					Account:        "nano_1trqphog5noig7z888asnjejcie8z1iopxyepcjdo1atps8whxiuwd51ehbw",
					Amount:         "45000000000000000000000000000",
					LocalTimestamp: "1573006938",
					Height:         "4",
					Hash:           "455C1A3E14E2A645EE4DFA120820614DE3A9876BA2673D0D38494396F3650227",
				},
				account: "nano_3ifzdoxn7keh7tn8zuwty1yr8k5pmaxoq6jpp3jidbjbfnz6hyanh89x6rwj",
			},
			wantTx: types.Tx{
				ID:     "455C1A3E14E2A645EE4DFA120820614DE3A9876BA2673D0D38494396F3650227",
				Coin:   165,
				Date:   1573006938,
				From:   "nano_3ifzdoxn7keh7tn8zuwty1yr8k5pmaxoq6jpp3jidbjbfnz6hyanh89x6rwj",
				To:     "nano_1trqphog5noig7z888asnjejcie8z1iopxyepcjdo1atps8whxiuwd51ehbw",
				Block:  4,
				Status: "completed",
				Fee:    "0",
				Meta: types.Transfer{
					Value:    types.Amount("45000000000000000000000000000"),
					Symbol:   "NANO",
					Decimals: 30,
				},
			},
		},
		{
			name: "Receive",
			args: args{
				srcTx: &Transaction{
					Type:           "receive",
					Account:        "nano_3jwrszth46rk1mu7rmb4rhm54us8yg1gw3ipodftqtikf5yqdyr7471nsg1k",
					Amount:         "90000000000000000000000000000",
					LocalTimestamp: "1570862429",
					Height:         "1",
					Hash:           "5D6B19DE75D8C1BAF7D91FBDA71AFC5F0FED68D483DEC5A51F0767A2384D0DE2",
				},
				account: "nano_3ifzdoxn7keh7tn8zuwty1yr8k5pmaxoq6jpp3jidbjbfnz6hyanh89x6rwj",
			},
			wantTx: types.Tx{
				ID:     "5D6B19DE75D8C1BAF7D91FBDA71AFC5F0FED68D483DEC5A51F0767A2384D0DE2",
				Coin:   165,
				Date:   1570862429,
				From:   "nano_3jwrszth46rk1mu7rmb4rhm54us8yg1gw3ipodftqtikf5yqdyr7471nsg1k",
				To:     "nano_3ifzdoxn7keh7tn8zuwty1yr8k5pmaxoq6jpp3jidbjbfnz6hyanh89x6rwj",
				Block:  1,
				Status: "completed",
				Fee:    "0",
				Meta: types.Transfer{
					Value:    types.Amount("90000000000000000000000000000"),
					Symbol:   "NANO",
					Decimals: 30,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTx, err := platform.Normalize(tt.args.srcTx, tt.args.account); !reflect.DeepEqual(gotTx, tt.wantTx) && err == nil {
				t.Errorf("Normalize() = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}
