package nano

const (
	BlockTypeSend    BlockType = "send"
	BlockTypeReceive BlockType = "receive"
)

type BlockType string

type AccountHistoryRequest struct {
	Action  string `json:"action"`
	Account string `json:"account"`
	Count   string `json:"count"`
}

type AccountHistory struct {
	Account string `json:"account"`
	History interface{} // NANO RPC returns string for address with 0 transactions
}

type Transaction struct {
	Type           BlockType `json:"type"`
	Account        string    `json:"account"`
	Amount         string    `json:"amount"`
	LocalTimestamp string    `json:"local_timestamp"`
	Height         string    `json:"height"`
	Hash           string    `json:"hash"`
}
