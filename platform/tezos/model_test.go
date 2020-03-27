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
		Sender:    "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
		Receiver:  "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
		//Errors: {{ID: "proto.005-PsBabyM1.delegate.unchanged"}, Kind: "temporary"}
	}
	normalizedTezosTransfer = blockatlas.Tx{
		ID:        "op6GzJ3a3wGJTu4KuD2WNCVJdwEU5WKDXV6EyjsBYMEjyPQWozF",
		Coin:      1729,
		From:      "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
		To:        "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
		Fee:       "0.001500",
		Date:      1582894746,
		Block:     843988,
		Status:    "completed",
		Type:      "transfer",
		Direction: "yourself",
		Memo:      "",
		Meta: blockatlas.Transfer{
			Value:    "0.000001",
			Symbol:   "XTZ",
			Decimals: 6,
		},
	}
	//delegationTransfer = Delegation{
	//	Delegate: "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
	//	GasLimit: "10600",
	//	Kind:     TxKindDelegation,
	//	Status:   TxStatusApplied,
	//	Fee:      "1500",
	//	Source:   "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	//}
	//op = Op{
	//	OpHash:         "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
	//	BlockLevel:     788568,
	//	BlockTimestamp: time.Time{},
	//}
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
		{"Error present", Transaction{IsSuccess: false}, "transaction error"},
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
		{"Delegation", Transaction{Delegate: "", Receiver: "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK"}, blockatlas.AnyActionDelegation},
		{"Undelegation", Transaction{Delegate: "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK", Receiver: ""}, blockatlas.AnyActionUndelegation},
		{"Delegation", Transaction{Delegate: "", Receiver: ""}, blockatlas.AnyActionDelegation},
		{"Delegation", Transaction{Delegate: "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK", Receiver: "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK"}, blockatlas.AnyActionDelegation},
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
		{"Delegation", Transaction{Time: "2020-02-04T12:27:59Z",}, 1580819279},
	}

	for _, tt := range testsBlockTimestamp {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.BlockTimestamp())
		})
	}

	testsKind := []struct {
		name string
		in   Transaction
		out  TxKind
	}{
		{"Type should be transaction", Transaction{Type: "transaction",}, TxKindTransaction},
		{"Type should be delegation", Transaction{Type: "delegation",}, TxKindDelegation},
		{"Type unsupported", Transaction{Type: "bake",}, ""},
		{"Type endorsement", Transaction{Type: "endorsement",}, ""},
	}

	for _, tt := range testsKind {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.out, tt.in.Kind())
		})
	}

	testsNormalize := []struct {
		name string
		in   Transaction
		out  blockatlas.Tx
	}{
		{"Normalize XTZ transfer", tezosTransfer, normalizedTezosTransfer},
	}

	for _, tt := range testsNormalize {
		t.Run(tt.name, func(t *testing.T) {
			normalized, ok := NormalizeTx(tt.in)
			if !ok {
				assert.False(t, ok, "issue to normalize")
			}
			assert.Equal(t, tt.out, normalized)
		})
	}

}
