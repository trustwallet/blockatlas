package blockbook

import (
	"encoding/json"
	"math/big"

	"github.com/trustwallet/golibs/types"
)

type NodeInfo struct {
	Blockbook *Blockbook `json:"blockbook"`
}

type Blockbook struct {
	BestHeight int64 `json:"bestHeight"`
	InSync     bool  `json:"inSync"`
}

type TokenTransfer struct {
	Decimals uint   `json:"decimals"`
	From     string `json:"from"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	To       string `json:"to"`
	Token    string `json:"token"`
	Type     string `json:"type"`
	Value    string `json:"value"`
}

// Token contains info about tokens held by an address
type Token struct {
	Balance  string          `json:"balance,omitempty"`
	Contract string          `json:"contract"`
	Decimals uint            `json:"decimals"`
	Name     string          `json:"name"`
	Symbol   string          `json:"symbol"`
	Type     types.TokenType `json:"type"`
}

// EthereumSpecific contains ethereum specific transaction data
type EthereumSpecific struct {
	Status   int      `json:"status"` // -1 pending, 0 Fail, 1 OK
	Nonce    uint64   `json:"nonce"`
	GasLimit *big.Int `json:"gasLimit"`
	GasUsed  *big.Int `json:"gasUsed"`
	GasPrice string   `json:"gasPrice"`
	Data     string   `json:"data,omitempty"`
}

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

type Transaction struct {
	ID               string            `json:"txid"`
	Version          uint64            `json:"version"`
	Vin              []Output          `json:"vin"`
	Vout             []Output          `json:"vout"`
	BlockHash        string            `json:"blockHash"`
	BlockHeight      int64             `json:"blockHeight"`
	Confirmations    uint64            `json:"confirmations"`
	BlockTime        int64             `json:"blockTime"`
	Value            string            `json:"value"`
	ValueOut         string            `json:"valueOut"`
	Fees             string            `json:"fees"`
	TokenTransfers   []TokenTransfer   `json:"tokenTransfers,omitempty"`
	EthereumSpecific *EthereumSpecific `json:"ethereumSpecific,omitempty"`
}

func (transaction Transaction) Amount() string {
	if len(transaction.Value) == 0 {
		return transaction.ValueOut
	}
	return transaction.Value
}

func (transaction Transaction) GetStatus() types.Status {
	if transaction.Confirmations == 0 {
		return types.StatusPending
	}
	return types.StatusCompleted
}

func (transaction Transaction) GetBlockHeight() uint64 {
	if transaction.BlockHeight > 0 {
		return uint64(transaction.BlockHeight)
	}
	return 0
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
