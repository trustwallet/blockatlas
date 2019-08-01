package cosmos

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const transferSrc = `
{
	"height": "151980",
	"txhash": "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
	"raw_log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
	"logs": [
	  {
		"msg_index": "0",
		"success": true,
		"log": ""
	  }
	],
	"gas_wanted": "100000",
	"gas_used": "27678",
	"tags": [
	  {
		"key": "action",
		"value": "send"
	  },
	  {
		"key": "sender",
		"value": "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl"
	  },
	  {
		"key": "recipient",
		"value": "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae"
	  }
	],
	"tx": {
	  "type": "auth/StdTx",
	  "value": {
		"msg": [
		  {
			"type": "cosmos-sdk/MsgSend",
			"value": {
			  "from_address": "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
			  "to_address": "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae",
			  "amount": [
				{
				  "denom": "uatom",
				  "amount": "2271999999"
				}
			  ]
			}
		  }
		],
		"fee": {
		  "amount": [
			{
			  "denom": "uatom",
			  "amount": "1"
			}
		  ],
		  "gas": "100000"
		},
		"signatures": [
		  {
			"pub_key": {
			  "type": "tendermint/PubKeySecp256k1",
			  "value": "A21fdP6IbVC9hER5smiim8I4EbFeIF/bW81IKwmmsdjH"
			},
			"signature": "MuR85p714L94tCenogRqzLh1bsbmhKTjs1L9JJPdhSVwQKh61EGlLqYGoUeN/n9xb+OOR9ESUOh2CAzVulKoVQ=="
		  }
		],
		"memo": ""
	  }
	},
	"timestamp": "2019-05-04T17:57:57Z"
  }
`

const delegateSrc = `
{  
   "height":"193844",
   "txhash":"35399455114314D7DB99B2FD913705A0FFB5899336BD579495A5C9057B209AB7",
   "raw_log":"[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
   "logs":[  
      {  
         "msg_index":"0",
         "success":true,
         "log":""
      }
   ],
   "gas_wanted":"200000",
   "gas_used":"99610",
   "tags":[  
      {  
         "key":"action",
         "value":"delegate"
      },
      {  
         "key":"delegator",
         "value":"cosmos1rzwuwnwnwu7j8qtw2c8qynygfluzlz5he634v3"
      },
      {  
         "key":"destination-validator",
         "value":"cosmosvaloper1ey69r37gfxvxg62sh4r0ktpuc46pzjrm873ae8"
      }
   ],
   "tx":{  
      "type":"auth/StdTx",
      "value":{  
         "msg":[  
            {  
               "type":"cosmos-sdk/MsgDelegate",
               "value":{  
                  "delegator_address":"cosmos1rzwuwnwnwu7j8qtw2c8qynygfluzlz5he634v3",
                  "validator_address":"cosmosvaloper1ey69r37gfxvxg62sh4r0ktpuc46pzjrm873ae8",
                  "amount":{  
                     "denom":"uatom",
                     "amount":"6789000000"
                  }
               }
            }
         ],
         "fee":{  
            "amount":null,
            "gas":"200000"
         },
         "signatures":[  
            {  
               "pub_key":{  
                  "type":"tendermint/PubKeySecp256k1",
                  "value":"ApcD+YmCvMKWG9rcodwUy5ueqDXpmN0csE7hUok6Ol1M"
               },
               "signature":"B7WgRyemnQFEP/zd7c7Mq7mDLDcr92OZuQeRG19ehaElxLKuXkGGIfjoA0xCXRikmQrKlwSxiZ5QgQ2wb3FsLg=="
            }
         ],
         "memo":""
      }
   },
   "timestamp":"2019-05-08T01:53:36Z"
}
`

