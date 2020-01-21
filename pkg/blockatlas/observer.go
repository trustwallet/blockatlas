package blockatlas

type Webhook struct {
	Subscriptions     map[string][]string `json:"subscriptions"`
	Webhook           string              `json:"webhook"`
}

type CoinStatus struct {
	Height int64  `json:"height"`
	Error  string `json:"error,omitempty"`
}

type Observer struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
