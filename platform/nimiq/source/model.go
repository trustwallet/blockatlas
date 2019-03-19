package source

import "errors"

type Tx struct {
	Hash          string `json:"hash"`
	BlockHash     string `json:"blockHash"`
	BlockNumber   uint64 `json:"blockNumber"`
	Timestamp     int64  `json:"timestamp"`
	Confirmations int    `json:"confirmations"`
	TxIndex       int    `json:"transactionIndex"`
	FromAddress   string `json:"fromAddress"`
	ToAddress     string `json:"toAddress"`
	Value         uint64 `json:"value"`
	Fee           uint64 `json:"fee"`
}

var ErrSourceConn  = errors.New("connection to servers failed")
var ErrInvalidAddr = errors.New("invalid address")
