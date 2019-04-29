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
	BlockTimestamp int64  `json:"blockTimestamp"`
	TxID           string `json:"txID"`
}

type TxReceipt struct {
	Paid    string            `json:paid`
	Meta    Meta              `json:meta`
	Outputs []TxReceiptOutput `json:outputs`
}

type TxReceiptOutput struct {
	Transfers []TxReceiptTransfer `json:transfers`
	Events    []interface{}       `json:events`
}
type TxReceiptTransfer struct {
	Sender    string `json:sender`
	Recipient string `json:recipient`
	Amount    string `json:amount`
}

