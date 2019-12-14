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
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Timestamp        string `json:"timestamp"`
}

type BlockInfo struct {
	Hash             string        `json:"hash"`
	Number           string        `json:"number"`
	Transactions     []Transaction `json:"transactions"`
}

