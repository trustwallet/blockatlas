package blockatlas

type Resolved struct {
	Result string `json:"result"`
	Error  string `json:"error,omitempty"`
	Coin   uint64 `json:"coin"`
}
