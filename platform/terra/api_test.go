package terra

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/trustwallet/blockatlas/coin"
)

const transferSrc = `
{
  "tx": {
      "type": "core/StdTx",
      "value": {
          "fee": {
              "gas": "78436",
              "amount": [
                  {
                      "denom": "ukrw",
                      "amount": "1177"
                  }
              ]
          },
          "msg": [
              {
                  "type": "bank/MsgSend",
                  "value": {
                      "amount": [
                          {
                              "denom": "ukrw",
                              "amount": "480000000000"
                          },
                          {
                              "denom": "uluna",
                              "amount": "1771645906"
                          }
                      ],
                      "to_address": "terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna",
                      "from_address": "terra1rf9xakxf97a49qa5svsf7yypjswzkutqfclur8"
                  }
              }
          ],
          "memo": "",
          "signatures": [
              {
                  "pub_key": {
                      "type": "tendermint/PubKeySecp256k1",
                      "value": "AyyQtraMl3kEYPXpEbdHUgUsjoixdJAbcIAO8AxAlCcN"
                  },
                  "signature": "Ysl1xngiESTCVVgrsENHuwmcfUPAZ03p1saFwQ7MuBUSW6lO9ZjXQVv9N3gDMPqHGNAignqYGofpAPuB6O/wNA=="
              }
          ]
      }
  },
  "logs": [
      {
          "log": {
              "tax": "1612000000ukrw"
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1rf9xakxf97a49qa5svsf7yypjswzkutqfclur8"
                      },
                      {
                          "key": "module",
                          "value": "bank"
                      },
                      {
                          "key": "action",
                          "value": "send"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna"
                      },
                      {
                          "key": "amount",
                          "value": "480000000000ukrw,1771645906uluna"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 0
      }
  ],
  "events": [
      {
          "type": "message",
          "attributes": [
              {
                  "key": "sender",
                  "value": "terra1rf9xakxf97a49qa5svsf7yypjswzkutqfclur8"
              },
              {
                  "key": "module",
                  "value": "bank"
              },
              {
                  "key": "action",
                  "value": "send"
              }
          ]
      },
      {
          "type": "transfer",
          "attributes": [
              {
                  "key": "recipient",
                  "value": "terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna"
              },
              {
                  "key": "amount",
                  "value": "480000000000ukrw,1771645906uluna"
              }
          ]
      }
  ],
  "height": "45478",
  "txhash": "06011507E4F3EF150C92F2DDB217499F2020801B7323C9C3ADDE1138916B0F98",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1rf9xakxf97a49qa5svsf7yypjswzkutqfclur8\"},{\"key\":\"module\",\"value\":\"bank\"},{\"key\":\"action\",\"value\":\"send\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna\"},{\"key\":\"amount\",\"value\":\"480000000000ukrw,1771645906uluna\"}]}]}]",
  "gas_used": "56017",
  "timestamp": "2019-12-17T04:03:56Z",
  "gas_wanted": "78436"
}`

const failedTransferSrc = `
{
  "tx": {
      "type": "core/StdTx",
      "value": {
          "fee": {
              "gas": "77672",
              "amount": [
                  {
                      "denom": "ukrw",
                      "amount": "1166"
                  }
              ]
          },
          "msg": [
              {
                  "type": "bank/MsgSend",
                  "value": {
                      "amount": [
                          {
                              "denom": "ukrw",
                              "amount": "430000000000"
                          }
                      ],
                      "to_address": "terra1dvghtnsqr6eusxxhqcmuhwmpw26rze8kgap823",
                      "from_address": "terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna"
                  }
              }
          ],
          "memo": "3642766313",
          "signatures": [
              {
                  "pub_key": {
                      "type": "tendermint/PubKeySecp256k1",
                      "value": "A4p3L23DzwwM6JnbLyY1xdgAl5ewiYPBQU+cD7Jzqwu7"
                  },
                  "signature": "+gXx9pLJBtuGfA+9tUkSg5a2WzVzx+5fneMuPhdIAiEwH0ITEN3iftueJAyi0k6WmZQ//d8O8mE7QoDXLVOn/g=="
              }
          ]
      }
  },
  "code": 10,
  "logs": [
      {
          "log": {
              "tax": "1597268616ukrw",
              "code": 10,
              "message": "insufficient account funds; 429833294565ukrw,99994738uluna,76621806umnt,14899518usdr,7168723uusd < 430000000000ukrw",
              "codespace": "sdk"
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "action",
                          "value": "send"
                      }
                  ]
              }
          ],
          "success": false,
          "msg_index": 0
      }
  ],
  "events": [
      {
          "type": "message",
          "attributes": [
              {
                  "key": "action",
                  "value": "send"
              }
          ]
      }
  ],
  "height": "138132",
  "txhash": "41B98DB9DE01BD0464718D1DAACD56C741C61A16ADD19FEDEBF7F6F1AAF57141",
  "raw_log": "[{\"msg_index\":0,\"success\":false,\"log\":\"{\\\"codespace\\\":\\\"sdk\\\",\\\"code\\\":10,\\\"message\\\":\\\"insufficient account funds; 429833294565ukrw,99994738uluna,76621806umnt,14899518usdr,7168723uusd \\u003c 430000000000ukrw\\\"}\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"send\"}]}]}]",
  "gas_used": "38634",
  "timestamp": "2019-12-24T03:29:45Z",
  "gas_wanted": "77672"
}`

