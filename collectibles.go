package blockatlas

type Collection struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	ImageUrl string `json:"image_url"`
	Description string `json:"description"`
	External_Link string `json:"external_link"`
	Total string `json:"total"`
	Category_Address string `json:"category_address"`
	Address string `json:"address"`
	Version string `json:"version"`
	Coin int `json:"coin"`
	Type string `json:"type"`
}

type Collectible struct {
	Token_ID string `json:"token_id"`
	Contract_Address string `json:"contract_address"`
	Category string `json:"category"`
	Image_URL string `json:"image_url"`
	External_Link string `json:"external_link"`
	Type string `json:"type"`
	Description string `json:"type"`
	Coin int `json:"coin"`
}
