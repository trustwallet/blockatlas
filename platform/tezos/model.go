package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

// Tx is a Tezos blockchain transaction
type Tx struct {
	Hash        string  `json:"hash"`
	BlockHash   string  `json:"block_hash"`
	NetworkHash string  `json:"network_hash"`
	Type        Manager `json:"type"`
}

// Manager contains transfer operations
type Manager struct {
	Kind       string      `json:"kind"`
	Src        Address     `json:"src"`
	Operations []Operation `json:"operations"`
}

// Operation is a Tezos transfer operation
type Operation struct {
	Kind         string            `json:"kind"`
	Src          Address           `json:"src"`
	Dest         Address           `json:"destination"`
	Amount       blockatlas.Amount `json:"amount"`
	Failed       bool              `json:"failed"`
	Internal     bool              `json:"internal"`
	Burn         string            `json:"burn"`
	Counter      int               `json:"counter"`
	Fee          blockatlas.Amount `json:"fee"`
	GasLimit     string            `json:"gas_limit"`
	StorageLimit string            `json:"storage_limit"`
	OpLevel      uint64            `json:"op_level"`
	Timestamp    string            `json:"timestamp"`
}

type Validator struct {
	Address string `json:"pkh"`
}

// Address is a Tezos address object
type Address struct {
	Tz string `json:"tz"`
}

type Head struct {
	Level int64 `json:"level"`
}