const delegateSrc = `
{
  "tx": {
      "type": "core/StdTx",
      "value": {
          "fee": {
              "gas": "219200",
              "amount": [
                  {
                      "denom": "uluna",
                      "amount": "3288"
                  }
              ]
          },
          "msg": [
              {
                  "type": "staking/MsgDelegate",
                  "value": {
                      "amount": {
                          "denom": "uluna",
                          "amount": "1000000"
                      },
                      "delegator_address": "terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c",
                      "validator_address": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
                  }
              }
          ],
          "memo": "",
          "signatures": [
              {
                  "pub_key": {
                      "type": "tendermint/PubKeySecp256k1",
                      "value": "A3+qIVFTL74hqg9OAbMvVBRN0QNVBxch3ZbPTmbtXEdy"
                  },
                  "signature": "vua/PpnaQbb+2uKw2iV4nIpf0hl6ZiEBEGMGmDb+YApZDFC0rNHqclLcpdpfs/YlI5B5ObB8doRfLRbLCV9bTA=="
              }
          ]
      }
  },
  "logs": [
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "delegate",
                  "attributes": [
                      {
                          "key": "validator",
                          "value": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
                      },
                      {
                          "key": "amount",
                          "value": "1000000"
                      }
                  ]
              },
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "staking"
                      },
                      {
                          "key": "sender",
                          "value": "terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c"
                      },
                      {
                          "key": "action",
                          "value": "delegate"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c"
                      },
                      {
                          "key": "amount",
                          "value": "75152537ukrw,86293uluna,6306umnt,15usdr"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 0
      }
  ],
  "events": [
      {
          "type": "delegate",
          "attributes": [
              {
                  "key": "validator",
                  "value": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
              },
              {
                  "key": "amount",
                  "value": "1000000"
              }
          ]
      },
      {
          "type": "message",
          "attributes": [
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "staking"
              },
              {
                  "key": "sender",
                  "value": "terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c"
              },
              {
                  "key": "action",
                  "value": "delegate"
              }
          ]
      },
      {
          "type": "transfer",
          "attributes": [
              {
                  "key": "recipient",
                  "value": "terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c"
              },
              {
                  "key": "amount",
                  "value": "75152537ukrw,86293uluna,6306umnt,15usdr"
              }
          ]
      }
  ],
  "height": "311067",
  "txhash": "EF3756D7911A1D534580976A9AC2539146955EFBFA562000034BBD4FBFB70D07",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"delegate\",\"attributes\":[{\"key\":\"validator\",\"value\":\"terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9\"},{\"key\":\"amount\",\"value\":\"1000000\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c\"},{\"key\":\"action\",\"value\":\"delegate\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c\"},{\"key\":\"amount\",\"value\":\"75152537ukrw,86293uluna,6306umnt,15usdr\"}]}]}]",
  "gas_used": "146058",
  "timestamp": "2020-01-06T05:25:41Z",
  "gas_wanted": "219200"
}`

