package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockMap(t *testing.T) {
	s := initStorage(t)
	assert.NotNil(t, s)

	block, err := s.GetLastParsedBlockNumber(60)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), block)

	newBlock := int64(1400)

	err = s.SetLastParsedBlockNumber(60, newBlock)
	assert.Nil(t, err)

	current, err := s.GetLastParsedBlockNumber(60)
	assert.Nil(t, err)
	assert.Equal(t, newBlock, current)
}
