package blockbook

import (
	"encoding/json"
	"testing"

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
					  }
					]}`,
				address:   "0x7d8bf18c7ce84b3e175b339c4ca93aed1dd166f1",
				token:     "0x6b175474e89094c44da98b954eedeac495271d0f",
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
					"type": "token_transfer",
					"direction": "incoming",
					"memo": "",
					"metadata": {
					  "name": "Dai Stablecoin",
					  "symbol": "DAI",
					  "token_id": "0x6B175474E89094C44Da98b954EedeAC495271d0F",
					  "decimals": 18,
					  "value": "2255656573089233195",
					  "from": "0x0000000000000000000000000000000000000000",
					  "to": "0x7d8bf18C7cE84b3E175b339c4Ca93aEd1dD166F1"
					}
				}]`,
		},
	}
	for _, tt := range tests {
		var page Page
		var txPage blockatlas.TxPage
		_ = json.Unmarshal([]byte(tt.args.srcPage), &page)
		_ = json.Unmarshal([]byte(tt.want), &txPage)
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePage(&page, tt.args.address, tt.args.token, tt.args.coinIndex)
			gotJson, _ := json.Marshal(got)
			gotTxPage, _ := json.Marshal(txPage)
			if string(gotJson) != string(gotTxPage) {
				t.Errorf("NormalizePage() = %v, want %v", string(gotJson), string(gotTxPage))
			}
		})
	}
}
