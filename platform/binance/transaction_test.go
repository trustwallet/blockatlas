package binance

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlatform_GetTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL, "<key>", "<staing_api>")
	txs, err := p.GetTxsByAddress("bnb136ns6lfw4zs5hg4n85vdthaad7hq5m4gtkgf23")
	assert.Nil(t, err)
	res, err := json.Marshal(txs)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTxs, string(res))
}

func TestPlatform_GetTokenTxsByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL, "<key>", "<staing_api>")
	txs, err := p.GetTokenTxsByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg", "AVA-645")
	assert.Nil(t, err)
	res, err := json.Marshal(txs)
	assert.Nil(t, err)
	assert.Len(t, res, 2)
}
