package ripple

import "github.com/trustwallet/golibs/types"

type TransactionType string
type TransactionResult string

const (
	transactionResultSuccess TransactionResult = "tesSUCCESS"

	transactionPayment TransactionType = "Payment"
)

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
	Meta        Meta    `json:"meta"`
}

type Payment struct {
	TransactionType TransactionType `json:"TransactionType"`
	Flags           uint64          `json:"Flags"`
	Sequence        uint64          `json:"Sequence"`
	Fee             types.Amount    `json:"Fee"`
	SigningPubKey   string          `json:"SigningPubKey"`
	TxnSignature    string          `json:"TxnSignature"`
	Account         string          `json:"Account"`
	Destination     string          `json:"Destination"`
	DestinationTag  int64           `json:"DestinationTag,omitempty"`
}

type Meta struct {
	TransactionResult TransactionResult `json:"TransactionResult,omitempty"`
	DeliveredAmount   interface{}       `json:"delivered_amount,omitempty"`
}

type DeliveredAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
}

type LedgerResponse struct {
	Ledger LedgerInfo `json:"ledger"`
}

type LedgerInfo struct {
	LedgerIndex  int64 `json:"ledger_index"`
	Transactions []Tx  `json:"transactions,omitempty"`
}
