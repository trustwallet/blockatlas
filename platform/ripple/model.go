package ripple

import (
	"encoding/json"

	"github.com/trustwallet/blockatlas/models"
)

type Amount struct {
	Value    models.Amount `json:"string"`
	Currency string        `json:"string"`
	Issuer   string        `json:"string"`
}

type Response struct {
	Result       string `json:"result"`
	Count        uint64 `json:"count"`
	Marker       string `json:"marker"`
	Transactions []Tx   `json:"transactions"`
}

type Tx struct {
	Hash        string  `json:"hash"`
	Date        string  `json:"date"`
	LedgerIndex uint64  `json:"ledger_index"`
	LedgerHash  string  `json:"ledger_hash"`
	Payment     Payment `json:"tx"`
}

type Payment struct {
	TransactionType string          `json:"string"`
	Flags           uint64          `json:"Flags"`
	Sequence        uint64          `json:"Sequence"`
	Amount          json.RawMessage `json:"Amount"`
	Fee             models.Amount   `json:"Fee"`
	SigningPubKey   string          `json:"SigningPubKey"`
	TxnSignature    string          `json:"TxnSignature"`
	Account         string          `json:"Account"`
	Destination     string          `json:"Destination"`
}
