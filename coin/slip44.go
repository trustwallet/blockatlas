package coin

type Coin struct {
	Index    uint   `json:"id"`
	Symbol   string `json:"symbol"`
	Title    string `json:"title"`
	Website  string `json:"website"`
	Decimals uint   `json:"decimals"`
}

const (
	ETH  = 60
	XRP  = 144
	XLM  = 148
	NIM  = 242
	AION = 425
	BNB  = 714
	XTZ  = 1729
	KIN  = 2017
)

var Coins = map[uint]Coin {
	XRP: {
		Index:    XRP,
		Symbol:   "XRP",
		Title:    "Ripple",
		Website:  "https://ripple.com",
		Decimals: 6,
	},
	XLM: {
		Index:    XLM,
		Symbol:   "XLM",
		Title:    "Stellar Lumens",
		Website:  "https://www.stellar.org/",
		Decimals: 7,
	},
	NIM: {
		Index:    NIM,
		Symbol:   "NIM",
		Title:    "Nimiq",
		Website:  "https://nimiq.com",
		Decimals: 5,
	},
	BNB: {
		Index:    BNB,
		Symbol:   "BNB",
		Title:    "Binance Coin",
		Website:  "https://binance.org",
		Decimals: 18,
	},
	KIN: {
		Index:    KIN,
		Symbol:   "KIN",
		Title:    "Kin",
		Website:  "https://www.kin.org",
		Decimals: 5,
	},
	XTZ: {
		Index:    XTZ,
		Symbol:   "XTZ",
		Title:    "Tezos",
		Website:  "https://tezos.com",
		Decimals: 6,
	},
	ETH: {
		Index:    ETH,
		Symbol:   "ETH",
		Title:    "Ether",
		Website:  "https://www.ethereum.org",
		Decimals: 18,
	},
	AION: {
		Index:    AION,
		Symbol:   "AION",
		Title:    "Aion",
		Website:  "https://aion.network",
		Decimals: 18,
	},
}
