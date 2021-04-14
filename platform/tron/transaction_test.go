package tron

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/mock"
	"github.com/trustwallet/golibs/types"
)

var (
	transferSrc, _                 = mock.JsonStringFromFilePath("mocks/" + "transfer.json")
	wantedTransactionsWithToken, _ = mock.JsonStringFromFilePath("mocks/" + "token_txs_response.json")
	wantedTransactionsOnly, _      = mock.JsonStringFromFilePath("mocks/" + "txs_response.json")

	transferDst = types.Tx{
		ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
		Coin:   coin.TRON,
		From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
		Fee:    "0", // TODO
		Date:   1564797900,
		Block:  0, // TODO
		Status: types.StatusCompleted,
		Meta: types.Transfer{
			Value:    "100666888000000",
			Symbol:   "TRX",
			Decimals: 6,
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
		name:        "transfer",
		apiResponse: transferSrc,
		expected:    &transferDst,
	})
}

func testNormalize(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	assert.NoError(t, err)
	assert.NotNil(t, srcTx)
	res, err := normalize(srcTx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, _test.expected, res)
}

func TestPlatform_GetTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()

	p := Init(server.URL, server.URL)
	res, err := p.GetTxsByAddress("TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R")
	assert.Nil(t, err)

	rawRes, err := json.Marshal(res)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTransactionsOnly, string(rawRes))
}

func TestPlatform_GetTokenTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()

	p := Init(server.URL, server.URL)
	res, err := p.GetTokenTxsByAddress("TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	assert.Nil(t, err)

	rawRes, err := json.Marshal(res)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTransactionsWithToken, string(rawRes))
}

func Test_getTokenType(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  types.TokenType
	}{
		{"default trc20", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", types.TRC20},
		{"default trc10", "1002001", types.TRC10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getTokenType(tt.token))
		})
	}
}
