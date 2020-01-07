package bitcoin

type TransactionsList struct {
	Page         int64         `json:"page"`
	TotalPages   int64         `json:"totalPages"`
	ItemsOnPage  int64         `json:"itemsOnPage"`
	Transactions []Transaction `json:"transactions"`
	Txs          []Transaction `json:"txs"`
	Tokens       []Token       `json:"tokens,omitempty"`
	TxCount      int64         `json:"txCount,omitempty"`
	Hash         string        `json:"hash,omitempty"`
}

func (tl *TransactionsList) TransactionList() []Transaction {
	if tl.Transactions == nil {
		return tl.Txs
	}
	return tl.Transactions
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
	BlockHeight   int64    `json:"blockHeight"`
	Confirmations uint64   `json:"confirmations"`
	BlockTime     uint64   `json:"blockTime"`
	Value         string   `json:"value"`
	ValueOut      string   `json:"valueOut"`
	Fees          string   `json:"fees"`
}

func (t Transaction) Amount() string {
	if len(t.Value) == 0 {
		return t.ValueOut
	}
	return t.Value
}

type Output struct {
	TxId         string       `json:"txid,omitempty"`
	Value        string       `json:"value,omitempty"`
	Addresses    []string     `json:"addresses,omitempty"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey,omitempty"`
}

type ScriptPubKey struct {
	Addresses []string `json:"addresses,omitempty"`
}

func (o Output) OutputAddress() []string {
	if len(o.Addresses) == 0 {
		return o.ScriptPubKey.Addresses
	}
	return o.Addresses
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
