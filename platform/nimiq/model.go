package nimiq

import (
	"encoding/json"

	"github.com/trustwallet/golibs/types"
)

type Tx struct {
	Hash          string       `json:"hash"`
	BlockHash     string       `json:"blockHash"`
	BlockNumber   uint64       `json:"blockNumber"`
	Timestamp     json.Number  `json:"timestamp"`
	Confirmations int          `json:"confirmations"`
	TxIndex       int          `json:"transactionIndex"`
	FromAddress   string       `json:"fromAddress"`
	ToAddress     string       `json:"toAddress"`
	Value         types.Amount `json:"value"`
	Fee           types.Amount `json:"fee"`
}

type Block struct {
	Number       int64  `json:"number"`
	Hash         string `json:"hash"`
	PoW          string `json:"pow"`
	ParentHash   string `json:"parentHash"`
	Nonce        uint32 `json:"nonce"`
	BodyHash     string `json:"bodyHash"`
	AccountsHash string `json:"accountsHash"`
	MinerHex     string `json:"miner"`
	Miner        string `json:"minerAddress"`
	Difficulty   string `json:"difficulty"`
	ExtraData    string `json:"extraData"`
	Size         int64  `json:"size"`
	Timestamp    int64  `json:"timestamp"`
	Txs          []Tx   `json:"transactions"`
}
