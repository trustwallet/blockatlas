package cosmos

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const basicSrc = `
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

var basicDst = blockatlas.Tx{
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

func TestNormalize(t *testing.T) {
	var srcTx Tx
	err := json.Unmarshal([]byte(basicSrc), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx := Normalize(&srcTx)
	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&basicDst)
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
