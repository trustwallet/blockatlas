package binance

import (
	"strconv"
	"strings"
	"time"

	"github.com/trustwallet/golibs/types"

	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
)

const (
	NewOrder    TxType = "NEW_ORDER"
	CancelOrder TxType = "CANCEL_ORDER"
	Transfer    TxType = "TRANSFER"
)

const (
	BNBAsset    = "BNB"
	tokensLimit = "1000"
)

type (
	NodeInfoResponse struct {
		SyncInfo struct {
			LatestBlockHeight int `json:"latest_block_height"`
		} `json:"sync_info"`
	}

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

	TransactionsByAddressAndAssetResponse struct {
		Txs []Tx `json:"tx"`
	}

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

func normalizeBlock(response TransactionsInBlockResponse) types.Block {
	result := types.Block{
		Number: int64(response.BlockHeight),
	}
	result.Txs = normalizeTransactions(response.Tx)
	return result
}

func normalizeTransactions(txs []Tx) types.Txs {
	totalTxs := make(types.Txs, 0, len(txs))
	for _, t := range txs {
		var txs types.Txs
		switch t.TxType {
		case CancelOrder, NewOrder:
			//txs = append(txs, normalizeOrderTransaction(t))
			continue
		case Transfer:
			if len(t.SubTransactions) > 0 {
				txs = normalizeMultiTransferTransaction(t)
			} else {
				txs = append(txs, normalizeTransferTransaction(t))
			}
		}
		totalTxs = append(totalTxs, txs...)
	}
	return totalTxs
}

func normalizeTransferTransaction(t Tx) types.Tx {
	tx := normalizeBaseOfTransaction(t)
	tx.To = t.ToAddr.(string)
	tx.From = t.FromAddr.(string)
	switch {
	case t.TxAsset == BNBAsset:
		tx.Type = types.TxTransfer
		tx.Meta = types.Transfer{
			Value:    normalizeAmount(t.Value),
			Symbol:   coin.Binance().Symbol,
			Decimals: coin.Binance().Decimals,
		}
	case t.TxAsset != "":
		tx.Type = types.TxNativeTokenTransfer
		tx.Meta = types.NativeTokenTransfer{
			Decimals: coin.Binance().Decimals,
			From:     t.FromAddr.(string),
			Symbol:   getTokenSymbolFromID(t.TxAsset),
			Name:     getTokenSymbolFromID(t.TxAsset),
			To:       t.ToAddr.(string),
			TokenID:  t.TxAsset,
			Value:    normalizeAmount(t.Value),
		}
	}
	return tx
}

func normalizeMultiTransferTransaction(t Tx) types.Txs {
	txs := make(types.Txs, 0, len(t.SubTransactions))
	for _, subTx := range t.SubTransactions {
		tx := types.Tx{
			ID:       subTx.TxHash,
			Coin:     coin.Binance().ID,
			From:     subTx.FromAddr,
			To:       subTx.ToAddr,
			Fee:      normalizeFee(subTx.TxFee),
			Date:     t.TimeStamp.Unix(),
			Block:    uint64(t.BlockHeight),
			Status:   types.StatusCompleted,
			Sequence: uint64(t.Sequence),
			Memo:     t.Memo,
		}
		switch {
		case subTx.TxAsset == BNBAsset:
			tx.Type = types.TxTransfer
			tx.Meta = types.Transfer{
				Value:    normalizeAmount(subTx.Value),
				Symbol:   coin.Binance().Symbol,
				Decimals: coin.Binance().Decimals,
			}
		case subTx.TxAsset != "":
			tx.Type = types.TxNativeTokenTransfer
			tx.Meta = types.NativeTokenTransfer{
				Decimals: coin.Binance().Decimals,
				Name:     getTokenSymbolFromID(subTx.TxAsset),
				From:     subTx.FromAddr,
				Symbol:   getTokenSymbolFromID(subTx.TxAsset),
				To:       subTx.ToAddr,
				TokenID:  subTx.TxAsset,
				Value:    normalizeAmount(subTx.Value),
			}
		default:
			continue
		}
		txs = append(txs, tx)
	}
	return txs
}

func normalizeBaseOfTransaction(t Tx) types.Tx {
	return types.Tx{
		ID:       t.TxHash,
		Coin:     coin.Binance().ID,
		From:     t.FromAddr.(string),
		Fee:      normalizeFee(t.TxFee),
		Date:     t.TimeStamp.Unix(),
		Block:    uint64(t.BlockHeight),
		Status:   types.StatusCompleted,
		Sequence: uint64(t.Sequence),
		Memo:     t.Memo,
	}
}

func normalizeTokens(srcBalance []TokenBalance, tokens Tokens) []types.Token {
	assetIds := make([]types.Token, 0)
	for _, srcToken := range srcBalance {
		if token, ok := normalizeToken(srcToken, tokens); ok {
			assetIds = append(assetIds, token)
		}
	}
	return assetIds
}

func normalizeToken(srcToken TokenBalance, tokens Tokens) (types.Token, bool) {
	var result types.Token
	if srcToken.isAllZeroBalance() {
		return result, false
	}

	token, ok := tokens.findTokenBySymbol(srcToken.Symbol)
	if !ok {
		return result, false
	}

	result = types.Token{
		Name:     token.Name,
		Symbol:   token.OriginalSymbol,
		TokenID:  token.Symbol,
		Coin:     coin.Binance().ID,
		Decimals: coin.Binance().Decimals,
		Type:     types.BEP2,
	}

	return result, true
}

func getTokenSymbolFromID(tokenID string) string {
	s := strings.Split(tokenID, "-")
	if len(s) > 1 {
		return s[0]
	}
	return tokenID
}

func normalizeAmount(amount string) types.Amount {
	val := numbers.DecimalExp(amount, int(coin.Binance().Decimals))
	return types.Amount(val)
}

func normalizeFee(amount string) types.Amount {
	a, err := numbers.StringNumberToFloat64(amount)
	if a != 0 && err == nil {
		return types.Amount(numbers.DecimalExp(amount, int(coin.Binance().Decimals)))
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
