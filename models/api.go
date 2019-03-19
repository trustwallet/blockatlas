package models

import "encoding/json"

const (
	TxBasic  = "basic"
	TxSwap   = "swap"
	TxLegacy = "legacy"
)

type Response struct {
	Total int        `json:"total"`
	Docs  []LegacyTx `json:"docs"`
}

type Balance struct {
	Amount uint64 `json:"amount"`
	Unit   string `json:"unit"`
}

type AccountInfo struct {
	Balances []Balance `json:"balances"`
	Txs      []Tx      `json:"txs"`
}

type Tx interface {
	Type() string
}

type BasicTx struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Fee   string `json:"fee"`
	Value string `json:"value"`
}

func (_ *BasicTx) Type() string {
	return TxBasic
}

type LegacyTx struct {
	Id          string `json:"id"`

	// Empty array
	Operations []json.RawMessage `json:"operations"`
	// Null pointer
	Contract   *json.RawMessage  `json:"contract"`

	BlockNumber uint64 `json:"blockNumber"`
	Timestamp   string `json:"timeStamp"`
	Nonce       int    `json:"nonce"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	GasUsed     string `json:"gasUsed"`
	Input       string `json:"input"`
	Coin        uint   `json:"coin"`
	Error       string `json:"error"`
}

func (_ *LegacyTx) Type() string {
	return TxLegacy
}
