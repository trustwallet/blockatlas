package blockatlas

type (
	CollectionV3 struct {
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
		Coin            uint   `json:"coin"`
		// Delete in the future version, as it's now part of Collectible
		Version string `json:"nft_version"`
		Type    string `json:"type"`
	}

	Collection struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		ImageUrl     string `json:"image_url"`
		Description  string `json:"description"`
		ExternalLink string `json:"external_link"`
		Total        int    `json:"total"`
		Address      string `json:"address"`
		Coin         uint   `json:"coin"`
		Type         string `json:"-"`
	}

	CollectionPageV3 []CollectionV3

	CollectionPage []Collection

	Collectible struct {
		ID              string `json:"id"`
		CollectionID    string `json:"collection_id"`
		TokenID         string `json:"token_id"`
		ContractAddress string `json:"contract_address"`
		Category        string `json:"category"`
		ImageUrl        string `json:"image_url"`
		ExternalLink    string `json:"external_link"`
		ProviderLink    string `json:"provider_link"`
		Type            string `json:"type"`
		Description     string `json:"description"`
		Coin            uint   `json:"coin"`
		Name            string `json:"name"`
		Version         string `json:"nft_version"`
	}

	CollectiblePage []Collectible

	CollectibleV3 struct {
		ID               string `json:"id"`
		CollectionID     string `json:"collection_id"`
		TokenID          string `json:"token_id"`
		CategoryContract string `json:"category_contract"`
		ContractAddress  string `json:"contract_address"`
		Category         string `json:"category"`
		ImageUrl         string `json:"image_url"`
		ExternalLink     string `json:"external_link"`
		ProviderLink     string `json:"provider_link"`
		Type             string `json:"type"`
		Description      string `json:"description"`
		Coin             uint   `json:"coin"`
		Name             string `json:"name"`
		Version          string `json:"nft_version"`
	}

	CollectiblePageV3 []CollectibleV3
)
