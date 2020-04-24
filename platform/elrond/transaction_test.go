package elrond

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

const rewardTxSrc = `
{
	"hash":"30d404cc7a42b0158b95f6adfbf9a517627d60f6c7e497c1442dfdb6460285df",
	"miniBlockHash":"13f575aa43185751b699dfc2696ca0969e4e996ece112f5ef01d20e506d47df0",
	"blockHash":"f54c2e5894dc9af137cbc2f5f22ecfd0821923013d432acee97f621e8a8f266d",
	"nonce":0,
	"round":35462,
	"value":"82516976060558456822",
	"receiver":"erd10yagg2vme2jns9zqf9xn8kl86fkc6dr063vnuj0mz2kk2jw0qwuqmfmaw0",
	"sender":"4294967295",
	"receiverShard":0,
	"senderShard":4294967295,
	"gasPrice":0,
	"gasLimit":0,
	"data":"",
	"signature":"",
	"timestamp":1587715632,
	"status":"Success"
}`

var rewardTxDst = blockatlas.Tx{
	ID:     "30d404cc7a42b0158b95f6adfbf9a517627d60f6c7e497c1442dfdb6460285df",
	Coin:   coin.ERD,
	Date:   int64(1587715632),
	From:   "metachain",
	To:     "erd10yagg2vme2jns9zqf9xn8kl86fkc6dr063vnuj0mz2kk2jw0qwuqmfmaw0",
	Fee:    "0",
	Status: blockatlas.StatusCompleted,
	Memo:   "reward transaction",
	Meta: blockatlas.Transfer{
		Value:    "82516976060558456822",
		Symbol:   coin.Elrond().Symbol,
		Decimals: coin.Elrond().Decimals,
	},
	Direction: blockatlas.DirectionOutgoing,
}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
	token       bool
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: rewardTxSrc,
		expected:    &rewardTxDst,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var tx Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &tx)
	if err != nil {
		t.Error(err)
		return
	}

	normalizedTx, ok := NormalizeTx(tx, tx.Sender)
	if !ok {
		t.Error(_test.name + ": cannot normalize tx")
	}

	resJSON, err := json.Marshal(&normalizedTx)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJSON, dstJSON) {
		println(string(resJSON))
		println(string(dstJSON))
		t.Error(_test.name + ": tx don't equal")
	}
}
