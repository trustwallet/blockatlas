package models

const (
	TxBasic = "basic"
	TxOrder = "order"
)

type Tx interface {
	Type() string
}
