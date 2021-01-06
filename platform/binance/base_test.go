package binance

import (
	"fmt"
	"net/http"

	"github.com/trustwallet/golibs/mock"
)

var (
	wantedBlockNoOrders, _       = mock.JsonFromFilePathToString("mocks/block_no_orders.json")
	wantedTxs, _                 = mock.JsonFromFilePathToString("mocks/txs.json")
	wantedTokens, _              = mock.JsonFromFilePathToString("mocks/tokens.json")
	wantedBlockMultiNoOrders, _  = mock.JsonFromFilePathToString("mocks/block_multi_no_orders.json")
	wantedTxsResponse, _         = mock.JsonFromFilePathToString("mocks/txs_response.json")
	wantedAccountMetaResponse, _ = mock.JsonFromFilePathToString("mocks/account_meta_response.json")
	wantedTokensResponse, _      = mock.JsonFromFilePathToString("mocks/tokens_response.json")
	wantedTxsResponseAva, _      = mock.JsonFromFilePathToString("mocks/txs_ava_response.json")
	wantedBlockResponseMulti, _  = mock.JsonFromFilePathToString("mocks/block_multi_response.json")
	mockedBlockResponse, _       = mock.JsonFromFilePathToString("mocks/block_response.json")
	mockedNodeInfo, _            = mock.JsonFromFilePathToString("mocks/node_info.json")
)

func createMockedAPI() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/api/v1/node-info", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedNodeInfo); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/api/v2/transactions-in-block/104867508", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, mockedBlockResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/api/v2/transactions-in-block/105529271", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, wantedBlockResponseMulti); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/api/v1/tokens", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, wantedTokensResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/api/v1/account/bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, wantedAccountMetaResponse); err != nil {
			panic(err)
		}
	})

	r.HandleFunc("/api/v1/transactions", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		var (
			address = r.URL.Query().Get("address")
			txAsset = r.URL.Query().Get("txAsset")

			response string
		)

		switch {
		case address == "bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23" && txAsset == "BNB":
			response = wantedTxsResponse
		case address == "bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg" && txAsset == "AVA-645":
			response = wantedTxsResponseAva
		default:
			response = ""
		}

		if _, err := fmt.Fprint(w, response); err != nil {
			panic(err)
		}
	})

	return r
}
