// build integration

package db_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
)

func TestDb_SetBlock(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	assert.Nil(t, database.SetLastParsedBlockNumber("ethereum", 0, context.Background()))

	block, err := database.GetLastParsedBlockNumber("ethereum", context.Background())
	assert.Nil(t, err)
	assert.Equal(t, block, int64(0))

	assert.Nil(t, database.SetLastParsedBlockNumber("ethereum", 110, context.Background()))

	newBlock, err := database.GetLastParsedBlockNumber("ethereum", context.Background())
	assert.Nil(t, err)
	assert.Equal(t, newBlock, int64(110))
}
