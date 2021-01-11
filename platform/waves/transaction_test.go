package waves

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	transferV1, _   = mock.JsonStringFromFilePath("mocks/" + "transfer.json")
	differentTxs, _ = mock.JsonStringFromFilePath("mocks/" + "different_txs.json")

	transferV1Obj = types.Tx{
		ID:     "7QoQc9qMUBCfY4QV35mgBsT8eTXybvGkM2HTumtAvBUL",
		Coin:   5741564,
		From:   "3PLrCnhKyX5iFbGDxbqqMvea5VAqxMcinPW",
		To:     "3PKWyVAmHom1sevggiXVfbGUc3kS85qT4Va",
		Fee:    "100000",
		Date:   1561048131,
		Block:  1580410,
		Status: types.StatusCompleted,
		Memo:   "",
		Meta: types.Transfer{
			Value:    types.Amount("9481600000"),
			Symbol:   "WAVES",
			Decimals: 8,
		},
	}

	differentTxsObj = types.Tx{
		ID:     "52GG9U2e6foYRKp5vAzsTQ86aDAABfRJ7synz7ohBp19",
		Coin:   5741564,
		From:   "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
		To:     "3NBVqYXrapgJP9atQccdBPAgJPwHDKkh6A8",
		Fee:    "100000",
		Date:   1479313236,
		Block:  7782,
		Memo:   "string",
		Status: types.StatusCompleted,
		Meta: types.Transfer{
			Value:    types.Amount("100000"),
			Symbol:   "WAVES",
			Decimals: 8,
		},
	}
)

type txParseTest struct {
	name        string
	apiResponse string
	expected    *types.Tx
}

type txFilterTest struct {
	name        string
	apiResponse string
	expected    types.Tx
}

func TestNormalize(t *testing.T) {
	testParseTx(t, &txParseTest{
		name:        "transfer",
		apiResponse: transferV1,
		expected:    &transferV1Obj,
	})
	testFilterTxs(t, &txFilterTest{
		name:        "filter transfer transactions txParseTest",
		apiResponse: differentTxs,
		expected:    differentTxsObj,
	})
}

func testParseTx(t *testing.T, _test *txParseTest) {
	var tx Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &tx)
	if err != nil {
		t.Error(err)
		return
	}

	res, _ := NormalizeTx(&tx)

	resJSON, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(resJSON), string(dstJSON))
}

func testFilterTxs(t *testing.T, _test *txFilterTest) {
	var txs [][]Transaction
	err := json.Unmarshal([]byte(_test.apiResponse), &txs)
	if err != nil {
		t.Error(err)
		return
	}
	var res types.Tx
	for _, tx := range txs[0] {
		if tx.Type == 4 {
			n, ok := NormalizeTx(&tx)
			if ok {
				res = n
			}
		}
	}

	resJSON, err := json.Marshal(&res)
	if err != nil {
		t.Fatal(err)
	}

	dstJSON, err := json.Marshal(&_test.expected)
	if err != nil {
		t.Fatal(err)
	}

	assert.JSONEq(t, string(resJSON), string(dstJSON))
}
