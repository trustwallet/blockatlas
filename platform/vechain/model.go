package vechain

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

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
	Receipt   *Receipt `json:"receipt"`
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

// ReceiptStatus function that describes transaction status
func ReceiptStatus(r bool) blockatlas.Status {
	if r {
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

// CurrentBlockInfo type is a model with current blockchain height
type CurrentBlockInfo struct {
	BestBlockNum int64 `json:"bestBlockNum"`
}

// Block type is a VeChain block model
type Block struct {
	ID           string   `json:"Id"`
	Transactions []string `json:"transactions"`
}

// Event type is a field in native transaction with contract call info
type Event struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

// Transfer type is a field in native transaction with VET transfer data
type Transfer struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
}

// Output type is a field in native transaction
type Output struct {
	Events    []Event    `json:"events"`
	Transfers []Transfer `json:"transfers"`
}

// TransactionReceipt type for parsing receipt info
type TransactionReceipt struct {
	Outputs  []Output `json:"outputs"`
	Paid     string   `json:"paid"`
	Reverted bool     `json:"reverted"`
}

// NativeTransaction type for Native VeChain transaction with full transfer info
type NativeTransaction struct {
	Block     uint64             `json:"block"`
	Clauses   []Clause           `json:"clauses"`
	ID        string             `json:"id"`
	Origin    string             `json:"origin"`
	Receipt   TransactionReceipt `json:"receipt"`
	Reverted  int64              `json:"reverted"`
	Timestamp int64              `json:"timestamp"`
}

// Error model for request error
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
