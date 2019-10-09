package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"math/big"
)

type Page struct {
	Total uint  `json:"total"`
	Docs  []Doc `json:"docs"`
}

type TokenPage struct {
	Total uint    `json:"total"`
	Docs  []Token `json:"docs"`
}

type Token struct {
	Balance  string   `json:"balance"`
	Contract Contract `json:"contract"`
}

type Doc struct {
	Ops         []Op   `json:"operations"`
	Contract    string `json:"contract"`
	ID          string `json:"id"`
	BlockNumber uint64 `json:"blockNumber"`
	Timestamp   int64  `json:"time"`
	Nonce       uint64 `json:"nonce"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
	GasUsed     string `json:"gasUsed"`
	Input       string `json:"input"`
	Error       string `json:"error"`
	Coin        uint   `json:"coin"`
}

type Op struct {
	TxID     string                     `json:"transactionId"`
	Contract *Contract                  `json:"contract"`
	From     string                     `json:"from"`
	To       string                     `json:"to"`
	Type     blockatlas.TransactionType `json:"type"`
	Value    string                     `json:"value"`
	Coin     uint                       `json:"coin"`
}

type Contract struct {
	Address     string `json:"address"`
	Symbol      string `json:"symbol"`
	Decimals    uint   `json:"decimals"`
	TotalSupply string `json:"totalSupply,omitempty"`
	Name        string `json:"name"`
	Contract    string `json:"contract,omitempty"`
}

type NodeInfo struct {
	LatestBlock int64 `json:"latest_block"`
}

type Collection struct {
	Name        string                 `json:"name"`
	ImageUrl    string                 `json:"image_url"`
	Description string                 `json:"description"`
	ExternalUrl string                 `json:"external_url"`
	Slug        string                 `json:"slug"`
	Total       *big.Int               `json:"owned_asset_count"`
	Contracts   []PrimaryAssetContract `json:"primary_asset_contracts"`
}

type PrimaryAssetContract struct {
	Name        string      `json:"name"`
	Address     string      `json:"address"`
	NftVersion  string      `json:"nft_version"`
	Symbol      string      `json:"symbol"`
	Description string      `json:"description"`
	Type        string      `json:"schema_name"`
	Data        DisplayData `json:"display_data"`
	Url         string      `json:"external_link"`
}

type DisplayData struct {
	Images []string `json:"images"`
}

type CollectiblePage struct {
	Collectibles []Collectible `json:"assets"`
}

type Collectible struct {
	TokenId         string              `json:"token_id"`
	AssetContract   CollectibleContract `json:"asset_contract"`
	ImageUrl        string              `json:"image_url"`
	ImagePreviewUrl string              `json:"image_preview_url"`
	Name            string              `json:"name"`
	ExternalLink    string              `json:"external_link"`
	Permalink       string              `json:"permalink"`
	Description     string              `json:"description"`
}

type CollectibleContract struct {
	Address      string `json:"address"`
	Category     string `json:"name"`
	ExternalLink string `json:"external_link"`
	Type         string `json:"schema_name"`
}
