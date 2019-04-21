package models

const (
	TxTransfer            = "transfer"
	TxNativeTokenTransfer = "native_token_transfer"
	TxTokenTransfer       = "token_transfer"
	TxCollectibleTransfer = "collectible_transfer"
	TxTokenSwap           = "token_swap"
	TxContractCall        = "contract_call"
)

const (
	StatusCompleted = "completed"
	StatusPending   = "pending"
	StatusFailed    = "failed"
)

const TxPerPage = 25

type Response []Tx

type Amount string

type Tx struct {
	Id       string      `json:"id"`
	Coin     uint        `json:"coin"`
	From     string      `json:"from"`
	To       string      `json:"to"`
	Fee      Amount      `json:"fee"`
	Date     int64       `json:"date"`
	Type     string      `json:"type"`
	Block    uint64      `json:"block,omitempty"`
	Status   string      `json:"status"`
	Error    string      `json:"error,omitempty"`
	Sequence uint64      `json:"sequence"`
	Meta     interface{} `json:"metadata"`
}

type Transfer struct {
	Value Amount `json:"value"`
}

type NativeTokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
}

type TokenTransfer struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	TokenID  string `json:"token_id"`
	Decimals uint   `json:"decimals"`
	Value    Amount `json:"value"`
	From     string `json:"from"`
	To       string `json:"to"`
}

type CollectibleTransfer struct {
	Name     string `json:"name"`
	Contract string `json:"contract"`
	ImageUrl string `json:"image_url"`
}

type TokenSwap struct {
	Input  TokenTransfer `json:"input"`
	Output TokenTransfer `json:"output"`
}

type ContractCall struct {
	Input string `json:"input"`
	Value string `json:"value"`
}
