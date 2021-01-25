package bitcoin

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/trustwallet/blockatlas/platform/bitcoin/blockbook"

	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	outgoingTx, _ = mock.JsonStringFromFilePath("mocks/" + "outgoing_tx.json")
	incomingTx, _ = mock.JsonStringFromFilePath("mocks/" + "incoming_tx.json")
	pendingTx, _  = mock.JsonStringFromFilePath("mocks/" + "pending_tx.json")

	expectedOutgoingTx = types.Tx{
		ID:   "df63ddab7d4eed2fb6cb40d4d0519e7e5ac7cf5ad556b2edbd45963ea1a2931c",
		Coin: coin.BITCOIN,
		From: "3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC",
		To:   "3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4",
		Inputs: []types.TxOutput{
			{
				Address: "3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC",
				Value:   "777200",
			},
		},
		Outputs: []types.TxOutput{
			{
				Address: "3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4",
				Value:   "677012",
			},
		},
		Fee:       "100188",
		Date:      1562945790,
		Type:      "transfer",
		Status:    types.StatusCompleted,
		Block:     585094,
		Sequence:  0,
		Direction: types.DirectionSelf,
		Meta: types.Transfer{
			Value:    "677012",
			Symbol:   "BTC",
			Decimals: 8,
		},
	}

	expectedIncomingTx = types.Tx{
		ID:   "a2d70bee124510c476f159fa83cdb34d663fc6020c81aad19b238601d679fed7",
		Coin: coin.ZCASH,
		From: "t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD",
		To:   "t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM",
		Inputs: []types.TxOutput{
			{
				Address: "t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD",
				Value:   "387582",
			},
		},
		Outputs: []types.TxOutput{
			{
				Address: "t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM",
				Value:   "200997",
			},
			{
				Address: "t1VzWtLj9CSAK3QnxA7uuiK6XhJrjGjKoy4",
				Value:   "186359",
			},
		},
		Fee:       "226",
		Date:      1549793065,
		Type:      "transfer",
		Status:    types.StatusCompleted,
		Block:     479017,
		Sequence:  0,
		Direction: types.DirectionIncoming,
		Meta: types.Transfer{
			Value:    "200997",
			Symbol:   "ZEC",
			Decimals: 8,
		},
	}

	expectedPendingTx = types.Tx{
		ID:   "a2d70bee124510c476f159fa83cdb34d663fc6020c81aad19b238601d679fed7",
		Coin: coin.ZCASH,
		From: "t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD",
		To:   "t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM",
		Inputs: []types.TxOutput{
			{
				Address: "t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD",
				Value:   "387582",
			},
		},
		Outputs: []types.TxOutput{
			{
				Address: "t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM",
				Value:   "200997",
			},
			{
				Address: "t1VzWtLj9CSAK3QnxA7uuiK6XhJrjGjKoy4",
				Value:   "186359",
			},
		},
		Fee:       "226",
		Date:      1549793065,
		Type:      "transfer",
		Status:    types.StatusCompleted,
		Block:     0,
		Sequence:  0,
		Direction: types.DirectionIncoming,
		Meta: types.Transfer{
			Value:    "200997",
			Symbol:   "ZEC",
			Decimals: 8,
		},
	}
)

