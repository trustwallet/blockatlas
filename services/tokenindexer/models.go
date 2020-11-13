package tokenindexer

type Request struct {
	From int64
	Coin int
}

type Response struct {
	Assets []Asset `json:"assets"`
}

type Asset struct {
	Asset  string `json:"asset"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Type   string `json:"type"`

	Decimals uint `json:"decimals"`
}
