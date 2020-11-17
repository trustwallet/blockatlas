package filecoin

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	chainHead, err := mock.JsonFromFilePathToString("mocks/ChainHead.json")
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

	p := Init(server.URL)
	block, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(243590), block)
}

func TestPlatform_GetBlockByNumber(t *testing.T) {
	p := Init("https://api.filscan.io:8700/rpc/v1")
	block, err := p.GetBlockByNumber(243590)
	assert.Nil(t, err)
	assert.NotNil(t, block)
}
