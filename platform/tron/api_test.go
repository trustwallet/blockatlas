package tron

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
	"testing"
)

const transferSrc = `
{
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"type_url": "type.googleapis.com/protocol.TransferContract",
					"value": {
						"amount": 100666888000000,
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferContract"
			}
		],
		"expiration": 1551357978000,
		"fee_limit": 0,
		"ref_block_bytes": "1c17",
		"ref_block_hash": "b25737f0375c676b",
		"timestamp": 1551357920889
	},
	"ret": [
		{
			"code": "SUCESS",
			"contractRet": "SUCCESS",
			"fee": 0
		}
	],
	"signature": [
		"f34c94cfa2ac0d9dddfa9febce257684138dcbdfb31886f249f14da7eb8d134331d07e70642cc841d6191269348bc8dd63cc9be217670453d27bfb789572e7e9009000"
	],
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

var transferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1551357920,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value: "100666888000000",
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    &transferDst,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	if err != nil {
		t.Error(err)
		return
	}
	res, ok := Normalize(&srcTx)
	if !ok {
		t.Errorf("%s: tx could not be normalized", _test.name)
		return
	}

	resJSON, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&transferDst)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Errorf("%s: tx don't equal", _test.name)
	}
}