const unDelegateSrc = `
{
  "tx": {
      "type": "core/StdTx",
      "value": {
          "fee": {
              "gas": "248333",
              "amount": [
                  {
                      "denom": "uluna",
                      "amount": "3725"
                  }
              ]
          },
          "msg": [
              {
                  "type": "staking/MsgUndelegate",
                  "value": {
                      "amount": {
                          "denom": "uluna",
                          "amount": "15780040401"
                      },
                      "delegator_address": "terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa",
                      "validator_address": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
                  }
              }
          ],
          "memo": "",
          "signatures": [
              {
                  "pub_key": {
                      "type": "tendermint/PubKeySecp256k1",
                      "value": "ApmzZf8lLhYPbyFUaJ3jcWC8NpT1oZPAy/JuWuZZSkjC"
                  },
                  "signature": "zWUKOsz5+Q3QuYlWrFntq/VH2Gc86yREP0sx5QsRyT0mKHEFB9AgoyvXfiIsdXwtmZbb7786rPOIQfzOvAsMVg=="
              }
          ]
      }
  },
  "data": "0C0892DAB2F10510C2C395E702",
  "logs": [
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "sender",
                          "value": "terra1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3nln0mh"
                      },
                      {
                          "key": "module",
                          "value": "staking"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa"
                      },
                      {
                          "key": "action",
                          "value": "begin_unbonding"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa"
                      },
                      {
                          "key": "amount",
                          "value": "733044ukrw,36uluna"
                      },
                      {
                          "key": "recipient",
                          "value": "terra1tygms3xhhs3yv487phx3dw4a95jn7t7l8l07dr"
                      },
                      {
                          "key": "amount",
                          "value": "15780040401uluna"
                      }
                  ]
              },
              {
                  "type": "unbond",
                  "attributes": [
                      {
                          "key": "validator",
                          "value": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
                      },
                      {
                          "key": "amount",
                          "value": "15780040401"
                      },
                      {
                          "key": "completion_time",
                          "value": "2020-01-25T21:03:14Z"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 0
      }
  ],
  "events": [
      {
          "type": "message",
          "attributes": [
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "sender",
                  "value": "terra1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3nln0mh"
              },
              {
                  "key": "module",
                  "value": "staking"
              },
              {
                  "key": "sender",
                  "value": "terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa"
              },
              {
                  "key": "action",
                  "value": "begin_unbonding"
              }
          ]
      },
      {
          "type": "transfer",
          "attributes": [
              {
                  "key": "recipient",
                  "value": "terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa"
              },
              {
                  "key": "amount",
                  "value": "733044ukrw,36uluna"
              },
              {
                  "key": "recipient",
                  "value": "terra1tygms3xhhs3yv487phx3dw4a95jn7t7l8l07dr"
              },
              {
                  "key": "amount",
                  "value": "15780040401uluna"
              }
          ]
      },
      {
          "type": "unbond",
          "attributes": [
              {
                  "key": "validator",
                  "value": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
              },
              {
                  "key": "amount",
                  "value": "15780040401"
              },
              {
                  "key": "completion_time",
                  "value": "2020-01-25T21:03:14Z"
              }
          ]
      }
  ],
  "height": "293214",
  "txhash": "BF763CB36F46A6E90092A219898BDB2785CD9E8E698808B04C5EB30BD414F239",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"sender\",\"value\":\"terra1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3nln0mh\"},{\"key\":\"module\",\"value\":\"staking\"},{\"key\":\"sender\",\"value\":\"terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa\"},{\"key\":\"action\",\"value\":\"begin_unbonding\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa\"},{\"key\":\"amount\",\"value\":\"733044ukrw,36uluna\"},{\"key\":\"recipient\",\"value\":\"terra1tygms3xhhs3yv487phx3dw4a95jn7t7l8l07dr\"},{\"key\":\"amount\",\"value\":\"15780040401uluna\"}]},{\"type\":\"unbond\",\"attributes\":[{\"key\":\"validator\",\"value\":\"terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9\"},{\"key\":\"amount\",\"value\":\"15780040401\"},{\"key\":\"completion_time\",\"value\":\"2020-01-25T21:03:14Z\"}]}]}]",
  "gas_used": "165581",
  "timestamp": "2020-01-04T21:03:14Z",
  "gas_wanted": "248333"
}`

