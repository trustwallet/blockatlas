package tron

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
)

func TestPlatform_GetTokenListByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()

	p := Init(server.URL, server.URL)
	res, err := p.GetTokenListIdsByAddress("TM1zzNDZD2DPASbKcgdVoTYhfmYgtfwx9R")
	assert.Nil(t, err)
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	rawRes, err := json.Marshal(res)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTokensResponse, string(rawRes))
}

var (
	wantedTokensResponse, _               = mock.JsonStringFromFilePath("mocks/tokens/tokens_response.json")
	mockedAccountsTransactionsResponse, _ = mock.JsonStringFromFilePath("mocks/tokens/accounts_txs_response.json")
	mockedTrc20Response, _                = mock.JsonStringFromFilePath("mocks/tokens/trc20_response.json")
	mockedTransactionsTrc20Response, _    = mock.JsonStringFromFilePath("mocks/tokens/txs_trc20_response.json")
	mockedTransactionsEmptyResponse, _    = mock.JsonStringFromFilePath("mocks/tokens/txs_empty_response.json")
	mockedAsset1000542Response, _         = mock.JsonStringFromFilePath("mocks/tokens/asset_1000542_response.json")
	mockedAsset1000567Response, _         = mock.JsonStringFromFilePath("mocks/tokens/asset_1000567_response.json")
	mockedAssetTR7NHResponse, _           = mock.JsonStringFromFilePath("mocks/tokens/asset_tr7nh_response.json")
	mockedAccountsResponse, _             = mock.JsonStringFromFilePath("mocks/tokens/accounts_response.json")
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
