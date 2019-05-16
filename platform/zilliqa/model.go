package zilliqa

type Tx struct {
	Hash           string `json:"hash"`
	BlockHeight    uint64 `json:"blockHeight"`
	From           string `json:"from"`
	To             string `json:"to"`
	Value          string `json:"value"`
	Fee            string `json:"fee"`
	Timestamp      int64  `json:"timestamp"`
	Signature      string `json:"signature"`
	Nonce          uint64 `json:"nonce"`
	ReceiptSuccess bool   `json:"receiptSuccess"`
}
