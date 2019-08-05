package vechain

import "github.com/trustwallet/blockatlas"

type TransferTx struct {
	Transactions []Tx `json:"transactions"`
}

type Tx struct {
	ID string `json:"id"`
}

type TransferReceipt struct {
	Block     uint64   `json:"block"`
	Clauses   []Clause `json:"clauses"`
	ID        string   `json:"id"`
	Nonce     string   `json:"nonce"`
	Origin    string   `json:"origin"`
	Receipt   Receipt  `json:"receipt"`
	Timestamp uint64   `json:"timestamp"`
}

type Clause struct {
	To    string `json:"to"`
	Value string `json:"value"`
}

type Meta struct {
	BlockID        string `json:"blockID"`
	BlockNumber    int    `json:"blockNumber"`
	BlockTimestamp int    `json:"blockTimestamp"`
	TxID           string `json:"txID"`
	TxOrigin       string `json:"txOrigin"`
}

type Receipt struct {
	Paid     string `json:"paid"`
	Reverted bool   `json:"reverted"`
}

func (r *Receipt) Status() string {
	if r.Reverted {
		return blockatlas.StatusFailed
	}
	return blockatlas.StatusCompleted
}

type TokenTransferTxs struct {
	TokenTransfers []TokenTransfer `json:"tokenTransfers"`
}

type TokenTransfer struct {
	Amount          string `json:"amount"`
	Block           uint64 `json:"block"`
	ContractAddress string `json:"contractAddress"`
	Origin          string `json:"origin"`
	Receiver        string `json:"receiver"`
	Timestamp       int64  `json:"timestamp"`
	TxID            string `json:"txId"`
}
