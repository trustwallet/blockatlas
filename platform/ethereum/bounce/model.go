package bounce

type Response struct {
	CodeStatus int    `json:"codeStatus"`
	Msg        string `json:"msg"`
}

type Collectible struct {
	ID           int    `json:"ID"`
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
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
}

type CollectionName struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CollectionDescription struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type CollectionImage struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Properties struct {
	Name        CollectionName        `json:"name"`
	Description CollectionDescription `json:"description"`
	Image       CollectionImage       `json:"image"`
}
