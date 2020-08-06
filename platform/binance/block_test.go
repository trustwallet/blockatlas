package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)
	number, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(104867535), number)
}

func TestPlatform_GetBlockByNumber(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)
	block, err := p.GetBlockByNumber(104867508)
	assert.Nil(t, err)
	res, err := json.Marshal(block)
	assert.Nil(t, err)
	assert.Equal(t, wantedBlock, string(res))
}
