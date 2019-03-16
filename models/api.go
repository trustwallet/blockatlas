package models

const (
	TxBasic = "basic"
	TxSwap  = "swap"
)

type Balance struct {
	Amount uint64 `json:"amount"`
	Unit   string `json:"unit"`
}

type AccountInfo struct {
	Balances  []Balance `json:"balances"`
	Txs       []Tx      `json:"txs"`
}

type Tx interface {
	Type() string
}

type BasicTx struct {
	Kind      string `json:"kind"`
	Id        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Asset     string `json:"asset"`
	Fee       string `json:"fee"`
	Value     string `json:"value"`
}

func (_ *BasicTx) Type() string {
	return TxBasic
}
