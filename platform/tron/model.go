package tron

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type BlockRequest struct {
	StartNum int64 `json:"startNum"`
	EndNum   int64 `json:"endNum"`
}

type Blocks struct {
	Blocks []Block `json:"block"`
}

type Block struct {
	BlockId     string `json:"blockID"`
	Txs         []Tx   `json:"transactions"`
	BlockHeader struct {
		Data BlockData `json:"raw_data"`
	} `json:"block_header"`
}

type BlockData struct {
	Number    int64 `json:"number"`
	Timestamp int64 `json:"timestamp"`
}

type Page struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Txs     []Tx   `json:"data"`
}

type Tx struct {
	ID        string `json:"txID"`
	BlockTime int64  `json:"block_timestamp"`
	Data      TxData `json:"raw_data"`
}

type TxData struct {
	Timestamp int64      `json:"timestamp"`
	Contracts []Contract `json:"contract"`
}

type ContractType string

const (
	TransferContract      ContractType = "TransferContract"
	TransferAssetContract ContractType = "TransferAssetContract"
)

type Contract struct {
	Type      ContractType `json:"type"`
	Parameter struct {
		Value TransferValue `json:"value"`
	} `json:"parameter"`
}

type TransferValue struct {
	Amount       blockatlas.Amount `json:"amount"`
	OwnerAddress string            `json:"owner_address"`
	ToAddress    string            `json:"to_address"`
	AssetName    string            `json:"asset_name,omitempty"`
}

type Account struct {
	Data []AccountData `json:"data"`
}

type AccountData struct {
	Balance  uint                `json:"balance"`
	AssetsV2 []AssetV2           `json:"assetV2"`
	Votes    []Votes             `json:"votes"`
	Frozen   []Frozen            `json:"frozen"`
	Trc20    []map[string]string `json:"trc20"`
}

type AssetV2 struct {
	Key string `json:"key"`
}

type Votes struct {
	VoteAddress string `json:"vote_address"`
	VoteCount   int    `json:"vote_count"`
}

type Frozen struct {
	ExpireTime    int64       `json:"expire_time"`
	FrozenBalance interface{} `json:"frozen_balance,string"` // nolint
}

type Asset struct {
	Data []AssetInfo `json:"data"`
}

type AssetInfo struct {
	Name     string `json:"name"`
	Symbol   string `json:"abbr"`
	ID       string `json:"id"`
	Decimals uint   `json:"precision"`
}

type Validators struct {
	Witnesses []Validator `json:"witnesses"`
}

type Validator struct {
	Address string `json:"address"`
}

type VotesRequest struct {
	Address string `json:"address"`
	Visible bool   `json:"visible"`
}

type TRC20Transactions struct {
	Data []TRC20Transaction `json:"data"`
}

type TRC20Transaction struct {
	From           string         `json:"from"`
	To             string         `json:"to"`
	BlockTimestamp int64          `json:"block_timestamp"`
	Value          string         `json:"value"`
	Type           string         `json:"type"`
	TransactionID  string         `json:"transaction_id"`
	TokenInfo      TRC20TokenInfo `json:"token_info"`
}

type TRC20TokenInfo struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
	Address  string `json:"address"`
}

type ExplorerResponse struct {
	ExplorerTrc20Tokens []ExplorerTrc20Tokens `json:"trc20token_balances"`
}

type ExplorerTrc20Tokens struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        int    `json:"decimals"`
	ContractAddress string `json:"contract_address"`
}
