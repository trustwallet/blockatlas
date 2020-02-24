package tezos

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	transaction = Tx{
		Destination: "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
		Amount:      "1751",
		GasLimit:    "15385",
		Kind:        TxKindTransaction,
		BlockHash:   "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
		Fee:         "0",
		Source:      "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
		Status:      TxStatusApplied,
	}
	del = Delegation{
		Delegate: "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93",
		GasLimit: "10600",
		Kind:     TxKindDelegation,
		Status:   TxStatusApplied,
		Fee:      "1500",
		Source:   "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK",
	}
	op = Op{
		OpHash:         "BMDYrXJ7GSwztzsy3ykJb43VXciNk1WY8EAaSoGbcUE7mA7HUzj",
		BlockLevel:     788568,
		BlockTimestamp: time.Time{},
	}
)

func TestTransaction_Status(t *testing.T) {
	type fields struct {
		Op         Op
		Tx         Tx
		Delegation Delegation
	}
	type want struct {
		Status      TxStatus
		Kind        TxKind
		Source      string
		Destination string
		Fee         string
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			"transaction",
			fields{op, transaction, Delegation{}},
			want{TxStatusApplied, TxKindTransaction, "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93", "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK", "0"},
		},
		{
			"delegation",
			fields{op, Tx{}, del},
			want{TxStatusApplied, TxKindDelegation, "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK", "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93", "1500"},
		},
		{
			"transaction and delegation",
			fields{op, transaction, del},
			want{TxStatusApplied, TxKindTransaction, "tz2FCNBrERXtaTtNX6iimR1UJ5JSDxvdHM93", "tz1WDujRWCYjLBDfZieXW6insg5EUbg1rCRK", "0"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &Transaction{
				Tx:         tt.fields.Tx,
				Op:         tt.fields.Op,
				Delegation: tt.fields.Delegation,
			}
			assert.Equal(t, tt.want.Status, tx.Status())
			assert.Equal(t, tt.want.Kind, tx.Kind())
			assert.Equal(t, tt.want.Source, tx.Source())
			assert.Equal(t, tt.want.Destination, tx.Destination())
			assert.Equal(t, tt.want.Fee, tx.Fee())
		})
	}
}
