package rpc

import "github.com/trustwallet/golibs/client"

type BlockTxs [][]string

func (b BlockTxs) txs() []string {
	txs := make([]string, 0)
	for _, ids := range b {
		txs = append(txs, ids...)
	}
	return txs
}

type ChainInfo struct {
	NumTxBlocks string `json:"NumTxBlocks"`
}

type Receipt struct {
	CumulativeGas string `json:"cumulative_gas"`
	EpochNum      string `json:"epoch_num"`
	Success       bool   `json:"success"`
}

type Tx struct {
	ID           string  `json:"ID"`
	Amount       string  `json:"amount"`
	GasLimit     string  `json:"gasLimit"`
	GasPrice     string  `json:"gasPrice"`
	Nonce        string  `json:"nonce"`
	Receipt      Receipt `json:"receipt"`
	SenderPubKey string  `json:"senderPubKey"`
	Signature    string  `json:"signature"`
	ToAddr       string  `json:"toAddr"`
	Version      string  `json:"version"`
}

type HashesResponse struct {
	ID      int        `json:"id"`
	Jsonrpc string     `json:"jsonrpc"`
	Result  [][]string `json:"result"`
}

func (h HashesResponse) Txs() []string {
	var result []string
	for _, subRes := range h.Result {
		result = append(result, subRes...)
	}
	return result
}

type BlockTxRpc struct {
	JsonRpc string           `json:"jsonrpc"`
	Error   *client.RpcError `json:"error,omitempty"`
	Result  BlockTxs         `json:"result,omitempty"`
	Id      string           `json:"id,omitempty"`
}

type Block struct {
	Header BlockHeader `json:"header"`
}

type BlockHeader struct {
	Number    string `json:"BlockNum"`
	Timestamp string `json:"Timestamp"`
}
