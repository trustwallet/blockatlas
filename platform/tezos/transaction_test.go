package tezos

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var (
	tezosTransfer = Transaction{
		Hash:      "op6GzJ3a3wGJTu4KuD2WNCVJdwEU5WKDXV6EyjsBYMEjyPQWozF",
		Type:      "transaction",
		Time:      "2020-02-28T12:59:06Z",
		Height:    843988,
		Stat:      "applied",
		IsSuccess: true,
		Volume:    0.000001,
		Fee:       0.0015,
		Sender:    addr1,
		Receiver:  addr1,
		//Errors: {{ID: "proto.005-PsBabyM1.delegate.unchanged"}, Kind: "temporary"}
	}

	normalizedTezosTransfer = blockatlas.Tx{
		ID:        "op6GzJ3a3wGJTu4KuD2WNCVJdwEU5WKDXV6EyjsBYMEjyPQWozF",
		Coin:      1729,
		From:      addr1,
		To:        addr1,
		Fee:       "1500",
		Date:      1582894746,
		Block:     843988,
		Status:    "completed",
		Type:      "transfer",
		Direction: "yourself",
		Memo:      "",
		Meta: blockatlas.Transfer{
			Value:    "1",
			Symbol:   "XTZ",
			Decimals: 6,
		},
	}
)

func TestNormalizeTx(t *testing.T) {
	tests := []struct {
		name    string
		in      Transaction
		out     blockatlas.Tx
		address string
	}{
		{"Normalize XTZ transfer", tezosTransfer, normalizedTezosTransfer, addr1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized, ok := NormalizeTx(tt.in, tt.address)
			if !ok {
				assert.False(t, ok, "issue to normalize")
			}
			assert.Equal(t, tt.out, normalized)
		})
	}
}
