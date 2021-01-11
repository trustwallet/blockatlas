package theta

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/txtype"
)

var (
	thetaTransfer, _ = mock.JsonStringFromFilePath("mocks/" + "theta_transfer.json")
	tFuelTransfer, _ = mock.JsonStringFromFilePath("mocks/" + "tfuel_transfer.json")

	expectedTransferTrx = txtype.Tx{
		ID:        "0x413d8423fd1e6df99fc57f425dfd58c791c877657b364c62c15905ade5114a70",
		Coin:      coin.THETA,
		From:      "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
		To:        "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
		Fee:       "2000000000000",
		Date:      1557136781,
		Type:      "transfer",
		Status:    txtype.StatusCompleted,
		Block:     700321,
		Sequence:  43,
		Direction: txtype.DirectionOutgoing,
		Meta: txtype.Transfer{
			Value:    "4000000000000000000",
			Symbol:   "THETA",
			Decimals: 18,
		},
	}

	expectedTfuelTransfer = txtype.Tx{
		ID:        "0x558cb5ec877119c2c84a677277efb5b3059adb830c6e74971b3dbe93221b7132",
		Coin:      coin.THETA,
		From:      "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
		To:        "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
		Fee:       "2000000000000",
		Date:      1557136821,
		Type:      txtype.TxNativeTokenTransfer,
		Status:    txtype.StatusCompleted,
		Sequence:  44,
		Block:     700327,
		Direction: txtype.DirectionIncoming,
		Meta: txtype.NativeTokenTransfer{
			Name:     "Theta Fuel",
			Symbol:   "TFUEL",
			TokenID:  "tfuel",
			Decimals: 18,
			Value:    "44324000000000000",
			From:     "0x0a7d7141e9abe5d1c760cffa1129c6eb94f35a2a",
			To:       "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f",
		},
	}
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		Transaction string
		Address     string
		Token       string
		Expected    txtype.Tx
	}{
		{thetaTransfer, "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f", "", expectedTransferTrx},
		{tFuelTransfer, "0xac0eeb6ee3e32e2c74e14ac74155063e4f4f981f", "tfuel", expectedTfuelTransfer},
	}

	for _, test := range tests {
		var trx Tx

		tErr := json.Unmarshal([]byte(test.Transaction), &trx)
		if tErr != nil {
			t.Fatal("THETA: Can't unmarshal transaction", tErr)
		}

		var readyTx txtype.Tx
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

		assert.JSONEq(t, string(actual), string(expected))
	}
}

func TestGetDirection(t *testing.T) {
	var addrChecksum = "0x42616C88c7076FbE6e1596b734c13356b5A508a4"
	var addr = "0x42616c88c7076fbe6e1596b734c13356b5a508a4"
	var otherAddr = "0x8665a3cbc02ff17cf9d712e8a20f3d7bb1444517"

	tests := []struct {
		address     string
		expectedDir txtype.Direction
		trxInput    Input
		trxOutput   Output
	}{
		{address: addrChecksum, expectedDir: txtype.DirectionSelf, trxInput: Input{Address: addr}, trxOutput: Output{Address: addr}},
		{address: addrChecksum, expectedDir: txtype.DirectionOutgoing, trxInput: Input{Address: addr}, trxOutput: Output{Address: otherAddr}},
		{address: addrChecksum, expectedDir: txtype.DirectionIncoming, trxInput: Input{Address: otherAddr}, trxOutput: Output{Address: addr}},
	}

	for _, test := range tests {
		actualDir, err := getDirection(test.address, test.trxInput, test.trxOutput)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, test.expectedDir, actualDir, test.expectedDir)
	}
}
