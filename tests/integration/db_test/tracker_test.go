// +build integration

package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
)

func TestDb_SetBlock(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assert.Nil(t, database.SetLastParsedBlockNumber("ethereum", 0))

	block, err := database.GetLastParsedBlockNumber("ethereum")
	assert.Nil(t, err)
	assert.Equal(t, block.Height, int64(0))

	assert.Nil(t, database.SetLastParsedBlockNumber("ethereum", 110))

	newBlock, err := database.GetLastParsedBlockNumber("ethereum")
	assert.Nil(t, err)
	assert.Equal(t, newBlock.Height, int64(110))
}
