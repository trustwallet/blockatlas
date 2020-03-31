// +build integration

package db_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
)

func TestDb_SetBlock(t *testing.T) {
	setup.CleanupPgContainer(dbConn)

	assert.Nil(t, db.SetLastParsedBlockNumber(dbConn, 60, 0))

	block, err := db.GetLastParsedBlockNumber(dbConn, 60)
	assert.Nil(t, err)
	assert.Equal(t, block, int64(0))

	assert.Nil(t, db.SetLastParsedBlockNumber(dbConn, 60, 110))

	newBlock, err := db.GetLastParsedBlockNumber(dbConn, 60)
	assert.Nil(t, err)
	assert.Equal(t, newBlock, int64(110))
}
