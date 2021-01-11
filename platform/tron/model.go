package tron

import "github.com/trustwallet/golibs/types"

type (
	BlockRequest struct {
		StartNum int64 `json:"startNum"`
		EndNum   int64 `json:"endNum"`
	}

	Blocks struct {
		Blocks []Block `json:"block"`
	}

	Block struct {
		BlockId     string `json:"blockID"`
		Txs         []Tx   `json:"transactions"`
		BlockHeader struct {
			Data BlockData `json:"raw_data"`
		} `json:"block_header"`
	}

	BlockData struct {
		Number    int64 `json:"number"`
		Timestamp int64 `json:"timestamp"`
	}

	Page struct {
		Success bool   `json:"success"`
		Error   string `json:"error,omitempty"`
		Txs     []Tx   `json:"data"`
	}

	Tx struct {
		ID        string `json:"txID"`
		BlockTime int64  `json:"block_timestamp"`
		Data      TxData `json:"raw_data"`
	}

	TxData struct {
		Timestamp int64      `json:"timestamp"`
		Contracts []Contract `json:"contract"`
	}

	ContractType string

	Contract struct {
		Type      ContractType `json:"type"`
		Parameter struct {
			Value TransferValue `json:"value"`
		} `json:"parameter"`
	}

	TransferValue struct {
		Amount       types.Amount `json:"amount"`
		OwnerAddress string       `json:"owner_address"`
		ToAddress    string       `json:"to_address"`
		AssetName    string       `json:"asset_name,omitempty"`
	}

	Account struct {
		Data []AccountData `json:"data"`
	}

	AccountData struct {
		Balance  uint                `json:"balance"`
		AssetsV2 []AssetV2           `json:"assetV2"`
		Votes    []Votes             `json:"votes"`
		Frozen   []Frozen            `json:"frozen"`
		Trc20    []map[string]string `json:"trc20"`
	}

	AssetV2 struct {
		Key string `json:"key"`
	}

	Votes struct {
		VoteAddress string `json:"vote_address"`
		VoteCount   int    `json:"vote_count"`
	}

	Frozen struct {
		ExpireTime    int64       `json:"expire_time"`
		FrozenBalance interface{} `json:"frozen_balance,string"` // nolint
	}

	Asset struct {
		Data []AssetInfo `json:"data"`
	}

	AssetInfo struct {
		Name     string `json:"name"`
		Symbol   string `json:"abbr"`
		ID       uint   `json:"id"`
		Decimals uint   `json:"precision"`
	}

	Validators struct {
		Witnesses []Validator `json:"witnesses"`
	}

	Validator struct {
		Address string `json:"address"`
	}

	VotesRequest struct {
		Address string `json:"address"`
		Visible bool   `json:"visible"`
	}

	TRC20Transactions struct {
		Data []TRC20Transaction `json:"data"`
	}

	TRC20Transaction struct {
		From           string         `json:"from"`
		To             string         `json:"to"`
		BlockTimestamp int64          `json:"block_timestamp"`
		Value          string         `json:"value"`
		Type           string         `json:"type"`
		TransactionID  string         `json:"transaction_id"`
		TokenInfo      TRC20TokenInfo `json:"token_info"`
	}

	TRC20TokenInfo struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		Address  string `json:"address"`
	}

	ExplorerResponse struct {
		ExplorerTrc20Tokens []ExplorerTrc20Tokens `json:"trc20token_balances"`
	}

	ExplorerTrc20Tokens struct {
		Name            string `json:"name"`
		Symbol          string `json:"symbol"`
		Decimals        int    `json:"decimals"`
		ContractAddress string `json:"contract_address"`
	}
)

const (
	TransferContract      ContractType = "TransferContract"
	TransferAssetContract ContractType = "TransferAssetContract"
)
