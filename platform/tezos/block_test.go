package tezos

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

func TestNormalizeRpcBlock(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *types.Block
		wantErr bool
	}{
		{
			name: "Test normalize block 1292516",
			args: args{
				filename: "rpc_block_operations.json",
			},
			want: &types.Block{
				Number: 1292516,
				Txs: []types.Tx{
					{
						ID:       "oo25sEdAT3YDb83WNdMSxRv4E6V2Rt6Jc8msgTio7R4FBnAiFmj",
						Coin:     1729,
						From:     "tz1SiPXX4MYGNJNDsRc7n8hkvUqFzg8xqF9m",
						To:       "tz1LGrjCS3Jj3YZyRx3mHMmEXRJVQeVnoYYi",
						Fee:      "2940",
						Date:     1610065014,
						Block:    1292516,
						Status:   "completed",
						Sequence: 0,
						Type:     "transfer",
						Memo:     "",
						Meta: types.Transfer{
							Value:    "5278500000",
							Symbol:   "XTZ",
							Decimals: 6,
						},
					},
					{
						ID:       "ooFFpMbVAJMkxNYFqxERDumptXS2kSEJQqTwoovVPinayJehB8f",
						Coin:     1729,
						From:     "tz1YbzpfsfRtSiYJWcvpWEDHPNe3kkqEvY56",
						To:       "tz1L1c5YoQMaciYGB8gpzWoHbscBmtsTsknF",
						Fee:      "30000",
						Date:     1610065014,
						Block:    1292516,
						Status:   "completed",
						Sequence: 0,
						Type:     "transfer",
						Memo:     "",
						Meta: types.Transfer{
							Value:    "290691675",
							Symbol:   "XTZ",
							Decimals: 6,
						},
					},
					{
						ID:       "opNgBb87Cnpjr29Uc4CRuXpWgTfhPx2h76S5txrPJELa5XYiidZ",
						Coin:     1729,
						From:     "tz1c5wM9826YcUNQ8a17z9eUYpKQ3oW3zfmJ",
						To:       "KT19kgnqC5VWoxktLRdRUERbyUPku9YioE8W",
						Fee:      "824",
						Date:     1610065014,
						Block:    1292516,
						Status:   "completed",
						Sequence: 0,
						Type:     "transfer",
						Memo:     "",
						Meta: types.Transfer{
							Value:    "0",
							Symbol:   "XTZ",
							Decimals: 6,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		var block RpcBlock
		_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &block)
		t.Run(tt.name, func(t *testing.T) {
			got, err := NormalizeRpcBlock(block)
			fmt.Println(got)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeRpcBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalizeRpcBlock() = %v, \nwant %v", got, tt.want)
			}
		})
	}
}
