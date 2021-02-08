package near

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

func TestNormalizeChunk(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     types.TxPage
	}{
		{
			name:     "Test normalize transfer",
			filename: "mocks/chunk_transfer.json",
			want: types.TxPage{
				types.Tx{
					ID:       "HnFPnVX9xpF8AGmkDqV7iNPfMN14pifmgTRjvkLuCmoM",
					Coin:     coin.NEAR,
					From:     "a7e956ffba0f7a1905445d107de74e32e8e84b31c6171caae381e6a1613e1b50",
					To:       "hewig.near",
					Fee:      "0",
					Date:     0,
					Block:    29445373,
					Status:   types.StatusCompleted,
					Sequence: 8,
					Type:     types.TxTransfer,
					Meta: types.Transfer{
						Value:    "100000000000000000000000",
						Symbol:   coin.Near().Name,
						Decimals: coin.Near().Decimals,
					},
				}},
		},
		{
			name:     "Test normalize funcation call",
			filename: "mocks/chunk_call.json",
			want:     types.TxPage{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chunk ChunkDetail
			err := mock.JsonModelFromFilePath(tt.filename, &chunk)
			assert.NoError(t, err)
			if got := NormalizeChunk(chunk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NormalizeChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
