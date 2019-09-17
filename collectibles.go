package blockatlas

type Collection struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Slug            string `json:"slug"`
	ImageUrl        string `json:"image_url"`
	Description     string `json:"description"`
	ExternalLink    string `json:"external_link"`
	Total           int    `json:"total"`
	CategoryAddress string `json:"category_address"`
	Address         string `json:"address"`
	Version         string `json:"nft_version"`
	Coin            uint   `json:"coin"`
	Type            string `json:"type"`
}

type CollectionPage []Collection

type Collectible struct {
	CollectionID     string `json:"collection_id"`
	TokenID          string `json:"token_id"`
	CategoryContract string `json:"category_contract"`
	// Deprecated: for support old client, ContractAddress eq CollectionID
	ContractAddress string `json:"contract_address"`
	Category        string `json:"category"`
	ImageUrl        string `json:"image_url"`
	ExternalLink    string `json:"external_link"`
	ProviderLink    string `json:"provider_link"`
	Type            string `json:"type"`
	Description     string `json:"description"`
	Coin            uint   `json:"coin"`
	Name            string `json:"name"`
}

type CollectiblePage []Collectible
