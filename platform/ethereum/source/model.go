package source

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Block struct {
	Hash         string         `json:"hash"`
	ParentHash   string         `json:"parentHash"`
	Number       hexutil.Uint64 `json:"number"`
	Difficulty   hexutil.Uint64 `json:"difficulty"`
	ReceiptsRoot string         `json:"receiptsRoot"`
	Timestamp    hexutil.Uint64 `json:"timestamp"`
	Transactions []Transaction  `json:"transactions"`
	GasUsed      hexutil.Uint64 `json:"gasUsed"`
	GasLimit     hexutil.Uint64 `json:"gasLimit"`
}

type Transaction struct {
	Hash        string         `json:"hash"`
	Gas         hexutil.Uint64 `json:"gas"`
	GasPrice    hexutil.Uint64 `json:"gasPrice"`
	From        string         `json:"from"`
	To          string         `json:"to"`
	Value       *hexutil.Big   `json:"value"`
	V           *hexutil.Big   `json:"v"`
	R           *hexutil.Big   `json:"r"`
	S           *hexutil.Big   `json:"s"`
	Payload     hexutil.Bytes  `json:"payload"`
}
