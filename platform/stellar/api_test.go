package stellar

import (
	"bytes"
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/models"
	"github.com/trustwallet/blockatlas/platform/stellar/source"
	"testing"
)

const createSrc = `
{
	"id": "25002129911451649",
	"paging_token": "25002129911451649",
	"transaction_successful": true,
	"source_account": "GBEZOC5U4TVH7ZY5N3FLYHTCZSI6VFGTULG7PBITLF5ZEBPJXFT46YZM",
	"type": "create_account",
	"type_i": 0,
	"created_at": "2016-08-10T17:30:20Z",
	"transaction_hash": "8b96cf3a660b85ef80b5a84c032cacdb93bb139cfe7e929b974ea9eaa0d29141",
	"starting_balance": "47326939370.0000000",
	"funder": "GBEZOC5U4TVH7ZY5N3FLYHTCZSI6VFGTULG7PBITLF5ZEBPJXFT46YZM",
	"account": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX"
}
`

var createDst = models.Tx{
	Id:    "8b96cf3a660b85ef80b5a84c032cacdb93bb139cfe7e929b974ea9eaa0d29141",
	Coin:  coin.XLM,
	From:  "GBEZOC5U4TVH7ZY5N3FLYHTCZSI6VFGTULG7PBITLF5ZEBPJXFT46YZM",
	To:    "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
	Fee:   "100",
	Date:  1470850220,
	Block: 25002129911451649,
	Meta: models.Transfer{
		Value: "473269393700000000",
	},
}

const transferSrc = `
{
	"id": "25008572362395649",
	"paging_token": "25008572362395649",
	"transaction_successful": true,
	"source_account": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
	"type": "payment",
	"type_i": 1,
	"created_at": "2016-08-10T19:39:01Z",
	"transaction_hash": "a596dc910bae20b5bbe64aa7aa3f42acbd55769b98307878f5ad095e994bc9cf",
	"asset_type": "native",
	"from": "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
	"to": "GAX3BRBNB5WTJ2GNEFFH7A4CZKT2FORYABDDBZR5FIIT3P7FLS2EFOZZ",
	"amount": "100000000.0000000"
}
`

var transferDst = models.Tx{
	Id:    "a596dc910bae20b5bbe64aa7aa3f42acbd55769b98307878f5ad095e994bc9cf",
	Coin:  coin.XLM,
	From:  "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
	To:    "GAX3BRBNB5WTJ2GNEFFH7A4CZKT2FORYABDDBZR5FIIT3P7FLS2EFOZZ",
	Fee:   "100",
	Date:  1470857941,
	Block: 25008572362395649,
	Meta: models.Transfer{
		Value: models.Amount("1000000000000000"),
	},
}

type test struct {
	name        string
	apiResponse string
	expected    *models.Tx
}

func TestNormalize(t *testing.T) {
	testNormalize(t, &test{
		name:        "create account",
		apiResponse: createSrc,
		expected:    &createDst,
	})
	testNormalize(t, &test{
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    &transferDst,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var payment source.Payment
	err := json.Unmarshal([]byte(_test.apiResponse), &payment)
	if err != nil {
		t.Error(err)
		return
	}
	tx, ok := Normalize(&payment, coin.XLM)
	if !ok {
		t.Errorf("%s: tx could not be normalized", _test.name)
	}

	resJson, err := json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	dstJson, err := json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(resJson, dstJson) {
		println(string(resJson))
		println(string(dstJson))
		t.Error(_test.name + ": tx don't equal")
	}
}