const claimRewardSrc1 = `
{
  "tx": {
      "type": "core/StdTx",
      "value": {
          "fee": {
              "gas": "1367533",
              "amount": [
                  {
                      "denom": "uluna",
                      "amount": "20513"
                  }
              ]
          },
          "msg": [
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1yad8pjqp93gvwkxa2aa5mh4vctzfs37ekjxr4s"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1gxa4zq407lx58ld5rxy6rqudg3yu4s0sknfc0m"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1gkumn82kkj3cww28yp53agy7aluxv06fsuynvd"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1fjuvyccn8hfmn5r7wc2t3kwqy09zzp6tyjcf50"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1vqnhgc6d0jyggtytzqrnsc40r4zez6tx99382w"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1dcrq2xwuhea9hm5xfuydjuwgz6gm7vdjz7e4uf"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper15zcjduavxc5mkp8qcqs9eyhwlqwdlrzy6jln3m"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1hg70rkal5d86fl57k0gc7de0rrk4klgs59r7jc"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1h0d5kq5p64jcyqysvja3h2gysxnfudk9h3a5rq"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1h6rf7y2ar5vz64q8rchz5443s3tqnswrpf4846"
                  }
              },
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
                      "validator_address": "terravaloper1c9ye54e3pzwm3e0zpdlel6pnavrj9qqvq89r3r"
                  }
              }
          ],
          "memo": "",
          "signatures": [
              {
                  "pub_key": {
                      "type": "tendermint/PubKeySecp256k1",
                      "value": "AyETa9Y9ihObzeRPWMP0MBAa0Mqune3I+5KonOCPTtkv"
                  },
                  "signature": "rGAAkY9YDxRZMFnPwVtlqBQRH2s4x3hjkv3M49Zq5JpmHeQwxijgHXYBAFZ7pjacVrUtJLQkhzfkKC409IyY6Q=="
              }
          ]
      }
  },
  "logs": [
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 0
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "11857971756ukrw,107799uluna,1181752umnt,3267usdr"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "11857971756ukrw,107799uluna,1181752umnt,3267usdr"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1yad8pjqp93gvwkxa2aa5mh4vctzfs37ekjxr4s"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 1
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1gxa4zq407lx58ld5rxy6rqudg3yu4s0sknfc0m"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 2
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1gkumn82kkj3cww28yp53agy7aluxv06fsuynvd"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 3
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "13849737864ukrw,127357uluna,1395749umnt,3889usdr,1uusd"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "13849737864ukrw,127357uluna,1395749umnt,3889usdr,1uusd"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1fjuvyccn8hfmn5r7wc2t3kwqy09zzp6tyjcf50"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 4
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "1267291473ukrw,10822uluna,132614umnt,367usdr"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "1267291473ukrw,10822uluna,132614umnt,367usdr"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1vqnhgc6d0jyggtytzqrnsc40r4zez6tx99382w"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 5
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "2408684318ukrw,22125uluna,243784umnt,675usdr"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "2408684318ukrw,22125uluna,243784umnt,675usdr"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1dcrq2xwuhea9hm5xfuydjuwgz6gm7vdjz7e4uf"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 6
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "145699518520ukrw,1030220uluna,14801645umnt,40336usdr,10uusd"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "145699518520ukrw,1030220uluna,14801645umnt,40336usdr,10uusd"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper15zcjduavxc5mkp8qcqs9eyhwlqwdlrzy6jln3m"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 7
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "34031556671ukrw,308031uluna,3534791umnt,9567usdr,2uusd"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "34031556671ukrw,308031uluna,3534791umnt,9567usdr,2uusd"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1hg70rkal5d86fl57k0gc7de0rrk4klgs59r7jc"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 8
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "38513911400ukrw,414uluna,3868050umnt,10909usdr,2uusd"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "38513911400ukrw,414uluna,3868050umnt,10909usdr,2uusd"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1h0d5kq5p64jcyqysvja3h2gysxnfudk9h3a5rq"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 9
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1h6rf7y2ar5vz64q8rchz5443s3tqnswrpf4846"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 10
      },
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
                      },
                      {
                          "key": "amount",
                          "value": "40243784823ukrw,356573uluna,4090704umnt,11480usdr,2uusd"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "40243784823ukrw,356573uluna,4090704umnt,11480usdr,2uusd"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1c9ye54e3pzwm3e0zpdlel6pnavrj9qqvq89r3r"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 11
      }
  ],
  "events": [
      {
          "type": "message",
          "attributes": [
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              },
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              }
          ]
      },
      {
          "type": "transfer",
          "attributes": [
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "11857971756ukrw,107799uluna,1181752umnt,3267usdr"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "13849737864ukrw,127357uluna,1395749umnt,3889usdr,1uusd"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "1267291473ukrw,10822uluna,132614umnt,367usdr"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "2408684318ukrw,22125uluna,243784umnt,675usdr"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "145699518520ukrw,1030220uluna,14801645umnt,40336usdr,10uusd"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "34031556671ukrw,308031uluna,3534791umnt,9567usdr,2uusd"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "38513911400ukrw,414uluna,3868050umnt,10909usdr,2uusd"
              },
              {
                  "key": "recipient",
                  "value": "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp"
              },
              {
                  "key": "amount",
                  "value": "40243784823ukrw,356573uluna,4090704umnt,11480usdr,2uusd"
              }
          ]
      },
      {
          "type": "withdraw_rewards",
          "attributes": [
              {
                  "key": "amount"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5"
              },
              {
                  "key": "amount",
                  "value": "11857971756ukrw,107799uluna,1181752umnt,3267usdr"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1yad8pjqp93gvwkxa2aa5mh4vctzfs37ekjxr4s"
              },
              {
                  "key": "amount"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1gxa4zq407lx58ld5rxy6rqudg3yu4s0sknfc0m"
              },
              {
                  "key": "amount"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1gkumn82kkj3cww28yp53agy7aluxv06fsuynvd"
              },
              {
                  "key": "amount",
                  "value": "13849737864ukrw,127357uluna,1395749umnt,3889usdr,1uusd"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1fjuvyccn8hfmn5r7wc2t3kwqy09zzp6tyjcf50"
              },
              {
                  "key": "amount",
                  "value": "1267291473ukrw,10822uluna,132614umnt,367usdr"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1vqnhgc6d0jyggtytzqrnsc40r4zez6tx99382w"
              },
              {
                  "key": "amount",
                  "value": "2408684318ukrw,22125uluna,243784umnt,675usdr"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1dcrq2xwuhea9hm5xfuydjuwgz6gm7vdjz7e4uf"
              },
              {
                  "key": "amount",
                  "value": "145699518520ukrw,1030220uluna,14801645umnt,40336usdr,10uusd"
              },
              {
                  "key": "validator",
                  "value": "terravaloper15zcjduavxc5mkp8qcqs9eyhwlqwdlrzy6jln3m"
              },
              {
                  "key": "amount",
                  "value": "34031556671ukrw,308031uluna,3534791umnt,9567usdr,2uusd"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1hg70rkal5d86fl57k0gc7de0rrk4klgs59r7jc"
              },
              {
                  "key": "amount",
                  "value": "38513911400ukrw,414uluna,3868050umnt,10909usdr,2uusd"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1h0d5kq5p64jcyqysvja3h2gysxnfudk9h3a5rq"
              },
              {
                  "key": "amount"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1h6rf7y2ar5vz64q8rchz5443s3tqnswrpf4846"
              },
              {
                  "key": "amount",
                  "value": "40243784823ukrw,356573uluna,4090704umnt,11480usdr,2uusd"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1c9ye54e3pzwm3e0zpdlel6pnavrj9qqvq89r3r"
              }
          ]
      }
  ],
  "height": "273300",
  "txhash": "BB1CD919307CFBE49A49C40863E2ADEB122FA37904EF108030F07933233B66CD",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5\"}]}]},{\"msg_index\":1,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"11857971756ukrw,107799uluna,1181752umnt,3267usdr\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"11857971756ukrw,107799uluna,1181752umnt,3267usdr\"},{\"key\":\"validator\",\"value\":\"terravaloper1yad8pjqp93gvwkxa2aa5mh4vctzfs37ekjxr4s\"}]}]},{\"msg_index\":2,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"terravaloper1gxa4zq407lx58ld5rxy6rqudg3yu4s0sknfc0m\"}]}]},{\"msg_index\":3,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"terravaloper1gkumn82kkj3cww28yp53agy7aluxv06fsuynvd\"}]}]},{\"msg_index\":4,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"13849737864ukrw,127357uluna,1395749umnt,3889usdr,1uusd\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"13849737864ukrw,127357uluna,1395749umnt,3889usdr,1uusd\"},{\"key\":\"validator\",\"value\":\"terravaloper1fjuvyccn8hfmn5r7wc2t3kwqy09zzp6tyjcf50\"}]}]},{\"msg_index\":5,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"1267291473ukrw,10822uluna,132614umnt,367usdr\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"1267291473ukrw,10822uluna,132614umnt,367usdr\"},{\"key\":\"validator\",\"value\":\"terravaloper1vqnhgc6d0jyggtytzqrnsc40r4zez6tx99382w\"}]}]},{\"msg_index\":6,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"2408684318ukrw,22125uluna,243784umnt,675usdr\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"2408684318ukrw,22125uluna,243784umnt,675usdr\"},{\"key\":\"validator\",\"value\":\"terravaloper1dcrq2xwuhea9hm5xfuydjuwgz6gm7vdjz7e4uf\"}]}]},{\"msg_index\":7,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"145699518520ukrw,1030220uluna,14801645umnt,40336usdr,10uusd\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"145699518520ukrw,1030220uluna,14801645umnt,40336usdr,10uusd\"},{\"key\":\"validator\",\"value\":\"terravaloper15zcjduavxc5mkp8qcqs9eyhwlqwdlrzy6jln3m\"}]}]},{\"msg_index\":8,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"34031556671ukrw,308031uluna,3534791umnt,9567usdr,2uusd\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"34031556671ukrw,308031uluna,3534791umnt,9567usdr,2uusd\"},{\"key\":\"validator\",\"value\":\"terravaloper1hg70rkal5d86fl57k0gc7de0rrk4klgs59r7jc\"}]}]},{\"msg_index\":9,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"38513911400ukrw,414uluna,3868050umnt,10909usdr,2uusd\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"38513911400ukrw,414uluna,3868050umnt,10909usdr,2uusd\"},{\"key\":\"validator\",\"value\":\"terravaloper1h0d5kq5p64jcyqysvja3h2gysxnfudk9h3a5rq\"}]}]},{\"msg_index\":10,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\"},{\"key\":\"validator\",\"value\":\"terravaloper1h6rf7y2ar5vz64q8rchz5443s3tqnswrpf4846\"}]}]},{\"msg_index\":11,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp\"},{\"key\":\"amount\",\"value\":\"40243784823ukrw,356573uluna,4090704umnt,11480usdr,2uusd\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"40243784823ukrw,356573uluna,4090704umnt,11480usdr,2uusd\"},{\"key\":\"validator\",\"value\":\"terravaloper1c9ye54e3pzwm3e0zpdlel6pnavrj9qqvq89r3r\"}]}]}]",
  "gas_used": "911745",
  "timestamp": "2020-01-03T08:59:01Z",
  "gas_wanted": "1367533"
}`

