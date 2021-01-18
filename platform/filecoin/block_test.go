package filecoin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	chainHead, err := mock.JsonStringFromFilePath("mocks/ChainHead.json")
	assert.Nil(t, err)

	data := make(map[string]func(http.ResponseWriter, *http.Request))
	data["/"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := fmt.Fprint(w, chainHead); err != nil {
			panic(err)
		}
	}

	server := httptest.NewServer(mock.CreateMockedAPI(data))
	defer server.Close()

	p := Init(server.URL, "")
	block, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(243590), block)
}

func TestPlatform_GetBlockByNumber(t *testing.T) {
	data := make(map[string]func(http.ResponseWriter, *http.Request))
	data["/"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		type Request map[string]interface{}
		var p Request

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			panic(err)
		}

		resp, ok := p["method"]
		if !ok {
			panic("bad json request")
		}
		var d string

		switch resp {
		case "Filecoin.ChainGetTipSetByHeight":
			chainHead, err := mock.JsonStringFromFilePath("mocks/ChainGetTipSetByHeight.json")
			if err != nil {
				panic(err)
			}
			d = chainHead
		case "Filecoin.ChainGetBlockMessages":
			blockMsg, err := mock.JsonStringFromFilePath("mocks/ChainGetBlockMessages.json")
			if err != nil {
				panic(err)
			}
			d = blockMsg
		}

		if _, err := fmt.Fprint(w, d); err != nil {
			panic(err)
		}
	}

	server := httptest.NewServer(mock.CreateMockedAPI(data))
	defer server.Close()

	p := Init(server.URL, "")
	block, err := p.GetBlockByNumber(243590)
	assert.Nil(t, err)

	blockJson, _ := json.Marshal(block)
	wanted, _ := mock.JsonStringFromFilePath("mocks/response.json")

	assert.JSONEq(t, string(blockJson), wanted)
}
