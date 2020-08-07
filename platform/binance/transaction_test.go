package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPlatform_GetTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)
	txs, err := p.GetTxsByAddress("bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23")
	res, err := json.Marshal(txs)
	assert.Nil(t, err)
	assert.Equal(t, wantedTxs, string(res))
}

func TestPlatform_GetTokenTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)
	txs, err := p.GetTokenTxsByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg", "AVA-645")
	res, err := json.Marshal(txs)
	assert.Nil(t, err)
	assert.Equal(t, wantedTxsAva, string(res))
}
