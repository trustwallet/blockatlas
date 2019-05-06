package icon

// Response describes the ripple transaction response
type Response struct {
	Data        []Tx   `json:"data"`
	ListSize    uint64 `json:"listSize"`
	TotalSize   uint64 `json:"totalSize"`
	Result      string `json:"result"`
	Description string `json:"description"`
}

// Tx describes the ripple transaction
type Tx struct {
	TxHash     string `json:"txHash"`
	Height     uint64 `json:"height"`
	CreateDate string `json:"createDate"`
	FromAddr   string `json:"fromAddr"`
	ToAddr     string `json:"toAddr"`
	TxType     string `json:"txType"`
	DataType   string `json:"dataType"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	State      uint64 `json:"status"`
}
