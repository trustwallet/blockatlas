package binance

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"math"
	"strconv"
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
	HasChildren   int         `json:"hasChildren"`
	SubTxsDto     SubTxsDto   `json:"subTxsDto"`
}

type SubTxsDto struct {
	TotalNum     uint   `json:"totalNum"`
	SubTxDtoList SubTxs `json:"subTxDtoList"`
}

type SubTx struct {
	Hash     string      `json:"hash"`
	Height   uint64      `json:"height"`
	Type     TxType      `json:"type"`
	Value    json.Number `json:"value"`
	Asset    string      `json:"asset"`
	FromAddr string      `json:"fromAddr"`
	ToAddr   string      `json:"toAddr"`
	Fee      json.Number `json:"fee"`
}

type SubTxs []SubTx

func (subTxs *SubTxs) getTxs() (txs []Tx) {
	mapTx := map[string]Tx{}
	for _, subTx := range *subTxs {
		key := subTx.ToAddr + subTx.Asset
		tx, ok := mapTx[key]
		if !ok {
			mapTx[key] = subTx.toTx()
			continue
		}
		txValue, err := tx.Value.Float64()
		if err != nil {
			txValue = 0
		}
		subTxValue, err := subTx.Value.Float64()
		if err != nil {
			subTxValue = 0
		}
		value := numbers.Float64toString(txValue + subTxValue)
		tx.Value = json.Number(value)
		mapTx[key] = tx
	}
	for _, tx := range mapTx {
		txs = append(txs, tx)
	}
	return
}

func (subTx *SubTx) toTx() Tx {
	return Tx{
		Hash:        subTx.Hash,
		BlockHeight: subTx.Height,
		Type:        TxTransfer,
		FromAddr:    subTx.FromAddr,
		ToAddr:      subTx.ToAddr,
		Asset:       subTx.Asset,
		Fee:         subTx.Fee,
		Value:       subTx.Value,
		HasChildren: 0,
	}
}

func (tx *Tx) containAddress(address string) bool {
	if len(address) == 0 {
		return true
	}
	if tx.FromAddr == address {
		return true
	}
	if tx.ToAddr == address {
		return true
	}
	return false
}

func (tx *Tx) getFee() string {
	fee := "0"
	feeNumber, err := tx.Fee.Float64()
	if err == nil && feeNumber > 0 {
		fee = numbers.DecimalExp(string(tx.Fee), 8)
	}
	return fee
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
	Symbol   string      `json:"symbol"`
	Base     string      `json:"-"`
	Quote    string      `json:"-"`
	Quantity interface{} `json:"quantity"`
	Price    interface{} `json:"price"`
}

func (od OrderData) GetVolume() (int64, bool) {
	price, ok := od.GetPrice()
	if !ok {
		return 0, false
	}
	quantity, ok := od.GetQuantity()
	if !ok {
		return 0, false
	}
	return removeFloatPoint(price * quantity), true
}

func (od OrderData) GetPrice() (float64, bool) {
	return convertValue(od.Price)
}

func (od OrderData) GetQuantity() (float64, bool) {
	return convertValue(od.Quantity)
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

func (balance *Balance) isAllZeroBalance() bool {
	balances := [3]string{balance.Frozen, balance.Free, balance.Locked}
	for _, value := range balances {
		value, err := strconv.ParseFloat(value, 64)
		if err != nil || value > 0 {
			return false
		}
	}
	return true

}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func removeFloatPoint(value float64) int64 {
	bnbCoin := coin.Coins[coin.BNB]
	pow := math.Pow(10, float64(bnbCoin.Decimals))
	return int64(value * pow)
}

func convertValue(value interface{}) (float64, bool) {
	result := 0.0
	switch v := value.(type) {
	case float64:
		result = v
	case int:
		result = float64(v)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err == nil {
			result = f
		}
	default:
		return result, false
	}
	return result, true
}
