package cosmos

// Tx - Base transaction object. Always returned as part of an array
type Tx struct {
	Block string `json:"height"`
	Date  string `json:"timestamp"`
	ID    string `json:"txhash"`
	Data  Data   `json:"tx"`
}

// Data - "tx" sub object
type Data struct {
	Contents Contents `json:"value"`
}

// Contents - amount, fee, and memo
type Contents struct {
	Message []Message `json:"msg"`
	Fee     Fee       `json:"fee"`
	Memo    string    `json:"memo"`
}

// Message - an array that holds multiple 'particulars' entries. Possibly used for multiple transfers in one transaction?
type Message struct {
	Particulars Particulars `json:"value"`
}

// Particulars - from, to, and amount
type Particulars struct {
	FromAddr string   `json:"from_address"`
	ToAddr   string   `json:"to_address"`
	Amount   []Amount `json:"amount"`
}

// Fee - also references the "amount" struct
type Fee struct {
	FeeAmount []Amount `json:"amount"`
}

// Amount - the asset & quantity. Always seems to be enclosed in an array/list for some reason.
// Perhaps used for multiple tokens transferred in a single sender/reciever transfer?
type Amount struct {
	Denom    string `json:"denom"`
	Quantity string `json:"amount"`
}

// # Staking

type CosmosValidator struct {
	Operator_Address    string `json:"operator_address"`
	Consensus_Pubkey    string `json:"consensus_pubkey"`
}
