package tezos

type Op struct {
	Txs []Tx `json:"ops"`
}

// Tx is a Tezos blockchain transaction
type Tx struct {
	Hash      string  `json:"hash"`
	BlockHash string  `json:"block"`
	Status    string  `json:"status"`
	Success   bool    `json:"is_success"`
	Time      string  `json:"time"`
	Height    uint64  `json:"height"`
	Type      string  `json:"type"`
	Sender    string  `json:"sender"`
	Volume    float64 `json:"volume"`
	Receiver  string  `json:"receiver"`
	Fee       int     `json:"gas_used"`
}

type Validator struct {
	Address string `json:"pkh"`
}

type Head struct {
	Height int64 `json:"height"`
}

type DelegateOptions struct {
	Setable bool `json:"setable"`
}

type Account struct {
	Name          Address         `json:"name"`
	Manager       Address         `json:"address"`
	Delegate      DelegateOptions `json:"delegate"`
	Balance       string          `json:"balance"`
	Spendable     bool            `json:"spendable"`
	Counter       string          `json:"counter"`
	NodeTimestamp string          `json:"node_timestamp"`
}
