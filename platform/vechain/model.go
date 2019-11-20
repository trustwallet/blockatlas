package vechain

const (
	filterPrefix = "0x000000000000000000000000%s"
	rangeUnit    = "block"
	blockRange   = 1000000
)

type LogRequest struct {
	Options     Options       `json:"options,omitempty"`
	CriteriaSet []CriteriaSet `json:"criteriaSet,omitempty"`
	Range       Range         `json:"range,omitempty"`
	Order       string        `json:"order,omitempty"`
}

type Range struct {
	Unit string `json:"unit"`
	From int64  `json:"from"`
	To   int64  `json:"to"`
}

type Options struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}

type CriteriaSet struct {
	TxOrigin string `json:"txOrigin,omitempty"`
	Address  string `json:"address,omitempty"`
	Topic0   string `json:"topic0,omitempty"`
	Topic1   string `json:"topic1,omitempty"`
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
	Address        string `json:"address"`
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
