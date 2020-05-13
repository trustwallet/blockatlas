package collection

import (
	"math/big"
)

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
	TokenId         string                 `json:"token_id"`
	AssetContract   AssetContract          `json:"asset_contract"`
	ImageUrl        string                 `json:"image_url"`
	ImagePreviewUrl string                 `json:"image_preview_url"`
	Name            string                 `json:"name"`
	ExternalLink    string                 `json:"external_link"`
	Permalink       string                 `json:"permalink"`
	Description     string                 `json:"description"`
	Collection      CollectibleCollections `json:"collection"`
}

type CollectibleCollections struct {
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	ExternalLink string `json:"external_url"`
}

type AssetContract struct {
	Address      string `json:"address"`
	Category     string `json:"name"`
	ExternalLink string `json:"external_link"`
	Type         string `json:"schema_name"`
	Version      string `json:"nft_version"`
}
