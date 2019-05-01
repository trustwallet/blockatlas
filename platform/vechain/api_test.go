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
	"sender": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	"recipient": "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
	"amount": "0x29bde5885d7ac80000",
	"meta": {
		"blockID": "0x0027fb06c4c12ee0116b8a525d8f4bf502562486abb70affa211d07c8cfba37b",
		"blockNumber": 2620166,
		"blockTimestamp": 1556569300,
		"txID": "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
		"txOrigin": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653"
	}
}
`

const transferReceipt = `
{
    "gasUsed": 21000,
    "gasPayer": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
    "paid": "0x1236efcbcbb340000",
    "reward": "0x576e189f04f60000",
    "reverted": false,
    "meta": {
        "blockID": "0x0027fb06c4c12ee0116b8a525d8f4bf502562486abb70affa211d07c8cfba37b",
        "blockNumber": 2620166,
        "blockTimestamp": 1556569300,
        "txID": "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
        "txOrigin": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653"
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
			"sender": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
			"recipient": "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
			"amount": "0x29bde5885d7ac80000"
		}
	]
}
`

var expectedTransferTrx = models.Tx{
	ID:    "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
	Coin:  coin.VET,
	From:  "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	To:    "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
	Fee:   "21000000000000000000",
	Date:  1556569300,
	Type:   "transfer",
	Status: "completed",
	Block: 2620166,
	Sequence: 1556569300,
	Meta:  models.Transfer{
		Value: "770000000000000000000",
	},
}

var expectedVeThorTrx = models.Tx{
	ID:     "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
	Coin:   coin.VET,
	From:   "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	To:     "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
	Fee:    "0",
	Date:   1556569300,
	Type:   "token_transfer",
	Status: "completed",
	Sequence: 1556569300,
	Block:  2620166,
	Meta:  models.NativeTokenTransfer{
		Name: "VeThor Token",
		Symbol: "VTHO",
		TokenID: "0x0000000000000000000000000000456e65726779",
		Decimals: 18,
		Value: "21000000000000000000",
		From: "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
		To: "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
	},
}
func TestNormalize(t *testing.T) {
	const address = "0xb853d6a965fbc047aaa9f04d774d53861d7ed653"

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
