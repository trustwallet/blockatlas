package tezos

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

var (
	addr1 = "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK"
	addr2 = "tz1egbN6RK2bM5vt4aAZw6r9j4nL8z49bPdS"
)

func TestTransaction_Status(t *testing.T) {
	testStatus := []struct {
		name string
		in   Transaction
		out  blockatlas.Status
	}{
		{"Status completed", Transaction{Stat: "applied"}, blockatlas.StatusCompleted},
		{"Status error", Transaction{Stat: "failed"}, blockatlas.StatusError},
		{"Status error", Transaction{Stat: ""}, blockatlas.StatusError},
		{"Status error", Transaction{Stat: "something else"}, blockatlas.StatusError},
	}

	for _, tt := range testStatus {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.Status())
		})
	}

	testsError := []struct {
		name string
		in   Transaction
		out  string
	}{
		{"Error present", Transaction{IsSuccess: false, Errors: []Error{{"unchanged", "temporary"}}}, "unchanged temporary"},
		{"Error no", Transaction{IsSuccess: true}, ""},
	}

	for _, tt := range testsError {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.ErrorMsg())
		})
	}

	testsTitle := []struct {
		name string
		in   Transaction
		out  blockatlas.KeyTitle
	}{
		{"Delegation", Transaction{Delegate: "", Receiver: addr1}, blockatlas.AnyActionDelegation},
		{"Undelegation", Transaction{Delegate: addr1, Receiver: ""}, blockatlas.AnyActionUndelegation},
		{"Delegation", Transaction{Delegate: "", Receiver: ""}, blockatlas.AnyActionDelegation},
		{"Delegation", Transaction{Delegate: addr1, Receiver: addr1}, blockatlas.AnyActionDelegation},
	}

	for _, tt := range testsTitle {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.Title())
		})
	}

	testsBlockTimestamp := []struct {
		name string
		in   Transaction
		out  int64
	}{
		{"Delegation", Transaction{Time: "2020-02-04T12:27:59Z"}, 1580819279},
	}

	for _, tt := range testsBlockTimestamp {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.BlockTimestamp())
		})
	}

	testsKind := []struct {
		name string
		in   Transaction
		out  blockatlas.TransactionType
	}{
		{"Type should be transaction", Transaction{Type: "transaction"}, blockatlas.TxTransfer},
		{"Type should be delegation", Transaction{Type: "delegation"}, blockatlas.TxAnyAction},
		{"Type unsupported", Transaction{Type: "bake"}, ""},
		{"Type endorsement", Transaction{Type: "endorsement"}, ""},
	}

	for _, tt := range testsKind {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.TransferType())
		})
	}

	testsDirection := []struct {
		name    string
		in      Transaction
		out     blockatlas.Direction
		address string
	}{
		{"Direction self", Transaction{Sender: addr1, Receiver: addr1}, blockatlas.DirectionSelf, addr1},
		{"Direction outgoing", Transaction{Sender: addr1, Receiver: addr2}, blockatlas.DirectionOutgoing, addr1},
		{"Direction incoming", Transaction{Sender: addr2, Receiver: addr1}, blockatlas.DirectionIncoming, addr1},
	}

	for _, tt := range testsDirection {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.Direction(tt.address))
		})
	}

	testsNormalize := []struct {
		name    string
		in      Transaction
		out     blockatlas.Tx
		address string
	}{
		{"Normalize XTZ transfer", tezosTransfer, normalizedTezosTransfer, addr1},
	}

	for _, tt := range testsNormalize {
		t.Run(tt.name, func(t *testing.T) {
			normalized, ok := NormalizeTx(tt.in, tt.address)
			if !ok {
				assert.False(t, ok, "issue to normalize")
			}
			assert.Equal(t, tt.out, normalized)
		})
	}

}
