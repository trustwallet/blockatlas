package tezos

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/golibs/txtype"
)

var (
	addr1 = "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK"
	addr2 = "tz1egbN6RK2bM5vt4aAZw6r9j4nL8z49bPdS"
)

func TestTransaction_Status(t *testing.T) {
	testStatus := []struct {
		name string
		in   Transaction
		out  txtype.Status
	}{
		{"Status completed", Transaction{Stat: "applied"}, txtype.StatusCompleted},
		{"Status error", Transaction{Stat: "failed"}, txtype.StatusError},
		{"Status error", Transaction{Stat: ""}, txtype.StatusError},
		{"Status error", Transaction{Stat: "something else"}, txtype.StatusError},
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
		out     txtype.KeyTitle
	}{
		{"Delegation title", addr1, Transaction{Sender: addr1, Delegate: addr2, Receiver: "", Type: TxTypeDelegation}, txtype.AnyActionDelegation},
		{"Undelegation title", addr1, Transaction{Sender: addr1, Delegate: "", Receiver: addr2, Type: TxTypeDelegation}, txtype.AnyActionUndelegation},
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
		out  txtype.TransactionType
	}{
		{"Type should be transaction", Transaction{Type: "transaction"}, txtype.TxTransfer},
		{"Type should be delegation", Transaction{Type: "delegation"}, txtype.TxAnyAction},
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
		out     txtype.Direction
		address string
	}{
		{"Direction self", Transaction{Sender: addr1, Receiver: addr1}, txtype.DirectionSelf, addr1},
		{"Direction outgoing", Transaction{Sender: addr1, Receiver: addr2}, txtype.DirectionOutgoing, addr1},
		{"Direction incoming", Transaction{Sender: addr2, Receiver: addr1}, txtype.DirectionIncoming, addr1},
	}

	for _, tt := range testsDirection {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.Direction(tt.address))
		})
	}

	testsNormalize := []struct {
		name    string
		in      Transaction
		out     txtype.Tx
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
