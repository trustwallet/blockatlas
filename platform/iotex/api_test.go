package iotex

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/trustwallet/blockatlas/coin"
)

const (
transfer = `
{
  "actionInfo":
  [
    {
      "action":
      {
        "core":
        {
          "version":1,
          "nonce":"3",
          "gasLimit":"10000",
          "gasPrice":"1000000000000",
          "transfer":
          {
            "amount":"21000000000000000000",
            "recipient":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
          }
        },
        "senderPubKey":"BKCXbZcntIxrdPFZdWratLOfKU2yUUc0LuF/ilB3JoQzd/mvXaUbPuBpIE/sUtxGo0YxcAcN0cylCo1EIPQwJqc=",
        "signature":"V4JBmqjFU+UmdVKQZ1+2CVElZ8sUMz1m0wfJEE5J7hFAG72nD2oI0wrLnTGBM0CaD1BjNGJvELYKi/g5IvO3AgE="
      },
      "actHash":"109b75cb688a5347268cbf11b20fa90fd0a14e92a42ba735c046bbf1a6e66ad7",
      "blkHash":"42ace162549ec8d44641d7da7184d1e12ebd4111b0d2888a2d97d88a7c4fa04b",
      "blkHeight":"96202",
      "sender":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
      "gasFee":"10000000000000000",
      "timestamp":"2019-05-03T06:09:00Z"
    },
    {
      "action":
      {
        "core":
        {
          "version":1,
          "nonce":"3",
          "gasLimit":"10000",
          "gasPrice":"1000000000000",
          "transfer":
          {
            "amount":"21000000000000000000",
            "recipient":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
          }
        },
        "senderPubKey":"BKCXbZcntIxrdPFZdWratLOfKU2yUUc0LuF/ilB3JoQzd/mvXaUbPuBpIE/sUtxGo0YxcAcN0cylCo1EIPQwJqc=",
        "signature":"V4JBmqjFU+UmdVKQZ1+2CVElZ8sUMz1m0wfJEE5J7hFAG72nD2oI0wrLnTGBM0CaD1BjNGJvELYKi/g5IvO3AgE="
      },
      "actHash":"109b75cb688a5347268cbf11b20fa90fd0a14e92a42ba735c046bbf1a6e66ad7",
      "blkHash":"42ace162549ec8d44641d7da7184d1e12ebd4111b0d2888a2d97d88a7c4fa04b",
      "blkHeight":"0",
      "sender":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
      "gasFee":"10000000000000000",
      "timestamp":"2019-05-03T06:09:00Z"
    },
    {
      "action":
      {
        "core":
        {
          "version":1,
          "nonce":"3.1",
          "gasLimit":"10000",
          "gasPrice":"1000000000000",
          "transfer":
          {
            "amount":"21000000000000000000",
            "recipient":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m"
          }
        },
        "senderPubKey":"BKCXbZcntIxrdPFZdWratLOfKU2yUUc0LuF/ilB3JoQzd/mvXaUbPuBpIE/sUtxGo0YxcAcN0cylCo1EIPQwJqc=",
        "signature":"V4JBmqjFU+UmdVKQZ1+2CVElZ8sUMz1m0wfJEE5J7hFAG72nD2oI0wrLnTGBM0CaD1BjNGJvELYKi/g5IvO3AgE="
      },
      "actHash":"109b75cb688a5347268cbf11b20fa90fd0a14e92a42ba735c046bbf1a6e66ad7",
      "blkHash":"42ace162549ec8d44641d7da7184d1e12ebd4111b0d2888a2d97d88a7c4fa04b",
      "blkHeight":"96202",
      "sender":"io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
      "gasFee":"10000000000000000",
      "timestamp":"2019-05-03T06:09:00Z"
    }
  ]
}
`
)

var expected = []blockatlas.Tx {
	{
		ID       : "109b75cb688a5347268cbf11b20fa90fd0a14e92a42ba735c046bbf1a6e66ad7",
		Coin     : coin.IOTX,
		From     : "io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
		To       : "io1mwekae7qqwlr23220k5n9z3fmjxz72tuchra3m",
		Fee      : blockatlas.Amount("10000000000000000"),
		Date     : int64(1556863740),
		Block    : 96202,
		Status   : blockatlas.StatusCompleted,
		Sequence : uint64(3),
		Type     : blockatlas.TxTransfer,
		Meta     : blockatlas.Transfer{
			Value : blockatlas.Amount("21000000000000000000"),
		},
	},
	{
		Coin   : coin.IOTX,
		Status : blockatlas.StatusFailed,
		Error  : "invalid block height",
	},
	{
		Coin   : coin.IOTX,
		Status : blockatlas.StatusFailed,
		Error  : "strconv.ParseInt: parsing \"3.1\": invalid syntax",
	},
}

func TestClient(t *testing.T) {
	assert := assert.New(t)

	c := Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    "https://pharos.iotex.io/v1",
	}
	uri := fmt.Sprintf("%s/actions/hash/%s",
		c.BaseURL,
		"109b75cb688a5347268cbf11b20fa90fd0a14e92a42ba735c046bbf1a6e66ad7")

	res, err := c.HTTPClient.Get(uri)
	assert.NoError(err)
	defer res.Body.Close()
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestNormalize(t *testing.T) {
	assert := assert.New(t)

	var act Response
	assert.NoError(json.Unmarshal([]byte(transfer), &act))
	assert.Equal(3, len(act.ActionInfo))

	for i, v := range act.ActionInfo {
		tx, _ := Normalize(v)
		assert.Equal(expected[i], tx)
	}
}
