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
		name    string
		address string
		in      Transaction
		out     blockatlas.KeyTitle
	}{
		{"Delegation title", addr1, Transaction{Sender: addr1, Delegate: addr2, Receiver: "", Type: TxTypeDelegation}, blockatlas.AnyActionDelegation},
		{"Undelegation title", addr1, Transaction{Sender: addr1, Delegate: "", Receiver: addr2, Type: TxTypeDelegation}, blockatlas.AnyActionUndelegation},
		{"Unsupported title", addr1, Transaction{Sender: addr1, Delegate: addr1, Receiver: addr1}, "unsupported title"},
		{"Unsupported title", addr1, Transaction{Sender: addr1, Delegate: addr2, Receiver: addr1}, "unsupported title"},
		{"Unsupported title", addr1, Transaction{Sender: addr1, Delegate: addr1, Receiver: addr2}, "unsupported title"},
		{"Unsupported title", addr1, Transaction{Sender: addr1, Delegate: addr2, Receiver: addr2}, "unsupported title"},
	}

	for _, tt := range testsTitle {
		t.Run(tt.name, func(t *testing.T) {
			title, _ := tt.in.Title(tt.address)
			assert.Equal(t, tt.out, title)
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

	testsTransferType := []struct {
		name string
		in   Transaction
		out  blockatlas.TransactionType
	}{
		{"Type should be transaction", Transaction{Type: "transaction",}, blockatlas.TxTransfer},
		{"Type should be delegation", Transaction{Type: "delegation",}, blockatlas.TxAnyAction},
		{"Type unsupported", Transaction{Type: "bake"}, "unsupported type"},
	}

	for _, tt := range testsTransferType {
		t.Run(tt.name, func(t *testing.T) {
			transferType, _ := tt.in.TransferType()
			assert.Equal(t, tt.out, transferType)
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

	testsGetReceiver := []struct {
		name string
		in   Transaction
		out  string
	}{
		{"Should get receiver when no delegate", Transaction{Receiver: addr1, Delegate: ""}, addr1},
		{"Should get receiver when delegate", Transaction{Receiver: "", Delegate: addr1}, addr1},
	}

	for _, tt := range testsGetReceiver {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.GetReceiver())
		})
	}

}
