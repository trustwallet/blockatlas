package vechain

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const transferReceipt = `{
    "block": 2620166,
    "id": "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
    "nonce": "0x3657a2025b11f27f",
    "origin": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	"timestamp": 1556569300,
	"receipt": {
		"paid": "0x1236efcbcbb340000"
	}
}`

const transferClause = `{
	"to": "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
	"value": "0x29bde5885d7ac80000"
}`

const tokenTransfer = `
{
	"amount": "0x00000000000000000000000000000000000000000000000d8d726b7177a80000",
	"block": 2465269,
	"origin": "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	"receiver": "0x9f3742c2c2fe66c7fca08d77d2262c22e3d56ac8",
	"timestamp": 1555009870,
	"txId": "0xd17dd968610fb4a39ab02a5d8827b26f4cdcd147fb4a4f7a5d5ab14066525d4b"
}
`

var expectedTransferTrx = blockatlas.Tx{
	ID:    "0x2b8776bd4679fa2afa28b55d66d4f6c7c77522fc878ce294d25e32475b704517",
	Coin:  coin.VET,
	From:  "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	To:    "0xda623049a13df5c8a24f0d7713f4add4ab136b1f",
	Fee:   "21000000000000000000",
	Date:  1556569300,
	Type:   "transfer",
	Status: "completed",
	Block: 2620166,
	Sequence: 2620166,
	Meta:  blockatlas.Transfer{
		Value: "770000000000000000000",
	},
}

var expectedVeThorTrx = blockatlas.Tx{
	ID:     "0xd17dd968610fb4a39ab02a5d8827b26f4cdcd147fb4a4f7a5d5ab14066525d4b",
	Coin:   coin.VET,
	From:   "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
	To:     "0x9f3742c2c2fe66c7fca08d77d2262c22e3d56ac8",
	Fee:    "0",
	Date:   1555009870,
	Type:   blockatlas.TxNativeTokenTransfer,
	Status: "completed",
	Sequence: 2465269,
	Block:  2465269,
	Meta:  blockatlas.NativeTokenTransfer{
		Name: "VeThor Token",
		Symbol: "VTHO",
		TokenID: VeThorContractLow,
		Decimals: 18,
		Value: "250000000000000000000",
		From: "0xb853d6a965fbc047aaa9f04d774d53861d7ed653",
		To: "0x9f3742c2c2fe66c7fca08d77d2262c22e3d56ac8",
	},
}
func TestNormalizeTransfer(t *testing.T) {
	var tests = []struct {
		Receipt  string
		Clause   string
		Expected blockatlas.Tx
	}{
		{transferReceipt, transferClause, expectedTransferTrx},
		// {transferTrx, transferReceipt, transferOutput, address, VeThorContract, expectedVeThorTrx},
	}

	for _, test := range tests {
		var receipt TransferReceipt
		var clause  Clause

		// Unmarshal(*t, test.Receipt, &receipt)
		rErr := json.Unmarshal([]byte(test.Receipt), &receipt)
		if rErr != nil {
			t.Fatal(rErr)
		}

		cErr  := json.Unmarshal([]byte(test.Clause), &clause)
		if cErr != nil {
			t.Fatal(cErr)
		}

		var readyTx blockatlas.Tx
		normTx, ok := NormalizeTransfer(&receipt, &clause)
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

func TestNormalizeTokenTransfer(t *testing.T) {
	var tests = []struct {
		Transfer  string
		Expected blockatlas.Tx
	}{
		{tokenTransfer, expectedVeThorTrx},
	}

	for _, test := range tests {
		var tt TokenTransfer

		ttErr := json.Unmarshal([]byte(test.Transfer), &tt)
		if ttErr != nil {
			t.Fatal(ttErr)
		}

		var readyTx blockatlas.Tx
		normTx, ok := NormalizeTokenTransfer(&tt)
		if !ok {
			t.Fatal("VeChain: Can't normalize token transaction", readyTx)
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

// func Unmarshal(t testing.T, input string, i *interface{}) {
// 	err := json.Unmarshal([]byte(input), i)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
