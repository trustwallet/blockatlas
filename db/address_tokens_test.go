package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/db/models"
)

func TestUpsertTokensMultiStmt(t *testing.T) {
	stmt, values := upsertTokensMultiStmt([]models.AddressToken{
		{Coin: 1, Address: "A", Token: "A"},
		{Coin: 1, Address: "A", Token: "B"},
		{Coin: 1, Address: "A", Token: "A"},
		{Coin: 1, Address: "B", Token: "A"},
		{Coin: 2, Address: "A", Token: "A"},
	})
	expectedStmt := `INSERT INTO address_tokens (coin, address, token) VALUES
  (?, ?, ?),
  (?, ?, ?),
  (?, ?, ?),
  (?, ?, ?),
  (?, ?, ?)
ON CONFLICT DO NOTHING;`
	expectedValues := []interface{}{
		uint(1), "A", "A",
		uint(1), "A", "B",
		uint(1), "A", "A",
		uint(1), "B", "A",
		uint(2), "A", "A",
	}
	assert.Equal(t, expectedStmt, stmt)
	assert.Equal(t, expectedValues, values)
}
