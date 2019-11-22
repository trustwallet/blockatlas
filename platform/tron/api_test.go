package tron

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

const transferSrc = `
{
	"block_timestamp": 1564797900000,
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"value": {
						"amount": 100666888000000,
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferContract"
			}
		]
	},
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

const tokenTransferSrc = `
{
	"block_timestamp": 1564797900000,
	"raw_data": {
		"contract": [
			{
				"parameter": {
					"value": {
						"amount": 2776267,
						"asset_name": "1002000",
						"owner_address": "4182dd6b9966724ae2fdc79b416c7588da67ff1b35",
						"to_address": "410583a68a3bcd86c25ab1bee482bac04a216b0261"
					}
				},
				"type": "TransferAssetContract"
			}
		]
	},
	"txID": "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df"
}
`

var transferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1564797900,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.Transfer{
		Value:    "100666888000000",
		Symbol:   "TRX",
		Decimals: 6,
	},
}

var tokenTransferDst = blockatlas.Tx{
	ID:     "24a10f7a503e78adc0d7e380b68005531b09e16b9e3f7b524e33f40985d287df",
	Coin:   coin.TRX,
	From:   "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
	To:     "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	Fee:    "0", // TODO
	Date:   1564797900,
	Block:  0, // TODO
	Status: blockatlas.StatusCompleted,
	Meta: blockatlas.TokenTransfer{
		Name:     "BitTorrent",
		Symbol:   "BTT",
		TokenID:  "1002000",
		Decimals: 6,
		Value:    "2776267",
		From:     "TMuA6YqfCeX8EhbfYEg5y7S4DqzSJireY9",
		To:       "TAUN6FwrnwwmaEqYcckffC7wYmbaS6cBiX",
	},
}

var assetInfo = AssetInfo{Name: "BitTorrent", Symbol: "BTT", Decimals: 6, ID: "1002000"}

type test struct {
	name        string
	apiResponse string
	expected    *blockatlas.Tx
}

func TestNormalizeTokenTransfer(t *testing.T) {
	testNormalizeTokenTransfer(t, &test{
		name:        "token transfer",
		apiResponse: tokenTransferSrc,
		expected:    &tokenTransferDst,
	})
}

func testNormalizeTokenTransfer(t *testing.T, _test *test) {
	var srcTx Tx
	err := json.Unmarshal([]byte(_test.apiResponse), &srcTx)
	assert.NoError(t, err)
	assert.NotNil(t, srcTx)
	res, err := Normalize(srcTx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	setTokenMeta(res, srcTx, assetInfo)
	assert.Equal(t, _test.expected, res)
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
	res, err := Normalize(srcTx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, _test.expected, res)
}

var tokenDst = blockatlas.Token{
	Name:     "Test",
	Symbol:   "TST",
	Decimals: 8,
	TokenID:  "1",
	Coin:     195,
	Type:     "TRC10",
}

func TestNormalizeToken(t *testing.T) {
	asset := AssetInfo{Name: "Test", Symbol: "TST", ID: "1", Decimals: 8}
	actual := NormalizeToken(asset)
	assert.Equal(t, tokenDst, actual)
}
