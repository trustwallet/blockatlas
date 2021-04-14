package binance

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL, "<key>", "<staing_api>")
	number, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(104867535), number)
}

func TestPlatform_GetBlockByNumber(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL, "<key>", "<staing_api>")
	block, err := p.GetBlockByNumber(104867508)
	assert.Nil(t, err)
	res, err := json.Marshal(block)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedBlockNoOrders, string(res))

	blockMulti, err := p.GetBlockByNumber(105529271)
	assert.Nil(t, err)
	resMulti, err := json.Marshal(blockMulti)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedBlockMultiNoOrders, string(resMulti))
}
