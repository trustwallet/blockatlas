// +build integration

package db_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
)

func TestDb_SetBlock(t *testing.T) {
	setup.CleanupPgContainer(dbInstance.DB)

	assert.Nil(t, dbInstance.SetLastParsedBlockNumber(60, 0))

	block, err := dbInstance.GetLastParsedBlockNumber(60)
	assert.Nil(t, err)
	assert.Equal(t, block, int64(0))

	assert.Nil(t, dbInstance.SetLastParsedBlockNumber(60, 110))

	newBlock, err := dbInstance.GetLastParsedBlockNumber(60)
	assert.Nil(t, err)
	assert.Equal(t, newBlock, int64(110))
}
