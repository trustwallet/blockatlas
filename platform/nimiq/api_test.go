package nimiq

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/nimiq/source"
	"testing"
)

const basicSrc = `
{
	"hash": "8b219949f4c1dfe9e7a9cdc5dbbc507e40dc16f44a1a5182ed6125c9a6891a50",
	"blockHash": "ab36a0909c6ed5761a984ef261d9c3456b7c1aea6a52d531c5bf2518526a32e6",
	"blockNumber": 252575,
	"timestamp": 1538924505,
	"confirmations": 271245,
	"transactionIndex": 37,
	"from": "4a88aaad038f9b8248865c4b9249efc554960e16",
	"fromAddress": "NQ69 9A4A MB83 HXDQ 4J46 BH5R 4JFF QMA9 C3GN",
	"to": "ad25610feb43d75307763d3f010822a757027429",
	"toAddress": "NQ15 MLJN 23YB 8FBM 61TN 7LYG 2212 LVBG 4V19",
	"value": 10000000000000,
	"fee": 138,
	"data": null,
	"flags": 0
}
`

var basicDst = models.Tx{
	Id:    "8b219949f4c1dfe9e7a9cdc5dbbc507e40dc16f44a1a5182ed6125c9a6891a50",
	Coin:  coin.NIM,
	From:  "NQ69 9A4A MB83 HXDQ 4J46 BH5R 4JFF QMA9 C3GN",
	To:    "NQ15 MLJN 23YB 8FBM 61TN 7LYG 2212 LVBG 4V19",
	Fee:   "138",
	Date:  1538924505,
	Block: 252575,
	Meta: models.Transfer{
		Value: "10000000000000",
	},
}

func TestNormalize(t *testing.T) {
	var srcTx source.Tx
	err := json.Unmarshal([]byte(basicSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx := Normalize(&srcTx)
	resJson, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(&basicDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		println(string(resJson))
		println(string(dstJson))
		t.Error("basic: tx don't equal")
	}
}
