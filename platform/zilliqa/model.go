package zilliqa

import (
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type BlockTxs [][]string

func (b BlockTxs) txs() []string {
	txs := make([]string, 0)
	for _, ids := range b {
		txs = append(txs, ids...)
	}
	return txs
}

type Tx struct {
	Hash           string      `json:"hash"`
	BlockHeight    uint64      `json:"blockHeight"`
	From           string      `json:"from"`
	To             string      `json:"to"`
	Value          string      `json:"value"`
	Fee            string      `json:"fee"`
	Timestamp      int64       `json:"timestamp"`
	Signature      string      `json:"signature"`
	Nonce          interface{} `json:"nonce"`
	ReceiptSuccess bool        `json:"receiptSuccess"`
}

func (tx Tx) NonceValue() uint64 {
	switch n := tx.Nonce.(type) {
	case string:
		r, err := strconv.Atoi(n)
		if err != nil {
			break
		}
		return uint64(r)
	case int:
		return uint64(n)
	case float64:
		return uint64(n)
	}
	return 0
}

type ChainInfo struct {
	NumTxBlocks string `json:"NumTxBlocks"`
}

type TxReceipt struct {
	CumulativeGas string `json:"cumulative_gas"`
	EpochNum      string `json:"epoch_num"`
	Success       bool   `json:"success"`
}

type TxRPC struct {
	ID           string    `json:"ID"`
	Amount       string    `json:"amount"`
	GasLimit     string    `json:"gasLimit"`
	GasPrice     string    `json:"gasPrice"`
	Nonce        string    `json:"nonce"`
	Receipt      TxReceipt `json:"receipt"`
	SenderPubKey string    `json:"senderPubKey"`
	Signature    string    `json:"signature"`
	ToAddr       string    `json:"toAddr"`
	Version      string    `json:"version"`
}

func (t *TxRPC) toTx() *Tx {
	to, err := hex.DecodeString(t.ToAddr)
	if err != nil {
		return nil
	}
	height, err := strconv.ParseUint(t.Receipt.EpochNum, 10, 64)
	if err != nil {
		return nil
	}
	gasLimt, ok := new(big.Int).SetString(t.GasLimit, 10)
	if !ok {
		return nil
	}
	gasPrice, ok := new(big.Int).SetString(t.GasPrice, 10)
	if !ok {
		return nil
	}
	fee := new(big.Int).Mul(gasLimt, gasPrice)

	return &Tx{
		Hash:           "0x" + t.ID,
		BlockHeight:    height,
		From:           EncodePublicKeyToAddress(t.SenderPubKey),
		To:             EncodeKeyHashToAddress(to),
		Value:          t.Amount,
		Fee:            fee.String(),
		Signature:      t.Signature,
		Nonce:          t.Nonce,
		ReceiptSuccess: t.Receipt.Success,
	}
}

type BlockTxRpc struct {
	JsonRpc string               `json:"jsonrpc"`
	Error   *blockatlas.RpcError `json:"error,omitempty"`
	Result  BlockTxs             `json:"result,omitempty"`
	Id      string               `json:"id,omitempty"`
}
