package oasis

type Block struct {
	Height int64 `json:"height"`
	Hash   string `json:"hash"`
	Time   int64 `json:"time"`
}

type BlockRequest struct {
	BlockIdentifier int64 `json:"block_identifier"`
}

type Fee struct {
	Amount string `json:"amount"`
	Gas    uint64 `json:"gas"`
}

type Transaction struct {
	Height    int64               `json:"height"`
	Nonce     uint64              `json:"nonce"`
	Timestamp int64               `json:"timestamp"`
	Fee       Fee                 `json:"fee,omitempty"`
	Method    string              `json:"method"`
	Gas       int64               `json:"gas"`
	Success   bool                `json:"success"`
	Metadata  TransactionMetadata `json:"Metadata"`
}

type TransactionMetadata struct {
	TxHash string `json:"tx_hash,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
}

type TransactionsByAddressRequest struct {
	Address string `json:"address"`
}
