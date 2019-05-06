package cosmos

// Tx - Base transaction object. Always returned as a list
type Tx struct {
	Block  uint64 `json:"height"`
	Date   string `json:"timestamp"`
	ID     string `json:"txhash"`
	TxData TxData `json:"tx"`
}

// TxData - "tx" sub object
type TxData struct {
	TxType     string     `json:"type"`
	TxContents TxContents `json:"value"`
}

// TxContents - amount, fee, and memo
type TxContents struct {
	TxMessage struct {
		TxParticulars TxParticulars `json:"value"`
	} `json:"msg"`
	TxFee  TxFee  `json:"fee"`
	TxMemo string `json:"memo"`
}

// TxParticulars - from, to, and amount
type TxParticulars struct {
	FromAddr string `json:"from_address"`
	ToAddr   string `json:"to_address"`
	TxAmount Amount `json:"amount"`
}

// TxFee - also references the "amount" struct
type TxFee struct {
	FeeAmount Amount `json:"amount"`
}

// Amount - the asset & quantity
type Amount struct {
	Denom    string `json:"denom"`
	Quantity uint64 `json:"amount"`
}
