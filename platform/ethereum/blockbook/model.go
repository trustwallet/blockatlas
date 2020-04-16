package blockbook

import (
	"math/big"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Page struct {
	Transactions []Transaction `json:"transactions,omitempty"`
	Tokens       []Token       `json:"tokens,omitempty"`
}

type NodeInfo struct {
	Blockbook *Blockbook `json:"blockbook"`
	Backend   *Backend   `json:"backend"`
}

type Blockbook struct {
	BestHeight int64 `json:"bestHeight"`
}

type Backend struct {
	Blocks int64 `json:"blocks"`
}

type Block struct {
	Transactions []Transaction `json:"txs"`
}

type Transaction struct {
	TxID             string            `json:"txid"`
	Vin              []Output          `json:"vin"`
	Vout             []Output          `json:"vout"`
	BlockHeight      int64             `json:"blockHeight"`
	BlockTime        int64             `json:"blockTime"`
	Value            string            `json:"value"`
	Fees             string            `json:"fees"`
	TokenTransfers   []TokenTransfer   `json:"tokenTransfers,omitempty"`
	EthereumSpecific *EthereumSpecific `json:"ethereumSpecific,omitempty"`
}

type Output struct {
	Value     string   `json:"value,omitempty"`
	Addresses []string `json:"addresses"`
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
	Contract string               `json:"contract"`
	Decimals uint                 `json:"decimals"`
	Name     string               `json:"name"`
	Symbol   string               `json:"symbol"`
	Type     blockatlas.TokenType `json:"type"`
}

// EthereumSpecific contains ethereum specific transaction data
type EthereumSpecific struct {
	Status   int      `json:"status"` // -1 pending, 0 Fail, 1 OK
	Nonce    uint64   `json:"nonce"`
	GasLimit *big.Int `json:"gasLimit"`
	GasUsed  *big.Int `json:"gasUsed"`
	GasPrice string   `json:"gasPrice"`
}
