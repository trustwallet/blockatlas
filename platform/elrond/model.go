package elrond

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type GenericResponse struct {
	Data  json.RawMessage `json:"data"`
	Code  string          `json:"code"`
	Error string          `json:"error"`
}

type NetworkStatus struct {
	Status StatusDetails `json:"status"`
}

type StatusDetails struct {
	Round float64 `json:"erd_current_round"`
	Epoch float64 `json:"erd_epoch_number"`
	Nonce float64 `json:"erd_nonce"`
}

type BlockResponse struct {
	Block Block `json:"hyperblock"`
}

type Block struct {
	Nonce        uint64        `json:"nonce"`
	Hash         string        `json:"hash"`
	Transactions []Transaction `json:"transactions"`
}

type TransactionsPage struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Type      string        `json:"type"`
	Hash      string        `json:"hash"`
	Nonce     uint64        `json:"nonce"`
	Value     string        `json:"value"`
	Receiver  string        `json:"receiver"`
	Sender    string        `json:"sender"`
	Data      string        `json:"data"`
	Timestamp time.Duration `json:"timestamp"`
	Status    string        `json:"status"`
	Fee       string        `json:"fee"`
	GasPrice  uint64        `json:"gasPrice,omitempty"`
	GasLimit  uint64        `json:"gasLimit,omitempty"`
}

func (tx *Transaction) TxFee() blockatlas.Amount {
	if tx.Fee != "0" && tx.Fee != "" {
		return blockatlas.Amount(tx.Fee)
	}

	txFee := big.NewInt(0).SetUint64(tx.GasPrice)
	txFee = txFee.Mul(txFee, big.NewInt(0).SetUint64(tx.GasLimit))

	return blockatlas.Amount(txFee.String())
}

func (tx *Transaction) TxStatus() blockatlas.Status {
	switch tx.Status {
	case "Success", "success":
		return blockatlas.StatusCompleted
	case "Pending", "pending":
		return blockatlas.StatusPending
	default:
		return blockatlas.StatusError
	}
}

func (tx *Transaction) Direction(address string) blockatlas.Direction {
	switch {
	case tx.Sender == address && tx.Receiver == address:
		return blockatlas.DirectionSelf
	case tx.Sender == address && tx.Receiver != address:
		return blockatlas.DirectionOutgoing
	default:
		return blockatlas.DirectionIncoming
	}
}
