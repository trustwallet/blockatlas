package zilliqa

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"

	"github.com/trustwallet/blockatlas/coin"
)

const transferTransaction = `
{
    "hash": "0xd44413c79e7518152f3b05ef1edff8ef59afd06119b16d09c8bc72e94fed7843",
    "blockHeight": 104282,
    "from": "0x88af5ba10796d9091d6893eed4db23ef0bbbca37",
    "to": "0x7fccacf066a5f26ee3affc2ed1fa9810deaa632c",
    "value": "7997000000000",
    "fee": "1000000000",
    "timestamp": 1557889788637,
    "signature": "0xF0F159C5B47079E36AABC7693E61FEE9D104BDE34F4FEADA62A5066F6363E05B382E65B9381CE8138CC6824A5B62CC60EDA8B7CF13A65264F8482279DF6F768B",
    "nonce": "3",
    "receiptSuccess": true,
    "events": []
}`

var transferDst = blockatlas.Tx{
	ID:       "0xd44413c79e7518152f3b05ef1edff8ef59afd06119b16d09c8bc72e94fed7843",
	Coin:     coin.ZIL,
	From:     "0x88af5ba10796d9091d6893eed4db23ef0bbbca37",
	To:       "0x7fccacf066a5f26ee3affc2ed1fa9810deaa632c",
	Fee:      "1000000000",
	Date:     1557889788,
	Block:    104282,
	Status:   blockatlas.StatusCompleted,
	Sequence: 3,
	Memo:     "",
	Meta: blockatlas.Transfer{
		Value:    "7997000000000",
		Symbol:   "ZIL",
		Decimals: 12,
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
	token       string
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferTransaction,
		expected:    &transferDst,
		token:       "",
	})
}

func testNormalize(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}

	tx := Normalize(&srcTx)

	resJSON, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error("transfer: tx don't equal")
	}
}
