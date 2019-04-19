package vechain

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"testing"
)

const trx = `
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

const clause = `
{
	"to": "0xa4adafaef9ec07bc4dc6de146934c7119341ee25",
	"value": "0x3b5dd5955f12cbf0000",
	"data": "0x"
}
`

const receipe = `
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
var finalTx = models.Tx{
	Id:    "0xe282bc330b660759ac4060e33f8d74b71a59dd8dea041d71046b32d2a84e8fdd",
	Coin:  coin.VET,
	From:  "0x7b41d6451eaa012c551b6ed3af9f871a00602e88",
	To:    "0xa4adafaef9ec07bc4dc6de146934c7119341ee25",
	Fee:   "0",
	Date:  1555688350,
	Type:   "transfer",
	Status: "completed",
	Block: 2533108,
	Sequence: 9223372036854775807,
	Meta:  models.Transfer{
		Value: "9223372036854775807",
	},
}

func TestNormalize(t *testing.T) {
	const address = "0x7b41d6451eAA012C551b6ED3af9F871a00602E88"
	var tx Tx
	var r  TxReceipt
	var cl Clause
	
	tErr := json.Unmarshal([]byte(trx), &tx)
	if tErr != nil {
		t.Fatal(tErr)
	}

	rErr  := json.Unmarshal([]byte(receipe), &r)
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
