package algorand

type TransactionType string

const (
	TransactionTypePay TransactionType = "pay"
)

type TransactionsResponse struct {
	Transactions []Transaction `json:"transactions"`
}

type BlockResponse struct {
	Timestamp    uint64            `json:"timestamp"`
	Transactions BlockTransactions `json:"txns"`
}

type BlockTransactions struct {
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Type      TransactionType    `json:"type"`
	Hash      string             `json:"tx"`
	From      string             `json:"from"`
	Fee       uint64             `json:"fee"`
	Round     uint64             `json:"round"`
	Payment   TransactionPayment `json:"payment"`
	Timestamp uint64             `json:"timestamp,omitempty"`
}

type TransactionPayment struct {
	To     string `json:"to"`
	Amount uint64 `json:"amount"`
}

type Status struct {
	LastRound int64 `json:"lastRound"`
}
