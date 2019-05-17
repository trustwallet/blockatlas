package iotex

const TransferFee = "10000000000000000"

type Response struct {
	Total      uint64        `json:"total"`
	ActionInfo []*ActionInfo `json:"actionInfo"`
}

type ActionInfo struct {
	Action    *Action `json:"action"`
	ActHash   string  `json:"actHash"`
	BlkHash   string  `json:"blkHash"`
	BlkHeight string  `json:"blkHeight"`
	Sender    string  `json:"sender"`
	Timestamp string  `json:"timestamp"`
}

type Action struct {
	Core         *ActionCore `json:"core"`
	SenderPubKey []byte      `json:"senderPubKey"`
	Signature    []byte      `json:"signature"`
}

type ActionCore struct {
	Version  uint32    `json:"version"`
	Nonce    string    `json:"nonce"`
	GasLimit string    `json:"gasLimit"`
	GasPrice string    `json:"gasPrice"`
	Transfer *Transfer `json:"transfer"`
}

type Transfer struct {
	Amount    string   `json:"amount"`
	Recipient string   `json:"recipient"`
	Payload   []byte   `json:"payload"`
}
