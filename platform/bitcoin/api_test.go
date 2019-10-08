package bitcoin

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	mapset "github.com/deckarep/golang-set"
	"github.com/trustwallet/blockatlas/coin"
)

const outgoingTx = `{
	"txid":"df63ddab7d4eed2fb6cb40d4d0519e7e5ac7cf5ad556b2edbd45963ea1a2931c",
	"version":1,
	"vin":[
		{
			"txid":"bf19be44d7dc3e1e6771801a1d250c7207fa9b09d8df9b0ee1b028b6c153475e",
			"sequence":4294967295,
			"n":0,
			"addresses":[
				"3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC"
			],
			"value":"777200",
			"hex":"00483045022100e9d0db3bb20a5828ab9dae7cd8373064ce087cc9c8e3def87034d5c2f6f3abb9022047d7c27b355c6487cff40bfbd45742d26d727f3135b2396d8f1abc371c51870c01473044022016280108af73079a69f378218ad4259f02c4e4b6f52c573729650cb3645bc9180220785973cb4e5c4ec6340dc77dc56cec3fb74ebd7296cf1d14344d4f3e157658bb014cc952410491bba2510912a5bd37da1fb5b1673010e43d2c6d812c514e91bfa9f2eb129e1c183329db55bd868e209aac2fbc02cb33d98fe74bf23f0c235d6126b1d8334f864104865c40293a680cb9c020e7b1e106d8c1916d3cef99aa431a56d253e69256dac09ef122b1a986818a7cb624532f062c1d1f8722084861c5c3291ccffef4ec687441048d2455d2403e08708fc1f556002f1b6cd83f992d085097f9974ab08a28838f07896fbab08f39495e15fa6fad6edbfb1e754e35fa1c7844c41f322a1863d4621353ae"}
	],
	"vout":[
		{
			"value":"677012",
			"n":0,
			"hex":"a91499fa965ad13a9580ed7a64ac24b2ecad30f1209a87",
			"addresses":["3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4"]
		}
	],
	"blockHash":"00000000000000000011b58c01ede5a602eec61ebaf097aaa6e682ef2819536e",
	"blockHeight":585094,
	"confirmations":1997,
	"blockTime":1562945790,
	"value":"677012",
	"valueIn":"777200",
	"fees":"100188",
	"hex":"01000000015e4753c1b628b0e10e9bdfd8099bfa07720c251d1a8071671e3edcd744be19bf00000000fd5d0100483045022100e9d0db3bb20a5828ab9dae7cd8373064ce087cc9c8e3def87034d5c2f6f3abb9022047d7c27b355c6487cff40bfbd45742d26d727f3135b2396d8f1abc371c51870c01473044022016280108af73079a69f378218ad4259f02c4e4b6f52c573729650cb3645bc9180220785973cb4e5c4ec6340dc77dc56cec3fb74ebd7296cf1d14344d4f3e157658bb014cc952410491bba2510912a5bd37da1fb5b1673010e43d2c6d812c514e91bfa9f2eb129e1c183329db55bd868e209aac2fbc02cb33d98fe74bf23f0c235d6126b1d8334f864104865c40293a680cb9c020e7b1e106d8c1916d3cef99aa431a56d253e69256dac09ef122b1a986818a7cb624532f062c1d1f8722084861c5c3291ccffef4ec687441048d2455d2403e08708fc1f556002f1b6cd83f992d085097f9974ab08a28838f07896fbab08f39495e15fa6fad6edbfb1e754e35fa1c7844c41f322a1863d4621353aeffffffff0194540a000000000017a91499fa965ad13a9580ed7a64ac24b2ecad30f1209a8700000000"
}`

