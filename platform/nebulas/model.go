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
	To        Address     `json: "to"`
	Timestamp int64       `json:"timestamp"`
	Status    int32       `json:"status"`
}

type Block struct {
	Height uint64 `json:"height"`
}

type Address struct {
	Hash string `json:"hash"`
}

type LatestIrreversibleBlockResponse struct {
	Result LatestIrreversibleBlock `json:"result"`
}

type LatestIrreversibleBlock struct {
	ChainID       int    `json:"chain_id"`
	Coinbase      string `json:"coinbase"`
	ConsensusRoot struct {
		DynastyRoot string `json:"dynasty_root"`
		Proposer    string `json:"proposer"`
		Timestamp   string `json:"timestamp"`
	} `json:"consensus_root"`
	EventsRoot   string `json:"events_root"`
	Hash         string `json:"hash"`
	Height       string `json:"height"`
	IsFinality   bool   `json:"is_finality"`
	Miner        string `json:"miner"`
	Nonce        string `json:"nonce"`
	ParentHash   string `json:"parent_hash"`
	RandomProof  string `json:"randomProof"`
	RandomSeed   string `json:"randomSeed"`
	StateRoot    string `json:"state_root"`
	Timestamp    string `json:"timestamp"`
	Transactions []struct {
		BlockHeight     string      `json:"block_height"`
		ChainID         int         `json:"chainId"`
		ContractAddress string      `json:"contract_address"`
		Data            interface{} `json:"data"`
		ExecuteError    string      `json:"execute_error"`
		ExecuteResult   string      `json:"execute_result"`
		From            string      `json:"from"`
		GasLimit        string      `json:"gas_limit"`
		GasPrice        string      `json:"gas_price"`
		GasUsed         string      `json:"gas_used"`
		Hash            string      `json:"hash"`
		Nonce           string      `json:"nonce"`
		Status          int         `json:"status"`
		Timestamp       string      `json:"timestamp"`
		To              string      `json:"to"`
		Type            string      `json:"type"`
		Value           string      `json:"value"`
	} `json:"transactions"`
	TxsRoot string `json:"txs_root"`
}
