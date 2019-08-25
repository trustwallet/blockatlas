package nebulas

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
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
		Value: "500000000000000000",
	},
}

const blockSrc = `{
    "result":{
        "hash":"c4a51d6241db372c1b8720e62c04426bd587e1f31054b7d04a3509f48ee58e9f",
        "nonce":"0",
        "height":"407",
        
        "transactions":[{
            "hash":"1e96493de6b5ebe686e461822ec22e73fcbfb41a6358aa58c375b935802e4145",
            "chainId":"100",
            "from":"n1Z6SbjLuAEXfhX1UJvXT6BB5osWYxVg3F3",
            "to":"n1orSeSMj7nn8KHHN4JcQEw3r52TVExu63r",
            "value":"10000000000000000000",
			"nonce":"34",
            "timestamp":"1522220087",
            "type":"binary",
            "data":null,
            "gas_price":"1000000",
            "gas_limit":"2000000",
            "contract_address":"",
            "status":1,
            "gas_used":"20000"
        }]
    }
}`

var tnxDst = blockatlas.Tx{
	ID:       "1e96493de6b5ebe686e461822ec22e73fcbfb41a6358aa58c375b935802e4145",
	Coin:     coin.NAS,
	From:     "n1Z6SbjLuAEXfhX1UJvXT6BB5osWYxVg3F3",
	To:       "n1orSeSMj7nn8KHHN4JcQEw3r52TVExu63r",
	Fee:      "20000000000",
	Sequence: 34,
	Date:     1522220087,
	Block:    407,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value: "10000000000000000000",
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
