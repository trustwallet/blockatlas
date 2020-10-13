// +build integration

package db_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/tests/integration/setup"
	"testing"
)

func TestDB_CreateTokenTypes(t *testing.T) {
	setup.CleanupPgContainer(database.Gorm)

	tokenTypes := []string{"ETH"}
	assert.NoError(t, database.CreateTokenTypes(context.Background(), tokenTypes))

	createdTokenTypes, err := database.GetTokenTypes(context.Background(), tokenTypes)
	assert.NoError(t, err)
	assert.Equal(t, tokenTypes[0], createdTokenTypes[0].Type)

	allTokenTypes, err := database.GetTokenTypes(context.Background(), []string{})
	assert.NoError(t, err)
	assert.Equal(t, tokenTypes[0], allTokenTypes[0].Type)

	tokenTypes = []string{"ETH", "ETH"}
	assert.NoError(t, database.CreateTokenTypes(context.Background(), tokenTypes))

	allTokenTypes, err = database.GetTokenTypes(context.Background(), []string{})
	assert.NoError(t, err)
	assert.Equal(t, len(allTokenTypes), 1)
}
