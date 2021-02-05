package viewblock

import (
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
