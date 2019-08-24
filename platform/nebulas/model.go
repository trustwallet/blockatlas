package nebulas

import "encoding/json"

type Response struct {
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	TxnList []Transaction `json:"txnList"`
}

type Transaction struct {
	Hash      string      `json:"hash"`
	Type      string      `json:"type"`
	Value     json.Number `json:"value"`
	TxFee     string      `json:"txFee"`
	Nonce     uint64      `json:"nonce"`
	Block     Block       `json:"block"`
	From      Address     `json:"from"`
	To        Address     `json:"to"`
	Timestamp int64       `json:"timestamp"`
	Status    int32       `json:"status"`
}

type NebulaTransaction struct {
	Hash            string      `json:"hash"`
	ChainID         uint64      `json:"chainId"`
	From            string      `json:"from"`
	To              string      `json:"to"`
	Value           string      `json:"value"`
	Nonce           string      `json:"nonce"`
	Timestamp       string      `json:"timestamp"`
	Type            string      `json:"type"`
	Data            interface{} `json:"data"`
	GasPrice        string      `json:"gas_price"`
	GasLimit        string      `json:"gas_limit"`
	ContractAddress string      `json:"contract_address"`
	Status          uint64      `json:"status"`
	GasUsed         string      `json:"gas_used"`
}

type Block struct {
	Height uint64 `json:"height"`
}

type Address struct {
	Hash string `json:"hash"`
}

type BlockResponse struct {
	Result NebulaBlock `json:"result"`
}

type ConsensusRoot struct {
	Timestamp   string `json:"timestamp"`
	Proposer    string `json:"proposer"`
	DynastyRoot string `json:"dynasty_root"`
}

// Different response gotten from the documentation added in the issue
// https://github.com/nebulasio/wiki/blob/master/rpc.md#latestirreversibleblock

type NebulaBlock struct {
	Hash          string              `json:"hash"`
	ParentHash    string              `json:"parent_hash"`
	Height        string              `json:"height"`
	Nonce         string              `json:"nonce"`
	Coinbase      string              `json:"coinbase"`
	Timestamp     string              `json:"timestamp"`
	ChainID       uint64              `json:"chain_id"`
	StateRoot     string              `json:"state_root"`
	TxsRoot       string              `json:"txs_root"`
	EventsRoot    string              `json:"events_root"`
	ConsensusRoot ConsensusRoot       `json:"consensus_root"`
	Miner         string              `json:"miner"`
	IsFinality    bool                `json:"is_finality"`
	Transactions  []NebulaTransaction `json:"transactions"`
}
