package binance

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlatform_GetTokenListByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL, "<key>", "<staing_api>")

	tokens, err := p.GetTokenListByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	assert.Nil(t, err)
	res, err := json.Marshal(tokens)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTokens, string(res))

	tokens, err = p.GetTokenListByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	assert.Nil(t, err)
	res, err = json.Marshal(tokens)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTokens, string(res))

	tokens, err = p.GetTokenListByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	assert.Nil(t, err)
	res, err = json.Marshal(tokens)
	assert.Nil(t, err)
	assert.JSONEq(t, wantedTokens, string(res))
}
