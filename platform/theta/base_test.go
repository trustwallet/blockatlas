package theta

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var transferReceipt = `{
	"hash": "0x413d8423fd1e6df99fc57f425dfd58c791c877657b364c62c15905ade5114a70",
	"data": {
		"fee": {
			"thetawei": "0",
			"tfuelwei": "2000000000000"
		},
		"inputs": [
			{
				"address": "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
				"coins": {
					"thetawei": "4000000000000000000",
					"tfuelwei": "2000000000000"
				},
				"sequence": "43"
			}
		],
		"outputs": [
			{
				"address": "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
				"coins": {
					"thetawei": "4000000000000000000",
					"tfuelwei": "0"
				}
			}
		]
	},
	"number": 7785228,
	"block_height": "700321",
	"timestamp": "1557136781",
	"status": "finalized"
}`

var tFuelTransfer = `
{
	"hash": "0x558cb5ec877119c2c84a677277efb5b3059adb830c6e74971b3dbe93221b7132",
	"type": 2,
	"data": {
		"fee": {
			"thetawei": "0",
			"tfuelwei": "2000000000000"
		},
		"inputs": [
			{
				"address": "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
				"coins": {
					"thetawei": "0",
					"tfuelwei": "44326000000000000"
				},
				"sequence": "44"
			}
		],
		"outputs": [
			{
				"address": "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
				"coins": {
					"thetawei": "0",
					"tfuelwei": "44324000000000000"
				}
			}
		]
	},
	"number": 7785266,
	"block_height": "700327",
	"timestamp": "1557136821",
	"status": "finalized"
}
`

var expectedTransferTrx = blockatlas.Tx{
	ID:       "0x413d8423fd1e6df99fc57f425dfd58c791c877657b364c62c15905ade5114a70",
	Coin:     coin.THETA,
	From:     "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
	To:       "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
	Fee:      "2000000000000",
	Date:     1557136781,
	Type:     "transfer",
	Status:   blockatlas.StatusCompleted,
	Block:    700321,
	Sequence: 43,
	Meta: blockatlas.Transfer{
		Value:    "4000000000000000000",
		Symbol:   "THETA",
		Decimals: 18,
	},
}

var expectedTfuelTransfer = blockatlas.Tx{
	ID:       "0x558cb5ec877119c2c84a677277efb5b3059adb830c6e74971b3dbe93221b7132",
	Coin:     coin.THETA,
	From:     "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
	To:       "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
	Fee:      "2000000000000",
	Date:     1557136821,
	Type:     blockatlas.TxNativeTokenTransfer,
	Status:   blockatlas.StatusCompleted,
	Sequence: 44,
	Block:    700327,
	Meta: blockatlas.NativeTokenTransfer{
		Name:     "Theta Fuel",
		Symbol:   "TFUEL",
		TokenID:  "tfuel",
		Decimals: 18,
		Value:    "44324000000000000",
		From:     "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
		To:       "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
	},
}

func TestNormalize(t *testing.T) {
	var tests = []struct {
		Transaction string
		Address     string
		Token       string
		Expected    blockatlas.Tx
	}{
		{transferReceipt, "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f", "", expectedTransferTrx},
		{tFuelTransfer, "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f", "tfuel", expectedTfuelTransfer},
	}

	for _, test := range tests {
		var trx Tx

		tErr := json.Unmarshal([]byte(test.Transaction), &trx)
		if tErr != nil {
			t.Fatal("THETA: Can't unmarshal transaction", tErr)
		}

		var readyTx blockatlas.Tx
		normTx, ok := Normalize(&trx, test.Address, test.Token)
		if !ok {
			t.Fatal("THETA: Can't normalize transaction", readyTx)
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
