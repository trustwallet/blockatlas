package models

const (
	TxTransfer = "transfer"
	TxSwap     = "swap"
)

type TransferTx struct {
	Kind      string `json:"kind"`
	Id        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Fee       uint64 `json:"fee"`
	FeeUnit   string `json:"fee_unit"`
	Value     uint64 `json:"value"`
	ValueUnit string `json:"value_unit"`
}
