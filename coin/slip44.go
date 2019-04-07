package coin

type Coin struct {
	Index    uint   `json:"id"`
	Symbol   string `json:"symbol"`
	Title    string `json:"title"`
	Website  string `json:"website"`
	Decimals uint   `json:"decimals"`
}

const (
	IndexETH = 60
	IndexXRP = 144
	IndexXLM = 148
	IndexNIM = 242
	IndexBNB = 714
	IndexXTZ = 1729
	IndexKIN = 2017
)

var Coins = map[uint]Coin {
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
		Index:    IndexBNB,
		Symbol:   "BNB",
		Title:    "Binance Coin",
		Website:  "https://binance.org",
		Decimals: 18,
	},
	IndexKIN: {
		Index:    IndexKIN,
		Symbol:   "KIN",
		Title:    "Kin",
		Website:  "https://www.kin.org",
		Decimals: 5,
	},
	IndexXTZ: {
		Index:    IndexXTZ,
		Symbol:   "XTZ",
		Title:    "Tezos",
		Website:  "https://tezos.com",
		Decimals: 6,
	},
	IndexETH: {
		Index:    IndexETH,
		Symbol:   "ETH",
		Title:    "Ether",
		Website:  "https://www.ethereum.org",
		Decimals: 18,
	},
}

var (
	ETH = Coins[IndexETH]
	XRP = Coins[IndexXRP]
	XLM = Coins[IndexXLM]
	NIM = Coins[IndexNIM]
	BNB = Coins[IndexBNB]
	KIN = Coins[IndexKIN]
	XTZ = Coins[IndexXTZ]
)
