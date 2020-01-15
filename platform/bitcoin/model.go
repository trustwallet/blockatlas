package bitcoin

import "encoding/json"

type TransactionsList struct {
	Page         int64         `json:"page"`
	TotalPages   int64         `json:"totalPages"`
	ItemsOnPage  int64         `json:"itemsOnPage"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Txs          interface{}   `json:"txs,omitempty"`
	Tokens       []Token       `json:"tokens,omitempty"`
	TxCount      int64         `json:"txCount,omitempty"`
	Hash         string        `json:"hash,omitempty"`
}

func (tl *TransactionsList) TransactionList() []Transaction {
	if tl.Transactions != nil {
		return tl.Transactions
	}
	b, err := json.Marshal(tl.Txs)
	if err != nil {
		return tl.Transactions
	}
	var txs []Transaction
	err = json.Unmarshal(b, &txs)
	if err != nil {
		return tl.Transactions
	}
	return txs
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

func (transaction Transaction) Amount() string {
	if len(transaction.Value) == 0 {
		return transaction.ValueOut
	}
	return transaction.Value
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

func (transaction Transaction) GetBlockHeight() uint64 {
	if transaction.BlockHeight > 0 {
		return uint64(transaction.BlockHeight)
	}
	return 0
}
