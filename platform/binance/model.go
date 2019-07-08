package binance

import (
	"encoding/json"
	"fmt"
)

// Binance cahin transfer types
const (
	TRANSFER  = "TRANSFER"
	NEW_ORDER = "NEW_ORDER"
)

type Account struct {
	AccountNumber int       `json:"account_number"`
	Address       string    `json:"address"`
	Balances      []Balance `json:"balances"`
	PublicKey     []byte    `json:"public_key"`
	Sequence      uint64    `json:"sequence"`
}

type Balance struct {
	Symbol string `json:"symbol"`
	Free   uint64 `json:"free"`
	Locked uint64 `json:"locked"`
	Frozen uint64 `json:"frozen"`
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type BlockDescriptor struct {
	BlockHeight int64 `json:"blockHeight"`
	BlockHash string `json:"blockHash"`
	TxNum int `json:"txNum"`
}

type BlockList struct {
	BlockArray []BlockDescriptor `json:"blockArray"`
}

type Tx struct {
	BlockHeight   uint64      `json:"blockHeight"`
	Type          string      `json:"txType"`
	Code          int         `json:"code"`
	ConfirmBlocks int         `json:"confirmBlocks"`
	Data          string      `json:"data"`
	FromAddr      string      `json:"fromAddr"`
	OrderID       string      `json:"orderId"`
	Timestamp     int64       `json:"timeStamp"`
	ToAddr        string      `json:"toAddr"`
	Age           int64       `json:"txAge"`
	MappedAsset   string      `json:"mappedTxAsset"`
	Asset         string      `json:"txAsset"`
	Fee           json.Number `json:"txFee"`
	Hash          string      `json:"txHash"`
	Value         json.Number `json:"value"`
	Memo          string      `json:"memo"`
}

type TxPage struct {
	Nums int  `json:"txNums"`
	Txs  []Tx `json:"txArray"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Transaction receipt structure
type Receipt struct {
	Hash        string    `json:"hash"`
	TxReceipts  TxReceipt `json:"tx"`
}

type TxReceipt struct {
	Value Value `json:"value"`
}

type Value struct {
	Msg []Msg `json:"msg"`
}

type Msg struct {
	MsgValue MsgValue `json:"value"`
}

type MsgValue struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}

type Input struct {
	Address string `json:"address"`
	Coins   []Coin `json:"coins"`
}

type Output struct {
	Address string `json:"address"`
	Coins   []Coin `json:"coins"`
}

type Coin struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}