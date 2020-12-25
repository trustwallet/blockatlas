package explorer

type Response struct {
	TotalCount int       `json:"totalCount"`
	Messages   []Message `json:"messages"`
}

type Receipt struct {
	ExitCode int `json:"exitCode"`
}

type Message struct {
	Cid       string  `json:"cid"`
	Height    uint64  `json:"height"`
	Timestamp int64   `json:"timestamp"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Nonce     uint64  `json:"nonce"`
	Value     string  `json:"value"`
	Method    string  `json:"method"`
	Receipt   Receipt `json:"receipt"`
}
