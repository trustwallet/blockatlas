package blockatlas

type Block struct {
	Number int64  `json:"number"`
	ID     string `json:"id,omitempty"`
	Txs    []Tx   `json:"txs"`
}
