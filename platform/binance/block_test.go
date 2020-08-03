package binance

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	server := httptest.NewServer(createMockedAPI())
	defer server.Close()
	p := Init(server.URL)
	block, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.Equal(t, int64(104867535), block)
}