func TestNormalizeTransfer(t *testing.T) {

	outgoingTxSet := mapset.NewSet()
	outgoingTxSet.Add("3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4")
	outgoingTxSet.Add("3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC")

	incomingTxSet := mapset.NewSet()
	incomingTxSet.Add("t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM")
	incomingTxSet.Add("t1ZBs9xvRypkjXmci2SS6zbNWVhuWH1h93L")
	incomingTxSet.Add("t1VZp67AK9zgdXwa35kwYrJ1Mh4NWjUENrM")

	tests := []struct {
		RawTx      string
		Expected   types.Tx
		AddressSet mapset.Set
	}{
		{outgoingTx, expectedOutgoingTx, outgoingTxSet},
		{incomingTx, expectedIncomingTx, incomingTxSet},
		{pendingTx, expectedPendingTx, incomingTxSet},
	}

	for _, test := range tests {
		var transaction blockbook.Transaction

		rErr := json.Unmarshal([]byte(test.RawTx), &transaction)
		if rErr != nil {
			t.Fatal(rErr)
		}

		var readyTx types.Tx
		normTx, ok := normalizeTransfer(transaction, test.Expected.Coin, test.AddressSet)
		if !ok {
			t.Fatal("Bitcoin: Can't normalize transaction", readyTx)
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

func TestTransactionStatus(t *testing.T) {
	tests := []struct {
		Tx       blockbook.Transaction
		Expected types.Status
	}{
		{blockbook.Transaction{Confirmations: 0}, types.StatusPending},
		{blockbook.Transaction{Confirmations: 1}, types.StatusCompleted},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Tx.GetStatus())
	}
}

func TestParseOutputs(t *testing.T) {
	tests := []struct {
		name    string
		outputs string
		want    []types.TxOutput
	}{
		{
			name: "Test Doge inputs from 0xb02977b96e5c65fd807e28230375c1267ded1de7c2c43292bf36552283bc5696",
			outputs: `[{
				"txid": "6f59c4d96566c84aecb8399884d781766fe39d3ec76c609e6c11b01192c07341",
				"sequence": 4294967295,
				"n": 0,
				"addresses": [
					"DPoYGk1wGQ3uWs5G3exd9WKvVyu8weKYVA"
				],
				"isAddress": true,
				"value": "10000000",
				"hex": "47304402207bb4fe5709874e5bc82c88945f8abb9aebdab996a1a619b0f01ec9a2ca3f862702204c02f986653bbf322ff5813322b72a7f7ac6cedfabf32d8e576be5eeee4acfc4012103565519e77659aae844889ae12609309f85a8d22bf815c4daa418e457c7cb01eb"
			},
			{
				"txid": "b476736ec08dc16942941103806ec16ce71ddde4d5422dcf6f05b4381c23b8b0",
				"sequence": 4294967295,
				"n": 1,
				"addresses": [
					"DPoYGk1wGQ3uWs5G3exd9WKvVyu8weKYVA"
				],
				"isAddress": true,
				"value": "500000000",
				"hex": "48304502210085426a63df0fa1343a3f7aa8903804a8e49815006f9ed65494b68a81588e605c02207a487df90dcc76e4cbe70358054084b97d407529912762fb784aac825c120125012103565519e77659aae844889ae12609309f85a8d22bf815c4daa418e457c7cb01eb"
			},
			{
				"txid": "239582019ca4dd5e4ba4441cdd949d67eb2461b6f0600d77245355f751bf9fb4",
				"sequence": 4294967295,
				"n": 2,
				"addresses": [
					"DPoYGk1wGQ3uWs5G3exd9WKvVyu8weKYVA"
				],
				"isAddress": true,
				"value": "500000000",
				"hex": "473044022047cad0afd2aa4ff9b3fc45a6afb40b9745c1b39499a96df803f899d189f3822c0220267464d2954de54825717b93e3597a0915c71aecf1215ed33021bef376813b9a012103565519e77659aae844889ae12609309f85a8d22bf815c4daa418e457c7cb01eb"
			}]`,
			want: []types.TxOutput{
				{
					Address: "DPoYGk1wGQ3uWs5G3exd9WKvVyu8weKYVA",
					Value:   "1010000000",
				},
			},
		},
		{
			name: "Test Doge outputs from 0xb02977b96e5c65fd807e28230375c1267ded1de7c2c43292bf36552283bc5696",
			outputs: `[{
				"value": "1000000000",
				"n": 0,
				"spent": true,
				"hex": "76a914e34df4959b71f9a06af1eeb4f836e521067b777988ac",
				"addresses": [
					"DRryKEukopEDv7cm6Y1Li6232VHEjnXptA"
				],
				"isAddress": true
			},
			{
				"value": "9973378",
				"n": 1,
				"hex": "76a914ccb78d11b3850ac3252c1cbda0a2ceeaa833feaf88ac",
				"addresses": [
					"DPoYGk1wGQ3uWs5G3exd9WKvVyu8weKYVA"
				],
				"isAddress": true
			}]`,
			want: []types.TxOutput{
				{
					Address: "DRryKEukopEDv7cm6Y1Li6232VHEjnXptA",
					Value:   "1000000000",
				},
				{
					Address: "DPoYGk1wGQ3uWs5G3exd9WKvVyu8weKYVA",
					Value:   "9973378",
				},
			},
		},
	}
	for _, tt := range tests {
		var outputs []blockbook.Output
		_ = json.Unmarshal([]byte(tt.outputs), &outputs)
		want := parseOutputs(outputs)
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(want, tt.want) {
				t.Errorf("parseOutputs() = %v, want %v", want, tt.want)
			}
		})
	}
}
