package bounce

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Collectible struct {
	ContractAddr string `json:"contract_addr"`
	ContractName string `json:"contract_name,omitempty"`
	TokenID      string `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	TokenURI     string `json:"token_uri"`
}

type CollectibleList struct {
	Collectibles []Collectible `json:"tokens"`
}

type CollectibleResponse struct {
	Response
	Data CollectibleList `json:"data"`
}

type Collection struct {
	ContractAddr string `json:"contract_addr"`
	TokenType    string `json:"token_type"`
	TokenID      string `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	Balance      string `json:"balance"`
	TokenURI     string `json:"token_uri"`
}

type CollectionList struct {
	Collections []Collection `json:"nfts721"`
}

type CollectionResponse struct {
	Response
	Data CollectionList `json:"data"`
}

type CollectionInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
