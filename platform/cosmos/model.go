package cosmos

// Tx - Base transaction object. Always returned as part of an array
type Tx struct {
	Block  string `json:"height"`
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
	TxMessage []TxMessage `json:"msg"`
	TxFee     TxFee       `json:"fee"`
	TxMemo    string      `json:"memo"`
}

// TxMessage - an array that holds multiple 'particulars' entries. Possibly used for multiple transfers in one transaction?
type TxMessage struct {
	TxParticulars TxParticulars `json:"value"`
}

// TxParticulars - from, to, and amount
type TxParticulars struct {
	FromAddr string   `json:"from_address"`
	ToAddr   string   `json:"to_address"`
	TxAmount []Amount `json:"amount"`
}

// TxFee - also references the "amount" struct
type TxFee struct {
	FeeAmount []Amount `json:"amount"`
}

// Amount - the asset & quantity. Always seems to be enclosed in an array/list for some reason.
// Perhaps used for multiple tokens transferred in a single sender/reciever transfer?
type Amount struct {
	Denom    string `json:"denom"`
	Quantity string `json:"amount"`
}
