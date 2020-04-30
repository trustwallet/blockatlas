package binance

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
	"time"
)

type TxType string

const (
	TxTransfer TxType = "TRANSFER"
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

type NodeInfo struct {
	SyncInfo SyncInfo `json:"sync_info"`
}

type SyncInfo struct {
	LatestBlockHeight int64 `json:"latest_block_height"`
}

type Transactions struct {
	Total int  `json:"total"`
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

type BlockTransactions struct {
	BlockHeight int64  `json:"blockHeight"`
	Txs         []TxV2 `json:"tx"`
}

type TxV2 struct {
	Tx
	OrderID         string  `json:"orderId"`         // Optional. Available when the transaction type is NEW_ORDER
	SubTransactions []SubTx `json:"subTransactions"` // Optional. Available when the transaction has sub-transactions, such as multi-send transaction or a transaction have multiple assets
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

type TokenList []Token

type Token struct {
	Name           string `json:"name"`
	OriginalSymbol string `json:"original_symbol"`
	Owner          string `json:"owner"`
	Symbol         string `json:"symbol"`
	TotalSupply    string `json:"total_supply"`
}

// From Explorer
type DexTxPage struct {
	Nums int     `json:"txNums"`
	Txs  []DexTx `json:"txArray"`
}

type DexTx struct {
	//Age           int64       `json:"txAge"`
	Asset         string      `json:"txAsset"`
	//Code          int         `json:"code"`
	//ConfirmBlocks int         `json:"confirmBlocks"`
	Data          string      `json:"data"`
	Fee           json.Number `json:"txFee"`
	FromAddr      string      `json:"fromAddr"`
	HasChildren   int         `json:"hasChildren"`
	Hash          string      `json:"txHash"`
	Memo          string      `json:"memo"`
	OrderID       string      `json:"orderId"`
	Timestamp     int64       `json:"timeStamp"`
	ToAddr        string      `json:"toAddr"`
	Type          TxType      `json:"txType"`
	Value         json.Number `json:"value"`
	BlockHeight   uint64      `json:"blockHeight"`
	//SubTxsDto     SubTxsDto   `json:"subTxsDto"`
}

//type SubTxsDto struct {
//	TotalNum     uint   `json:"totalNum"`
//	SubTxDtoList SubTxs `json:"subTxDtoList"`
//}
//
//type SubTxs []SubTrx
//
//type SubTrx struct {
//	Hash     string      `json:"hash"`
//	Height   uint64      `json:"height"`
//	Type     TxType      `json:"type"`
//	Value    json.Number `json:"value"`
//	Asset    string      `json:"asset"`
//	FromAddr string      `json:"fromAddr"`
//	ToAddr   string      `json:"toAddr"`
//	Fee      json.Number `json:"fee"`
//}

type TxHash struct {
	Code   int      `json:"code"`
	Hash   string   `json:"hash"`
	Height string   `json:"height"`
	Log    string   `json:"log"`
	Ok     bool     `json:"ok"`
	Tx     TxHashTx `json:"tx"`
}

type TxHashTx struct {
	V Value `json:"value"`
}

type Value struct {
	Messages []Msg `json:"msg"`
}

type Msg struct {
	MV MsgValue `json:"msg"`
}

type MsgValue struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}

type Input struct {
	Address string `json:"address"`
}

type Output struct {
	Address string `json:"address"`
	Coins   []struct {
		Amount string `json:"amount"`
		Denom  string `json:"denom"`
	} `json:"coins"`
}

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

func (tx *Tx) getFee() string {
	fee := "0"
	if _, err := strconv.ParseFloat(tx.Fee, 64); err == nil {
		fee = numbers.DecimalExp(tx.Fee, 8)
	}
	return fee
}

func (tx *DexTx) getDexFee() string {
	fee := "0"
	feeNumber, err := tx.Fee.Float64()
	if err == nil && feeNumber > 0 {
		fee = numbers.DecimalExp(string(tx.Fee), 8)
	}
	return fee
}

func (tx *Tx) getStatus() blockatlas.Status {
	switch tx.Code {
	case 0:
		return blockatlas.StatusCompleted
	default:
		return blockatlas.StatusError
	}
}

func (tx *Tx) getError() string {
	switch tx.getStatus() {
	case blockatlas.StatusCompleted:
		return ""
	default:
		return "error"
	}
}

func (tx *Tx) blockTimestamp() int64 {
	unix := int64(0)
	date, err := time.Parse(time.RFC3339, tx.Timestamp)
	if err == nil {
		unix = date.Unix()
	}
	return unix
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

// findToken find a token into a token list
func (page TokenList) findToken(symbol string) *Token {
	for _, t := range page {
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

// Add test
func (tx *Tx) Direction(address string) blockatlas.Direction {
	if tx.FromAddr == address && tx.ToAddr == address {
		return blockatlas.DirectionSelf
	}
	if tx.FromAddr == address && tx.ToAddr != address {
		return blockatlas.DirectionOutgoing
	}

	return blockatlas.DirectionIncoming
}
