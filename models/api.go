package models

import "errors"

// Types of transaction metadata
const (
	TxTransfer            = "transfer"
	TxNativeTokenTransfer = "native_token_transfer"
	TxTokenTransfer       = "token_transfer"
	TxCollectibleTransfer = "collectible_transfer"
	TxTokenSwap           = "token_swap"
	TxContractCall        = "contract_call"
)

// Types of transaction statuses
const (
	StatusCompleted = "completed"
	StatusPending   = "pending"
	StatusFailed    = "failed"
)

// TxPerPage says how many transactions to return per page
const TxPerPage = 25

// Response is a page of transactions
type Response []Tx

// Amount is a positive decimal integer string.
// It is written in the smallest possible unit (e.g. Wei, Satoshis)
type Amount string

// Tx is a generic
type Tx struct {
	// Unique identifier
	ID string `json:"id"`
	// SLIP-44 coin index of the platform
	Coin uint `json:"coin"`
	// Address of the transaction sender
	From string `json:"from"`
	// Address of the transaction recipient
	To string `json:"to"`
	// Transaction fee (native currency)
	Fee Amount `json:"fee"`
	// Unix timestamp of the block the transaction was included in
	Date int64 `json:"date"`
	// Height of the block the transaction was included in
	Block uint64 `json:"block,omitempty"`
	// Status of the transaction
	Status string `json:"status"`
	// Empty if the transaction was successful,
	// else error explaining why the transaction failed (optional)
	Error string `json:"error,omitempty"`
	// Transaction nonce or sequence
	Sequence uint64 `json:"sequence,omitempty"`
	// Type of metadata
	Type string `json:"type"`
	// Meta data object
	Meta interface{} `json:"metadata"`
}

// Transfer describes the transfer of currency native to the platform
type Transfer struct {
	Value Amount `json:"value"`
}

// NativeTokenTransfer describes the transfer of native tokens.
// Example: Stellar Tokens, TRC10
type NativeTokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
}

// TokenTransfer describes the transfer of non-native tokens.
// Examples: ERC-20, TRC20
type TokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
	From     string `json:"from"`
	To       string `json:"to"`
}

// CollectibleTransfer describes the transfer of a
// "collectible", unique token.
type CollectibleTransfer struct {
	Name     string `json:"name"`
	Contract string `json:"contract"`
	ImageURL string `json:"image_url"`
}

// TokenSwap describes the exchange of two different tokens
type TokenSwap struct {
	Input  TokenTransfer `json:"input"`
	Output TokenTransfer `json:"output"`
}

// ContractCall describes a
type ContractCall struct {
	Input string `json:"input"`
	Value string `json:"value"`
}

// ErrSourceConn signals that the connection to the source API failed
var ErrSourceConn  = errors.New("connection to servers failed")

// ErrInvalidAddr signals that the requested address is invalid
var ErrInvalidAddr = errors.New("invalid address")

// ErrNotFound signals that the resource has not been found
var ErrNotFound = errors.New("not found")
