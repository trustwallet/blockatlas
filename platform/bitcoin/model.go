package bitcoin

type TransactionsList struct {
	Page         int64         `json:"page"`
	TotalPages   int64         `json:"totalPages"`
	ItemsOnPage  int64         `json:"itemsOnPage"`
	Transactions []Transaction `json:"transactions"`
	Tokens       []Token       `json:"tokens,omitempty"`
}

type Block struct {
	Page         int64         `json:"page"`
	TotalPages   int64         `json:"totalPages"`
	ItemsOnPage  int64         `json:"itemsOnPage"`
	Transactions []Transaction `json:"txs"`
	TxCount      int64         `json:"txCount"`
	Hash         string        `json:"hash"`
}

type Tx struct {
	ID string `json:"id"`
}

type Transaction struct {
	ID            string   `json:"txid"`
	Version       uint64   `json:"version"`
	Vin           []Output `json:"vin"`
	Vout          []Output `json:"vout"`
	BlockHash     string   `json:"blockHash"`
	BlockHeight   uint64   `json:"blockHeight"`
	Confirmations uint64   `json:"confirmations"`
	BlockTime     uint64   `json:"blockTime"`
	Value         string   `json:"value"`
	ValueIn       string   `json:"valueIn"`
	Fees          string   `json:"fees"`
	Hex           string   `json:"hex"`
}

type Output struct {
	TxId      string   `json:"txid,omitempty"`
	Value     string   `json:"value,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
	Hex       string   `json:"hex,omitempty"`
}

type Token struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Transfers int    `json:"transfers"`
	Balance   string `json:"balance"`
}

type BlockchainStatus struct {
	Backend Backend `json:"backend"`
}

type Backend struct {
	Chain  string `json:"chain"`
	Blocks int64  `json:"blocks"`
}
