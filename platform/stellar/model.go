package stellar

// PaymentsPage of payments returned by Horizon
type PaymentsPage struct {
	Embedded struct {
		Records []Payment
	} `json:"_embedded"`
}

type LedgersPage struct {
	Embedded struct {
		Records []Ledger
	} `json:"_embedded"`
}

type Ledger struct {
	Sequence int64  `json:"sequence"`
	Id       string `json:"id"`
}

type Block struct {
	Ledger   Ledger
	Payments []Payment
}

// Payment model returned by Horizon
type Payment struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	SourceAccount   string `json:"source_account"`
	CreatedAt       string `json:"created_at"`
	Account         string `json:"account"`
	Funder          string `json:"funder"`
	StartingBalance string `json:"starting_balance"`
	Into            string `json:"into"`
	From            string `json:"from"`
	To              string `json:"to"`
	AssetType       string `json:"asset_type"`
	Amount          string `json:"amount"`
	TransactionHash string `json:"transaction_hash"`
}
