package vechain

const (
	filterPrefix = "0x000000000000000000000000%s"
	rangeUnit    = "block"
)

const (
	gasTokenName     = "VeThor"
	gasTokenSymbol   = "VTHO"
	gasTokenAddress  = "0x0000000000000000000000000000456E65726779"
	gasTokenDecimals = 18
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
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Address   string `json:"address,omitempty"`
	Topic0    string `json:"topic0,omitempty"` // Raw transaction hash
	Topic1    string `json:"topic1,omitempty"` // Sender
	Topic2    string `json:"topic2,omitempty"` // Receiver
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
	Meta    LogMeta  `json:"meta"`
}

type TxReceipt struct {
	Reverted bool     `json:"reverted"`
	Paid     string   `json:"paid"`
	Outputs  []Output `json:"outputs"`
}

type Output struct {
	Events []Event `json:"events"`
}

type Event struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type Clause struct {
	To   string `json:"to"`
	Data string `json:"data"`
}

type LogTransfer struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    string  `json:"amount"`
	Meta      LogMeta `json:"meta"`
}

type LogEvent struct {
	Meta LogMeta `json:"meta"`
}

type LogMeta struct {
	TxId           string `json:"txID,omitempty"`
	TxOrigin       string `json:"txOrigin,omitempty"`
	BlockId        string `json:"blockID,omitempty"`
	BlockNumber    uint64 `json:"blockNumber,omitempty"`
	BlockTimestamp int64  `json:"blockTimestamp,omitempty"`
}

type Account struct {
	Balance string `json:"balance"`
}
