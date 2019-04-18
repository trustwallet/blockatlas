package vechain

type Tx struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
	Meta      Meta   `json:"meta"`
}

type Meta struct {
	BlockID        string `json:"blockID"`
	BlockNumber    uint64 `json:"blockNumber"`
	BlockTimestamp uint64 `json:"blockTimestamp"`
	TxID           string `json:"txID"`
}

type TxReceipt struct {
	Id      string   `json:"id"`
	Clauses []Clause `json:"clauses"`
	Nonce   string   `json:"nonce"`
	Gas     uint64   `json:"gas"`
 }

 type Clause struct {
	 to    string `json:"to"`
	 value string `json:"value"`
	 data  string `json:"data"`

 }