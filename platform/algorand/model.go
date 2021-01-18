package algorand

type TransactionType string

const (
	TransactionTypePay TransactionType = "pay"
)

type Account struct {
	Amount                      uint64 `json:"amount"`
	Pendingrewards              uint64 `json:"pendingrewards"`
	Address                     string `json:"address"`
	Round                       uint64 `json:"round"`
	Amountwithoutpendingrewards uint64 `json:"amountwithoutpendingrewards"`
	Rewards                     uint64 `json:"rewards"`
	Status                      string `json:"status"`
}

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type BlockResponse struct {
	Timestamp    uint64        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Type      TransactionType    `json:"tx-type"`
	Hash      string             `json:"id"`
	From      string             `json:"sender"`
	Fee       uint64             `json:"fee"`
	Round     uint64             `json:"confirmed-round"`
	Payment   TransactionPayment `json:"payment-transaction"`
	Timestamp uint64             `json:"round-time"`
}

type TransactionPayment struct {
	Receiver string `json:"receiver"`
	Amount   uint64 `json:"amount"`
}

type Status struct {
	Block string `json:"message"`
}
