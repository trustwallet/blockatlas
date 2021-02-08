package near

type Block struct {
	Header BlockHeader `json:"header"`
	Chunks []Chunk     `json:"chunks"`
}

type BlockHeader struct {
	Height    uint64 `json:"height"`
	Timestamp uint64 `json:"timestamp"`
}

type Chunk struct {
	Hash      string `json:"chunk_hash"`
	Height    uint64 `json:"height_created"`
	Timestamp uint64
}

type ChunkDetail struct {
	Header       Chunk `json:"header"`
	Transactions []Tx  `json:"transactions,omitempty"`
}

type Tx struct {
	SignerID   string        `json:"signer_id"`
	Nonce      int           `json:"nonce"`
	ReceiverID string        `json:"receiver_id"`
	Actions    []interface{} `json:"actions"`
	Hash       string        `json:"hash"`
}

type TransferAction struct {
	Transfer Transfer `json:"Transfer"`
}

type Transfer struct {
	Deposit string `json:"deposit"`
}

type FunctionCall struct {
	MethodName string `json:"method_name"`
	Args       string `json:"args"`
	Gas        int64  `json:"gas"`
	Deposit    string `json:"deposit"`
}
