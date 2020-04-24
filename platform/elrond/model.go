package elrond

import (
	"math/big"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type LatestNonce struct {
	Nonce uint64 `json:"nonce"`
}

type Block struct {
	Nonce        uint64        `json:"nonce"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
}

type BulkTransactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Hash          string        `json:"hash"`
	MBHash        string        `json:"miniBlockHash"`
	BlockHash     string        `json:"blockHash"`
	Nonce         uint64        `json:"nonce"`
	Round         uint64        `json:"round"`
	Value         string        `json:"value"`
	Receiver      string        `json:"receiver"`
	Sender        string        `json:"sender"`
	ReceiverShard uint32        `json:"receiverShard"`
	SenderShard   uint32        `json:"senderShard"`
	GasPrice      uint64        `json:"gasPrice"`
	GasLimit      uint64        `json:"gasLimit"`
	Data          string        `json:"data"`
	Signature     string        `json:"signature"`
	Timestamp     time.Duration `json:"timestamp"`
	Status        string        `json:"status"`
}

func (t *Transaction) Fee() string {
	fee := big.NewInt(0).SetUint64(t.GasPrice)
	fee.Mul(fee, big.NewInt(0).SetUint64(t.GasLimit))

	return fee.String()
}

func (t *Transaction) Stratus() blockatlas.Status {
	switch t.Status {
	case "Success":
		return blockatlas.StatusCompleted
	case "Pending":
		return blockatlas.StatusPending
	default:
		return blockatlas.StatusError
	}
}

func (t *Transaction) Direction(address string) blockatlas.Direction {
	switch {
	case t.Sender == address && t.Receiver == address:
		return blockatlas.DirectionSelf
	case t.Sender == address && t.Receiver != address:
		return blockatlas.DirectionOutgoing
	default:
		return blockatlas.DirectionIncoming
	}
}
