package solana

type EpochInfo struct {
	AbsoluteSlot uint64 `json:"absoluteSlot"`
	BlockHeight  uint64 `json:"blockHeight"`
	Epoch        uint64 `json:"epoch"`
	SlotIndex    uint64 `json:"slotIndex"`
	SlotsInEpoch uint64 `json:"slotsInEpoch"`
}

type Block struct {
	BlockTime    int64                  `json:"blockTime"`
	Transactions []ConfirmedTransaction `json:"transactions"`
}

type ConfirmedSignature struct {
	Memo      string `json:"memo"`
	Signature string `json:"signature"`
	Slot      uint64 `json:"slot"`
}

type ConfirmedTransaction struct {
	Meta        Meta        `json:"meta"`
	BlockTime   int64       `json:"blockTime,omitempty"`
	Slot        uint64      `json:"slot,omitempty"`
	Transaction Transaction `json:"transaction"`
}

type Meta struct {
	Err interface{} `json:"err"`
	Fee uint64      `json:"fee"`
}

type TransferInfo struct {
	Destination string `json:"destination"`
	Lamports    uint64 `json:"lamports"`
	Source      string `json:"source"`
}

type Parsed struct {
	Info interface{} `json:"info"`
	Type string      `json:"type"`
}

type TokenTransferInfo struct {
	Destination string      `json:"destination"`
	Mint        string      `json:"mint"`
	Source      string      `json:"source"`
	TokenAmount TokenAmount `json:"tokenAmount"`
}

type TokenAmount struct {
	Amount   string `json:"amount"`
	Decimals uint   `json:"decimals"`
}

type Instruction struct {
	Parsed  interface{} `json:"parsed"`
	Program string      `json:"program"`
}

type Message struct {
	Instructions []Instruction `json:"instructions"`
}

type Transaction struct {
	Message    Message  `json:"message"`
	Signatures []string `json:"signatures"`
}
