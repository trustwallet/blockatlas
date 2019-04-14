package binance

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"testing"
)

const transferSrc = `
{
	"txHash": "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	"blockHeight": 7761368,
	"txType": "TRANSFER",
	"timeStamp": 1555049867552,
	"fromAddr": "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	"toAddr": "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	"value": 100000,
	"txAsset": "BNB",
	"mappedTxAsset": "BNB",
	"txFee": 0.00125,
	"txAge": 222901,
	"code": 0,
	"log": "Msg 0: ",
	"confirmBlocks": 0,
	"hasChildren": 0
}`

var transferDst = models.Tx{
	Id: "1681EE543FB4B5A628EF21D746E031F018E226D127044A4F9BA5EE2542A44555",
	Coin: coin.BNB,
	From: "tbnb1fhr04azuhcj0dulm7ka40y0cqjlafwae9k9gk2",
	To: "tbnb1sylyjw032eajr9cyllp26n04300qzzre38qyv5",
	Fee: "125",
	Date: 1555049867,
	Block: 7761368,
	Status: models.StatusCompleted,
	Meta: models.Transfer{
		Value: "10000000000",
	},
}


func TestNormalize(t *testing.T) {
	var srcTx Tx
	err := json.Unmarshal([]byte(transferSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx, ok := Normalize(&srcTx)
	if !ok {
		t.Errorf("transfer: tx could not be normalized")
	}

	resJson, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		println(string(resJson))
		println(string(dstJson))
		t.Error("transfer: tx don't equal")
	}
}

