package icon

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const basicSrc = `
{
	"txHash": "0x34b8b6ec3a52710c24074f5e298f4a9c67bb61a0a1dde20e695efaeb30ff3754",
	"height": 357832,
	"createDate": "2019-04-16T06:36:34.000+0000",
	"fromAddr": "hx1b8959dd5c57d2c502e22ee0a887d33baec09091",
	"toAddr": "cx334db6519871cb2bfd154cec0905ced4ea142de1",
	"txType": "1",
	"dataType": "call",
	"amount": "0.00347",
	"fee": "0.0017476",
	"state": 1,
	"targetContractAddr": "cx334db6519871cb2bfd154cec0905ced4ea142de1",
	"id": 730841
}
`

var basicDst = blockatlas.Tx{
	ID:    "0x34b8b6ec3a52710c24074f5e298f4a9c67bb61a0a1dde20e695efaeb30ff3754",
	Coin:  coin.ICX,
	From:  "hx1b8959dd5c57d2c502e22ee0a887d33baec09091",
	To:    "cx334db6519871cb2bfd154cec0905ced4ea142de1",
	Fee:   "1747600000000000",
	Date:  1555396594,
	Block: 357832,
	Meta: blockatlas.Transfer{
		Value:    "3470000000000000",
		Symbol:   "ICX",
		Decimals: 18,
	},
}

func TestNormalize(t *testing.T) {
	var srcTx Tx
	err := json.Unmarshal([]byte(basicSrc), &srcTx)

	if err != nil {
		t.Error(err)
		return
	}

	tx, _ := Normalize(&srcTx)
	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&basicDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("Transactions not equal")
	}
}
