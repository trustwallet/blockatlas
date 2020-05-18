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
	bep2MultiTransfer = `
{
    "blockHeight": 80167666,
    "tx": [
        {
            "txHash": "FAD8C1C5E450BE5E0913B12007AAEACC307F8CFFAFFB0844A9F83155E1235C25",
            "blockHeight": 80167666,
            "txType": "TRANSFER",
            "timeStamp": "2020-04-09T20:34:12.922Z",
            "fromAddr": null,
            "toAddr": null,
            "value": null,
            "txAsset": null,
            "txFee": null,
            "code": 0,
            "data": null,
            "memo": "multisend",
            "source": 1,
            "sequence": 72,
            "subTransactions": [
                {
                    "txHash": "FAD8C1C5E450BE5E0913B12007AAEACC307F8CFFAFFB0844A9F83155E1235C25",
                    "blockHeight": 80167666,
                    "txType": "TRANSFER",
                    "fromAddr": "bnb1ds83nt2tz2s9m7kkdcu53t3ccjc07un4xvdld7",
                    "toAddr": "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl",
                    "txAsset": "TWT-8C2",
                    "txFee": "0.29970000",
                    "value": "2800.00000000"
                }
            ]
        }
    ]
}
`
)

var (
	expectBnbSingleRPCV2TransferResponse = ExplorerTxs{
		BlockHeight:        84191216,
		Code:               0,
		FromAddr:           "bnb1mr5f97rx5wnkfcakx9fcpvljmx2s6kwqc08yur",
		HasChildren:        0,
		Memo:               "",
		MultisendTransfers: []MultiTransfer{},
		Timestamp:          1588086357,
		ToAddr:             "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl",
		TxFee:              0.000375,
		TxHash:             "4577CB3B5B202696E9E0B093A6DA973C7DD9CBC6808DA1326872745C35F3C089",
		TxType:             "TRANSFER",
		Value:              0.0004,
		TxAsset:            "BNB",
	}

	expectBEP2MultiRPCV2TransferResponse = ExplorerTxs{
		BlockHeight:        80167666,
		Code:               0,
		HasChildren:        1,
		Memo:               "multisend",
		MultisendTransfers: bep2MultisendTransfer,
		Timestamp:          1586464452,
		TxHash:             "FAD8C1C5E450BE5E0913B12007AAEACC307F8CFFAFFB0844A9F83155E1235C25",
		TxType:             "TRANSFER",
		Value:              0,
	}

	sender   = "bnb1ds83nt2tz2s9m7kkdcu53t3ccjc07un4xvdld7"
	receiver = "bnb14cjy0yl23xkf0hnw3ql295v8nghqstvlzkvqpl"

	bep2MultisendTransfer = []MultiTransfer{
		{Amount: "2800.00000000", Asset: "TWT-8C2", From: sender, To: receiver},
	}
)

func Test_normalizeBlockSubTx(t *testing.T) {
	tests := []struct {
		name       string
		V2Response string
		expected   ExplorerTxs
	}{
		{name: "Normalize single BNB transfer", V2Response: bnbSingleTransfer, expected: expectBnbSingleRPCV2TransferResponse},
		{name: "Normalize multiple BEP2 transfer", V2Response: bep2MultiTransfer, expected: expectBEP2MultiRPCV2TransferResponse},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var blockTxs BlockTransactions
			err := json.Unmarshal([]byte(tt.V2Response), &blockTxs)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, normalizeTxsToExplorer(blockTxs.Txs[0]), "tx don't equal")
		})
	}
}
