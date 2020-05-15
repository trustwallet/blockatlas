package binance

import (
	"fmt"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
)

const (
	TxTransfer              TxType                  = "TRANSFER"       // e.g: BNB, TWT-8C2
	SingleTransferOperation ExplorerTransactionType = "singleTransfer" // e.g: BNB, TWT-8C2
	MultiTransferOperation  ExplorerTransactionType = "multiTransfer"  // e.g [BNB, BNB], [TWT-8C2, TWT-8C2]
)

type (
	TxType                  string
	ExplorerTransactionType string

	Account struct {
		AccountNumber int       `json:"account_number"`
		Address       string    `json:"address"`
		Balances      []Balance `json:"balances"`
		PublicKey     []byte    `json:"public_key"`
		Sequence      uint64    `json:"sequence"`
	}

	Balance struct {
		Free   string `json:"free"`
		Frozen string `json:"frozen"`
		Locked string `json:"locked"`
		Symbol string `json:"symbol"`
	}

	Error struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	}

	NodeInfo struct {
		SyncInfo SyncInfo `json:"sync_info"`
	}

	SyncInfo struct {
		LatestBlockHeight int64 `json:"latest_block_height"`
	}

	Transactions struct {
		Total int  `json:"total"`
		Txs   []Tx `json:"tx"`
	}

	Tx struct {
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

	BlockTransactions struct {
		BlockHeight int64  `json:"blockHeight"`
		Txs         []TxV2 `json:"tx"`
	}

	TxV2 struct {
		Tx
		OrderID         string  `json:"orderId"`         // Optional. Available when the transaction type is NEW_ORDER
		SubTransactions []SubTx `json:"subTransactions"` // Optional. Available when the transaction has sub-transactions, such as multi-send transaction or a transaction have multiple assets
	}

	SubTx struct {
		Asset    string `json:"txAsset"`
		Height   uint64 `json:"blockHeight"`
		Fee      string `json:"txFee"`
		FromAddr string `json:"fromAddr"`
		Hash     string `json:"txHash"`
		ToAddr   string `json:"toAddr"`
		Type     TxType `json:"txType"`
		Value    string `json:"value"`
	}

	TokenList []Token

	Token struct {
		Name           string `json:"name"`
		OriginalSymbol string `json:"original_symbol"`
		Owner          string `json:"owner"`
		Symbol         string `json:"symbol"`
		TotalSupply    string `json:"total_supply"`
	}

	// Transaction response from Explorer
	ExplorerResponse struct {
		Nums int           `json:"txNums"`
		Txs  []ExplorerTxs `json:"txArray"`
	}

	ExplorerTxs struct {
		BlockHeight        uint64          `json:"blockHeight"`
		Code               int             `json:"code"`
		FromAddr           string          `json:"fromAddr"`
		HasChildren        int             `json:"hasChildren"`
		Memo               string          `json:"memo"`
		MultisendTransfers []MultiTransfer `json:"subTxsDto"` // Not part of response, added from hash info tx for simplifying logic
		Timestamp          int64           `json:"timeStamp"`
		ToAddr             string          `json:"toAddr"`
		TxFee              float64         `json:"txFee"`
		TxHash             string          `json:"txHash"`
		TxType             TxType          `json:"txType"`
		Value              float64         `json:"value"`
		TxAsset            string          `json:"txAsset"`
	}

	TxHashRPC struct {
		Hash string   `json:"hash"`
		Tx   TxHashTx `json:"tx"`
	}

	TxHashTx struct {
		Value Value `json:"value"`
	}

	Value struct {
		Msg []Msg `json:"msg"`
	}

	Msg struct {
		Value MsgValue `json:"value"`
	}

	MsgValue struct {
		Inputs  []Input  `json:"inputs"`
		Outputs []Output `json:"outputs"`
	}

	Input struct {
		Address string `json:"address"`
	}

	Output struct {
		Address string `json:"address"`
		Coins   []struct {
			Amount string `json:"amount"`
			Denom  string `json:"denom"`
		} `json:"coins"`
	}

	MultiTransfer struct {
		Amount string `json:"amount"` // Float string ind decimal point
		Asset  string `json:"asset"`
		From   string `json:"from"`
		To     string `json:"to"`
	}
)

func extractMultiTransfers(messages Value) (extracted []MultiTransfer) {
	for _, msg := range messages.Msg {
		var tr MultiTransfer
		tr.From = msg.Value.Inputs[0].Address // Assumed multisend transfer has one input, never seen multiple
		for _, output := range msg.Value.Outputs {
			tr.Amount = output.Coins[0].Amount
			tr.Asset = output.Coins[0].Denom
			tr.To = output.Address

			extracted = append(extracted, tr)
		}
	}
	return
}

// Get explorer transfer fee converted to decimal expression
func (tx *Tx) getFee() string {
	if _, err := strconv.ParseFloat(tx.Fee, 64); err == nil {
		return numbers.DecimalExp(tx.Fee, int(coin.Binance().Decimals))
	}
	return "0"
}

// Converts explorer transfer fee to amount in decimal expression
func (tx *ExplorerTxs) getDexFee() blockatlas.Amount {
	if tx.TxFee > 0 {
		return blockatlas.Amount(numbers.DecimalExp(numbers.Float64toString(tx.TxFee), int(coin.Binance().Decimals)))
	} else {
		return blockatlas.Amount(0)
	}
}

// Get Explorer transfer status based on transfer code
func (tx *ExplorerTxs) getStatus() blockatlas.Status {
	switch tx.Code {
	case 0:
		return blockatlas.StatusCompleted
	default:
		return blockatlas.StatusError
	}
}

func (tx *ExplorerTxs) getDexValue() blockatlas.Amount {
	val := numbers.DecimalExp(numbers.Float64toString(tx.Value), int(coin.Binance().Decimals))
	return blockatlas.Amount(val)
}

// Determines transaction status
func (tx *Tx) getStatus() blockatlas.Status {
	switch tx.Code {
	case 0:
		return blockatlas.StatusCompleted
	default:
		return blockatlas.StatusError
	}
}

// Get explorer transfer error message if transaction failed
func (tx *ExplorerTxs) getError() string {
	switch tx.getStatus() {
	case blockatlas.StatusCompleted:
		return ""
	default:
		return "error"
	}
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

// Determines Explorer transaction direction relatively to address
func (tx *ExplorerTxs) getDirection(address string) blockatlas.Direction {
	if address == "" {
		return ""
	}

	if tx.FromAddr == address && tx.ToAddr == address {
		return blockatlas.DirectionSelf
	}
	if tx.FromAddr == address && tx.ToAddr != address {
		return blockatlas.DirectionOutgoing
	}

	return blockatlas.DirectionIncoming
}

// Determines Explorer transaction type
func (tx *ExplorerTxs) getTransactionType() ExplorerTransactionType {
	var txType ExplorerTransactionType
	if tx.HasChildren == 1 {
		txType = MultiTransferOperation
	} else {
		txType = SingleTransferOperation
	}
	return txType
}
