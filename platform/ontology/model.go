package ontology

type TxPage struct {
	Result Result `json:"Result"`
}

type Result struct {
	TxnList []Tx `json:"TxnList"`
}

type Transfer struct {
	Amount      string `json:"Amount"`
	FromAddress string `json:"FromAddress"`
	ToAddress   string `json:"ToAddress"`
}

type Tx struct {
	TxnHash     string `json:"TxnHash"`
	ConfirmFlag uint64 `json:"ConfirmFlag"`
	TxnType     uint64 `json:"TxnType"`
	TxnTime     int64  `json:"TxnTime"`
	Height      uint64 `json:"Height"`
	Fee         string `json:"Fee"`
	BlockIndex  uint64 `json:"BlockIndex"`

	TransferList []Transfer `json:"TransferList"`
}

type BlockResults struct {
	Error  int     `json:"Error"`
	Result []Block `json:"Result"`
}

type BlockResult struct {
	Error  int   `json:"Error"`
	Result Block `json:"Result"`
}

type Block struct {
	Height  int    `json:"Height"`
	TxnList []Tx   `json:"TxnList"`
	Hash    string `json:"Hash"`
}

type TxResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Result TxV2   `json:"Result"`
}

type TxV2 struct {
	Hash        string             `json:"tx_hash"`
	Type        int                `json:"tx_type"`
	Time        int64              `json:"tx_time"`
	BlockHeight uint64             `json:"block_height"`
	Fee         string             `json:"fee"`
	Description string             `json:"description"`
	BlockIndex  int                `json:"block_index"`
	ConfirmFlag int                `json:"confirm_flag"`
	EventType   int                `json:"event_type"`
	Details     TransactionDetails `json:"detail"`
}

type TransactionDetails struct {
	Transfers []TransferDetails `json:"transfers"`
}

type TransferDetails struct {
	Amount      string `json:"amount"`
	AssetName   string `json:"asset_name"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
}
