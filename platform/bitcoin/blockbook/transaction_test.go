package blockbook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	srcPage, _ = mock.JsonStringFromFilePath("mocks/blockbook_txs.json")
	want, _    = mock.JsonStringFromFilePath("mocks/expected_txs.json")
)

func TestNormalizeTxs(t *testing.T) {
	type args struct {
		srcPage   string
		address   string
		token     string
		coinIndex uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test normalize blockbook txs",
			args: args{
				srcPage:   srcPage,
				address:   "0x7d8bf18c7ce84b3e175b339c4ca93aed1dd166f1",
				token:     "",
				coinIndex: 60,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		var page TransactionsList
		var txs types.Txs
		err := json.Unmarshal([]byte(tt.args.srcPage), &page)
		assert.Nil(t, err)
		err = json.Unmarshal([]byte(tt.want), &txs)
		assert.Nil(t, err)
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePage(page, tt.args.address, tt.args.token, tt.args.coinIndex)
			gotJson, err := json.Marshal(got)
			assert.Nil(t, err)
			wantTxs, err := json.Marshal(txs)
			assert.Nil(t, err)
			assert.JSONEq(t, string(gotJson), string(wantTxs))
		})
	}
}
