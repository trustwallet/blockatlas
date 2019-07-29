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
