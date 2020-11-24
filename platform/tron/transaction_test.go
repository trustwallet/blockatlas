package tron

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/tokentype"
	"net/http/httptest"
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
	res, err := normalize(srcTx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	addTokenMeta(res, srcTx, assetInfo)
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
	assert.Equal(t, wantedTransactionsOnly, string(rawRes))
}

func TestPlatform_GetTokenTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()

	p := Init(server.URL, server.URL)
	res, err := p.GetTokenTxsByAddress("TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	assert.Nil(t, err)

	rawRes, err := json.Marshal(res)
	assert.Nil(t, err)
	assert.Equal(t, wantedTransactionsWithToken, string(rawRes))
}

func Test_getTokenType(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  tokentype.Type
	}{
		{"default trc20", "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", tokentype.TRC20},
		{"default trc10", "1002001", tokentype.TRC10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getTokenType(tt.token))
		})
	}
}

var (
	wantedTransactionsWithToken = `[{"id":"fb078403adfee637608c3906d9d21dd158611aba149b9993f43d0f292ce543a0","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGg7zHY9qd36aN3jLVDDRuiFeJjaaAtx8A","fee":"0","date":1592757117,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"500000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGg7zHY9qd36aN3jLVDDRuiFeJjaaAtx8A"}},{"id":"c4052b526e5cd21e1f023c31cce6b6a13eb9d8aeae3ae80fcefe6038dfbeb022","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TXJTFuXzfoPbWKCnw47AYxMzgVPUyhJGRd","fee":"0","date":1592757066,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"50000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TXJTFuXzfoPbWKCnw47AYxMzgVPUyhJGRd"}},{"id":"c4052b526e5cd21e1f023c31cce6b6a13eb9d8aeae3ae80fcefe6038dfbeb022","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TXJTFuXzfoPbWKCnw47AYxMzgVPUyhJGRd","fee":"0","date":1592757066,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"50000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TXJTFuXzfoPbWKCnw47AYxMzgVPUyhJGRd"}},{"id":"0b52a4ef9fb8c13fbfae2b8c3506333ec1d718f307062a15f170562818a01d0a","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TFNEJYAKBVgc17X6fPppZ8ayaf9yswmMYV","fee":"0","date":1592756784,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"3988000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TFNEJYAKBVgc17X6fPppZ8ayaf9yswmMYV"}},{"id":"19d2ec6174bf64beb1061475f6429cba03b64944a763686cc3551447d0e8d9d5","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGYETHZr2MTkFDe8GqwdVFPfadofTVk4am","fee":"0","date":1592756763,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"640990000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGYETHZr2MTkFDe8GqwdVFPfadofTVk4am"}},{"id":"efb7d44305759cfb189c9fd22720609a2ddeb7fbd7c8afe1dd8851342471da8d","coin":195,"from":"TAxbLztoanFhYu4TuS5RabaJYGnUkfzNKG","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","fee":"0","date":1592756631,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"1062000000","from":"TAxbLztoanFhYu4TuS5RabaJYGnUkfzNKG","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D"}},{"id":"48bd90dc3f12086178e65b9389caa8b3c74683937b86d4d61cdec77f0095994a","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGg7zHY9qd36aN3jLVDDRuiFeJjaaAtx8A","fee":"0","date":1592756610,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"2000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGg7zHY9qd36aN3jLVDDRuiFeJjaaAtx8A"}},{"id":"afd5ae7e2462c9cc899c7f730b90fd2a5e4c1315e836c92468b504ed85f0b798","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TN6Wy4j37wn3vxynynrKemWhDsUBHYje3R","fee":"0","date":1592756589,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"1000000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TN6Wy4j37wn3vxynynrKemWhDsUBHYje3R"}},{"id":"3d613031f4b2a0e19deeea030d1d18599b6d9799d2dd530005ead9712c6d219d","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TXP7prwMqugLFWZRwcJAWuKZ4UN4wz3ifq","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"21200000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TXP7prwMqugLFWZRwcJAWuKZ4UN4wz3ifq"}},{"id":"cbe359c2574efbdc8fc6a892ffc54812837295067c9816d41734126c82d0c141","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TBStJt5wDtLqeUvGEasqd55uo1CbDTCsf5","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"125000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TBStJt5wDtLqeUvGEasqd55uo1CbDTCsf5"}},{"id":"c87248b02a4caaa6f443c1b8c4d4588c8dd281a4687b73e6afecfba6741b50d8","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TJ2qhZSQ9g5YqEAJgYfPZZxn1djbf5ogkC","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"5277600000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TJ2qhZSQ9g5YqEAJgYfPZZxn1djbf5ogkC"}},{"id":"8584f1b6a70ead8232fed19bd653ba13e4c2a8befd070f4a9a06eca3a2e3e548","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TRqyhrttStrn1o7gS3mmgKrmkVw6qyw23W","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"485342000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TRqyhrttStrn1o7gS3mmgKrmkVw6qyw23W"}},{"id":"2b28b69e6747db68647acc3a62c45da5355b97acd8d2c260ee752aa9bd63a624","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TLdYhJeKCLKxVm33JL8GAyi6i6zrSz8VFr","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"1000000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TLdYhJeKCLKxVm33JL8GAyi6i6zrSz8VFr"}},{"id":"1da6576dec0bd303f56cbfb5712f782e0a56a8713cb661f8afd2f2533e5c6209","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TDpRQi5HguNpasa9Cn7AJyrn646nRDAH6x","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"2000000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TDpRQi5HguNpasa9Cn7AJyrn646nRDAH6x"}},{"id":"f9c86cce1873cb816d6cd8718e76df8839172add293bbc8d11a5f98c80f9e322","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGy3K5iDbxm8SM34UTWWniNsS13FtLnHkK","fee":"0","date":1592756583,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"24216600000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TGy3K5iDbxm8SM34UTWWniNsS13FtLnHkK"}},{"id":"75eb35734857daa79c38ef923a7e7eb2e3dfb23d2722762fd2b180651021643f","coin":195,"from":"TGMTZMty79L9psKi5b4vwXZPJaiCb9k6mV","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","fee":"0","date":1592756541,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"8241997837","from":"TGMTZMty79L9psKi5b4vwXZPJaiCb9k6mV","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D"}},{"id":"adee73dadce006ff848ff30d8c5c41f033be2e5e1a8b875f6dbbaab524177d08","coin":195,"from":"TUwgGpDrVBc3uDZg3Tj9BZZN8xkLK29yzH","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","fee":"0","date":1592756220,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"18000000000","from":"TUwgGpDrVBc3uDZg3Tj9BZZN8xkLK29yzH","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D"}},{"id":"3655a1156c9adcb876c6c9c9e0f5f1f39704ac4c7296fea05fedf5ca8f6b1a19","coin":195,"from":"TMaDtMFGJ8BBiNXchGBQmRBWi2mpfi2kdV","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","fee":"0","date":1592755962,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"863399098","from":"TMaDtMFGJ8BBiNXchGBQmRBWi2mpfi2kdV","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D"}},{"id":"03574741eb0016050a19f181e4acc4b20b70f41e11e63140c9556c31eae09fba","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TH7AaBSjS4NYuF3r8vXQcuwVXmGrv9iwYQ","fee":"0","date":1592755740,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"20000000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TH7AaBSjS4NYuF3r8vXQcuwVXmGrv9iwYQ"}},{"id":"f3aa00595996e31dbe9528a3cb21bff987f333bf1f675420ba4fa2ad43c8205f","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TA1VFEzYiU8oB9P1xdhMaFJ7BZ6FvUTyug","fee":"0","date":1592755722,"block":0,"status":"completed","sequence":0,"type":"token_transfer","memo":"","metadata":{"name":"Tether USD","symbol":"USDT","token_id":"TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t","decimals":6,"value":"21161340000","from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D","to":"TA1VFEzYiU8oB9P1xdhMaFJ7BZ6FvUTyug"}}]`
	wantedTransactionsOnly      = `[{"id":"3fca53c08ccb48bb625439a58998713d8ecc3dc1348cc3cfab912e0815b62b1a","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","to":"TVmkAmaQrY6raatYozLtcQCGqWP6VaPnHU","fee":"0","date":1592755098,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"13195916000","symbol":"TRX","decimals":6}},{"id":"b38fb6328e1fa622b7762eed856778551845c33723491e36baf357f00cc48002","coin":195,"from":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","to":"TDfVk6U7i6m82ZCRprbrfz7QE3sTEnN1Xs","fee":"0","date":1592754717,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"8737000000","symbol":"TRX","decimals":6}},{"id":"82efc8456a3c38a0919af416a53363405ced78db7c13e1b94a79ebcea98f9909","coin":195,"from":"TVqx5Dx54HgBQFfpN7KN4MWiHEnRXbch7a","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","fee":"0","date":1592754447,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"2538461","symbol":"TRX","decimals":6}},{"id":"a336bd174c127d38bf2325bc9c927059af099e8cfb91159750a1b1be16dd0bd4","coin":195,"from":"TYCwQ4bC1mHR6heAe1qgHrktFtyJ8mKkC3","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","fee":"0","date":1592754444,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"30000000","symbol":"TRX","decimals":6}},{"id":"007bbcc3855f4bf51bd76e63d7776160c115c803e227f7c44c7d1fd1bd587611","coin":195,"from":"TSUCQKEKhXREEaod5WgSKETKjYUhL2TUV7","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","fee":"0","date":1592754444,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"111337320","symbol":"TRX","decimals":6}},{"id":"9351e87b129142844f000a52911daf36fc95677dfe2846abcd28ea0d8fe2e2ea","coin":195,"from":"TFUP7BdBj61oyTHt52McZC5Q1w6CKzyNCN","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","fee":"0","date":1592754444,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"200000000","symbol":"TRX","decimals":6}},{"id":"008ebda5749c38e26a69717faa66e6f4fd8a0d358c9b947192765cc9843cff5a","coin":195,"from":"TGkLtfPuPhkG4RzewwkgNHfPxTwb5YRq6b","to":"TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R","fee":"0","date":1592754444,"block":0,"status":"completed","sequence":0,"type":"transfer","memo":"","metadata":{"value":"511000000","symbol":"TRX","decimals":6}}]`
)
