package bounce

type Response struct {
	CodeStatus int    `json:"codeStatus"`
	Msg        string `json:"msg"`
}

type Collectible struct {
	ContractAddr string `json:"contract_addr"`
	TokenID      int    `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	ChainID      int    `json:"chain_id"`
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
	TokenID      int    `json:"token_id"`
	OwnerAddr    string `json:"owner_addr"`
	ChainID      int    `json:"chain_id"`
	Balance      string `json:"balance"`
	TokenURI     string `json:"token_uri"`
}

type CollectionList struct {
	Collections []Collection `json:"nfts"`
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
