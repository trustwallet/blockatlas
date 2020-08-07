package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
	"time"
)

const (
	NewOrder    TxType = "NEW_ORDER"
	CancelOrder TxType = "CANCEL_ORDER"
	Transfer    TxType = "TRANSFER"
)

const BNBAsset = "BNB"
const tokensLimit = "1000"

type (
	NodeInfoResponse struct {
		SyncInfo struct {
			LatestBlockHeight int `json:"latest_block_height"`
		} `json:"sync_info"`
	}
)

type (
	TxType string

	TransactionsInBlockResponse struct {
		BlockHeight int  `json:"blockHeight"`
		Tx          []Tx `json:"tx"`
	}

	Tx struct {
		TxHash          string            `json:"txHash"`
		BlockHeight     int               `json:"blockHeight"`
		TxType          TxType            `json:"txType"`
		TimeStamp       time.Time         `json:"timeStamp"`
		FromAddr        interface{}       `json:"fromAddr"`
		ToAddr          interface{}       `json:"toAddr"`
		Value           string            `json:"value"`
		TxAsset         string            `json:"txAsset"`
		TxFee           string            `json:"txFee"`
		OrderID         string            `json:"orderId,omitempty"`
		Code            int               `json:"code"`
		Data            string            `json:"data"`
		Memo            string            `json:"memo"`
		Source          int               `json:"source"`
		SubTransactions []SubTransactions `json:"subTransactions,omitempty"`
		Sequence        int               `json:"sequence"`
	}

	TransactionData struct {
		OrderData struct {
			Symbol      string `json:"symbol"`
			OrderType   string `json:"orderType"`
			Side        string `json:"side"`
			Price       string `json:"price"`
			Quantity    string `json:"quantity"`
			TimeInForce string `json:"timeInForce"`
			OrderID     string `json:"orderId"`
		} `json:"orderData"`
	}

	SubTransactions struct {
		TxHash      string `json:"txHash"`
		BlockHeight int    `json:"blockHeight"`
		TxType      string `json:"txType"`
		FromAddr    string `json:"fromAddr"`
		ToAddr      string `json:"toAddr"`
		TxAsset     string `json:"txAsset"`
		TxFee       string `json:"txFee"`
		Value       string `json:"value"`
	}
)

type (
	TransactionsByAddressAndAssetResponse struct {
		Txs []Tx `json:"tx"`
	}
)

type (
	AccountMeta struct {
		Balances []TokenBalance `json:"balances"`
	}

	TokenBalance struct {
		Free   string `json:"free"`
		Frozen string `json:"frozen"`
		Locked string `json:"locked"`
		Symbol string `json:"symbol"`
	}

	Tokens []Token

	Token struct {
		Name           string `json:"name"`
		OriginalSymbol string `json:"original_symbol"`
		Owner          string `json:"owner"`
		Symbol         string `json:"symbol"`
		TotalSupply    string `json:"total_supply"`
	}
)

func normalizeAmount(amount string) blockatlas.Amount {
	val := numbers.DecimalExp(amount, int(coin.Binance().Decimals))
	return blockatlas.Amount(val)
}

func normalizeFee(amount string) blockatlas.Amount {
	a, err := numbers.StringNumberToFloat64(amount)
	if a != 0 && err == nil {
		return blockatlas.Amount(numbers.DecimalExp(amount, int(coin.Binance().Decimals)))
	} else {
		return "0"
	}
}

func (balance TokenBalance) isAllZeroBalance() bool {
	balances := [3]string{balance.Frozen, balance.Free, balance.Locked}
	for _, value := range balances {
		value, err := strconv.ParseFloat(value, 64)
		if err != nil || value > 0 {
			return false
		}
	}
	return true
}

func (page Tokens) findTokenBySymbol(symbol string) (Token, bool) {
	for _, t := range page {
		if t.Symbol == symbol {
			return t, true
		}
	}
	return Token{}, false
}