const unDelegateSrc = `
{  
   "height":"1257037",
   "txhash":"A1EC36741FEF681F4A77B8F6032AD081100EE5ECB4CC76AEAC2174BC6B871CFE",
   "data":"0C0889ECF7EA0510FB9D8CAD03",
   "raw_log":"[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
   "logs":[  
      {  
         "msg_index":"0",
         "success":true,
         "log":""
      }
   ],
   "gas_wanted":"200000",
   "gas_used":"107804",
   "tags":[  
      {  
         "key":"action",
         "value":"begin_unbonding"
      },
      {  
         "key":"delegator",
         "value":"cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94"
      },
      {  
         "key":"source-validator",
         "value":"cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v"
      },
      {  
         "key":"end-time",
         "value":"2019-08-22T01:55:21Z"
      }
   ],
   "tx":{  
      "type":"auth/StdTx",
      "value":{  
         "msg":[  
            {  
               "type":"cosmos-sdk/MsgUndelegate",
               "value":{  
                  "delegator_address":"cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94",
                  "validator_address":"cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
                  "amount":{  
                     "denom":"uatom",
                     "amount":"5100000000"
                  }
               }
            }
         ],
         "fee":{  
            "amount":[  
               {  
                  "denom":"uatom",
                  "amount":"5000"
               }
            ],
            "gas":"200000"
         },
         "signatures":[  
            {  
               "pub_key":{  
                  "type":"tendermint/PubKeySecp256k1",
                  "value":"A+tPzMXCW7vxmW5VN9Q/CO+fxnEXYlSMOklDVgaFutQD"
               },
               "signature":"rh25A/RTm8TUTUGOGhufqxn9vLFef/04xEKMJLUD5QhBVabRADvEgAP1J842XTDtVBS0SpVD/MrPduqRp0nNzg=="
            }
         ],
         "memo":""
      }
   },
   "timestamp":"2019-08-01T01:55:21Z"
}
`

const validatorSrc = `
{
    "operator_address": "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
    "consensus_pubkey": "cosmosvalconspub1zcjduepqwrjpn0slu86e32zfu5xxg8l42uk40guuw6er44vw2yl6s7wc38est6l0ux",
    "jailed": false,
    "status": 2,
    "tokens": "9538882295763",
    "delegator_shares": "9538882295763.000000000000000000",
    "description": {
      "moniker": "Certus One",
      "identity": "ABD51DF68C0D1ECF",
      "website": "https://certus.one",
      "details": "Stake and earn rewards with the most secure and stable validator."
    },
    "unbonding_height": "0",
    "unbonding_time": "1970-01-01T00:00:00Z",
    "commission": {
      "rate": "0.125000000000000000",
      "max_rate": "0.300000000000000000",
      "max_change_rate": "0.010000000000000000",
      "update_time": "2019-03-13T23:00:00Z"
    },
    "min_self_delegation": "1"
  }
`

var transferDst = blockatlas.Tx{
	ID:    "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
	Coin:  coin.ATOM,
	From:  "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
	To:    "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae",
	Fee:   "1",
	Date:  1556992677,
	Block: 151980,
	Meta: blockatlas.Transfer{
		Value: "2271999999",
	},
}

var delegateDst = blockatlas.Tx{
	ID:    "35399455114314D7DB99B2FD913705A0FFB5899336BD579495A5C9057B209AB7",
	Coin:  coin.ATOM,
	From:  "cosmos1rzwuwnwnwu7j8qtw2c8qynygfluzlz5he634v3",
	To:    "cosmosvaloper1ey69r37gfxvxg62sh4r0ktpuc46pzjrm873ae8",
	Fee:   "0",
	Date:  1557280416,
	Block: 193844,
	Meta: blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    blockatlas.AnyActionDelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    "6789000000",
	},
}

var unDelegateDst = blockatlas.Tx{
	ID:    "A1EC36741FEF681F4A77B8F6032AD081100EE5ECB4CC76AEAC2174BC6B871CFE",
	Coin:  coin.ATOM,
	From:  "cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94",
	To:    "cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
	Fee:   "5000",
	Date:  1564624521,
	Block: 1257037,
	Meta: blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    blockatlas.AnyActionUndelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    "5100000000",
	},
}

func TestNormalize(t *testing.T) {
	testNormalize(t, transferSrc, &transferDst)
	testNormalize(t, delegateSrc, &delegateDst)
	testNormalize(t, unDelegateSrc, &unDelegateDst)
}

func testNormalize(t *testing.T, src string, dst *blockatlas.Tx) {
	var srcTx Tx
	err := json.Unmarshal([]byte(src), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx, _ := Normalize(&srcTx)
	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&dst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("basic: tx don't equal")
	}
}

func TestNormalizeValidator(t *testing.T) {
	var v CosmosValidator
	_ = json.Unmarshal([]byte(validatorSrc), &v)

	expected := blockatlas.StakeValidator{
		Status: true,
		Info: blockatlas.StakeValidatorInfo{
			Name:        v.Description.Moniker,
			Description: v.Description.Description,
			Website:     v.Description.Website,
		},
		Address:   v.Operator_Address,
		PublicKey: v.Consensus_Pubkey,
	}

	result := normalizeValidator(v)

	assert.Equal(t, result, expected)
}
