package ripple

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
)

const paymentSrc = `
{
	"hash": "40279A3DE51148BD41409DADF29DE8DCCD50F5AEE30840827B2C4C81C4E36505",
	"ledger_index": 34698103,
	"date": "2017-12-01T22:45:30+00:00",
	"tx": {
		"TransactionType": "Payment",
		"Flags": 2147483648,
		"Sequence": 21,
		"LastLedgerSequence": 34698105,
		"DestinationTag": 2500,
		"Amount": "100000000",
		"Fee": "3115",
		"SigningPubKey": "03807050F9E271B2E49B0FF658362EF37DBFDD31435E610B6E11C52879DF8A9907",
		"TxnSignature": "3045022100D14057AA2A868F54FC7CA2E44C8310D9A944446580EAA45936A75CFFDD00425602205CCBFACB55AB0F5B02659F1EBE619FC04DE75B0227C8EB148DC6D08CABBAB072",
		"Account": "rGSxFjoqmWz54PycrgQBQ5dB6e7TUpMxzq",
		"Destination": "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
		"Memos": [
			{
				"Memo": {
					"MemoType": "636C69656E74",
					"MemoFormat": "7274312E342E332D31332D6735383261336135"
				}
			}
		]
	},
	"meta": {
		"TransactionIndex": 20,
		"TransactionResult": "tesSUCCESS",
		"delivered_amount": "100000000"
	}
}
`

var paymentDst = blockatlas.Tx{
	ID:     "40279A3DE51148BD41409DADF29DE8DCCD50F5AEE30840827B2C4C81C4E36505",
	Coin:   coin.XRP,
	From:   "rGSxFjoqmWz54PycrgQBQ5dB6e7TUpMxzq",
	To:     "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1",
	Fee:    "3115",
	Date:   1512168330,
	Block:  34698103,
	Status: blockatlas.StatusCompleted,
	Memo:   "2500",
	Meta: blockatlas.Transfer{
		Value: "100000000",
	},
}

const paymentSrc2 = `
{
   "hash":"3D8512E02414EF5A6BC00281D945735E85DED9EF739B1DCA9EABE04D9EEC72C1",
   "ledger_index":49163909,
   "date":"2019-08-06T17:58:01+00:00",
   "tx":{
      "TransactionType":"Payment",
      "Flags":2147614720,
      "Sequence":115,
      "DestinationTag":0,
      "LastLedgerSequence":49163911,
      "Amount":"1000000000",
      "Fee":"120",
      "SendMax":{
         "value":"0.001",
         "currency":"USD",
         "issuer":"rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
      },
      "SigningPubKey":"030E4853E7D0B0E2D3C1233EADCB1B1C35DE75AD4AECD94AC534B3057537753B94",
      "TxnSignature":"3045022100EBBDDB5D2F59472463CA03429DDDED4F06648FF097662697CCFF3C5C9C36091202205367A18FE65F767D6C6D256B2F7058BBA3C5D35655AD881A94EFC4BA2C2422DF",
      "Account":"raz97dHvnyBcnYTbXGYxhV8bGyr1aPrE5w",
      "Destination":"rna8qC8Y9uLd2vzYtSEa1AJcdD3896zQ9S",
      "Memos":[
         {
            "Memo":{
               "MemoType":"636C69656E74",
               "MemoData":"726D2D312E322E34"
            }
         }
      ]
   },
   "meta":{
      "TransactionIndex":24,
      "DeliveredAmount":"3100",
      "TransactionResult":"tesSUCCESS",
      "delivered_amount":"3100"
   }
}
`

var paymentDst2 = blockatlas.Tx{
	ID:     "3D8512E02414EF5A6BC00281D945735E85DED9EF739B1DCA9EABE04D9EEC72C1",
	Coin:   coin.XRP,
	From:   "raz97dHvnyBcnYTbXGYxhV8bGyr1aPrE5w",
	To:     "rna8qC8Y9uLd2vzYtSEa1AJcdD3896zQ9S",
	Fee:    "120",
	Date:   1565114281,
	Block:  49163909,
	Status: blockatlas.StatusCompleted,
	Memo:   "",
	Meta: blockatlas.Transfer{
		Value: "3100",
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "payment",
		apiResponse: paymentSrc,
		expected:    &paymentDst,
	})

	testNormalize(t, &test{
		name:        "payment",
		apiResponse: paymentSrc2,
		expected:    &paymentDst2,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var payment Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &payment)
	if err != nil {
		t.Error(err)
		return
	}
	tx, ok := NormalizeTx(&payment)
	if !ok {
		t.Errorf("%s: tx could not be normalized", _test.name)
	}

	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error(_test.name + ": tx don't equal")
	}
}
