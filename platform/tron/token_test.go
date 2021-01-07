package tron

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/mock"
)

var (
	tokenDst = blockatlas.Token{
		Name:     "Test",
		Symbol:   "TST",
		Decimals: 8,
		TokenID:  "1",
		Coin:     195,
		Type:     "TRC10",
	}
)

func TestNormalizeToken(t *testing.T) {
	asset := AssetInfo{Name: "Test", Symbol: "TST", ID: 1, Decimals: 8}
	actual := NormalizeToken(asset)
	assert.Equal(t, tokenDst, actual)
}

func TestPlatform_GetTokenListByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()

	p := Init(server.URL, server.URL)
	res, err := p.GetTokenListByAddress("TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R")
	assert.Nil(t, err)
	sort.Slice(res, func(i, j int) bool {
		return res[i].TokenID < res[j].TokenID
	})
	rawRes, err := json.Marshal(res)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTokensResponse, string(rawRes))
}

var (
	wantedTokensResponse, _               = mock.JsonFromFilePathToString("mocks/tokens/tokens_response.json")
	mockedAccountsTransactionsResponse, _ = mock.JsonFromFilePathToString("mocks/tokens/accounts_txs_response.json")
	mockedTrc20Response, _                = mock.JsonFromFilePathToString("mocks/tokens/trc20_response.json")
	mockedTransactionsTrc20Response, _    = mock.JsonFromFilePathToString("mocks/tokens/txs_trc20_response.json")
	mockedTransactionsEmptyResponse, _    = mock.JsonFromFilePathToString("mocks/tokens/txs_empty_response.json")
	mockedAsset1000542Response, _         = mock.JsonFromFilePathToString("mocks/tokens/asset_1000542_response.json")
	mockedAsset1000567Response, _         = mock.JsonFromFilePathToString("mocks/tokens/asset_1000567_response.json")
	mockedAssetTR7NHResponse, _           = mock.JsonFromFilePathToString("mocks/tokens/asset_tr7nh_response.json")
	mockedAccountsResponse, _             = mock.JsonFromFilePathToString("mocks/tokens/accounts_response.json")
)

func createMockedAPI() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/v1/accounts/TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedAccountsResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/v1/assets/1000542", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedAsset1000542Response); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/v1/assets/1000567", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedAsset1000567Response); err != nil {
			panic(err)
		}
	})
	r.HandleFunc("/v1/assets/TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedAssetTR7NHResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/api/account", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedTrc20Response); err != nil {
			panic(err)
		}
	})
	r.HandleFunc("/v1/accounts/TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R/transactions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedAccountsTransactionsResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/v1/accounts/TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D/transactions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedTransactionsEmptyResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/v1/accounts/TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9D/transactions/trc20", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedTransactionsTrc20Response); err != nil {
			panic(err)
		}
	})

	return r
}
