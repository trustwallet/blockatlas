package bitcoin

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

func TestTransactionStatus(t *testing.T) {
	var tests = []struct {
		Tx       Transaction
		Expected blockatlas.Status
	}{
		{Transaction{Confirmations: 0}, blockatlas.StatusPending},
		{Transaction{Confirmations: 1}, blockatlas.StatusCompleted},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, test.Tx.getStatus())
	}
}
