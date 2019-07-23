package waves

type Transaction struct {
	Hash        string `json:"hash"`
	BlockHeight uint64 `json:"block_height"`
	Timestamp   uint64 `json:"time"`
	TxValue     Tx     `json:"tx"`
}

type Tx struct {
	Sender      string `json:"sender_id"`
	Recipient   string `json:"recipient_id"`
	Amount      uint64 `json:"amount"`
	Fee         uint64 `json:"fee"`
	BlockHeight uint64 `json:"block_height"`
	Type        string `json:"type"`
}
