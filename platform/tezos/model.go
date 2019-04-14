package tezos

import (
	"errors"
	"github.com/trustwallet/blockatlas/models"
)

type Tx struct {
	Hash        string  `json:"hash"`
	BlockHash   string  `json:"block_hash"`
	NetworkHash string  `json:"network_hash"`
	Type        Manager `json:"type"`
}

type Manager struct {
	Kind       string      `json:"kind"`
	Src        Address     `json:"src"`
	Operations []Operation `json:"operations"`
}

type Operation struct {
	Kind         string        `json:"kind"`
	Src          Address       `json:"src"`
	Dest         Address       `json:"destination"`
	Amount       models.Amount `json:"amount"`
	Failed       bool          `json:"failed"`
	Internal     bool          `json:"internal"`
	Burn         int           `json:"burn"`
	Counter      int           `json:"counter"`
	Fee          models.Amount `json:"fee"`
	GasLimit     string        `json:"gas_limit"`
	StorageLimit string        `json:"storage_limit"`
	OpLevel      uint64        `json:"op_level"`
	Timestamp    string        `json:"timestamp"`
}

type Address struct {
	Tz string `json:"tz"`
}

var ErrSourceConn = errors.New("connection to servers failed")
