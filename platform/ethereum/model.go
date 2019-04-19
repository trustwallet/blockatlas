package ethereum

type Page struct {
	Total uint  `json:"total"`
	Docs  []Doc `json:"docs"`
}

type Doc struct {
	Ops         []Op      `json:"operations"`
	Contract    *Contract `json:"contract"`
	ID          string    `json:"id"`
	BlockNumber uint64    `json:"blockNumber"`
	TimeStamp   string    `json:"timeStamp"`
	Nonce       uint64    `json:"nonce"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Value       string    `json:"value"`
	Gas         string    `json:"gas"`
	GasPrice    string    `json:"gasPrice"`
	GasUsed     string    `json:"gasUsed"`
	Input       string    `json:"input"`
	Error       string    `json:"error"`
	Coin        uint      `json:"coin"`
}

type Op struct {
	TxID     string    `json:"transactionId"`
	Contract *Contract `json:"contract"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Type     string    `json:"type"`
	Value    string    `json:"value"`
	Coin     uint      `json:"coin"`
}

type Contract struct {
	Address     string `json:"address"`
	Symbol      string `json:"symbol"`
	Decimals    uint   `json:"decimals"`
	TotalSupply string `json:"totalSupply"`
	Name        string `json:"name"`
}
