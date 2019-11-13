package blockatlas

type Block struct {
	Number int64  `json:"number"`
	ID     string `json:"id,omitempty"`
	Txs    []Tx   `json:"txs"`
}

type Subscription struct {
	Coin    uint   `json:"coin"`
	Address string `json:"address"`
	Webhook string `json:"webhook"`
}