const claimRewardSrc2 = `
{
  "tx": {
      "type": "core/StdTx",
      "value": {
          "fee": {
              "gas": "201000",
              "amount": [
                  {
                      "denom": "uluna",
                      "amount": "3015"
                  }
              ]
          },
          "msg": [
              {
                  "type": "distribution/MsgWithdrawDelegationReward",
                  "value": {
                      "delegator_address": "terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk",
                      "validator_address": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
                  }
              }
          ],
          "memo": "",
          "signatures": [
              {
                  "pub_key": {
                      "type": "tendermint/PubKeySecp256k1",
                      "value": "Axy8q6tC162ln2UpRxS5hRC34gWLlVmV/I0j1eV7K6yt"
                  },
                  "signature": "f6RIrYdP/+BllEXP67CwIVJr4FrdgTtlnwsQI1JYd8klilkcuchFojPDc5ob/Q+2UonlKXveMUsGghjytldaLw=="
              }
          ]
      }
  },
  "logs": [
      {
          "log": {
              "tax": ""
          },
          "events": [
              {
                  "type": "message",
                  "attributes": [
                      {
                          "key": "sender",
                          "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
                      },
                      {
                          "key": "module",
                          "value": "distribution"
                      },
                      {
                          "key": "sender",
                          "value": "terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk"
                      },
                      {
                          "key": "action",
                          "value": "withdraw_delegator_reward"
                      }
                  ]
              },
              {
                  "type": "transfer",
                  "attributes": [
                      {
                          "key": "recipient",
                          "value": "terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk"
                      },
                      {
                          "key": "amount",
                          "value": "32181037226453ukrw,66292771689uluna,2697515175umnt,7903733usdr,943uusd"
                      }
                  ]
              },
              {
                  "type": "withdraw_rewards",
                  "attributes": [
                      {
                          "key": "amount",
                          "value": "32181037226453ukrw,66292771689uluna,2697515175umnt,7903733usdr,943uusd"
                      },
                      {
                          "key": "validator",
                          "value": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
                      }
                  ]
              }
          ],
          "success": true,
          "msg_index": 0
      }
  ],
  "events": [
      {
          "type": "message",
          "attributes": [
              {
                  "key": "sender",
                  "value": "terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl"
              },
              {
                  "key": "module",
                  "value": "distribution"
              },
              {
                  "key": "sender",
                  "value": "terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk"
              },
              {
                  "key": "action",
                  "value": "withdraw_delegator_reward"
              }
          ]
      },
      {
          "type": "transfer",
          "attributes": [
              {
                  "key": "recipient",
                  "value": "terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk"
              },
              {
                  "key": "amount",
                  "value": "32181037226453ukrw,66292771689uluna,2697515175umnt,7903733usdr,943uusd"
              }
          ]
      },
      {
          "type": "withdraw_rewards",
          "attributes": [
              {
                  "key": "amount",
                  "value": "32181037226453ukrw,66292771689uluna,2697515175umnt,7903733usdr,943uusd"
              },
              {
                  "key": "validator",
                  "value": "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9"
              }
          ]
      }
  ],
  "height": "177275",
  "txhash": "A023BFC68EB4529DDFB27A64E51A49EAC1FD1344A777957B4A047AE4415F4985",
  "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"sender\",\"value\":\"terra1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8pm7utl\"},{\"key\":\"module\",\"value\":\"distribution\"},{\"key\":\"sender\",\"value\":\"terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk\"},{\"key\":\"action\",\"value\":\"withdraw_delegator_reward\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk\"},{\"key\":\"amount\",\"value\":\"32181037226453ukrw,66292771689uluna,2697515175umnt,7903733usdr,943uusd\"}]},{\"type\":\"withdraw_rewards\",\"attributes\":[{\"key\":\"amount\",\"value\":\"32181037226453ukrw,66292771689uluna,2697515175umnt,7903733usdr,943uusd\"},{\"key\":\"validator\",\"value\":\"terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9\"}]}]}]",
  "gas_used": "133834",
  "timestamp": "2019-12-27T02:30:12Z",
  "gas_wanted": "201000"
}`

