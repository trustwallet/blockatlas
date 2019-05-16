package ontology

type TxPage struct {
	Result Result `json:"result"`
}

type Result struct {
	TxnList []Tx `json:txn_list`
}

type Transfer struct {
	Amount      string `json:amount`
	FromAddress string `json:from_address`
	ToAddress   string `json:to_address`
}

type Tx struct {
	TxnHash     string `json:txn_hash`
	ConfirmFlag uint64 `json:confirm_flag`
	TxnType     uint64 `json:txn_type`
	TxnTime     int64  `json:txn_time`
	Height      uint64 `json:height`
	Fee         string `json:fee`
	BlockIndex  uint64 `json:block_index`

	TransferList []Transfer `json:transfer_list`
}