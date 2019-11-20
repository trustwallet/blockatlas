package vechain

type LogRequest struct {
	Options     Options       `json:"options"`
	CriteriaSet []CriteriaSet `json:"criteriaSet"`
	Order       string        `json:"order"`
}

type Options struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}
type CriteriaSet struct {
	TxOrigin string `json:"txOrigin"`
}

type Block struct {
	Id           string   `json:"id"`
	Number       int64    `json:"number"`
	Transactions []string `json:"transactions"`
}

type Tx struct {
	Id      string   `json:"id"`
	Origin  string   `json:"origin"`
	Clauses []Clause `json:"clauses"`
	Gas     int      `json:"gas"`
	Nonce   string   `json:"nonce"`
	Meta    TxMeta   `json:"meta"`
}

type Clause struct {
	To    string `json:"to"`
	Value string `json:"value"`
}

type LogTx struct {
	Id        string `json:"id"`
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Amount    string `json:"amount"`
	Meta      TxMeta `json:"meta"`
}

type TxMeta struct {
	TxId           string `json:"txID,omitempty"`
	TxOrigin       string `json:"txOrigin,omitempty"`
	BlockId        string `json:"blockID,omitempty"`
	BlockNumber    uint64 `json:"blockNumber,omitempty"`
	BlockTimestamp int64  `json:"blockTimestamp,omitempty"`
}
