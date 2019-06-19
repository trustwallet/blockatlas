package semux

import "github.com/trustwallet/blockatlas"

type Tx struct {
	BlockNumber string        `json:"blockNumber"`
	Data        string        `json:"data"`
	Fee         blockatlas.Amount `json:"fee"`
	From        string        `json:"from"`
	Hash        string        `json:"hash"`
	Nonce       string        `json:"nonce"`
	Timestamp   string        `json:"timestamp"`
	To          string        `json:"to"`
	Type        string        `json:"type"`
	Value       blockatlas.Amount `json:"value"`
}

type Account struct {
	Available               blockatlas.Amount `json:"available"`
	Locked                  blockatlas.Amount `json:"locked"`
	Nonce                   string        `json:"nonce"`
	PendingTransactionCount uint64        `json:"pendingTransactionCount"`
	TransactionCount        uint64        `json:"transactionCount"`
}

type GetAccountTransactionsResponse struct {
	Result  []Tx   `json:"result"`
}

type GetAccountResponse struct {
	Result  Account `json:"result"`
}
