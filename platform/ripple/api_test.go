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
      "AffectedNodes":[
         {
            "ModifiedNode":{
               "LedgerEntryType":"Offer",
               "PreviousTxnLgrSeq":49163824,
               "PreviousTxnID":"35FCE78C3747B2E289A95CB0836CD3AB70029AD058B0F30D72F7768BF97A2E60",
               "LedgerIndex":"0687311959AF1E758610E8E449B5EDF0B9400ECD5BA52B4BDEF6420A0DCD38B7",
               "PreviousFields":{
                  "TakerPays":{
                     "value":"3.218003992015969",
                     "currency":"USD",
                     "issuer":"rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
                  },
                  "TakerGets":"9993801"
               },
               "FinalFields":{
                  "Flags":131072,
                  "Sequence":32207,
                  "BookNode":"0000000000000000",
                  "OwnerNode":"0000000000000022",
                  "BookDirectory":"79C54A4EBD69AB2EADCE313042F36092BE432423CC6A4F784E0B7092AC2D4000",
                  "TakerPays":{
                     "value":"3.217005988023954",
                     "currency":"USD",
                     "issuer":"rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
                  },
                  "TakerGets":"9990701",
                  "Account":"ra5J6KL9fbt6EeNt6c1eea3J7BsQJBPApi"
               }
            }
         },
         {
            "ModifiedNode":{
               "LedgerEntryType":"AccountRoot",
               "PreviousTxnLgrSeq":49163824,
               "PreviousTxnID":"35FCE78C3747B2E289A95CB0836CD3AB70029AD058B0F30D72F7768BF97A2E60",
               "LedgerIndex":"2F8B26EAEA18184A1648D8C5456D452D41B32AECB62666FA3A5B537384F613A2",
               "PreviousFields":{
                  "Sequence":115,
                  "Balance":"110066899"
               },
               "FinalFields":{
                  "Flags":0,
                  "Sequence":116,
                  "OwnerCount":2,
                  "Balance":"110066779",
                  "Account":"raz97dHvnyBcnYTbXGYxhV8bGyr1aPrE5w"
               }
            }
         },
         {
            "ModifiedNode":{
               "LedgerEntryType":"RippleState",
               "PreviousTxnLgrSeq":49163824,
               "PreviousTxnID":"35FCE78C3747B2E289A95CB0836CD3AB70029AD058B0F30D72F7768BF97A2E60",
               "LedgerIndex":"5BDDF7E599C2DFBE7EE798D40F50FFBCF6BA5DF3FF825A66D2C8E3618FB48491",
               "PreviousFields":{
                  "Balance":{
                     "value":"-1593.040776451775",
                     "currency":"USD",
                     "issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"
                  }
               },
               "FinalFields":{
                  "Flags":2228224,
                  "LowNode":"0000000000001750",
                  "HighNode":"0000000000000000",
                  "Balance":{
                     "value":"-1593.041774455767",
                     "currency":"USD",
                     "issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"
                  },
                  "LowLimit":{
                     "value":"0",
                     "currency":"USD",
                     "issuer":"rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
                  },
                  "HighLimit":{
                     "value":"1000000000",
                     "currency":"USD",
                     "issuer":"ra5J6KL9fbt6EeNt6c1eea3J7BsQJBPApi"
                  }
               }
            }
         },
         {
            "ModifiedNode":{
               "LedgerEntryType":"AccountRoot",
               "PreviousTxnLgrSeq":48920623,
               "PreviousTxnID":"A234278D7B95426036995FC1B564AF15F7015AEA3EACB7EE3C4C700182B854C0",
               "LedgerIndex":"79882F4BD34221999EAD8921FDE8FEF0F381F4357B55ADCF7DB443A2A7D1A2F5",
               "PreviousFields":{
                  "Balance":"32404286"
               },
               "FinalFields":{
                  "Flags":0,
                  "Sequence":6,
                  "OwnerCount":0,
                  "Balance":"32407386",
                  "Account":"rna8qC8Y9uLd2vzYtSEa1AJcdD3896zQ9S"
               }
            }
         },
         {
            "ModifiedNode":{
               "LedgerEntryType":"RippleState",
               "PreviousTxnLgrSeq":49163824,
               "PreviousTxnID":"35FCE78C3747B2E289A95CB0836CD3AB70029AD058B0F30D72F7768BF97A2E60",
               "LedgerIndex":"B0468ED19407ECE973CCF6B110390634CD47ECF269EFCA7DA5D1840CADDC8457",
               "PreviousFields":{
                  "Balance":{
                     "value":"-0.4289",
                     "currency":"USD",
                     "issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"
                  }
               },
               "FinalFields":{
                  "Flags":2228224,
                  "LowNode":"0000000000001B3F",
                  "HighNode":"0000000000000000",
                  "Balance":{
                     "value":"-0.4279",
                     "currency":"USD",
                     "issuer":"rrrrrrrrrrrrrrrrrrrrBZbvji"
                  },
                  "LowLimit":{
                     "value":"0",
                     "currency":"USD",
                     "issuer":"rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
                  },
                  "HighLimit":{
                     "value":"0",
                     "currency":"USD",
                     "issuer":"raz97dHvnyBcnYTbXGYxhV8bGyr1aPrE5w"
                  }
               }
            }
         },
         {
            "ModifiedNode":{
               "LedgerEntryType":"AccountRoot",
               "PreviousTxnLgrSeq":49163904,
               "PreviousTxnID":"1DBA32E68D05AF401727324878B56A34B602DCE68B346D2BB5DC30A5AD339BC4",
               "LedgerIndex":"F18B050AFC1118348416EA55F3DA471D8F81DC222B1FEA4A738CEF477D68B874",
               "PreviousFields":{
                  "Balance":"66409650297"
               },
               "FinalFields":{
                  "Flags":0,
                  "Sequence":32210,
                  "OwnerCount":93,
                  "Balance":"66409647197",
                  "Account":"ra5J6KL9fbt6EeNt6c1eea3J7BsQJBPApi"
               }
            }
         }
      ],
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
