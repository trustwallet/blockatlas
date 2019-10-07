package zilliqa

import (
	"encoding/hex"
	"math/big"
	"strconv"
)

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
	n, ok := tx.Nonce.(string)
	if ok {
		r, _ := strconv.Atoi(n)
		return uint64(r)
	}
	nu, ok := tx.Nonce.(int)
	if ok {
		return uint64(nu)
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

func (t *TxRPC) toTx() Tx {
	to, _ := hex.DecodeString(t.ToAddr)
	height, _ := strconv.ParseUint(t.Receipt.EpochNum, 10, 64)
	gasLimt, _ := new(big.Int).SetString(t.GasLimit, 10)
	gasPrice, _ := new(big.Int).SetString(t.GasPrice, 10)
	fee := new(big.Int).Mul(gasLimt, gasPrice)

	tx := Tx{
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
	return tx
}
