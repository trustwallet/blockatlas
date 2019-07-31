package zilliqa

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcutil/bech32"
	"math/big"
	"strconv"
	"strings"
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
	NumTxBlocks   string `json:"NumTxBlocks"`
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

func (t *TxRPC)toTx() Tx {
	to, _ := hex.DecodeString(t.ToAddr)
	nonce, _ := strconv.ParseUint(t.Nonce, 10, 64)
	height, _ := strconv.ParseUint(t.Receipt.EpochNum, 10, 64)
	gasLimt, _ := new(big.Int).SetString(t.GasLimit, 10)
	gasPrice, _ := new(big.Int).SetString(t.GasPrice, 10)
	fee := new(big.Int).Mul(gasLimt, gasPrice)

	tx := Tx{
		Hash: t.ID,
		BlockHeight: height,
		From: encodePublicKeyToAddress(t.SenderPubKey),
		To: encodeKeyHashToAddress(to),
		Value: t.Amount,
		Fee: fee.String(),
		Signature: t.Signature,
		Nonce: nonce,
		ReceiptSuccess: t.Receipt.Success,
	}
	return tx
}

func encodePublicKeyToAddress(hexString string) string {
	if strings.HasPrefix(hexString,"0x") {
		hexString = hexString[2:]
	}
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return ""
	}
	keyHash := sha256.Sum256(bytes)
	return encodeKeyHashToAddress(keyHash[12:])
}

func encodeKeyHashToAddress(keyHash []byte) string {
	conv, err := bech32.ConvertBits(keyHash, 8, 5, true)
	if err != nil {
		return ""
	}
	encoded, err := bech32.Encode("zil", conv)
	if err != nil {
		return ""
	}
	return encoded
}
