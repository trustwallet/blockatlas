package tron

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

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
	Contracts []Contract `json:"contract"`
}

type Contract struct {
	Type      string      `json:"type"`
	Parameter interface{} `json:"parameter"`
}

type TransferContract struct {
	Value TransferValue `json:"value"`
}

type TransferValue struct {
	Amount       blockatlas.Amount `json:"amount"`
	OwnerAddress string            `json:"owner_address"`
	ToAddress    string            `json:"to_address"`
}

// Type for token transfer
type TransferAssetContract struct {
	Value TransferAssetValue `json:"value"`
}

type TransferAssetValue struct {
	TransferValue
	AssetName string `json:"asset_name"`
}

type Account struct {
	Data []AccountData `json:"data"`
}

type AccountData struct {
	Balance  uint      `json:"balance"`
	AssetsV2 []AssetV2 `json:"assetV2"`
	Votes    []Votes   `json:"votes"`
	Frozen   []Frozen  `json:"frozen"`
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
	FrozenBalance interface{} `json:"frozen_balance,string"`
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

func (c *Contract) UnmarshalJSON(buf []byte) error {
	var contractInternal struct {
		Type      string          `json:"type"`
		Parameter json.RawMessage `json:"parameter"`
	}
	err := json.Unmarshal(buf, &contractInternal)
	if err != nil {
		return err
	}
	switch contractInternal.Type {
	case "TransferContract":
		var transfer TransferContract
		err = json.Unmarshal(contractInternal.Parameter, &transfer)
		c.Parameter = transfer
	case "TransferAssetContract":
		var tokenTransfer TransferAssetContract
		err = json.Unmarshal(contractInternal.Parameter, &tokenTransfer)
		c.Parameter = tokenTransfer
	}
	return err
}
