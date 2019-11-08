package nano

type AccountHistoryRequest struct {
	Action  string `json:"action"`
	Account string `json:"account"`
	Count   string `json:"count"`
	Raw     bool   `json:"raw,omitempty"`
}

type AccountHistory struct {
	Account string        `json:"account"`
	History []Transaction `json:"history"`
}

type Transaction struct {
	Type           string `json:"type"`
	Account        string `json:"account"`
	Amount         string `json:"amount"`
	LocalTimestamp string `json:"local_timestamp"`
	Height         string `json:"height"`
	Hash           string `json:"hash"`
}
