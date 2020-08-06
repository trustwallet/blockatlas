package binance

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.FetchLatestBlockNumber()
	if err != nil {
		return 0, err
	}

	return block, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	transactionInBlockResponse, err := p.client.FetchTransactionsInBlock(num)
	if err != nil {
		return nil, err
	}
	block := normalizeBlock(transactionInBlockResponse)
	return &block, nil
}

func normalizeBlock(response TransactionsInBlockResponse) blockatlas.Block {
	result := blockatlas.Block{
		Number: int64(response.BlockHeight),
	}
	totalTxs := make([]blockatlas.Tx, 0, len(response.Tx))
	for _, t := range response.Tx {
		var txs []blockatlas.Tx
		switch t.TxType {
		case CancelOrder, NewOrder:
			txs = append(txs, normalizeOrderTransactionForBlock(t))
		case Transfer:
			if len(t.SubTransactions) > 0 {
				txs = normalizeMultiTransferTransactionForBlock(t)
				continue
			}
			txs = append(txs, normalizeTransferTransactionForBlock(t))
		}
		totalTxs = append(totalTxs, txs...)
	}
	result.Txs = totalTxs
	return result
}

func normalizeTransferTransactionForBlock(t BlockTx) blockatlas.Tx {
	tx := getBaseTxBodyForBlock(t)
	tx.To = t.ToAddr.(string)
	switch {
	case t.TxAsset == BNBAsset:
		tx.Type = blockatlas.TxTransfer
		tx.Meta = blockatlas.Transfer{
			Value:    normalizeAmount(t.Value),
			Symbol:   coin.Binance().Symbol,
			Decimals: coin.Binance().Decimals,
		}
	case t.TxAsset != "":
		tx.Type = blockatlas.TxNativeTokenTransfer
		tx.Meta = blockatlas.NativeTokenTransfer{
			Decimals: coin.Binance().Decimals,
			From:     t.FromAddr,
			Symbol:   tokenSymbol(t.TxAsset),
			To:       t.ToAddr.(string),
			TokenID:  t.TxAsset,
			Value:    normalizeAmount(t.Value),
		}
	}
	return tx
}

func normalizeMultiTransferTransactionForBlock(t BlockTx) []blockatlas.Tx {
	txs := make([]blockatlas.Tx, 0, len(t.SubTransactions))
	for _, subTx := range t.SubTransactions {
		tx := blockatlas.Tx{
			ID:       subTx.TxHash,
			Coin:     coin.Binance().ID,
			From:     subTx.FromAddr,
			To:       subTx.ToAddr,
			Fee:      normalizeFee(subTx.TxFee),
			Date:     t.TimeStamp.Unix(),
			Block:    uint64(t.BlockHeight),
			Status:   blockatlas.StatusCompleted,
			Sequence: uint64(t.Sequence),
			Memo:     t.Memo,
		}
		switch {
		case subTx.TxAsset == BNBAsset:
			tx.Type = blockatlas.TxTransfer
			tx.Meta = blockatlas.Transfer{
				Value:    blockatlas.Amount(subTx.Value),
				Symbol:   coin.Binance().Symbol,
				Decimals: coin.Binance().Decimals,
			}
		case subTx.TxAsset != "":
			tx.Type = blockatlas.TxNativeTokenTransfer
			tx.Meta = blockatlas.NativeTokenTransfer{
				Decimals: coin.Binance().Decimals,
				From:     subTx.FromAddr,
				Symbol:   tokenSymbol(subTx.TxAsset),
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

func getBaseTxBodyForBlock(t BlockTx) blockatlas.Tx {
	return blockatlas.Tx{
		ID:       t.TxHash,
		Coin:     coin.Binance().ID,
		From:     t.FromAddr,
		Fee:      normalizeFee(t.TxFee),
		Date:     t.TimeStamp.Unix(),
		Block:    uint64(t.BlockHeight),
		Status:   blockatlas.StatusCompleted,
		Sequence: uint64(t.Sequence),
		Memo:     t.Memo,
	}
}

func normalizeOrderTransactionForBlock(t BlockTx) blockatlas.Tx {
	tx := getBaseTxBodyForBlock(t)

	tx.Type = blockatlas.TxAnyAction
	meta := blockatlas.AnyAction{
		Coin:     coin.Binance().ID,
		Decimals: coin.Binance().Decimals,
	}

	data, err := parseOrderData(t.Data)
	if err == nil {
		base, _ := parseOrderDataSymbol(data.OrderData.Symbol)
		meta.TokenID = base
		meta.Value = blockatlas.Amount(data.OrderData.Quantity)
		meta.Name = data.OrderData.Side
		meta.Symbol = tokenSymbol(base)
	}
	switch t.TxType {
	case CancelOrder:
		meta.Title = blockatlas.KeyTitleCancelOrder
		meta.Key = blockatlas.KeyCancelOrder
	case NewOrder:
		meta.Title = blockatlas.KeyTitlePlaceOrder
		meta.Key = blockatlas.KeyPlaceOrder
	}

	tx.Meta = meta
	tx.Direction = blockatlas.DirectionOutgoing
	return tx
}

func parseOrderData(rawOrderData string) (TransactionData, error) {
	var result TransactionData
	err := json.Unmarshal([]byte(rawOrderData), &result)
	return result, err
}

func parseOrderDataSymbol(symbol string) (string, string) {
	result := strings.Split(symbol, "_")
	if len(result) == 0 {
		return symbol, symbol
	}
	return result[0], result[1]
}

func tokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}
