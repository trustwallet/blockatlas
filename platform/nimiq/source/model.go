package source

import "errors"

type Tx struct {
	Hash          string `json:"hash"`
	FromAddress   string `json:"fromAddress"`
	ToAddress     string `json:"toAddress"`
	Value         uint64 `json:"value"`
	Fee           uint64 `json:"fee"`
}

var ErrSourceConn  = errors.New("connection to servers failed")
var ErrInvalidAddr = errors.New("invalid address")