const incomingTx = `{
    "txid": "a2d70bee124510c476f159fa83cdb34d663fc6020c81aad19b238601d679fed7",
    "version": 4,
    "vin": [{
        "txid": "5a3664328ac4e1c0688729573296c2ec69dd9a7cf98d49967b41520be794229b",
        "n": 0,
        "addresses": ["t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD"],
        "isAddress": true,
        "value": "387582",
        "hex": "483045022100ec29a476dac49578339a92e6c20451aaf3ff6691efaf7d4d3113d07589771ca702203c0c173bdc356300edbd64cdfaa868b97c13ebc403026b283eb5e1fca398db8b012103729cc4211cf70f87c70c3cef90e0ca9b91e99b42364b8c600d5781277647de5f"
    }],
    "vout": [{
        "value": "200997",
        "n": 0,
        "spent": true,
        "hex": "76a9146fd73e7c147d8ccc15fda31d8429e70f302b843988ac",
        "addresses": ["t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM"],
        "isAddress": true
    }, {
        "value": "186359",
        "n": 1,
        "spent": true,
        "hex": "76a91484f0258cb7974993e6af928921b7f699c51a309488ac",
        "addresses": ["t1VzWtLj9CSAK3QnxA7uuiK6XhJrjGjKoy4"],
        "isAddress": true
    }],
    "blockHash": "0000000000a8248c4a14a2dcb74d92855bf9440da9b7b1e6d4baa14ee7e3081c",
    "blockHeight": 479017,
    "confirmations": 116233,
    "blockTime": 1549793065,
    "value": "387356",
    "valueIn": "387582",
    "fees": "226",
    "hex": "0400008085202f89019b2294e70b52417b96498df97c9add69ecc2963257298768c0e1c48a3264365a000000006b483045022100ec29a476dac49578339a92e6c20451aaf3ff6691efaf7d4d3113d07589771ca702203c0c173bdc356300edbd64cdfaa868b97c13ebc403026b283eb5e1fca398db8b012103729cc4211cf70f87c70c3cef90e0ca9b91e99b42364b8c600d5781277647de5f000000000225110300000000001976a9146fd73e7c147d8ccc15fda31d8429e70f302b843988acf7d70200000000001976a91484f0258cb7974993e6af928921b7f699c51a309488ac00000000000000000000000000000000000000"
}`

var expectedOutgoingTx = blockatlas.Tx{
	ID:   "df63ddab7d4eed2fb6cb40d4d0519e7e5ac7cf5ad556b2edbd45963ea1a2931c",
	Coin: coin.BTC,
	From: "3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC",
	To:   "3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4",
	Inputs: []blockatlas.TxOutput{
		{
			Address: "3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC",
			Value:   "777200",
		},
	},
	Outputs: []blockatlas.TxOutput{
		{
			Address: "3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4",
			Value:   "677012",
		},
	},
	Fee:       "100188",
	Date:      1562945790,
	Type:      "transfer",
	Status:    blockatlas.StatusCompleted,
	Block:     585094,
	Sequence:  0,
	Direction: blockatlas.DirectionSelf,
	Meta: blockatlas.Transfer{
		Value:    "677012",
		Symbol:   "BTC",
		Decimals: 8,
	},
}

