package tezos

type RpcBlockHeader struct {
	Level     int64  `json:"level"`
	Timestamp string `json:"timestamp"`
}

type RpcOperationContent struct {
	Hash     string        `json:"hash"`
	Contents []interface{} `json:"contents"`
}

type RpcTransaction struct {
	Kind         string      `json:"kind"`
	Source       string      `json:"source"`
	Fee          string      `json:"fee"`
	Counter      string      `json:"counter"`
	GasLimit     string      `json:"gas_limit"`
	StorageLimit string      `json:"storage_limit"`
	Amount       string      `json:"amount,omitempty"`
	Destination  string      `json:"destination,omitempty"`
	Delegate     string      `json:"delegate,omitempty"`
	Metadata     RpcMetadata `json:"metadata"`
}

type RpcOperationContents []RpcOperationContent

type RpcBlock struct {
	Header     RpcBlockHeader         `json:"header"`
	Operations []RpcOperationContents `json:"operations"`
}

type RpcMetadata struct {
	OperationResult OperationResult `json:"operation_result"`
}

type OperationResult struct {
	Status string `json:"status"`
}
