package solana

import (
	"bytes"
	"fmt"
	"testing"

	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/mock"
)

func TestPlatform_GetTxsByAddress(t *testing.T) {
	wanted, err := mock.JsonFromFilePathToString("mocks/GetTxsByAddress.json")
	if err != nil {
		panic(err)
	}
	data := make(map[string]func(http.ResponseWriter, *http.Request))
	data["/"] = func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)

		var r blockatlas.RpcRequest
		var rs []blockatlas.RpcRequest
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
				signatures, err := mock.JsonFromFilePathToString("mocks/getConfirmedSignaturesForAddress2.json")
				if err != nil {
					panic(err)
				}
				response = signatures
			}
		} else if err := json.Unmarshal([]byte(requestBody), &rs); err == nil {
			switch rs[0].Method {
			case "getConfirmedTransaction":
				signatures, err := mock.JsonFromFilePathToString("mocks/getConfirmedTransaction.json")
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
	assert.Equal(t, wanted, string(raw))
}

func TestPlatform_EstimateTimestamp(t *testing.T) {
	assert.Equal(t, int64(1606944859), EstimateTimestamp(52838300))
	assert.Equal(t, int64(1588062639), EstimateTimestamp(5632752))
	assert.Equal(t, int64(1588026961), EstimateTimestamp(5543556))
	assert.Equal(t, int64(1586007052), EstimateTimestamp(493784))
	assert.Equal(t, int64(1585809539), EstimateTimestamp(0))
}
