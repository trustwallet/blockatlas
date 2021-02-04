package oasis

type Block struct {
	Height    int64  `json:"height"`
	Hash      string `json:"hash"`
	Timestamp int64  `json:"time"`
}

type BlockRequest struct {
	BlockIdentifier int64 `json:"block_identifier"`
}

type Transaction struct {
	Hash     string `json:"tx_hash"`
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   string `json:"amount"`
	Fee      string `json:"fee"`
	Date     int64  `json:"date"`
	Block    uint64 `json:"block"`
	Success  bool   `json:"success"`
	ErrorMsg string `json:"error_message,omitempty"`
	Sequence uint64 `json:"sequence"`
}

type TransactionsByAddressRequest struct {
	Address string `json:"address"`
}