var transferDst = blockatlas.Tx{
	ID:     "06011507E4F3EF150C92F2DDB217499F2020801B7323C9C3ADDE1138916B0F98",
	Coin:   coin.LUNA,
	From:   "terra1rf9xakxf97a49qa5svsf7yypjswzkutqfclur8",
	To:     "terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna",
	Fee:    "0",
	Date:   1576555436,
	Block:  45478,
	Status: blockatlas.StatusCompleted,
	Type:   blockatlas.TxMultiCurrencyTransfer,
	Meta: blockatlas.MultiCurrencyTransfer{
		Currencies: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["ukrw"],
				Value:      "480000000000",
				CurrencyID: "ukrw",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "1771645906",
				CurrencyID: "uluna",
			},
		},
		Fees: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["ukrw"],
				Value:      "1177",
				CurrencyID: "ukrw",
			},
		},
	},
}

var failedTransferDst = blockatlas.Tx{
	ID:     "41B98DB9DE01BD0464718D1DAACD56C741C61A16ADD19FEDEBF7F6F1AAF57141",
	Coin:   coin.LUNA,
	From:   "terra1t849fxw7e8ney35mxemh4h3ayea4zf77dslwna",
	To:     "terra1dvghtnsqr6eusxxhqcmuhwmpw26rze8kgap823",
	Fee:    "0",
	Date:   1577158185,
	Block:  138132,
	Status: blockatlas.StatusFailed,
	Type:   blockatlas.TxMultiCurrencyTransfer,
	Memo:   "3642766313",
	Meta: blockatlas.MultiCurrencyTransfer{
		Currencies: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["ukrw"],
				Value:      "430000000000",
				CurrencyID: "ukrw",
			},
		},
		Fees: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["ukrw"],
				Value:      "1166",
				CurrencyID: "ukrw",
			},
		},
	},
}

