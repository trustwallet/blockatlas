package coin

type Coin struct {
	Index   int    `json:"id"`
	Symbol  string `json:"symbol"`
	Title   string `json:"title"`
	Website string `json:"website"`
}

const (
	IndexXRP = 144
	IndexXLM = 148
	IndexNIM = 242
	IndexBNB = 714
)

var Coins = map[int]Coin {
	IndexXRP: {
		Index:   IndexXRP,
		Symbol:  "XRP",
		Title:   "Ripple",
		Website: "https://ripple.com",
	},
	IndexXLM: {
		Index:   IndexXLM,
		Symbol:  "XLM",
		Title:   "Stellar Lumens",
		Website: "https://www.stellar.org/",
	},
	IndexNIM: {
		Index:   IndexNIM,
		Symbol:  "NIM",
		Title:   "Nimiq",
		Website: "https://nimiq.com",
	},
	IndexBNB: {
		Index:   IndexBNB,
		Symbol:  "BNB",
		Title:   "Binance Coin",
		Website: "https://binance.org",
	},
}

var (
	XRP = Coins[IndexXRP]
	XLM = Coins[IndexXLM]
	NIM = Coins[IndexNIM]
	BNB = Coins[IndexBNB]
)
