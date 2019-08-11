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

func ReceiptStatus(r bool) string {
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

type CurrentBlockInfo struct {
	BestBlockNum int64 `json:"bestBlockNum"`
}

type Block struct {
	Id           string   `json:"Id"`
	Transactions []string `json:"transactions"`
}

type Event struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type Transfer struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type Output struct {
	Events    []Event    `json:"events"`
	Transfers []Transfer `json:"transfers"`
}

type TransactionReceipt struct {
	Outputs  []Output `json:"outputs"`
	Paid     string   `json:"paid"`
	Reverted bool     `json:"reverted"`
}

type NativeTransaction struct {
	Block     uint64             `json:"block"`
	Clauses   []Clause           `json:"clauses"`
	ID        string             `json:"id"`
	Origin    string             `json:"origin"`
	Receipt   TransactionReceipt `json:"receipt"`
	Reverted  int64              `json:"reverted"`
	Timestamp int64              `json:"timestamp"`
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
