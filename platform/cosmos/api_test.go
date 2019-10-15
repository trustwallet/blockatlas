package cosmos

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
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
   "height":"1258202",
   "txhash":"11078091D1D5BD84F4275B6CEE02170428944DB0E8EEC37E980551435F6D04C7",
   "raw_log":"[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
   "logs":[  
      {  
         "msg_index":"0",
         "success":true,
         "log":""
      }
   ],
   "gas_wanted":"200000",
   "gas_used":"103206",
   "tags":[  
      {  
         "key":"action",
         "value":"delegate"
      },
      {  
         "key":"delegator",
         "value":"cosmos1237l0vauhw78qtwq045jd24ay4urpec6r3xfn3"
      },
      {  
         "key":"destination-validator",
         "value":"cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2"
      }
   ],
   "tx":{  
      "type":"auth/StdTx",
      "value":{  
         "msg":[  
            {  
               "type":"cosmos-sdk/MsgDelegate",
               "value":{  
                  "delegator_address":"cosmos1237l0vauhw78qtwq045jd24ay4urpec6r3xfn3",
                  "validator_address":"cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
                  "amount":{  
                     "denom":"uatom",
                     "amount":"49920"
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
                  "value":"AsZL4GaIEGW6ogh1rEasxHtmirpeBnycLz4VR0rSVr9p"
               },
               "signature":"w6sNVzTSsE32ERbBdYYySSp6nj+4xNODuq5GKRVb8q04jMHUbx9AhuZeAhYrkvdkzOl3bD7vRYGx9P1V6yHj0A=="
            }
         ],
         "memo":""
      }
   },
   "timestamp":"2019-08-01T04:10:16Z"
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

const delegationsSrc = `
[
  {
    "delegator_address": "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
    "validator_address": "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
    "shares": "1999999.999931853807876751"
  }
]`

const unbondingDelegationsSrc = `
[
  {
    "delegator_address": "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
    "validator_address": "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
    "entries": [
      {
        "creation_height": "1780365",
        "completion_time": "2019-10-03T05:37:26.350018207Z",
        "initial_balance": "5000000",
        "balance": "5000000"
      }
    ]
  }
]`

var transferDst = blockatlas.Tx{
	ID:     "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
	Coin:   coin.ATOM,
	From:   "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
	To:     "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae",
	Fee:    "1",
	Date:   1556992677,
	Block:  151980,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "2271999999",
		Symbol:   "ATOM",
		Decimals: 6,
	},
}

var delegateDst = blockatlas.Tx{
	ID:     "11078091D1D5BD84F4275B6CEE02170428944DB0E8EEC37E980551435F6D04C7",
	Coin:   coin.ATOM,
	From:   "cosmos1237l0vauhw78qtwq045jd24ay4urpec6r3xfn3",
	To:     "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
	Fee:    "5000",
	Date:   1564632616,
	Block:  1258202,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    blockatlas.AnyActionDelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     "ATOM",
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    "49920",
	},
}

var unDelegateDst = blockatlas.Tx{
	ID:     "A1EC36741FEF681F4A77B8F6032AD081100EE5ECB4CC76AEAC2174BC6B871CFE",
	Coin:   coin.ATOM,
	From:   "cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94",
	To:     "cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
	Fee:    "5000",
	Date:   1564624521,
	Block:  1257037,
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.AnyAction{
		Coin:     coin.ATOM,
		Title:    blockatlas.AnyActionUndelegation,
		Key:      blockatlas.KeyStakeDelegate,
		Name:     "ATOM",
		Symbol:   coin.Coins[coin.ATOM].Symbol,
		Decimals: coin.Coins[coin.ATOM].Decimals,
		Value:    "5100000000",
	},
}

var stakingPool = StakingPool{"1222", "200"}

var cosmosValidator = Validator{Commission: CosmosCommission{Rate: "0.4"}}

var inflation = 0.7

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
	var v Validator
	_ = json.Unmarshal([]byte(validatorSrc), &v)
	coin := coin.Coin{}
	expected := blockatlas.Validator{
		Status: true,
		ID:     v.Address,
		Details: blockatlas.StakingDetails{
			Reward:        blockatlas.StakingReward{Annual: 435.48749999999995},
			LockTime:      1814400,
			MinimumAmount: "0",
			Type:          blockatlas.DelegationTypeDelegate,
		},
	}

	result := normalizeValidator(v, stakingPool, inflation, coin)

	assert.Equal(t, result, expected)
}

func TestCalculateAnnualReward(t *testing.T) {
	result := CalculateAnnualReward(StakingPool{"1222", "200"}, inflation, cosmosValidator)
	assert.Equal(t, result, 298.61999703347686)
}

var validator1 = blockatlas.StakeValidator{
	ID:     "cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys",
	Status: true,
	Info: blockatlas.StakeValidatorInfo{
		Name:        "Certus One",
		Description: "Stake and earn rewards with the most secure and stable validator. Winner of the Game of Stakes. Operated by Certus One Inc. By delegating, you confirm that you are aware of the risk of slashing and that Certus One Inc is not liable for any potential damages to your investment.",
		Image:       "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/cosmos/validators/assets/cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys/logo.png",
		Website:     "https://certus.one",
	},
	Details: blockatlas.StakingDetails{
		Reward: blockatlas.StakingReward{
			Annual: 9.259735525366604,
		},
		LockTime:      1814400,
		MinimumAmount: "0",
	},
}

var validatorMap = blockatlas.ValidatorMap{
	"cosmosvaloper1qwl879nx9t6kef4supyazayf7vjhennyh568ys": validator1,
}

func TestNormalizeDelegations(t *testing.T) {
	var delegations []Delegation
	err := json.Unmarshal([]byte(delegationsSrc), &delegations)
	assert.NoError(t, err)
	assert.NotNil(t, delegations)

	expected := []blockatlas.Delegation{
		{
			Delegator: validator1,
			Value:     "1999999",
			Status:    blockatlas.DelegationStatusActive,
		},
	}
	result := NormalizeDelegations(delegations, validatorMap)
	assert.Equal(t, result, expected)
}

func TestNormalizeUnbondingDelegations(t *testing.T) {
	var delegations []UnbondingDelegation
	err := json.Unmarshal([]byte(unbondingDelegationsSrc), &delegations)
	assert.NoError(t, err)
	assert.NotNil(t, delegations)

	expected := []blockatlas.Delegation{
		{
			Delegator: validator1,
			Value:     "5000000",
			Status:    blockatlas.DelegationStatusPending,
			Metadata: blockatlas.DelegationMetaDataPending{
				AvailableDate: 1570081046,
			},
		},
	}
	result := NormalizeUnbondingDelegations(delegations, validatorMap)
	assert.Equal(t, result, expected)
}
