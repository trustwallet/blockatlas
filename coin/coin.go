package coin

//go:generate rm -f slip44.go
//go:generate go run gen.go

type Coin struct {
	Index    uint   `json:"index"`
	Symbol   string `json:"symbol"`
	Title    string `json:"name"`
	Website  string `json:"link"`
	Decimals uint   `json:"decimals"`
}
