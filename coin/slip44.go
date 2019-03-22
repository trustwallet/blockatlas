package coin

type Coin struct {
	Index    int    `json:"id"`
	Symbol   string `json:"symbol"`
	Title    string `json:"title"`
	Website  string `json:"website"`
	Decimals uint   `json:"decimals"`
}

const (
	IndexXRP = 144
	IndexXLM = 148
	IndexNIM = 242
	IndexBNB = 714
	IndexKIN = 2017
)

var Coins = map[int]Coin {
	IndexXRP: {
		Index:    IndexXRP,
		Symbol:   "XRP",
		Title:    "Ripple",
		Website:  "https://ripple.com",
		Decimals: 6,
	},
	IndexXLM: {
		Index:    IndexXLM,
		Symbol:   "XLM",
		Title:    "Stellar Lumens",
		Website:  "https://www.stellar.org/",
		Decimals: 7,
	},
	IndexNIM: {
		Index:    IndexNIM,
		Symbol:   "NIM",
		Title:    "Nimiq",
		Website:  "https://nimiq.com",
		Decimals: 5,
	},
	IndexBNB: {
		Index:   IndexBNB,
		Symbol:    "BNB",
		Title:    "Binance Coin",
		Website:  "https://binance.org",
		Decimals: 18,
	},
	IndexKIN: {
		Index:   IndexKIN,
		Symbol:  "KIN",
		Title:   "Kin",
		Website: "https://www.kin.org",
	},
}

var (
	XRP = Coins[IndexXRP]
	XLM = Coins[IndexXLM]
	NIM = Coins[IndexNIM]
	BNB = Coins[IndexBNB]
	KIN = Coins[IndexKIN]
)
