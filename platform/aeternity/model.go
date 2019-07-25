package aeternity

import "encoding/json"

type Transaction struct {
	Hash        string `json:"hash"`
	BlockHeight uint64 `json:"block_height"`
	Timestamp   int64  `json:"time"`
	TxValue     Tx     `json:"tx"`
}

type Tx struct {
	Sender    string      `json:"sender_id"`
	Recipient string      `json:"recipient_id"`
	Amount    json.Number `json:"amount"`
	Fee       json.Number `json:"fee"`
	Type      string      `json:"type"`
	Payload   string      `json:"payload"`
	Nonce     uint64      `json:"nonce"`
}
