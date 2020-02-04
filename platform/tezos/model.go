package tezos

import "time"

type TxType string
type TxKind string
type TxStatus string

const (
	TxTransactions TxType = "transactions"
	TxDelegations  TxType = "delegations"

	TxKindTransaction TxKind = "transaction"
	TxKindDelegation  TxKind = "delegation"

	TxStatusApplied TxStatus = "applied"
	TxStatusSkipped TxStatus = "skipped"
)

type Op struct {
	OpHash         string    `json:"opHash"`
	BlockLevel     uint64    `json:"blockLevel"`
	BlockTimestamp time.Time `json:"blockTimestamp"`
}

type Transaction struct {
	Tx         Tx         `json:"tx"`
	Op         Op         `json:"op"`
	Delegation Delegation `json:"delegation"`
}

func (t *Transaction) Status() TxStatus {
	if len(t.Tx.Status) > 0 {
		return t.Tx.Status
	}
	return t.Delegation.Status
}

func (t *Transaction) Kind() TxKind {
	if len(t.Tx.Kind) > 0 {
		return t.Tx.Kind
	}
	return t.Delegation.Kind
}

func (t *Transaction) Source() string {
	if len(t.Tx.Source) > 0 {
		return t.Tx.Source
	}
	return t.Delegation.Source
}

func (t *Transaction) Destination() string {
	if len(t.Tx.Destination) > 0 {
		return t.Tx.Destination
	}
	return t.Delegation.Delegate
}

func (t *Transaction) Fee() string {
	if len(t.Tx.Fee) > 0 {
		return t.Tx.Fee
	}
	return t.Delegation.Fee
}

type Tx struct {
	Destination string   `json:"destination"`
	Amount      string   `json:"amount"`
	GasLimit    string   `json:"gasLimit"`
	Kind        TxKind   `json:"kind"`
	BlockHash   string   `json:"blockHash"`
	Fee         string   `json:"fee"`
	Source      string   `json:"source"`
	Status      TxStatus `json:"operationResultStatus"`
}

type Delegation struct {
	Delegate string   `json:"delegate"`
	GasLimit string   `json:"gasLimit"`
	Kind     TxKind   `json:"kind"`
	Status   TxStatus `json:"operationResultStatus"`
	Fee      string   `json:"fee"`
	Source   string   `json:"source"`
}

type Validator struct {
	Address string `json:"pkh"`
}

type Account struct {
	Balance  string `json:"balance"`
	Delegate string `json:"delegate"`
}
