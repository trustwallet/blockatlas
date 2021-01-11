package iotex

import (
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/txtype"
)

func TestNormalize(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want []*txtype.Tx
	}{
		{
			name: "Test normalize actions",
			args: args{
				filename: "transfer.json",
			},
			want: []*txtype.Tx{
				{
					ID:       "109b75cb688a5347268cbf11b20fa90fd0a14e92a42ba735c046bbf1a6e66ad7",
					Coin:     coin.IOTX,
					From:     "io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
					To:       "io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
					Fee:      txtype.Amount("10000000000000000"),
					Date:     int64(1556863740),
					Block:    96202,
					Status:   txtype.StatusCompleted,
					Sequence: uint64(3),
					Type:     txtype.TxTransfer,
					Meta: txtype.Transfer{
						Value:    txtype.Amount("21000000000000000000"),
						Symbol:   "IOTX",
						Decimals: 18,
					},
				},
				nil,
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response Response
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &response)
			for i, v := range response.ActionInfo {
				if got := Normalize(v); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Normalize() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
