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
