package vechain

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"testing"
)

const transferTrx = `
{
	"sender": "0xb6b6c3ad63192cadd9064432242f3a52329302f3",
	"recipient": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	"amount": "0xaf6751326c8d4a80000",
	"meta": {
		"blockID": "0x0003cae121113a2e3c7bf510629ceab7ead48b6cf2cf1a18747e7fe058b3e190",
		"blockNumber": 248545,
		"blockTimestamp": 1532802660,
		"txID": "0xd989ab9e91e9af849666d507758f75ea2d4dcbe0a658284d8196ac527dbee181",
		"txOrigin": "0xb6b6c3ad63192cadd9064432242f3a52329302f3"
	}
}
`

const transferReceipt = `
{
    "gasUsed": 21000,
    "gasPayer": "0xb6b6c3ad63192cadd9064432242f3a52329302f3",
    "paid": "0x1236efcbcbb340000",
    "reward": "0x576e189f04f60000",
    "reverted": false,
    "meta": {
        "blockID": "0x0003cae121113a2e3c7bf510629ceab7ead48b6cf2cf1a18747e7fe058b3e190",
        "blockNumber": 248545,
        "blockTimestamp": 1532802660,
        "txID": "0xd989ab9e91e9af849666d507758f75ea2d4dcbe0a658284d8196ac527dbee181",
        "txOrigin": "0xb6b6c3ad63192cadd9064432242f3a52329302f3"
    },
    "outputs": []
}
`

const transferOutput = `
{
	"contractAddress": null,
	"events": [],
	"transfers": [
		{
			"sender": "0xb6b6c3ad63192cadd9064432242f3a52329302f3",
			"recipient": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
			"amount": "0xaf6751326c8d4a80000"
		}
	]
}
`

var expectedTransferTrx = models.Tx{
	ID:    "0xd989ab9e91e9af849666d507758f75ea2d4dcbe0a658284d8196ac527dbee181",
	Coin:  coin.VET,
	From:  "0xb6b6c3ad63192cadd9064432242f3a52329302f3",
	To:    "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	Fee:   "21000000000000000000",
	Date:  1532802660,
	Type:   "transfer",
	Status: "completed",
	Block: 248545,
	Sequence: 1532802660,
	Meta:  models.Transfer{
		Value: "51770000000000000000000",
	},
}

var expectedVeThorTrx = models.Tx{
	ID:     "0xd989ab9e91e9af849666d507758f75ea2d4dcbe0a658284d8196ac527dbee181",
	Coin:   coin.VET,
	From:   "0xb6b6c3ad63192cadd9064432242f3a52329302f3",
	To:     "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	Fee:    "0",
	Date:   1532802660,
	Type:   "transfer",
	Status: "completed",
	Sequence: 1532802660,
	Block:  248545,
	Meta:  models.Transfer{
		Value: "21000000000000000000",
	},
}
func TestNormalize(t *testing.T) {
	const address = "0xb853d6a965fbc047aaa9f04d774d53861d7ed653"
	const VeThorContract = "0x0000000000000000000000000000456e65726779" 

	var tests = []struct {
		Tx 		  string
		Receipt   string
		Output    string
		Address   string
		Token     string
		Expected  models.Tx
	}{
		{transferTrx, transferReceipt, transferOutput, address, "", expectedTransferTrx},
		{transferTrx, transferReceipt, transferOutput, address, VeThorContract, expectedVeThorTrx},
	}

	for _, test := range tests {
		var tx      Tx
		var receipt TxReceipt
		var output  TxReceiptOutput


		tErr := json.Unmarshal([]byte(test.Tx), &tx)
		if tErr != nil {
			t.Fatal(tErr)
		}

		rErr  := json.Unmarshal([]byte(test.Receipt), &receipt)
		if rErr != nil {
			t.Fatal(rErr)
		}

		olErr := json.Unmarshal([]byte(test.Output), &output)
		if olErr != nil {
			t.Fatal(olErr)
		}

		var readyTx models.Tx
		normTx, ok := Normalize(&tx, &receipt, &output, test.Address, test.Token)
		if !ok {
			t.Fatal("VeChain: Can't normalize transaction", readyTx)
		}
		readyTx = normTx
	
		actual, err := json.Marshal(&readyTx)
		if err != nil {
			t.Fatal(err)
		}
	
		expected, err := json.Marshal(&test.Expected)
		if err != nil {
			t.Fatal(err)
		}
	
		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}
