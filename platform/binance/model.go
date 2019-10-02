package binance

import (
	"encoding/json"
	"fmt"
)

type Account struct {
	AccountNumber int       `json:"account_number"`
	Address       string    `json:"address"`
	Balances      []Balance `json:"balances"`
	PublicKey     []byte    `json:"public_key"`
	Sequence      uint64    `json:"sequence"`
}

type Balance struct {
	Free   string `json:"free"`
	Frozen string `json:"frozen"`
	Locked string `json:"locked"`
	Symbol string `json:"symbol"`
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type BlockDescriptor struct {
	BlockHeight int64  `json:"blockHeight"`
	BlockHash   string `json:"blockHash"`
	TxNum       int    `json:"txNum"`
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

type Token struct {
	Mintable       bool   `json:"mintable"`
	Name           string `json:"name"`
	OriginalSymbol string `json:"original_symbol"`
	Owner          string `json:"owner"`
	Symbol         string `json:"symbol"`
	TotalSupply    string `json:"total_supply"`
}

type TokenPage []Token

// findToken find a token into a token list
func (a TokenPage) findToken(symbol string) *Token {
	for _, t := range a {
		if t.Symbol == symbol {
			return &t
		}
	}
	return nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
