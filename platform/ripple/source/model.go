package source

import (
	"encoding/json"
	"errors"
)

type Amount struct {
	Value    string `json:"string"`
	Currency string `json:"string"`
	Issuer   string `json:"string"`
}

type Response struct {
	Result       string        `json:"success"`
	Count        uint64        `json:"count"`
	Marker       string        `json:"marker"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash        string    `json:"hash"`
	Date        string    `json:"date"`
	LedgerIndex string    `json:"ledger_index"`
	LedgerHash  string    `json:"ledger_hash"`
	Tx          PaymentTx `json:"tx"`
}

type PaymentTx struct {
	TransactionType string `json:"string"`
	Flags           uint64 `json:"Flags"`
	Sequence        uint64 `json:"Sequence"`
	Amount          json.RawMessage `json:"Amount"`
	Fee             string `json:"Fee"`
	SigningPubKey   string `json:"SigningPubKey"`
	TxnSignature    string `json:"TxnSignature"`
	Account         string `json:"Account"`
	Destination     string `json:"Destination"`
}

var ErrSourceConn  = errors.New("connection to servers failed")
var ErrInvalidAddr = errors.New("invalid address")
