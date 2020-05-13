package blockbook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func TestNormalizePage(t *testing.T) {
	type args struct {
		srcPage   string
		address   string
		token     string
		coinIndex uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test normalize blockbook txs",
			args: args{
				srcPage: `{
					"transactions": [
					  {
						"txid": "0xb1a32935f9b015bcfdda1b2e3d281b3780d1a6f7a2d4406e05ec2b826b2349cb",
						"vin": [
						  {
							"n": 0,
							"addresses": [
							  "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
							],
							"isAddress": true
						  }
						],
						"vout": [
						  {
							"value": "0",
							"n": 0,
							"addresses": [
							  "0xc73e0383F3Aff3215E6f04B0331D58CeCf0Ab849"
							],
							"isAddress": true
						  }
						],
						"blockHash": "0x90d6b2e0fb0f983a15738206c8e9951db53624f5e9b29628fd8b71c5400430cb",
						"blockHeight": 8958320,
						"confirmations": 825021,
						"blockTime": 1574107019,
						"value": "0",
						"fees": "227056700000000",
						"tokenTransfers": [
						  {
							"type": "ERC20",
							"from": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
							"to": "0xc73e0383F3Aff3215E6f04B0331D58CeCf0Ab849",
							"token": "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359",
							"name": "Dai Stablecoin v1.0",
							"symbol": "DAI",
							"decimals": 18,
							"value": "2255656573089233195"
						  },
						  {
							"type": "ERC20",
							"from": "0xc73e0383F3Aff3215E6f04B0331D58CeCf0Ab849",
							"to": "0xad37fd42185Ba63009177058208dd1be4b136e6b",
							"token": "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359",
							"name": "Dai Stablecoin v1.0",
							"symbol": "DAI",
							"decimals": 18,
							"value": "2255656573089233195"
						  },
						  {
							"type": "ERC20",
							"from": "0x0000000000000000000000000000000000000000",
							"to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
							"token": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
							"name": "Dai Stablecoin",
							"symbol": "DAI",
							"decimals": 18,
							"value": "2255656573089233195"
						  }
						],
						"ethereumSpecific": {
						  "status": 1,
						  "nonce": 378,
						  "gasLimit": 3703313,
						  "gasUsed": 174659,
						  "gasPrice": "1300000000"
						}
					  },
					  {
						"txid": "0x17bb2b5e61f34119d4d4fbfae406ad3d854f0a00f13013d77de9aab7179f183f",
						"vin": [
							{
								"n": 0,
								"addresses": [
									"0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
								],
								"isAddress": true
							}
						],
						"vout": [
							{
								"value": "0",
								"n": 0,
								"addresses": [
									"0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359"
								],
								"isAddress": true
							}
						],
						"blockHash": "0xcee3a57858e3629785fb6e7ca34e68605fe3d2f149b73138f38314a3ef935723",
						"blockHeight": 9852430,
						"confirmations": 67071,
						"blockTime": 1586627561,
						"value": "0",
						"fees": "87378000000000",
						"tokenTransfers": [
							{
								"type": "ERC20",
								"from": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
								"to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
								"token": "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359",
								"name": "Dai Stablecoin v1.0",
								"symbol": "DAI",
								"decimals": 18,
								"value": "100000000000000"
							}
						],
						"ethereumSpecific": {
							"status": 1,
							"nonce": 523,
							"gasLimit": 43323,
							"gasUsed": 29126,
							"gasPrice": "3000000000"
						}
					},
					{
						"txid": "0x7a3929f2fad5e61f535ed5c1317f34e739655d582bc1b0161b9869b0957df6af",
						"vin": [
							{
								"n": 0,
								"addresses": [
									"0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
								],
								"isAddress": true
							}
						],
						"vout": [
							{
								"value": "567000000000000",
								"n": 0,
								"addresses": [
									"0x47331175b23C2f067204B506CA1501c26731C990"
								],
								"isAddress": true
							}
						],
						"blockHash": "0xf08fd4b1d6ace92bf9516bbed37de100025f8b0879a80a92359a08f37e788b95",
						"blockHeight": 10050786,
						"confirmations": 43,
						"blockTime": 1589278824,
						"value": "567000000000000",
						"fees": "407799043328112",
						"ethereumSpecific": {
							"status": 1,
							"nonce": 535,
							"gasLimit": 33000,
							"gasUsed": 21064,
							"gasPrice": "19360000158",
							"data": "0xdeadbeef"
						}
					},
					{
						"txid": "0xb1a56570bcb072d376630b987bd1f44ecc8f2c20ece52f02c9245296d3e3da39",
						"vin": [
							{
								"n": 0,
								"addresses": [
									"0x2A0A572d77F6d6Ce62C6539E679d943824c3b218"
								],
								"isAddress": true
							}
						],
						"vout": [
							{
								"value": "0",
								"n": 0,
								"addresses": [
									"0xdAC17F958D2ee523a2206206994597C13D831ec7"
								],
								"isAddress": true
							}
						],
						"blockHeight": -1,
						"confirmations": 0,
						"blockTime": 1589279659,
						"value": "0",
						"fees": "0",
						"rbf": true,
						"tokenTransfers": [
							{
								"type": "ERC20",
								"from": "0x2A0A572d77F6d6Ce62C6539E679d943824c3b218",
								"to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
								"token": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
								"name": "Tether USD",
								"symbol": "USDT",
								"decimals": 6,
								"value": "23000000"
							}
						],
						"ethereumSpecific": {
							"status": -1,
							"nonce": 15647,
							"gasLimit": 100000,
							"gasUsed": null,
							"gasPrice": "28560000000",
							"data": "0xa9059cbb000000000000000000000000595031ff9bf6e0c1de20349e1377f43934f8951400000000000000000000000000000000000000000000000000000000015ef3c0"
						}
					},
					{
						"txid": "0xfe7cce9928450e356f3332485e611781e407425b5888b8b2c7c66afaa4787cb1",
						"vin": [
							{
								"n": 0,
								"addresses": [
									"0x267be1C1D684F78cb4F6a176C4911b741E4Ffdc0"
								],
								"isAddress": true
							}
						],
						"vout": [
							{
								"value": "295000000000000000",
								"n": 0,
								"addresses": [
									"0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
								],
								"isAddress": true
							}
						],
						"blockHeight": -1,
						"confirmations": 0,
						"blockTime": 1589287339,
						"value": "295000000000000000",
						"fees": "0",
						"rbf": true,
						"ethereumSpecific": {
							"status": -1,
							"nonce": 1282636,
							"gasLimit": 30000,
							"gasUsed": null,
							"gasPrice": "24255000245",
							"data": "0x"
						}
					}
					]}`,
				address:   "0x7d8bf18c7ce84b3e175b339c4ca93aed1dd166f1",
				token:     "",
				coinIndex: 60,
			},
			want: `[{
					"id": "0xb1a32935f9b015bcfdda1b2e3d281b3780d1a6f7a2d4406e05ec2b826b2349cb",
					"coin": 60,
					"from": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
					"to": "0xc73e0383F3Aff3215E6f04B0331D58CeCf0Ab849",
					"fee": "227056700000000",
					"date": 1574107019,
					"block": 8958320,
					"status": "completed",
					"sequence": 378,
					"type": "contract_call",
					"direction": "outgoing",
					"memo": "",
					"metadata": {
						"input": "0x",
						"value": "0"
					}
				  },{
					"id": "0x17bb2b5e61f34119d4d4fbfae406ad3d854f0a00f13013d77de9aab7179f183f",
					"coin": 60,
					"from": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
					"to": "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359",
					"fee": "87378000000000",
					"date": 1586627561,
					"block": 9852430,
					"status": "completed",
					"sequence": 523,
					"type": "token_transfer",
					"direction": "yourself",
					"memo": "",
					"metadata": {
						"name": "Dai Stablecoin v1.0",
						"symbol": "DAI",
						"token_id": "0x89d24A6b4CcB1B6fAA2625fE562bDD9a23260359",
						"decimals": 18,
						"value": "100000000000000",
						"from": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
						"to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
					}
				  },{
					"id": "0x7a3929f2fad5e61f535ed5c1317f34e739655d582bc1b0161b9869b0957df6af",
					"coin": 60,
					"from": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
					"to": "0x47331175b23C2f067204B506CA1501c26731C990",
					"fee": "407799043328112",
					"date": 1589278824,
					"block": 10050786,
					"status": "completed",
					"sequence": 535,
					"type": "contract_call",
					"direction": "outgoing",
					"memo": "",
					"metadata": {
						"input": "0xdeadbeef",
						"value": "567000000000000"
					}
				  },{
					"id": "0xb1a56570bcb072d376630b987bd1f44ecc8f2c20ece52f02c9245296d3e3da39",
					"coin": 60,
					"from": "0x2A0A572d77F6d6Ce62C6539E679d943824c3b218",
					"to": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
					"fee": "2856000000000000",
					"date": 1589279659,
					"block": 0,
					"status": "pending",
					"sequence": 15647,
					"type": "token_transfer",
					"direction": "incoming",
					"memo": "",
					"metadata": {
						"name": "Tether USD",
						"symbol": "USDT",
						"token_id": "0xdAC17F958D2ee523a2206206994597C13D831ec7",
						"decimals": 6,
						"value": "23000000",
						"from": "0x2A0A572d77F6d6Ce62C6539E679d943824c3b218",
						"to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
					}
				  },{
					"id": "0xfe7cce9928450e356f3332485e611781e407425b5888b8b2c7c66afaa4787cb1",
					"coin": 60,
					"from": "0x267be1C1D684F78cb4F6a176C4911b741E4Ffdc0",
					"to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1",
					"fee": "727650007350000",
					"date": 1589287339,
					"block": 0,
					"status": "pending",
					"sequence": 1282636,
					"type": "transfer",
					"direction": "incoming",
					"memo": "",
					"metadata": {
						"value": "295000000000000000",
						"symbol": "ETH",
						"decimals": 18
					}
				  }
				  ]`,
		},
	}
	for _, tt := range tests {
		var page Page
		var txPage blockatlas.TxPage
		err := json.Unmarshal([]byte(tt.args.srcPage), &page)
		assert.Nil(t, err)
		err = json.Unmarshal([]byte(tt.want), &txPage)
		assert.Nil(t, err)
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePage(&page, tt.args.address, tt.args.token, tt.args.coinIndex)
			gotJson, err := json.Marshal(got)
			assert.Nil(t, err)
			gotTxPage, err := json.Marshal(txPage)
			assert.Nil(t, err)
			if string(gotJson) != string(gotTxPage) {
				t.Errorf("NormalizePage() = %v, want %v", string(gotJson), string(gotTxPage))
			}
		})
	}
}
