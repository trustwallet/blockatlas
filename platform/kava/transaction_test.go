package kava

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
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
                "denom": "ukava",
                "amount": "2271999999"
              }
            ]
          }
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "ukava",
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
}`

const transferSrcKava = `
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
      "value": "kava17wcggpjx007uc09s8y4hwrj8f228mlwe0n0upn"
    },
    {
      "key": "recipient",
      "value": "kava1z89utvygweg5l56fsk8ak7t6hh88fd0agl98n0"
    }
  ],
  "tx": {
    "type": "auth/StdTx",
    "value": {
      "msg": [
        {
          "type": "cosmos-sdk/MsgSend",
          "value": {
            "from_address": "kava17wcggpjx007uc09s8y4hwrj8f228mlwe0n0upn",
            "to_address": "kava1z89utvygweg5l56fsk8ak7t6hh88fd0agl98n0",
            "amount": [
              {
                "denom": "ukava",
                "amount": "2271999999"
              }
            ]
          }
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "ukava",
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
}`

const transferSrcKavaToken = `
{
    "height": "645538",
    "txhash": "514D43780335A1C516850FEE5692F59E9A9DF1D8D986FC62AC434BEB58EDB8E2",
    "raw_log": "[{\"msg_index\":0,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"},{\"key\":\"sender\",\"value\":\"kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7\"},{\"key\":\"module\",\"value\":\"bank\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"kava1wj2swfmeakxdrlqemvpzx2a4ljux9l4xq6qmcn\"},{\"key\":\"sender\",\"value\":\"kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7\"},{\"key\":\"amount\",\"value\":\"10000000000hard\"}]}]}]",
    "logs": [
    {
        "msg_index": 0,
        "log": "",
        "events": [
        {
            "type": "message",
            "attributes": [
            {
                "key": "action",
                "value": "send"
            },
            {
                "key": "sender",
                "value": "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7"
            },
            {
                "key": "module",
                "value": "bank"
            }
            ]
        },
        {
            "type": "transfer",
            "attributes": [
            {
                "key": "recipient",
                "value": "kava1wj2swfmeakxdrlqemvpzx2a4ljux9l4xq6qmcn"
            },
            {
                "key": "sender",
                "value": "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7"
            },
            {
                "key": "amount",
                "value": "10000000000hard"
            }
            ]
        }
        ]
    }
    ],
    "gas_wanted": "200000",
    "gas_used": "74794",
    "tx": {
    "type": "cosmos-sdk/StdTx",
    "value": {
    "msg": [
    {
        "type": "cosmos-sdk/MsgSend",
        "value": {
        "from_address": "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7",
        "to_address": "kava1wj2swfmeakxdrlqemvpzx2a4ljux9l4xq6qmcn",
        "amount": [
        {
            "denom": "hard",
            "amount": "10000000000"
        }
        ]
    }
    }
    ],
    "fee": {
    "amount": [
    {
        "denom": "ukava",
        "amount": "30"
    }
    ],
    "gas": "200000"
},
    "signatures": [
    {
        "pub_key": {
        "type": "tendermint/PubKeySecp256k1",
        "value": "A/Uf73aHoxZXS7SvKweW4Ijo723YDLuhwhir6JSmVmWe"
    },
        "signature": "Cuse3VunNLXi19GxASZD0fFMWeelEKe8DnC7nIYyS6FGqXzN+KP9hAdct7g9jczGT+4emVsQtiLvuM9Racrnyg=="
    }
    ],
    "memo": ""
}
},
    "timestamp": "2020-12-19T10:49:24Z"
}`

const failedTransferSrc = `
{
  "height": "5552",
  "txhash": "5E78C65A8C1A6C8239EBBBBF2E42020E6ADBA8037EDEA83BF88E1A9159CF13B8",
  "code": 12,
  "raw_log": "{\"codespace\":\"sdk\",\"code\":12,\"message\":\"out of gas in location: WritePerByte; gasWanted: 40000, gasUsed: 40480\"}",
  "gas_wanted": "40000",
  "gas_used": "40480",
  "tx": {
    "type": "cosmos-sdk/StdTx",
    "value": {
      "msg": [
        {
          "type": "cosmos-sdk/MsgSend",
          "value": {
            "from_address": "cosmos1shpfyt7psrff2ux7nznxvj6f7gq59fcqng5mku",
            "to_address": "cosmos1za4pu5gxm80fg6sx0956f88l2sx7jfg2vf7nlc",
            "amount": [
              {
                "denom": "ukava",
                "amount": "100000"
              }
            ]
          }
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "ukava",
            "amount": "2000"
          }
        ],
        "gas": "40000"
      },
      "signatures": [
        {
          "pub_key": {
            "type": "tendermint/PubKeySecp256k1",
            "value": "A0IDIokqw01U2YcdylvqD/sJHW5w9puS5vZWSf2GUaqL"
          },
          "signature": "1Kwp4dBZUbVV6Fk8AFcmNfSqi7MXFfqyLvHexFZXoqcKh+sNuezry89RhDAWgSMNLyaK20hI2XcUyks+Vo4QEQ=="
        }
      ],
      "memo": "UniCoins registration rewards"
    }
  },
  "timestamp": "2019-12-12T03:21:42Z"
}`

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
                     "denom":"ukava",
                     "amount":"49920"
                  }
               }
            }
         ],
         "fee":{  
            "amount":[  
               {  
                  "denom":"ukava",
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
}`

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
                     "denom":"ukava",
                     "amount":"5100000000"
                  }
               }
            }
         ],
         "fee":{  
            "amount":[  
               {  
                  "denom":"ukava",
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
}`

const claimRewardSrc1 = `
{
  "height": "79678",
  "txhash": "C382DCFDC30E2DA294421DAEAD5862F118592A7B000EE91F6BEF8452A1F525D7",
  "gas_wanted": "1600000",
  "gas_used": "492252",
  "tx": {
    "type": "cosmos-sdk/StdTx",
    "value": {
      "msg": [
        {
          "type": "cosmos-sdk/MsgWithdrawDelegationReward",
          "value": {
            "delegator_address": "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
            "validator_address": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5"
          }
        },
        {
          "type": "cosmos-sdk/MsgWithdrawDelegationReward",
          "value": {
            "delegator_address": "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
            "validator_address": "cosmosvaloper1fhr7e04ct0zslmkzqt9smakg3sxrdve6ulclj2"
          }
        },
        {
          "type": "cosmos-sdk/MsgWithdrawDelegationReward",
          "value": {
            "delegator_address": "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
            "validator_address": "cosmosvaloper1we6knm8qartmmh2r0qfpsz6pq0s7emv3e0meuw"
          }
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "ukava",
            "amount": "1000"
          }
        ],
        "gas": "1600000"
      },
      "memo": ""
    }
  },
  "timestamp": "2019-12-18T03:04:33Z",
  "events": [
    {
      "type": "transfer",
      "attributes": [
        {
          "key": "recipient",
          "value": "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug"
        }
      ]
    },
    {
      "type": "withdraw_rewards",
      "attributes": [
        {
          "key": "amount",
          "value": "1138ukava"
        },
        {
          "key": "validator",
          "value": "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5"
        },
        {
          "key": "amount",
          "value": "40612ukava"
        },
        {
          "key": "validator",
          "value": "cosmosvaloper1fhr7e04ct0zslmkzqt9smakg3sxrdve6ulclj2"
        },
        {
          "key": "amount",
          "value": "954ukava"
        },
        {
          "key": "validator",
          "value": "cosmosvaloper1we6knm8qartmmh2r0qfpsz6pq0s7emv3e0meuw"
        },
        {
          "key": "amount",
          "value": "43574ukava"
        },
        {
          "key": "amount"
        }
      ]
    }
  ]
}`

const claimRewardSrc2 = `
{
  "height": "54561",
  "txhash": "082BA88EC055A7C343A353297EAC104CE87C659E0DDD84621C9AC3C284232800",
  "gas_wanted": "300000",
  "gas_used": "156772",
  "tx": {
    "type": "cosmos-sdk/StdTx",
    "value": {
      "msg": [
        {
          "type": "cosmos-sdk/MsgWithdrawDelegationReward",
          "value": {
            "delegator_address": "cosmos1y6yvdel7zys8x60gz9067fjpcpygsn62ae9x46",
            "validator_address": "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2"
          }
        },
        {
          "type": "cosmos-sdk/MsgDelegate",
          "value": {
            "delegator_address": "cosmos1y6yvdel7zys8x60gz9067fjpcpygsn62ae9x46",
            "validator_address": "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
            "amount": {
              "denom": "ukava",
              "amount": "2692326"
            }
          }
        }
      ],
      "fee": {
        "amount": [
          {
            "denom": "ukava",
            "amount": "0"
          }
        ],
        "gas": "300000"
      },
      "memo": "复投"
    }
  },
  "timestamp": "2019-12-16T02:21:03Z",
  "events": [
    {
      "type": "delegate",
      "attributes": [
        {
          "key": "validator",
          "value": "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2"
        },
        {
          "key": "amount",
          "value": "2692326"
        }
      ]
    },
    {
      "type": "message",
      "attributes": [
        {
          "key": "sender",
          "value": "cosmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd88lyufl"
        },
        {
          "key": "module",
          "value": "distribution"
        }
      ]
    },
    {
      "type": "transfer",
      "attributes": [
        {
          "key": "recipient",
          "value": "cosmos1y6yvdel7zys8x60gz9067fjpcpygsn62ae9x46"
        },
        {
          "key": "amount",
          "value": "2692701ukava"
        }
      ]
    },
    {
      "type": "withdraw_rewards",
      "attributes": [
        {
          "key": "amount",
          "value": "2692701ukava"
        },
        {
          "key": "validator",
          "value": "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2"
        }
      ]
    }
  ]
}`

var transferDst = types.Tx{
	ID:     "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
	Coin:   coin.COSMOS,
	From:   "cosmos1rw62phusuv9vzraezr55k0vsqssvz6ed52zyrl",
	To:     "cosmos1nynns8ex9fq6sjjfj8k79ymkdz4sqth06xexae",
	Fee:    "1",
	Date:   1556992677,
	Block:  151980,
	Status: types.StatusCompleted,
	Type:   types.TxTransfer,
	Meta: types.Transfer{
		Value:    "2271999999",
		Symbol:   coin.Cosmos().Symbol,
		Decimals: 6,
	},
}

var transferDstKava = types.Tx{
	ID:     "E19B011D20D862DA0BEA7F24E3BC6DFF666EE6E044FCD9BD95B073478086DBB6",
	Coin:   coin.KAVA,
	From:   "kava17wcggpjx007uc09s8y4hwrj8f228mlwe0n0upn",
	To:     "kava1z89utvygweg5l56fsk8ak7t6hh88fd0agl98n0",
	Fee:    "1",
	Date:   1556992677,
	Block:  151980,
	Status: types.StatusCompleted,
	Type:   types.TxTransfer,
	Meta: types.Transfer{
		Value:    "2271999999",
		Symbol:   coin.Kava().Symbol,
		Decimals: 6,
	},
}

var transferDstKavaToken = types.Tx{
	ID:     "514D43780335A1C516850FEE5692F59E9A9DF1D8D986FC62AC434BEB58EDB8E2",
	Coin:   coin.KAVA,
	From:   "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7",
	To:     "kava1wj2swfmeakxdrlqemvpzx2a4ljux9l4xq6qmcn",
	Fee:    "30",
	Date:   1608374964,
	Block:  645538,
	Status: types.StatusCompleted,
	Type:   types.TxNativeTokenTransfer,
	Meta: types.NativeTokenTransfer{
		Value:    "10000000000",
		Symbol:   "HARD",
		Decimals: 6,
		From:     "kava1ys70jvnajkv88529ys6urjcyle3k2j9r24g6a7",
		To:       "kava1wj2swfmeakxdrlqemvpzx2a4ljux9l4xq6qmcn",
		TokenID:  "hard",
		Name:     "hard",
	},
}

var delegateDst = types.Tx{
	ID:        "11078091D1D5BD84F4275B6CEE02170428944DB0E8EEC37E980551435F6D04C7",
	Coin:      coin.COSMOS,
	From:      "cosmos1237l0vauhw78qtwq045jd24ay4urpec6r3xfn3",
	To:        "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
	Fee:       "5000",
	Date:      1564632616,
	Block:     1258202,
	Status:    types.StatusCompleted,
	Type:      types.TxAnyAction,
	Direction: types.DirectionOutgoing,
	Meta: types.AnyAction{
		Coin:     coin.COSMOS,
		Title:    types.AnyActionDelegation,
		Key:      types.KeyStakeDelegate,
		Name:     coin.Cosmos().Name,
		Symbol:   coin.Coins[coin.COSMOS].Symbol,
		Decimals: coin.Coins[coin.COSMOS].Decimals,
		Value:    "49920",
	},
}

var unDelegateDst = types.Tx{
	ID:        "A1EC36741FEF681F4A77B8F6032AD081100EE5ECB4CC76AEAC2174BC6B871CFE",
	Coin:      coin.COSMOS,
	From:      "cosmos137rrp4p8n0nqcft0mwc62tdnyhhzf80knv5t94",
	To:        "cosmosvaloper1te8nxpc2myjfrhaty0dnzdhs5ahdh5agzuym9v",
	Fee:       "5000",
	Date:      1564624521,
	Block:     1257037,
	Status:    types.StatusCompleted,
	Type:      types.TxAnyAction,
	Direction: types.DirectionIncoming,
	Meta: types.AnyAction{
		Coin:     coin.COSMOS,
		Title:    types.AnyActionUndelegation,
		Key:      types.KeyStakeDelegate,
		Name:     coin.Cosmos().Name,
		Symbol:   coin.Coins[coin.COSMOS].Symbol,
		Decimals: coin.Coins[coin.COSMOS].Decimals,
		Value:    "5100000000",
	},
}

var claimRewardDst2 = types.Tx{
	ID:        "082BA88EC055A7C343A353297EAC104CE87C659E0DDD84621C9AC3C284232800",
	Coin:      coin.COSMOS,
	From:      "cosmos1y6yvdel7zys8x60gz9067fjpcpygsn62ae9x46",
	To:        "cosmosvaloper12w6tynmjzq4l8zdla3v4x0jt8lt4rcz5gk7zg2",
	Fee:       "0",
	Date:      1576462863,
	Block:     54561,
	Status:    types.StatusCompleted,
	Type:      types.TxAnyAction,
	Direction: types.DirectionIncoming,
	Memo:      "复投",
	Meta: types.AnyAction{
		Coin:     coin.COSMOS,
		Title:    types.AnyActionClaimRewards,
		Key:      types.KeyStakeClaimRewards,
		Name:     coin.Cosmos().Name,
		Symbol:   coin.Coins[coin.COSMOS].Symbol,
		Decimals: coin.Coins[coin.COSMOS].Decimals,
		Value:    "2692701",
	},
}

var claimRewardDst1 = types.Tx{
	ID:        "C382DCFDC30E2DA294421DAEAD5862F118592A7B000EE91F6BEF8452A1F525D7",
	Coin:      coin.COSMOS,
	From:      "cosmos1cxehfdhfm96ljpktdxsj0k6xp9gtuheghwgqug",
	To:        "cosmosvaloper1ptyzewnns2kn37ewtmv6ppsvhdnmeapvtfc9y5",
	Fee:       "1000",
	Date:      1576638273,
	Block:     79678,
	Status:    types.StatusCompleted,
	Type:      types.TxAnyAction,
	Direction: types.DirectionIncoming,
	Memo:      "",
	Meta: types.AnyAction{
		Coin:     coin.COSMOS,
		Title:    types.AnyActionClaimRewards,
		Key:      types.KeyStakeClaimRewards,
		Name:     coin.Cosmos().Name,
		Symbol:   coin.Coins[coin.COSMOS].Symbol,
		Decimals: coin.Coins[coin.COSMOS].Decimals,
		Value:    "86278",
	},
}

var failedTransferDst = types.Tx{
	ID:     "5E78C65A8C1A6C8239EBBBBF2E42020E6ADBA8037EDEA83BF88E1A9159CF13B8",
	Coin:   coin.COSMOS,
	From:   "cosmos1shpfyt7psrff2ux7nznxvj6f7gq59fcqng5mku",
	To:     "cosmos1za4pu5gxm80fg6sx0956f88l2sx7jfg2vf7nlc",
	Fee:    "2000",
	Date:   1576120902,
	Block:  5552,
	Status: types.StatusError,
	Type:   types.TxTransfer,
	Memo:   "UniCoins registration rewards",
	Meta: types.Transfer{
		Value:    "100000",
		Symbol:   coin.Cosmos().Symbol,
		Decimals: 6,
	},
}

type test struct {
	name     string
	platform Platform
	Data     string
	want     types.Tx
}

type filterTest struct {
	name     string
	platform Platform
	Data     []string
	want     types.Txs
}

func TestNormalize(t *testing.T) {

	cosmos := Platform{CoinIndex: coin.COSMOS}
	kava := Platform{CoinIndex: coin.KAVA}

	tests := []test{
		{
			"test transfer tx",
			cosmos,
			transferSrc,
			transferDst,
		},
		{
			"test delegate tx",
			cosmos,
			delegateSrc,
			delegateDst,
		},
		{
			"test undelegate tx",
			cosmos,
			unDelegateSrc,
			unDelegateDst,
		},
		{
			"test claimReward tx 1",
			cosmos,
			claimRewardSrc1,
			claimRewardDst1,
		},
		{
			"test claimReward tx 2",
			cosmos,
			claimRewardSrc2,
			claimRewardDst2,
		},
		{
			"test failed tx",
			cosmos,
			failedTransferSrc,
			failedTransferDst,
		},
		{
			"test kava transfer tx",
			kava,
			transferSrcKava,
			transferDstKava,
		},
		{
			"test kava transfer token tx",
			kava,
			transferSrcKavaToken,
			transferDstKavaToken,
		},
	}
	for _, tt := range tests {
		testNormalize(t, tt)
	}
}

func TestFilterByDenom(t *testing.T) {

	kava := Platform{CoinIndex: coin.KAVA}

	test := filterTest{
		"test transfer tx",
		kava,
		[]string{
			transferSrc,
			delegateSrc,
			unDelegateSrc,
			claimRewardSrc1,
			claimRewardSrc2,
			failedTransferSrc,
			transferSrcKava,
			transferSrcKavaToken,
		},
		types.Txs{transferDstKavaToken},
	}

	testFilter(t, test)
}

func testFilter(t *testing.T, test filterTest) {
	t.Run(test.name, func(t *testing.T) {
		srcTxs := make([]Tx, 0)
		for _, tx := range test.Data {
			var srcTx Tx
			err := json.Unmarshal([]byte(tx), &srcTx)
			assert.Nil(t, err)
			srcTxs = append(srcTxs, srcTx)
		}
		tx := test.platform.FilterTxsByDenom(srcTxs, "hard")
		assert.Equal(t, len(test.want), len(tx), "length: filtered to expected 1 tx")
		normalized, ok := test.platform.Normalize(&tx[0])
		assert.True(t, ok)
		assert.Equal(t, test.want[0], normalized, "denom: filtered txs are equal")
	})
}

func testNormalize(t *testing.T, tt test) {
	t.Run(tt.name, func(t *testing.T) {
		var srcTx Tx
		err := json.Unmarshal([]byte(tt.Data), &srcTx)
		assert.Nil(t, err)
		tx, ok := tt.platform.Normalize(&srcTx)
		assert.True(t, ok)
		assert.Equal(t, tt.want, tx, "transfer: tx don't equal")
	})
}