var expectedIncomingTx = blockatlas.Tx{
	ID:   "a2d70bee124510c476f159fa83cdb34d663fc6020c81aad19b238601d679fed7",
	Coin: coin.ZEC,
	From: "t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD",
	To:   "t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM",
	Inputs: []blockatlas.TxOutput{
		{
			Address: "t1T7cLkvDVScjw95WguoAZbbT8mrdqVtpiD",
			Value:   "387582",
		},
	},
	Outputs: []blockatlas.TxOutput{
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
	Status:    blockatlas.StatusCompleted,
	Block:     479017,
	Sequence:  0,
	Direction: blockatlas.DirectionIncoming,
	Meta: blockatlas.Transfer{
		Value:    "200997",
		Symbol:   "ZEC",
		Decimals: 8,
	},
}

func TestNormalizeTransfer(t *testing.T) {

	outgoingTxSet := mapset.NewSet()
	outgoingTxSet.Add("3FjBW1KL9L8aYtdKzJ8FhCNxmXB7dXDRw4")
	outgoingTxSet.Add("3QJmV3qfvL9SuYo34YihAf3sRCW3qSinyC")

	incomingTxSet := mapset.NewSet()
	incomingTxSet.Add("t1U4xs3qMxc2TL8wwYufmBngA5mewLHRwhM")
	incomingTxSet.Add("t1ZBs9xvRypkjXmci2SS6zbNWVhuWH1h93L")
	incomingTxSet.Add("t1VZp67AK9zgdXwa35kwYrJ1Mh4NWjUENrM")

	var tests = []struct {
		RawTx      string
		Expected   blockatlas.Tx
		AddressSet mapset.Set
	}{
		{outgoingTx, expectedOutgoingTx, outgoingTxSet},
		{incomingTx, expectedIncomingTx, incomingTxSet},
	}

	for _, test := range tests {
		var transaction Transaction

		rErr := json.Unmarshal([]byte(test.RawTx), &transaction)
		if rErr != nil {
			t.Fatal(rErr)
		}

		p := &Platform{CoinIndex: test.Expected.Coin}
		var readyTx blockatlas.Tx
		normTx, ok := p.NormalizeTransfer(&transaction, test.Expected.Coin, test.AddressSet)
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

		if !bytes.Equal(actual, expected) {
			println(string(actual))
			println(string(expected))
			t.Error("Transactions not equal")
		}
	}
}

// zpub: zpub6r9CEhEkruYbEcu2yQCaRKQ1qufTa4zLrx6ezs31P627UpAepVNBE2td3d3mHnSaXyRbwksRwDJGzLBWQeZPFMut8N3BvXpcwRwEWGEwAnq
var btcSet = mapset.NewSet("bc1qfrrncxmf7skye2glyef95xlpmrlmf2e8qlav2l", "bc1qxm90n0rxkadhdkvglev56k60qths73luzlnn7a",
	"bc1q2sykr9c342mjpm9mwnps8ksk6e35lz75rpdlfe", "bc1qs86ucvr3unce2grvfp77433npy66nzha9w0e3c")
var btcInputs1 = []blockatlas.TxOutput{{Address: "bc1q2sykr9c342mjpm9mwnps8ksk6e35lz75rpdlfe"}}
var btcOutputs1 = []blockatlas.TxOutput{{Address: "bc1q6wf7tj62f0uwr6almah3666th2ejefdg72ek6t"}}
var btcInputs2 = []blockatlas.TxOutput{{
	Address: "3CgvDkzcJ7yMZe75jNBem6Bj6nkMAWwMEf"},
	{Address: "3LyzYcB54pm9EAMmzXpFfb1kzEDAFvqBgT"},
	{Address: "3Q6DYour5q5WdMhyXsyPgBeAqPCXchzCsF"},
	{Address: "3JZZM1rwst7G5izxbFL7KNvy7ZiZ47SVqG"}}
var btcOutputs2 = []blockatlas.TxOutput{
	{Address: "139f1CrnLWvVajGzs3ZtpQhbGWxM599sho"},
	{Address: "3LyzYcB54pm9EAMmzXpFfb1kzEDAFvqBgT"},
	{Address: "bc1q9mx5tm66zs7epa4skvyuf2vfuwmtnlttj74cnl"},
	{Address: "3JZZM1rwst7G5izxbFL7KNvy7ZiZ47SVqG"}}

var dogeSet = mapset.NewSet("DB49sNjVdxyREXEBEzUV54TrQYYpvi3Be7")
var dogeInputs = []blockatlas.TxOutput{{Address: "DAukM5pPtGdbPxMX1u2LYHoyhbDhEFHbnH"}}
var dogeOutputs = []blockatlas.TxOutput{{Address: "DB49sNjVdxyREXEBEzUV54TrQYYpvi3Be7"}, {Address: "DAukM5pPtGdbPxMX1u2LYHoyhbDhEFHbnH"}}

func TestInferDirection(t *testing.T) {
	var tests = []struct {
		AddressSet mapset.Set
		Inputs     []blockatlas.TxOutput
		Outputs    []blockatlas.TxOutput
		Expected   blockatlas.Direction
		Coin       uint
	}{
		{
			btcSet,
			btcInputs1,
			btcOutputs1,
			blockatlas.DirectionOutgoing,
			coin.BTC,
		},
		{
			btcSet,
			btcInputs2,
			btcOutputs2,
			blockatlas.DirectionIncoming,
			coin.BTC,
		},
		{
			dogeSet,
			dogeInputs,
			dogeOutputs,
			blockatlas.DirectionIncoming,
			coin.DOGE,
		},
	}

	for _, test := range tests {
		p := &Platform{CoinIndex: test.Coin}
		tx := blockatlas.Tx{
			Inputs:  test.Inputs,
			Outputs: test.Outputs,
		}

		direction := p.InferDirection(&tx, test.AddressSet)
		if direction != test.Expected {
			t.Errorf("direction is not %s", test.Expected)
		}
	}
}
