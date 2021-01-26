package aeternity

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
				ID:       "th_oJfBC6KZKaKsL4WXTq1ZtFiSE8Wp2PQYEnwyZqtudyHcU3Qg6",
				Coin:     coin.AETERNITY,
				From:     "ak_nv5B93FPzRHrGNmMdTDfGdd5xGZvep3MVSpJqzcQmMp59bBCv",
				To:       "ak_ZWrS6xGhzxBasKmMbVSACfRioWqPyM5jNqMpBQ5ngP75RS6pS",
				Fee:      "20500000000000",
				Date:     1563848658,
				Block:    113579,
				Status:   types.StatusCompleted,
				Memo:     "Hello, Miner! /Yours Beepool./",
				Sequence: 251291,
				Meta: types.Transfer{
					Value:    "252550000000000000000",
					Symbol:   "AE",
					Decimals: 18,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var srcTx Transaction
			_ = mock.JsonModelFromFilePath("mocks/"+tt.args.filename, &srcTx)
			gotTx, err := NormalizeTx(&srcTx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NormalizeTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTx, tt.wantTx) {
				t.Errorf("NormalizeTx() gotTx = %v, want %v", gotTx, tt.wantTx)
			}
		})
	}
}

func TestPayloadEncoding(t *testing.T) {
	assert.Equal(t, getPayload("ba_SGVsbG8sIE1pbmVyISAvWW91cnMgQmVlcG9vbC4vKXcQag=="), "Hello, Miner! /Yours Beepool./")
	assert.Equal(t, getPayload("xvass-///BadEncoding///Test"), "")
	assert.Equal(t, getPayload(""), "")
}
