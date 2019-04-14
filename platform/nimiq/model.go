package nimiq

import (
	"errors"
	"github.com/trustwallet/blockatlas/models"
)

type Tx struct {
	Hash          string `json:"hash"`
	BlockHash     string `json:"blockHash"`
	BlockNumber   uint64 `json:"blockNumber"`
	Timestamp     int64  `json:"timestamp"`
	Confirmations int    `json:"confirmations"`
	TxIndex       int    `json:"transactionIndex"`
	FromAddress   string `json:"fromAddress"`
	ToAddress     string `json:"toAddress"`
	Value         models.Amount `json:"value"`
	Fee           models.Amount `json:"fee"`
}

var ErrSourceConn  = errors.New("connection to servers failed")
var ErrInvalidAddr = errors.New("invalid address")
