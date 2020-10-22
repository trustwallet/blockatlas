package trustray

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"testing"
)

const tokenTransferSrc = `
{
    "operations": [
        {
            "transactionId": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2-0",
            "contract": {
                "address": "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
                "symbol": "KBC",
                "decimals": 7,
                "totalSupply": "120000000000000000",
                "name": "KaratBank Coin"
            },
            "from": "0xd35F30D194684a391C63A6decEd7d3dd5207c265",
            "to": "0xaA4D790076f1Bf7511a0A0AC498C89e13e1eFE17",
            "type": "token_transfer",
            "value": "4291000000",
            "coin": 60
        }
    ],
    "contract": null,
    "_id": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
    "blockNumber": 7491945,
    "time": 1554248437,
    "nonce": 88,
    "from": "0xd35f30d194684a391c63a6deced7d3dd5207c265",
    "to": "0xf3586684107ce0859c44aa2b2e0fb8cd8731a15a",
    "value": "0",
    "gas": "67497",
    "gasPrice": "6900000256",
    "gasUsed": "51921",
    "input": "0xa9059cbb000000000000000000000000aa4d790076f1bf7511a0a0ac498c89e13e1efe1700000000000000000000000000000000000000000000000000000000ffc376c0",
    "error": "",
    "id": "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
    "coin": 60
}`

const contractCallSrc = `
{
	"addresses": [
		"0x09862ed5908c0a336f9f92e5ffeb9768deac6091"
	],
	"operations": [],
	"contract": "0xe4dc0f23b1a3f2c47dc288a22f72e100e9b1cd70",
	"_id": "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	"blockNumber": 7522627,
	"time": 1554661737,
	"nonce": 534,
	"from": "0xc9a16a82c284efc3cb0fe8c891ab85d6eba0eefb",
	"to": "0xc67f9c909c4d185e4a5d21d642c27d05a145a76c",
	"value": "1800000000000000000",
	"gas": "1000000",
	"gasPrice": "2000000000",
	"gasUsed": "21340",
	"input": "0xfffdefefed",
	"error": "",
	"id": "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	"coin": 60
}
`

const transferSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	"blockNumber": 7522781,
	"time": 1554663642,
	"nonce": 88,
	"from": "0xf5aea47e57c058881b31ee8fce1002c409188f06",
	"to": "0x0ae933a89d9e249d0873cfc7ca022fcb3f1280ce",
	"value": "1999895000000000000",
	"gas": "21000",
	"gasPrice": "5000000000",
	"gasUsed": "21000",
	"input": "0x",
	"error": "",
	"id": "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	"coin": 60
}`

const failedSrc = `
{
	"operations": [],
	"contract": null,
	"_id": "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	"blockNumber": 7522678,
	"time": 1554662399,
	"nonce": 1,
	"from": "0x4b55af7ce28a113d794f9a9940fe1506f37aa619",
	"to": "0xe65f787c8561a4b15771111bb427274dedfe85d7",
	"value": "59859820000000000",
	"gas": "21000",
	"gasPrice": "3000000000",
	"gasUsed": "21000",
	"input": "0x",
	"error": "Error",
	"id": "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	"coin": 60
}`

var (
	addr1     = "0xd35F30D194684a391C63A6decEd7d3dd5207c265"
	addr2     = "0xaA4D790076f1Bf7511a0A0AC498C89e13e1eFE17"
	contract1 = "0xf3586684107CE0859c44aa2b2E0fB8cd8731a15a"
)
var tokenTransferDst = blockatlas.Tx{
	ID:       "0x7777854580f273df61e0162e1a41b3e1e05ab8b9f553036fa9329a90dd7e9ab2",
	Coin:     coin.ETH,
	From:     addr1,
	To:       contract1,
	Fee:      "358254913291776",
	Date:     1554248437,
	Block:    7491945,
	Sequence: 88,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.TokenTransfer{
		Name:     "KaratBank Coin",
		Symbol:   "KBC",
		TokenID:  contract1,
		Decimals: 7,
		Value:    "4291000000",
		From:     addr1,
		To:       addr2,
	},
}

var contractCallDst = blockatlas.Tx{
	ID:       "0x34ab0028a9aa794d5cc12887e7b813cec17889948276b301028f24a408da6da4",
	Coin:     coin.ETH,
	From:     "0xc9a16a82c284EFC3cB0fE8C891ab85d6EBa0EeFB",
	To:       "0xc67f9C909C4d185E4A5d21D642c27D05A145a76c",
	Fee:      "42680000000000",
	Date:     1554661737,
	Block:    7522627,
	Sequence: 534,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.ContractCall{
		Input: "0xfffdefefed",
		Value: "1800000000000000000",
	},
}

var transferDst = blockatlas.Tx{
	ID:       "0x77f8a3b2203933493d103a1637de814b4853410b1fb2981c4d2cff4d7a3071ab",
	Coin:     coin.ETH,
	From:     "0xf5AeA47E57c058881B31EE8fcE1002C409188F06",
	To:       "0x0Ae933A89D9E249D0873cfc7CA022fCB3F1280Ce",
	Fee:      "105000000000000",
	Date:     1554663642,
	Block:    7522781,
	Sequence: 88,
	Status:   blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "1999895000000000000",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

var failedDst = blockatlas.Tx{
	ID:       "0x8dfe7e859f7bdcea4e6f4ada18567d96a51c3aa29e618ef09b80ae99385e191e",
	Coin:     coin.ETH,
	From:     "0x4b55af7cE28A113D794F9A9940fe1506f37aA619",
	To:       "0xE65f787c8561A4b15771111bb427274deDfe85D7",
	Fee:      "63000000000000",
	Date:     1554662399,
	Block:    7522678,
	Sequence: 1,
	Status:   blockatlas.StatusError,
	Error:    "Error",
	Meta: blockatlas.Transfer{
		Value:    "59859820000000000",
		Symbol:   "ETH",
		Decimals: 18,
	},
}

func TestNormalize(t *testing.T) {
	var (
		doc   Doc
		tests = []struct {
			name, apiResponse string
			expected          *blockatlas.Tx
		}{
			{"transfer", transferSrc, &transferDst},
			{"token transfer", tokenTransferSrc, &tokenTransferDst},
			{"contract call", contractCallSrc, &contractCallDst},
			{"failed transaction", failedSrc, &failedDst},
		}
	)

	t.Run("TestNormalize", func(t *testing.T) {
		for _, tt := range tests {
			err := json.Unmarshal([]byte(tt.apiResponse), &doc)
			if err != nil {
				t.Error(err)
				return
			}
			res := AppendTxs(nil, &doc, coin.ETH)

			resJSON, err := json.Marshal(res)
			if err != nil {
				t.Fatal(err)
			}

			dstJSON, err := json.Marshal([]blockatlas.Tx{*tt.expected})
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(resJSON, dstJSON) {
				println("\n", "Test failed ", tt.name)
				println("resJSON", string(resJSON))
				println("dstJSON", string(dstJSON))
				t.Error(tt.name + ": tx don't equal")
			}
		}
	})
}
