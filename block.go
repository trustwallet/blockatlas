package blockatlas

type Block struct {
	Number int64 `json:"number"`
	ID string    `json:"id"`
	Txs []Tx     `json:"txs"`
}
