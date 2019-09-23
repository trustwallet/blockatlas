package nebulas

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `
{
  "hash": "96bd280d60447b7dbcdb3fa76a99856e0422a76304e9d01d0c87e1dfceb6d952",
  "block": {
    "height": 2848548
  },
  "from": {
    "hash": "n1Yv9xJJcH4UjoJPVDGdUCL2CxK29asFuyV"
  },
  "to": {
    "hash": "n1TFrmLUDTe5ggQaWJiXHSqNSRzKYdaV6hQ"
  },
  "value": "500000000000000000",
  "nonce": 7,
  "status": 1,
  "timestamp": 1565213205000,
  "type": "binary",
  "currentTimestamp": 1565361175536,
  "txFee": "400000000000000"
}`

var transferDst = blockatlas.Tx{
	ID:       "96bd280d60447b7dbcdb3fa76a99856e0422a76304e9d01d0c87e1dfceb6d952",
	Coin:     coin.NAS,
	From:     "n1Yv9xJJcH4UjoJPVDGdUCL2CxK29asFuyV",
	To:       "n1TFrmLUDTe5ggQaWJiXHSqNSRzKYdaV6hQ",
	Fee:      "400000000000000",
	Sequence: 7,
	Date:     1565213205,
	Block:    2848548,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "500000000000000000",
		Symbol:   "NAS",
		Decimals: 18,
	},
}

func TestNormalize(t *testing.T) {
	var srcTx Transaction
	err := json.Unmarshal([]byte(transferSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	resTx := NormalizeTx(srcTx)

	resJSON, err := json.Marshal(&resTx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("tx don't equal")
	}
}

func TestNormalizeTxs(t *testing.T) {
	txs := []Transaction{
		Transaction{},
		Transaction{},
	}

	assert.Equal(t, 2, len(NormalizeTxs(txs)))
}
