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
	ID          string `json:"id"`
	Type        string `json:"type"`
	PagingToken string `json:"paging_token"`

	Links struct {
		Effects struct {
			Href string `json:"href"`
		} `json:"effects"`
		Transaction struct {
			Href string `json:"href"`
		} `json:"transaction"`
	} `json:"_links"`

	SourceAccount string `json:"source_account"`
	CreatedAt     string `json:"created_at"`

	// create_account and account_merge field
	Account string `json:"account"`

	// create_account fields
	Funder          string `json:"funder"`
	StartingBalance string `json:"starting_balance"`

	// account_merge fields
	Into string `json:"into"`

	// payment/path_payment fields
	From        string `json:"from"`
	To          string `json:"to"`
	AssetType   string `json:"asset_type"`
	AssetCode   string `json:"asset_code"`
	AssetIssuer string `json:"asset_issuer"`
	Amount      string `json:"amount"`

	// transaction fields
	TransactionHash string `json:"transaction_hash"`
	Memo            struct {
		Type  string `json:"memo_type"`
		Value string `json:"memo"`
	}
}
