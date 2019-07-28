package ripple

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
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
		"AffectedNodes": [
			{
				"CreatedNode": {
					"LedgerEntryType": "AccountRoot",
					"LedgerIndex": "564241023DCB6F74760910F17F78B179AEC159C701BBACD99A1D3259D77D3CFF",
					"NewFields": {
						"Sequence": 1,
						"Balance": "100000000",
						"Account": "rMQ98K56yXJbDGv49ZSmW51sLn94Xe1mu1"
					}
				}
			},
			{
				"ModifiedNode": {
					"LedgerEntryType": "AccountRoot",
					"PreviousTxnLgrSeq": 34698098,
					"PreviousTxnID": "7040A099F51E0DC386B909FB4C01DCCF23CB61D3D05B0EC562C01359FB60C754",
					"LedgerIndex": "D242D4E3501E5829AB003BA788CF361D4717419D9653304E556A14C6166847E8",
					"PreviousFields": {
						"Sequence": 21,
						"Balance": "1999935364"
					},
					"FinalFields": {
						"Flags": 0,
						"Sequence": 22,
						"OwnerCount": 0,
						"Balance": "1899932249",
						"Account": "rGSxFjoqmWz54PycrgQBQ5dB6e7TUpMxzq"
					}
				}
			}
		],
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
	Meta: blockatlas.Transfer{
		Value: "100000000",
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
}

func testNormalize(t *testing.T, _test *test) {
	var payment Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &payment)
	if err != nil {
		t.Error(err)
		return
	}
	tx, ok := Normalize(&payment)
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
