package vechain

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `{
  "id": "0xe75d6f28297d910faf31d7aaff9bc57faf14895ffc65da056a90b5a258c17784",
  "chainTag": 74,
  "blockRef": "0x000bae2bae707d76",
  "expiration": 720,
  "clauses": [
    {
      "to": "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
      "value": "0x2b5e3af16b1880000",
      "data": "0x"
    }
  ],
  "gasPriceCoef": 0,
  "gas": 21000,
  "origin": "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
  "delegator": null,
  "nonce": "0x603ca6b1879375dc",
  "dependsOn": null,
  "size": 130,
  "meta": {
    "blockID": "0x000bae2cd3cdea4c79d0d7df3c16f2012b93877eba111803e50d3e0c5f17ed0d",
    "blockNumber": 765484,
    "blockTimestamp": 1537983090
  }
}`

var expectedTransfer = blockatlas.Tx{
	ID:       "0xe75d6f28297d910faf31d7aaff9bc57faf14895ffc65da056a90b5a258c17784",
	Coin:     coin.VET,
	From:     "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
	To:       "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
	Date:     1537983090,
	Type:     blockatlas.TxTransfer,
	Fee:      blockatlas.Amount("21000"),
	Status:   blockatlas.StatusCompleted,
	Block:    765484,
	Sequence: 6934600807657731548,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("50000000000000000000"),
		Decimals: 18,
		Symbol:   "VET",
	},
}

func TestNormalizeTransaction(t *testing.T) {
	tests := []struct {
		name   string
		txData string
		want   blockatlas.TxPage
	}{
		{"test normalize tx", transferSrc, blockatlas.TxPage{expectedTransfer}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tx Tx
			err := json.Unmarshal([]byte(tt.txData), &tx)
			assert.Nil(t, err)
			got, err := NormalizeTransaction(tx)
			assert.Nil(t, err)
			assert.Equal(t, len(got), 1, "tx could not be normalized")
			assert.Equal(t, tt.want, got, "tx don't equal")
		})
	}
}

const transferLogSrc = `{
  "sender": "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
  "recipient": "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
  "amount": "0x38d7ea4c68000",
  "meta": {
    "blockID": "0x000ab61020bcd82739b9956541ccfff3e8aa4f3891e0cbe7d073a0491fc1f87d",
    "blockNumber": 701968,
    "blockTimestamp": 1537342090,
    "txID": "0xc89b224842f3b5edcf4e6950194148be7e83cc951ff4fdfa2490e4c0cf12c80b",
    "txOrigin": "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
    "clauseIndex": 0
  }
}`

var expectedTransferLog = blockatlas.Tx{
	ID:     "0xc89b224842f3b5edcf4e6950194148be7e83cc951ff4fdfa2490e4c0cf12c80b",
	Coin:   coin.VET,
	From:   "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
	To:     "0x15f4d9bed894e2e426d65e3df1480c61fb131a57",
	Date:   1537342090,
	Type:   blockatlas.TxTransfer,
	Fee:    blockatlas.Amount("0"),
	Status: blockatlas.StatusCompleted,
	Block:  701968,
	Meta: blockatlas.Transfer{
		Value:    blockatlas.Amount("1000000000000000"),
		Decimals: 18,
		Symbol:   "VET",
	},
}

func TestNormalizeLogTransaction(t *testing.T) {
	tests := []struct {
		name   string
		txData string
		want   blockatlas.Tx
	}{
		{"test normalize log tx", transferLogSrc, expectedTransferLog},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tx LogTx
			err := json.Unmarshal([]byte(tt.txData), &tx)
			assert.Nil(t, err)
			got, err := NormalizeLogTransaction(tx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got, "tx don't equal")
		})
	}
}

func Test_hexToInt(t *testing.T) {
	tests := []struct {
		name    string
		hex     string
		want    int64
		wantErr bool
	}{
		{"value 1", "0x603ca6b1879375dc", 6934600807657731548, false},
		{"value 2", "0x38d7ea4c68000", 1000000000000000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hexToInt(tt.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("hexToInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("hexToInt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
