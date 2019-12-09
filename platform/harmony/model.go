package harmony

type TxResponse struct {
	Result TxResult `json:"result"`
}

type TxResult struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	ShardID          int    `json:"shardID"`
	ToShardID        int    `json:"toShardID"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type BlockInfo struct {
	Result BlockResult `json:"result"`
}

type BlockResult struct {
	Difficulty       int           `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            int           `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactionsRoot"`
	Uncles           []interface{} `json:"uncles"`
}

