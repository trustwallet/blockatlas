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

type BlockResponse struct {
	Result Result `json:"result"`
}

type Result struct {
	ChainID       int    `json:"chain_id"`
	Coinbase      string `json:"coinbase"`
	ConsensusRoot struct {
		DynastyRoot string `json:"dynasty_root"`
		Proposer    string `json:"proposer"`
		Timestamp   string `json:"timestamp"`
	} `json:"consensus_root"`
	EventsRoot   string        `json:"events_root"`
	Hash         string        `json:"hash"`
	Height       string        `json:"height"`
	IsFinality   bool          `json:"is_finality"`
	Miner        string        `json:"miner"`
	Nonce        string        `json:"nonce"`
	ParentHash   string        `json:"parent_hash"`
	RandomProof  string        `json:"randomProof"`
	RandomSeed   string        `json:"randomSeed"`
	StateRoot    string        `json:"state_root"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions,omitempty"`
	TxsRoot      string        `json:"txs_root"`
}
