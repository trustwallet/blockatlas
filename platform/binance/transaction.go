package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strings"
	"sync"

	"github.com/trustwallet/blockatlas/coin"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// Endpoint supports queries without token query parameter
	return p.GetTokenTxsByAddress(address, p.Coin().Symbol)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	srcTxs, err := p.client.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}
	txs, err := p.getTxChildChan(srcTxs.Txs)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(txs, token, address), nil
}

// getTxChildChan get all child assets from a tx
func (p *Platform) getTxChildChan(srcTxs []Tx) ([]Tx, error) {
	txs := make([]Tx, 0)
	var wg sync.WaitGroup
	out := make(chan Tx, len(srcTxs))
	for _, srcTx := range srcTxs {
		if srcTx.HasChildren != 1 {
			// Return the same transaction if doesn't have a child
			txs = append(txs, srcTx)
			continue
		}
		wg.Add(1)
		go func(srcTx Tx, out chan Tx, wg *sync.WaitGroup) {
			defer wg.Done()
			tx, err := p.client.GetTx(srcTx.Hash)
			if err != nil {
				// Return the same transaction if an error occurs
				out <- srcTx
				logger.Error("GetTransactionsByBlockChan", err, logger.Params{"hash": srcTx.Hash})
				return
			}
			out <- tx
		}(srcTx, out, &wg)
	}
	wg.Wait()
	close(out)
	for r := range out {
		txs = append(txs, r)
	}
	return txs, nil
}

func normalizeTransfer(tx blockatlas.Tx, srcTx Tx, token, address string) (blockatlas.TxPage, bool) {
	// Verify if the tx has more them one asset
	if srcTx.HasChildren == 1 {
		txs := make(blockatlas.TxPage, 0)
		// Parse all assets as a transaction
		for _, subTx := range srcTx.SubTxsDto.SubTxDtoList.getTxs() {
			// If this is not called from a block observer_test, only get the user txs/assets
			if !subTx.containAddress(address) {
				continue
			}
			// Recursive call to normalize the tx
			newTxs, ok := normalizeTransfer(tx, subTx, token, address)
			if !ok {
				continue
			}
			txs = append(txs, newTxs...)
		}
		if len(txs) == 0 {
			return txs, false
		}
		return txs, true
	}

	// Verify if this is the same asset we are looking for
	if len(token) > 0 && srcTx.Asset != token {
		return blockatlas.TxPage{tx}, false
	}

	tx.From = srcTx.FromAddr
	tx.To = srcTx.ToAddr

	bnbCoin := coin.Coins[coin.BNB]
	value := numbers.DecimalExp(string(srcTx.Value), 8)
	if srcTx.Asset == bnbCoin.Symbol {
		// Condition for native transfer (BNB)
		tx.Meta = blockatlas.Transfer{
			Value:    blockatlas.Amount(value),
			Symbol:   bnbCoin.Symbol,
			Decimals: bnbCoin.Decimals,
		}
		return blockatlas.TxPage{tx}, true
	}

	// Condition for native token transfer
	tx.Meta = blockatlas.NativeTokenTransfer{
		TokenID:  srcTx.Asset,
		Symbol:   TokenSymbol(srcTx.Asset),
		Value:    blockatlas.Amount(value),
		Decimals: bnbCoin.Decimals,
		From:     srcTx.FromAddr,
		To:       srcTx.ToAddr,
	}
	return blockatlas.TxPage{tx}, true
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx Tx, token, address string) (blockatlas.TxPage, bool) {
	tx := blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.BNB,
		From:   srcTx.FromAddr,
		To:     srcTx.ToAddr,
		Fee:    blockatlas.Amount(srcTx.getFee()),
		Date:   srcTx.Timestamp / 1000,
		Block:  srcTx.BlockHeight,
		Status: blockatlas.StatusCompleted,
		Memo:   srcTx.Memo,
	}

	switch srcTx.Type {
	case TxTransfer:
		return normalizeTransfer(tx, srcTx, token, address)
	}
	//case TxCancelOrder, TxNewOrder:
	//	return tx, false
	//	dt, err := srcTx.getData()
	//	if err != nil {
	//		return tx, false
	//	}
	//
	//	symbol := dt.OrderData.Quote
	//	if len(token) > 0 && symbol != token {
	//		return tx, false
	//	}
	//
	//	key := blockatlas.KeyPlaceOrder
	//	title := blockatlas.KeyTitlePlaceOrder
	//	if srcTx.Type == TxCancelOrder {
	//		key = blockatlas.KeyCancelOrder
	//		title = blockatlas.KeyTitleCancelOrder
	//	}
	//	volume, ok := dt.OrderData.GetVolume()
	//	if ok {
	//		value = strconv.Itoa(int(volume))
	//	}
	//
	//	tx.Meta = blockatlas.AnyAction{
	//		Coin:     coin.BNB,
	//		TokenID:  dt.OrderData.Symbol,
	//		Symbol:   TokenSymbol(symbol),
	//		Name:     symbol,
	//		Value:    blockatlas.Amount(value),
	//		Decimals: coin.Coins[coin.BNB].Decimals,
	//		Title:    title,
	//		Key:      key,
	//	}
	//}
	return blockatlas.TxPage{tx}, false
}

func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []Tx, token, adress string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, token, adress)
		if !ok {
			continue
		}
		txs = append(txs, tx...)
	}
	return
}
