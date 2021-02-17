package solana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/client"
	"github.com/trustwallet/golibs/mock"
)

func TestPlatform_GetTxsByAddress(t *testing.T) {
	wanted, err := mock.JsonStringFromFilePath("mocks/GetTxsByAddress.json")
	if err != nil {
		panic(err)
	}
	data := make(map[string]func(http.ResponseWriter, *http.Request))
	data["/"] = func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)

		var r client.RpcRequest
		var rs []client.RpcRequest
		var response string

		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(req.Body)
		if err != nil {
			panic(err)
		}
		requestBody := buf.String()

		if err := json.Unmarshal([]byte(requestBody), &r); err == nil {
			switch r.Method {
			case "getConfirmedSignaturesForAddress2":
				signatures, err := mock.JsonStringFromFilePath("mocks/getConfirmedSignaturesForAddress2.json")
				if err != nil {
					panic(err)
				}
				response = signatures
			}
		} else if err := json.Unmarshal([]byte(requestBody), &rs); err == nil {
			switch rs[0].Method {
			case "getConfirmedTransaction":
				signatures, err := mock.JsonStringFromFilePath("mocks/getConfirmedTransaction.json")
				if err != nil {
					panic(err)
				}
				response = signatures
			}
		} else {
			panic("not valid json rpc request")
		}

		if _, err := fmt.Fprint(w, response); err != nil {
			panic(err)
		}
	}

	server := httptest.NewServer(mock.CreateMockedAPI(data))
	defer server.Close()

	p := Init(server.URL)
	txs, err := p.GetTxsByAddress("AHy6YZA8BsHgQfVkk7MbwpAN94iyN7Nf1zN4nPqUN32Q")
	assert.Nil(t, err)
	raw, err := json.Marshal(txs)
	assert.Nil(t, err)
	assert.JSONEq(t, wanted, string(raw))
}
