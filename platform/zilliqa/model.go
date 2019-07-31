package zilliqa

import (
	"encoding/hex"
	"math/big"
	"strconv"
)

type Tx struct {
	Hash           string `json:"hash"`
	BlockHeight    uint64 `json:"blockHeight"`
	From           string `json:"from"`
	To             string `json:"to"`
	Value          string `json:"value"`
	Fee            string `json:"fee"`
	Timestamp      int64  `json:"timestamp"`
	Signature      string `json:"signature"`
	Nonce          uint64 `json:"nonce"`
	ReceiptSuccess bool   `json:"receiptSuccess"`
}

type ChainInfo struct {
	NumTxBlocks string `json:"NumTxBlocks"`
}

type TxRPC struct {
	ID       string `json:"ID"`
	Amount   string `json:"amount"`
	GasLimit string `json:"gasLimit"`
	GasPrice string `json:"gasPrice"`
	Nonce    string `json:"nonce"`
	Receipt  struct {
		CumulativeGas string `json:"cumulative_gas"`
		EpochNum      string `json:"epoch_num"`
		Success       bool   `json:"success"`
	} `json:"receipt"`
	SenderPubKey string `json:"senderPubKey"`
	Signature    string `json:"signature"`
	ToAddr       string `json:"toAddr"`
	Version      string `json:"version"`
}

func (t *TxRPC) toTx() Tx {
	to, _ := hex.DecodeString(t.ToAddr)
	nonce, _ := strconv.ParseUint(t.Nonce, 10, 64)
	height, _ := strconv.ParseUint(t.Receipt.EpochNum, 10, 64)
	gasLimt, _ := new(big.Int).SetString(t.GasLimit, 10)
	gasPrice, _ := new(big.Int).SetString(t.GasPrice, 10)
	fee := new(big.Int).Mul(gasLimt, gasPrice)

	tx := Tx{
		Hash:           t.ID,
		BlockHeight:    height,
		From:           EncodePublicKeyToAddress(t.SenderPubKey),
		To:             EncodeKeyHashToAddress(to),
		Value:          t.Amount,
		Fee:            fee.String(),
		Signature:      t.Signature,
		Nonce:          nonce,
		ReceiptSuccess: t.Receipt.Success,
	}
	return tx
}
