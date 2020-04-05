package binance

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strings"
	"sync"

	"github.com/trustwallet/blockatlas/coin"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.GetTokenTxsByAddress(address, "")
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	var transferTypes []TxType
	if token == "" {
		transferTypes = []TxType{TxTransfer, TxNewOrder, TxCancelOrder}
	} else {
		transferTypes = []TxType{TxTransfer}
	}

	var wg sync.WaitGroup
	out := make(chan []TxV1, len(transferTypes))
	wg.Add(len(transferTypes))
	for _, t := range transferTypes {
		go func(txType TxType, address, token string, wg *sync.WaitGroup) {
			defer wg.Done()
			txs, err := p.client.GetTxsOfAddress(address, token, string(txType))
			if err != nil {
				log.Error("GetTxsOfAddress : ", err, logger.Params{"txType": txType, "address": address, "token": token})
				return
			}
			out <- txs
		}(t, address, token, &wg)
	}
	wg.Wait()
	close(out)

	srcTx := make([]TxV1, 0)
	for r := range out {
		srcTx = append(srcTx, r...)
	}

	return NormalizeTxs(srcTx, address, token), nil
}

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []TxV1, address, token string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, address, token)
		if !ok {
			continue
		}
		txs = append(txs, tx...)
	}
	return
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(t TxV1, address, token string) (blockatlas.TxPage, bool) {
	tBase := blockatlas.Tx{
		ID:       t.TxHash,
		Coin:     coin.BNB,
		From:     t.FromAddr,
		To:       t.ToAddr,
		Fee:      blockatlas.Amount(t.getFee()),
		Date:     t.BlockTimestamp(),
		Block:    t.BlockHeight,
		Status:   blockatlas.StatusCompleted, // FIX
		Error:    "",
		Sequence: t.Sequence,
		//Type: add infunc
		Direction: t.Direction(address),
		Memo:      t.Memo,
		//Meta: add in func
	}

	switch t.Type {
	case TxTransfer:
		normalized, ok := normalizeTransfer(tBase, t, address, token)
		if !ok {
			return blockatlas.TxPage{}, false
		}
		return blockatlas.TxPage{normalized}, true
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
	return blockatlas.TxPage{tBase}, false
}

func normalizeTransfer(t blockatlas.Tx, srcTx TxV1, address, token string) (blockatlas.Tx, bool) {
	t.Type = blockatlas.TxTransfer
	bnbCoin := coin.Coins[coin.BNB]
	value := blockatlas.Amount(numbers.DecimalExp(srcTx.Value, 8))

	// Condition for native transfer (BNB)
	if srcTx.Asset == bnbCoin.Symbol {
		t.Meta = blockatlas.Transfer{
			Value:    value,
			Symbol:   bnbCoin.Symbol,
			Decimals: bnbCoin.Decimals,
		}
		return t, true
	}

	// Condition for BEP2 token transfer e.g. TWT-8C2
	t.Meta = blockatlas.NativeTokenTransfer{
		Name:     "",
		TokenID:  srcTx.Asset,
		Symbol:   TokenSymbol(srcTx.Asset),
		Value:    value,
		Decimals: bnbCoin.Decimals,
		From:     srcTx.FromAddr,
		To:       srcTx.ToAddr,
	}
	return t, true
}

// Extract BEP2 token symbol from asset name e.g: TWT-8C2 => TWT
func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}