var delegateDst = blockatlas.Tx{
	ID:        "EF3756D7911A1D534580976A9AC2539146955EFBFA562000034BBD4FBFB70D07",
	Coin:      coin.LUNA,
	From:      "terra1f5slu5vhxtxfdh6zg66rpg627d3r7lpsfaq55c",
	To:        "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9",
	Fee:       "0",
	Date:      1578288341,
	Block:     311067,
	Status:    blockatlas.StatusCompleted,
	Type:      blockatlas.TxMultiCurrencyAnyAction,
	Direction: blockatlas.DirectionOutgoing,
	Meta: blockatlas.MultiCurrencyAnyAction{
		Title: blockatlas.AnyActionDelegation,
		Key:   blockatlas.KeyStakeDelegate,
		Currencies: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "1000000",
				CurrencyID: "uluna",
			},
		},
		Fees: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "3288",
				CurrencyID: "uluna",
			},
		},
	},
}

var unDelegateDst = blockatlas.Tx{
	ID:        "BF763CB36F46A6E90092A219898BDB2785CD9E8E698808B04C5EB30BD414F239",
	Coin:      coin.LUNA,
	From:      "terra1e4jspgz5fteppglvy4a0xxn3uqlejsswysy3qa",
	To:        "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9",
	Fee:       "0",
	Date:      1578171794,
	Block:     293214,
	Status:    blockatlas.StatusCompleted,
	Type:      blockatlas.TxMultiCurrencyAnyAction,
	Direction: blockatlas.DirectionIncoming,
	Meta: blockatlas.MultiCurrencyAnyAction{
		Title: blockatlas.AnyActionUndelegation,
		Key:   blockatlas.KeyStakeDelegate,
		Currencies: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "15780040401",
				CurrencyID: "uluna",
			},
		},
		Fees: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "3725",
				CurrencyID: "uluna",
			},
		},
	},
}

