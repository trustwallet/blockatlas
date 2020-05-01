package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
	"time"
)

type TxType string
type QuantityTransfer string

const (
	TxTransfer     TxType           = "TRANSFER"
	SingleTransfer QuantityTransfer = "singleTransfer" // e.g: BNB => BNB, TWT-8C2 => TWT-8C2
	MultiTransfer  QuantityTransfer = "multiTransfer"
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
	Asset              string  `json:"txAsset"`
	Code               int     `json:"code"`
	Data               string  `json:"data"`
	Fee                float64 `json:"txFee"`
	FromAddr           string  `json:"fromAddr"`
	HasChildren        int     `json:"hasChildren"`
	Hash               string  `json:"txHash"`
	Memo               string  `json:"memo"`
	OrderID            string  `json:"orderId"`
	Timestamp          int64   `json:"timeStamp"`
	ToAddr             string  `json:"toAddr"`
	Type               TxType  `json:"txType"`
	Value              float64 `json:"value"`
	BlockHeight        uint64  `json:"blockHeight"`
	MultisendTransfers []Msg   `json:"subTxsDto"` // Added from hash info tx
}

type TxHashRPC struct {
	//Code   int      `json:"code"`
	Hash string `json:"hash"`
	//Height string   `json:"height"`
	//Log    string   `json:"log"`
	//Ok     bool     `json:"ok"`
	Tx TxHashTx `json:"tx"`
}

type TxHashTx struct {
	Value Value `json:"value"`
}

type Value struct {
	Messages []Msg `json:"msg"`
}

type Msg struct {
	Value MsgValue `json:"msg"`
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

type Extracted struct {
	Amount string
	Asset  string
	From   string
	To     string
}

func (srcTx *DexTx) extractMultiTransfers(address string) (extracted []Extracted) {
	for _, msg := range srcTx.MultisendTransfers {
		var tr Extracted
		tr.From = msg.Value.Inputs[0].Address // Assumed multisend transfer has one input, never seen multiple
		for _, output := range msg.Value.Outputs {
			if output.Address == address {
				tr.Amount = output.Coins[0].Amount
				tr.Asset = output.Coins[0].Denom
				tr.To = output.Address

				extracted = append(extracted, tr)
			}
			continue
		}
	}
	return
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
	if len(address) == 0 || tx.FromAddr == address || tx.ToAddr == address {
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
func (srcTx *DexTx) Direction(address string) blockatlas.Direction {
	if srcTx.FromAddr == address && srcTx.ToAddr == address {
		return blockatlas.DirectionSelf
	}
	if srcTx.FromAddr == address && srcTx.ToAddr != address {
		return blockatlas.DirectionOutgoing
	}

	return blockatlas.DirectionIncoming
}

func (srcTx *DexTx) QuantityTransferType() QuantityTransfer {
	if srcTx.HasChildren == 1 {
		return MultiTransfer
	} else {
		return SingleTransfer
	}
}
