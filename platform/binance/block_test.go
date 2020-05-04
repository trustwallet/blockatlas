package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	bnbSingleTransfer = `
{
    "blockHeight": 84191216,
    "tx": [
        {
            "txHash": "4577CB3B5B202696E9E0B093A6DA973C7DD9CBC6808DA1326872745C35F3C089",
            "blockHeight": 84191216,
            "txType": "TRANSFER",
            "timeStamp": "2020-04-28T15:05:57.686Z",
            "fromAddr": "bnb1mr5f97rx5wnkfcakx9fcpvljmx2s6kwqc08yur",
            "toAddr": "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl",
            "value": "0.00040000",
            "txAsset": "BNB",
            "txFee": "0.00037500",
            "code": 0,
            "data": null,
            "memo": "",
            "source": 0,
            "sequence": 95
        }
    ]
}`
)

var (
	expectBnbSingleRPCV2TransferResponse = DexTx{
		BlockHeight: 84191216,
		Code: 0,
		FromAddr: "bnb1mr5f97rx5wnkfcakx9fcpvljmx2s6kwqc08yur",
		HasChildren: 0,
		Memo: "",
		Timestamp: 0,
		ToAddr: "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl",
		TxFee: 0.000375,
		TxHash: "4577CB3B5B202696E9E0B093A6DA973C7DD9CBC6808DA1326872745C35F3C089",
		TxType: "TRANSFER",
		Value: 0.0004,
		TxAsset: "BNB",
	}
)

func Test_normalizeBlockSubTx(t *testing.T) {
	tests := []struct {
		name, V2Response string
		expected         DexTx
	}{
		{name: "Should normalize RPC trx response v2 to dex tr", V2Response: bnbSingleTransfer, expected: expectBnbSingleRPCV2TransferResponse},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var blockTxs BlockTransactions
			err := json.Unmarshal([]byte(tt.V2Response), &blockTxs)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, normalizeBlockSubTx(&blockTxs.Txs[0]), "tx don't equal")
		})
	}
}
