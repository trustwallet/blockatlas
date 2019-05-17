package iotex

type Response struct {
	ActionInfo []*ActionInfo `json:"actionInfo"`
}

type ActionInfo struct {
	Action    *Action `json:"action"`
	ActHash   string  `json:"actHash"`
	BlkHeight string  `json:"blkHeight"`
	Sender    string  `json:"sender"`
	GasFee    string  `json:"gasFee"`
	Timestamp string  `json:"timestamp"`
}

type Action struct {
	Core         *ActionCore `json:"core"`
}

type ActionCore struct {
	Nonce    string    `json:"nonce"`
	Transfer *Transfer `json:"transfer"`
}

type Transfer struct {
	Amount    string   `json:"amount"`
	Recipient string   `json:"recipient"`
}
