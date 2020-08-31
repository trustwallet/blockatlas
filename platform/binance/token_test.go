package binance

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPlatform_GetTokenListByAddress(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)

	tokens, err := p.GetTokenListByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	assert.Nil(t, err)
	res, err := json.Marshal(tokens)
	assert.Nil(t, err)
	assert.Equal(t, wantedTokens, string(res))

	tokens, err = p.GetTokenListByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	assert.Nil(t, err)
	res, err = json.Marshal(tokens)
	assert.Nil(t, err)
	assert.Equal(t, wantedTokens, string(res))

	tokens, err = p.GetTokenListByAddress("bnb1w7puzjxu05ktc5zvpnzkndt6tyl720nsutzvpg")
	assert.Nil(t, err)
	res, err = json.Marshal(tokens)
	assert.Nil(t, err)
	assert.Equal(t, wantedTokens, string(res))
}
