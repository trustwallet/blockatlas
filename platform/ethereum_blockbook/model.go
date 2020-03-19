package ethereum_blockbook

import "math/big"

type Page struct {
	Transactions []Transaction `json:"transactions"`
	//Tokens       []Token       `json:"tokens"`
}

type TokenPage struct {
	Tokens []Token `json:"tokens"`
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
	VIN              []Vin             `json:"vin"`
	VOUT             []Vout            `json:"vout"`
	BlockHeight      uint64            `json:"blockHeight"`
	BlockTime        int64             `json:"blockTime"`
	Value            string            `json:"value"`
	Fees             string            `json:"fees"`
	TokenTransfers   []TokenTransfer   `json:"tokenTransfers"`
	EthereumSpecific *EthereumSpecific `json:"ethereumSpecific,omitempty"`
}

// Vin contains information about single transaction input
type Vin struct {
	Addresses []string `json:addresses`
}

// Vout contains information about single transaction output
type Vout struct {
	Value     string   `json:"value"`
	Addresses []string `json:addresses`
}

// TokenType specifies type of token
type TokenType string

// Amount is datatype holding amounts
type Amount big.Int

// ERC20TokenType is Ethereum ERC20 token
const ERC20TokenType TokenType = "ERC20"

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
	Balance  string    `json:"balance"`
	Contract string    `json:"contract"`
	Decimals uint      `json:"decimals"`
	Name     string    `json:"name"`
	Symbol   string    `json:"symbol"`
	Type     TokenType `json:"type"`
}

// EthereumSpecific contains ethereum specific transaction data
type EthereumSpecific struct {
	Status   int      `json:"status"` // -1 pending, 0 Fail, 1 OK
	Nonce    uint64   `json:"nonce"`
	GasLimit *big.Int `json:"gasLimit"`
	GasUsed  *big.Int `json:"gasUsed"`
	GasPrice *Amount  `json:"gasPrice"`
}
