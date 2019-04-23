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
	"sender": "0x7b41d6451eaa012c551b6ed3af9f871a00602e88",
	"recipient": "0xa4adafaef9ec07bc4dc6de146934c7119341ee25",
	"amount": "0x3b5dd5955f12cbf0000",
	"meta": {
		"blockID": "0x0026a6f4a7673e9268e1ee5cfdec899d751b93b611ffff5ec4fa16492b56c55a",
		"blockNumber": 2533108,
		"blockTimestamp": 1555688350,
		"txID": "0xe282bc330b660759ac4060e33f8d74b71a59dd8dea041d71046b32d2a84e8fdd"
	}
}
`

const tokenTrx = `
{
	"sender": "0xb69ded9f0da15d240ee6803dacd7fcf68744e8ff",
	"recipient": "0x70e0a09fa23019f83b42f018891611935c1459d0",
	"amount": "0x4405277d0cb1630000",
	"meta": {
		"blockID": "0x002712dcaa19f9837e85579dfa28ad4ea54ebe2cc8cdf945ef6eb4d9b72aa0f5",
		"blockNumber": 2560732,
		"blockTimestamp": 1555964860,
		"txID": "0x52ae85810f8568c0dc13bc8afd6f144d7a8e2bb949b76a5d14ab529c2438f32f",
		"txOrigin": "0x70e0a09fa23019f83b42f018891611935c1459d0"
	}
}
`


const clause = `
{
	"to": "0xa4adafaef9ec07bc4dc6de146934c7119341ee25",
	"value": "0x3b5dd5955f12cbf0000",
	"data": "0x"
}
`

const transactionId = `
{
    "id": "0xe282bc330b660759ac4060e33f8d74b71a59dd8dea041d71046b32d2a84e8fdd",
    "clauses": [
        {
            "to": "0xa4adafaef9ec07bc4dc6de146934c7119341ee25",
            "value": "0x3b5dd5955f12cbf0000",
            "data": "0x"
        }
    ],
    "gasPriceCoef": 0,
    "gas": 21000,
    "nonce": "0xfd15e25fbb381811",
    "meta": {
        "blockID": "0x0026a6f4a7673e9268e1ee5cfdec899d751b93b611ffff5ec4fa16492b56c55a",
        "blockNumber": 2533108,
        "blockTimestamp": 1555688350
    }
}
`

const transactionReceipt = `
{
    "paid": "0x1966cc5d60c170000",
    "meta": {
        "blockID": "0x002712dcaa19f9837e85579dfa28ad4ea54ebe2cc8cdf945ef6eb4d9b72aa0f5",
        "blockNumber": 2560732,
        "blockTimestamp": 1555964860,
        "txID": "0x52ae85810f8568c0dc13bc8afd6f144d7a8e2bb949b76a5d14ab529c2438f32f"
    },
    "outputs": [
        {
            "transfers": [
                {
                    "sender": "0xb69ded9f0da15d240ee6803dacd7fcf68744e8ff",
                    "recipient": "0x70e0a09fa23019f83b42f018891611935c1459d0",
                    "amount": "0x4405277d0cb1630000"
                }
            ]
        }
    ]
}
`

const output = `
{
	"transfers": [
		{
			"sender": "0xb69ded9f0da15d240ee6803dacd7fcf68744e8ff",
			"recipient": "0x70e0a09fa23019f83b42f018891611935c1459d0",
			"amount": "0x4405277d0cb1630000"
		}
	]
}
`

var finalTx = models.Tx{
	ID:    "0xe282bc330b660759ac4060e33f8d74b71a59dd8dea041d71046b32d2a84e8fdd",
	Coin:  coin.VET,
	From:  "0x7b41d6451eaa012c551b6ed3af9f871a00602e88",
	To:    "0xa4adafaef9ec07bc4dc6de146934c7119341ee25",
	Fee:   "0",
	Date:  1555688350,
	Type:   "transfer",
	Status: "completed",
	Block: 2533108,
	Sequence: 18236731166897477649,
	Meta:  models.Transfer{
		Value: "17521910000000000000000",
	},
}

var expectedTokenTrx = models.Tx{
	ID:     "0x52ae85810f8568c0dc13bc8afd6f144d7a8e2bb949b76a5d14ab529c2438f32f",
	Coin:   coin.VET,
	From:   "0xb69ded9f0da15d240ee6803dacd7fcf68744e8ff",
	To:     "0x70e0a09fa23019f83b42f018891611935c1459d0",
	Fee:    "0",
	Date:   1555964860,
	Type:   "transfer",
	Status: "completed",
	Sequence: 2560732,
	Block:  2560732,
	Meta:  models.Transfer{
		Value: "29286000000000000000",
	},
}
func TestNormalize(t *testing.T) {
	const address = "0x7b41d6451eAA012C551b6ED3af9F871a00602E88"
	var tx Tx
	var r  TxId
	var cl Clause
	
	tErr := json.Unmarshal([]byte(transferTrx), &tx)
	if tErr != nil {
		t.Fatal(tErr)
	}

	rErr  := json.Unmarshal([]byte(transactionId), &r)
	if rErr != nil {
		t.Fatal(rErr)
	}

	clErr := json.Unmarshal([]byte(clause), &cl)
	if clErr != nil {
		t.Fatal(clErr)
	}

	var readyTx models.Tx
	normTx, ok := Normalize(&tx, &r, &cl, address)
	if !ok {
		t.Fatal("VeChain: Can't normalize transaction", readyTx)
	}
	readyTx = normTx

	actual, err := json.Marshal(&readyTx)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := json.Marshal(&finalTx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		println(string(actual))
		println(string(expected))
		t.Error("Transactions not equal")
	}
}

func TestNormalizeToken(t *testing.T) {
	var tx Tx
	var receipt TxReceipt
	var out TxReceiptOutput
	
	tErr := json.Unmarshal([]byte(tokenTrx), &tx)
	if tErr != nil {
		t.Fatal(tErr)
	}

	rErr  := json.Unmarshal([]byte(transactionReceipt), &receipt)
	if rErr != nil {
		t.Fatal(rErr)
	}

	clErr := json.Unmarshal([]byte(output), &out)
	if clErr != nil {
		t.Fatal(clErr)
	}

	var readyTx models.Tx
	normTx, ok := NormalizeToken(&out, &receipt)
	if !ok {
		t.Fatal("VeChain: Can't normalize transaction", readyTx)
	}
	readyTx = normTx

	actual, err := json.Marshal(&readyTx)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := json.Marshal(&expectedTokenTrx)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(actual, expected) {
		println(string(actual))
		println(string(expected))
		t.Error("Transactions not equal")
	}
}
