package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"math"
	"strconv"
	"time"
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

//type BlockDescriptor struct {
//	BlockHeight int64  `json:"blockHeight"`
//	BlockHash   string `json:"blockHash"`
//	TxNum       int    `json:"txNum"`
//}
//
//type BlockList struct {
//	BlockArray []BlockDescriptor `json:"blockArray"`
//}

type NodeInfo struct {
	SyncInfo SyncInfo `json:"sync_info"`
}

type SyncInfo struct {
	LatestBlockHeight int64 `json:"latest_block_height"`
}

type TxType string

const (
	TxTransfer TxType = "TRANSFER"
	//TxNewOrder    TxType = "NEW_ORDER"
	//TxCancelOrder TxType = "CANCEL_ORDER"
)

type TxHash struct {
	Hash int `json:"hash"`

	Height uint64 `json:"height"`
	Tx     []Txx  `json:"tx"`
}

type Txx struct {
	Type string `json:"type"`
	V    Value  `json:"value"`
}

type Value struct {
	Memo string `json:"memo"`
	Msgs []Msg  `json:"msg"`
}

type Msg struct {
}

//type SubTxsDto struct {
//	TotalNum     uint   `json:"totalNum"`
//	SubTxDtoList SubTxs `json:"subTxDtoList"`
//}

//type SubTx struct {
//	Hash     string      `json:"hash"`
//	Height   uint64      `json:"height"`
//	Type     TxType      `json:"type"`
//	Value    json.Number `json:"value"`
//	Asset    string      `json:"asset"`
//	FromAddr string      `json:"fromAddr"`
//	ToAddr   string      `json:"toAddr"`
//	Fee      json.Number `json:"fee"`
//}

//type SubTxs []SubTx
//
//func (subTxs *SubTxs) getTxs() (txs []Tx) {
//	mapTx := map[string]Tx{}
//	for _, subTx := range *subTxs {
//		key := subTx.ToAddr + subTx.Asset
//		tx, ok := mapTx[key]
//		if !ok {
//			mapTx[key] = subTx.toTx()
//			continue
//		}
//		txValue, err := tx.Value.Float64()
//		if err != nil {
//			txValue = 0
//		}
//		subTxValue, err := subTx.Value.Float64()
//		if err != nil {
//			subTxValue = 0
//		}
//		value := numbers.Float64toString(txValue + subTxValue)
//		tx.Value = json.Number(value)
//		mapTx[key] = tx
//	}
//	for _, tx := range mapTx {
//		txs = append(txs, tx)
//	}
//	return
//}
//
//func (subTx *SubTx) toTx() Tx {
//	return Tx{
//		Hash:        subTx.Hash,
//		BlockHeight: subTx.Height,
//		Type:        TxTransfer,
//		FromAddr:    subTx.FromAddr,
//		ToAddr:      subTx.ToAddr,
//		Asset:       subTx.Asset,
//		Fee:         subTx.Fee,
//		Value:       subTx.Value,
//		HasChildren: 0,
//	}
//}

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
	return numbers.DecimalExp(tx.Fee, 8)
	//fee := "0"
	//feeNumber, err := t.Fee.Float64()
	//if err == nil && feeNumber > 0 {
	//	fee = numbers.DecimalExp(string(t.Fee), 8)
	//}
	//return fee
}

func (t *Tx) BlockTimestamp() int64 {
	unix := int64(0)
	date, err := time.Parse(time.RFC3339, t.Timestamp)
	if err == nil {
		unix = date.Unix()
	}
	return unix
}

//func (t *Tx) getData() (Data, error) {
//	rawIn := json.RawMessage(t.Data)
//	b, err := rawIn.MarshalJSON()
//	if err != nil {
//		return Data{}, errors.E(err, "getData MarshalJSON", errors.Params{"data": tx.Data})
//	}
//
//	var data Data
//	err = json.Unmarshal(b, &data)
//	if err != nil {
//		return Data{}, errors.E(err, "getData Unmarshal", errors.Params{"data": string(b)})
//	}
//
//	symbols := strings.Split(data.OrderData.Symbol, "_")
//	if len(symbols) < 2 {
//		return data, nil
//	}
//
//	data.OrderData.Base = symbols[0]
//	data.OrderData.Quote = symbols[1]
//	return data, nil
//}

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

//func (od OrderData) GetVolume() (int64, bool) {
//	price, ok := od.GetPrice()
//	if !ok {
//		return 0, false
//	}
//	quantity, ok := od.GetQuantity()
//	if !ok {
//		return 0, false
//	}
//	return removeFloatPoint(price * quantity), true
//}

//func (od OrderData) GetPrice() (float64, bool) {
//	return convertValue(od.Price)
//}

//func (od OrderData) GetQuantity() (float64, bool) {
//	return convertValue(od.Quantity)
//}

type TxPage struct {
	Nums int    `json:"txNums"`
	Txs  []Tx `json:"txArray"`
}

type TransactionsV1 struct {
	Total int    `json:"total"`
	Txs   []Tx `json:"tx"`
}

type Tx struct {
	Asset       string `json:"txAsset"`
	BlockHeight uint64 `json:"blockHeight"`
	Code        int    `json:"code"`
	Data        string `json:"data"`
	Fee         string `json:"txFee"`
	FromAddr    string `json:"fromAddr"`
	Memo        string `json:"memo"`
	OrderID     string `json:"orderId"`
	Sequence    uint64 `json:"sequence"`
	Source      int    `json:"source"`
	Timestamp   string `json:"timeStamp"`
	ToAddr      string `json:"toAddr"`
	TxHash      string `json:"txHash"`
	Type        TxType `json:"txType"`
	Value       string `json:"value"`
}

type TxV2 struct {
	Tx
	OrderID         string  `json:"orderId"`         // Optional. Available when the transaction type is NEW_ORDER
	SubTransactions []SubTx `json:"subTransactions"` // 	Optional. Available when the transaction has sub-transactions, such as multi-send transaction or a transaction have multiple assets
}

type SubTx struct {
	Asset    string `json:"txAsset"`
	Height   uint64 `json:"blockHeight"`
	Fee      string `json:"txFee"`
	FromAddr string `json:"fromAddr"`
	Hash     string `json:"txHash"`
	ToAddr   string `json:"toAddr"`
	Type     TxType `json:"txType"`
	Value    string `json:"value"`
}

type BlockTxV2 struct {
	BlockHeight int64  `json:"blockHeight"`
	Txs         []TxV2 `json:"tx"`
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

// Add test
func (t *Tx) Direction(address string) blockatlas.Direction {
	if t.FromAddr == address && t.ToAddr == address {
		return blockatlas.DirectionSelf
	}
	if t.FromAddr == address && t.ToAddr != address {
		return blockatlas.DirectionOutgoing
	}

	return blockatlas.DirectionIncoming
}

//func convertValue(value interface{}) (float64, bool) {
//	result := 0.0
//	switch v := value.(type) {
//	case float64:
//		result = v
//	case int:
//		result = float64(v)
//	case string:
//		f, err := strconv.ParseFloat(v, 64)
//		if err == nil {
//			result = f
//		}
//	default:
//		return result, false
//	}
//	return result, true
//}
