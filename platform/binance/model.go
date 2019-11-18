package binance

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strings"
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

type TxType string

const (
	TxTransfer    TxType = "TRANSFER"
	TxNewOrder    TxType = "NEW_ORDER"
	TxCancelOrder TxType = "CANCEL_ORDER"
)

type Tx struct {
	BlockHeight   uint64      `json:"blockHeight"`
	Type          TxType      `json:"txType"`
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

func (tx *Tx) getData() (Data, error) {
	rawIn := json.RawMessage(tx.Data)
	b, err := rawIn.MarshalJSON()
	if err != nil {
		return Data{}, errors.E(err, "getData MarshalJSON", errors.Params{"data": tx.Data})
	}

	var data Data
	err = json.Unmarshal(b, &data)
	if err != nil {
		return Data{}, errors.E(err, "getData Unmarshal", errors.Params{"data": string(b)})
	}

	symbols := strings.Split(data.OrderData.Symbol, "_")
	if len(symbols) < 2 {
		return data, nil
	}

	data.OrderData.Base = symbols[0]
	data.OrderData.Quote = symbols[1]
	return data, nil
}

type Data struct {
	OrderData OrderData `json:"orderData"`
}

type OrderData struct {
	Symbol string      `json:"symbol"`
	Base   string      `json:"-"`
	Quote  string      `json:"-"`
	Price  interface{} `json:"price"`
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