var claimRewardDst1 = blockatlas.Tx{
	ID:        "BB1CD919307CFBE49A49C40863E2ADEB122FA37904EF108030F07933233B66CD",
	Coin:      coin.LUNA,
	From:      "terra1e82da9n6jz4t42eh0wn5hrt6hdmf7jyq8sufkp",
	To:        "terravaloper1rf9xakxf97a49qa5svsf7yypjswzkutqfhnpn5",
	Fee:       "0",
	Date:      1578041941,
	Block:     273300,
	Status:    blockatlas.StatusCompleted,
	Type:      blockatlas.TxMultiCurrencyAnyAction,
	Direction: blockatlas.DirectionIncoming,
	Memo:      "",
	Meta: blockatlas.MultiCurrencyAnyAction{
		Title: blockatlas.AnyActionClaimRewards,
		Key:   blockatlas.KeyStakeClaimRewards,
		Currencies: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["ukrw"],
				Value:      "287872456825",
				CurrencyID: "ukrw",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "1963341",
				CurrencyID: "uluna",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["umnt"],
				Value:      "29249089",
				CurrencyID: "umnt",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["usdr"],
				Value:      "80490",
				CurrencyID: "usdr",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uusd"],
				Value:      "17",
				CurrencyID: "uusd",
			},
		},
		Fees: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "20513",
				CurrencyID: "uluna",
			},
		},
	},
}

var claimRewardDst2 = blockatlas.Tx{
	ID:        "A023BFC68EB4529DDFB27A64E51A49EAC1FD1344A777957B4A047AE4415F4985",
	Coin:      coin.LUNA,
	From:      "terra1p54hc4yy2ajg67j645dn73w3378j6k05v52cnk",
	To:        "terravaloper1p54hc4yy2ajg67j645dn73w3378j6k05vmx9r9",
	Fee:       "0",
	Date:      1577413812,
	Block:     177275,
	Status:    blockatlas.StatusCompleted,
	Type:      blockatlas.TxMultiCurrencyAnyAction,
	Direction: blockatlas.DirectionIncoming,
	Memo:      "",
	Meta: blockatlas.MultiCurrencyAnyAction{
		Title: blockatlas.AnyActionClaimRewards,
		Key:   blockatlas.KeyStakeClaimRewards,
		Currencies: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["ukrw"],
				Value:      "32181037226453",
				CurrencyID: "ukrw",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "66292771689",
				CurrencyID: "uluna",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["umnt"],
				Value:      "2697515175",
				CurrencyID: "umnt",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["usdr"],
				Value:      "7903733",
				CurrencyID: "usdr",
			},
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uusd"],
				Value:      "943",
				CurrencyID: "uusd",
			},
		},
		Fees: []blockatlas.Currency{
			blockatlas.Currency{
				Decimals:   coin.Terra().Decimals,
				Symbol:     DenomMap["uluna"],
				Value:      "3015",
				CurrencyID: "uluna",
			},
		},
	},
}

type test struct {
	name     string
	platform Platform
	Data     string
	want     blockatlas.Tx
}

func TestNormalize(t *testing.T) {

	platformTerra := Platform{}

	tests := []test{
		{
			"test transfer tx",
			platformTerra,
			transferSrc,
			transferDst,
		},
		{
			"test failed tx",
			platformTerra,
			failedTransferSrc,
			failedTransferDst,
		},
		{
			"test delegate tx",
			platformTerra,
			delegateSrc,
			delegateDst,
		},
		{
			"test undelegate tx",
			platformTerra,
			unDelegateSrc,
			unDelegateDst,
		},
		{
			"test claimReward tx 1",
			platformTerra,
			claimRewardSrc1,
			claimRewardDst1,
		},
		{
			"test claimReward tx 2",
			platformTerra,
			claimRewardSrc2,
			claimRewardDst2,
		},
	}
	for _, tt := range tests {
		testNormalize(t, tt)
	}
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
