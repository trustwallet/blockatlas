package nebulas

import "encoding/json"

type Response struct {
	Data ResponseData `json:"data"`
}

type NasResponse struct {
	Result Block `json:"result"`
}


type ResponseData struct {
	TxnList []Transaction `json:"txnList"`
}

type Transaction struct {
	Hash        string      `json:"hash"`
	Type        string      `json:"type"`
	Value       json.Number `json:"value"`
	TxFee       string      `json:"txFee"`
	Nonce       uint64      `json:"nonce"`
	Block       Block       `json:"block"`
	From        Address     `json:"from"`
	To          Address     `json: "to"`
	Timestamp   int64       `json:"timestamp"`
	Status      int32       `json:"status"`
	GasPrice	string		`json:"gas_price"`
	GasLimit	string		`json:"gas_limit"`
	GasUsed		string		`json:"gas_used"`
}

type Block struct {
	Height uint64 `json:"height"`
	Nonce  uint64 `json:"nonce"`
	TxnList []Transaction `json:"transactions"`
}

type Address struct {
	Hash string `json:"hash"`
}
