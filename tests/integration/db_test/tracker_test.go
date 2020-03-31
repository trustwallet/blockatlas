// +build integration

package db_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
)

func TestDb_SetBlock(t *testing.T) {
	setup.CleanupPgContainer()

	assert.Nil(t, db.SetLastParsedBlockNumber(60, 0))

	block, err := db.GetLastParsedBlockNumber(60)
	assert.Nil(t, err)
	assert.Equal(t, block, int64(0))

	assert.Nil(t, db.SetLastParsedBlockNumber(60, 110))

	newBlock, err := db.GetLastParsedBlockNumber(60)
	assert.Nil(t, err)
	assert.Equal(t, newBlock, int64(110))
}
