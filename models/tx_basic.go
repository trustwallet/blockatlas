package models

type BasicTx struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	From  string `json:"from"`
	To    string `json:"to"`
	Fee   string `json:"fee"`
	Value string `json:"value"`
}

func (_ *BasicTx) Type() string {
	return TxBasic
}
