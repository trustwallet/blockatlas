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

type TxId struct {
	Id           string   `json:"id"`
	Clauses      []Clause `json:"clauses"`
	Nonce        string   `json:"nonce"`
	Gas          uint64   `json:"gas"`
	GasPriceCoef uint64   `json:"gasPriceCoef"`
}

type Clause struct {
	To    string `json:"to"`
	Value string `json:"value"`
	Data  string `json:"data"`
}

type TxReceipt struct {
	Paid    string            `json:paid`
	Meta    Meta              `json:meta`
	Outputs []TxReceiptOutput `json:outputs`
}

type TxReceiptOutput struct {
	Transfers []TxReceiptTransfer `json:transfers`
}
type TxReceiptTransfer struct {
	Sender    string `json:sender`
	Recipient string `json:recipient`
	Amount    string `json:amount`
}

