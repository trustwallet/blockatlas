package stellar

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	createSrc, _   = mock.JsonStringFromFilePath("mocks/" + "create_tx.json")
	transferSrc, _ = mock.JsonStringFromFilePath("mocks/" + "transfer_tx.json")

	createDst = types.Tx{
		ID:    "8b96cf3a660b85ef80b5a84c032cacdb93bb139cfe7e929b974ea9eaa0d29141",
		Coin:  coin.STELLAR,
		From:  "GBEZOC5U4TVH7ZY5N3FLYHTCZSI6VFGTULG7PBITLF5ZEBPJXFT46YZM",
		To:    "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
		Fee:   "100",
		Date:  1470850220,
		Block: 0,
		Meta: types.Transfer{
			Value:    "473269393700000000",
			Symbol:   "XLM",
			Decimals: 7,
		},
	}

	transferDst = types.Tx{
		ID:    "a596dc910bae20b5bbe64aa7aa3f42acbd55769b98307878f5ad095e994bc9cf",
		Coin:  coin.STELLAR,
		From:  "GDKIJJIKXLOM2NRMPNQZUUYK24ZPVFC6426GZAEP3KUK6KEJLACCWNMX",
		To:    "GAX3BRBNB5WTJ2GNEFFH7A4CZKT2FORYABDDBZR5FIIT3P7FLS2EFOZZ",
		Fee:   "100",
		Date:  1470857941,
		Memo:  "testing",
		Block: 123,
		Meta: types.Transfer{
			Value:    types.Amount("1000000000000000"),
			Symbol:   "XLM",
			Decimals: 7,
		},
	}
)

type test struct {
	name        string
	apiResponse string
	expected    *types.Tx
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
	var payment Payment
	err := json.Unmarshal([]byte(_test.apiResponse), &payment)
	if err != nil {
		t.Error(err)
		return
	}
	tx, ok := Normalize(&payment, coin.STELLAR)
	if !ok {
		t.Errorf("%s: tx could not be normalized", _test.name)
	}

	_, err = json.Marshal(&tx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, tx, *_test.expected)
}
