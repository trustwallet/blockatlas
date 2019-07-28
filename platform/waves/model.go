package waves

type Transaction struct {
	Id         string `json:"id"`
	Sender     string `json:"sender"`
	AssetId    string `json:"assetId"`
	Recipient  string `json:"recipient"`
	Amount     uint64 `json:"amount"`
	FeeAssetId string `json:"feeAssetId"`
	Fee        uint64 `json:"fee"`
	Timestamp  uint64 `json:"timestamp"`
	Attachment string `json:"attachment"`
	Block      uint64 `json:"height"`
	Type       uint64 `json:"type"`
}

type CurrentBlock struct {
	Height int64 `json:"height"`
}

type Block struct {
	Transactions []Transaction `json:"transactions"`
}
