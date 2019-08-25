package nebulas

import "encoding/json"

type Response struct {
	Data ResponseData `json:"data"`
}

type NewBlockResponse struct {
	Data []NewBlock `json:"data"`
}

type ResponseData struct {
	Transactions []Transaction `json:"txnList"`
}

type Transaction struct {
	Hash      string      `json:"hash"`
	Type      string      `json:"type"`
	Value     json.Number `json:"value"`
	TxFee     string      `json:"txFee"`
	Nonce     uint64      `json:"nonce"`
	Block     Block       `json:"block"`
	From      Address     `json:"from"`
	To        Address     `json: "to"`
	Timestamp int64       `json:"timestamp"`
	Status    int32       `json:"status"`
}

type Block struct {
	Height uint64 `json:"height"`
}

type NewBlock struct {
	Height int64 `json:"height"`
}

type Address struct {
	Hash string `json:"hash"`
}
