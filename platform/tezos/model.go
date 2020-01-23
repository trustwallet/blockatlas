package tezos

import "time"

type Op struct {
	OpHash         string    `json:"opHash"`
	BlockLevel     uint64    `json:"blockLevel"`
	BlockTimestamp time.Time `json:"blockTimestamp"`
}

type Transaction struct {
	Tx Tx `json:"tx"`
	Op Op `json:"op"`
}

type Tx struct {
	Destination string `json:"destination"`
	Amount      string `json:"amount"`
	GasLimit    string `json:"gasLimit"`
	Kind        string `json:"kind"`
	BlockHash   string `json:"blockHash"`
	Fee         string `json:"fee"`
	Source      string `json:"source"`
	Status      string `json:"operationResultStatus"`
}

type TxDelegation struct {
	Delegation Delegation `json:"delegation"`
	Op         Op         `json:"op"`
}

type Delegation struct {
	Delegate string `json:"delegate"`
	GasLimit string `json:"gasLimit"`
	Kind     string `json:"kind"`
	Status   string `json:"operationResultStatus"`
	Fee      string `json:"fee"`
	Source   string `json:"source"`
}

type Balance struct {
	Balance          string `json:"balance"`
	FrozenBalance    string `json:"frozen_balance"`
	StakingBalance   string `json:"staking_balance"`
	DelegatedBalance string `json:"delegated_balance"`
}

type Validator struct {
	Address string `json:"pkh"`
}
