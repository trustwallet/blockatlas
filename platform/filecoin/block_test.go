package filecoin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlatform_CurrentBlockNumber(t *testing.T) {
	p := Init("https://api.filscan.io:8700/rpc/v1")
	block, err := p.CurrentBlockNumber()
	assert.Nil(t, err)
	assert.NotNil(t, block)
}
