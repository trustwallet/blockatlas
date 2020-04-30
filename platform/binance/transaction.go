package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strings"
	"sync"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.GetTokenTxsByAddress(address, "")
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	txPage, err := p.dexClient.GetTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}
	txs, err := p.getTxChildChan(txPage.Txs)
	if err != nil {
		return nil, err
	}
	return NormalizeTxs(txs, address, token), nil
}

// getTxChildChan get all child assets from a tx
func (p *Platform) getTxChildChan(srcTxs []DexTx) ([]DexTx, error) {
	txs := make([]DexTx, 0)
	var wg sync.WaitGroup
	outChan := make(chan TxHash, len(srcTxs))
	wg.Add(len(srcTxs))
	for _, srcTx := range srcTxs {
		if srcTx.HasChildren == 0 {
			// Return the same transaction if doesn't have a child
			txs = append(txs, srcTx)
			continue
		}
		wg.Add(1)
		go func(srcTx DexTx, out chan TxHash, wg *sync.WaitGroup) {
			defer wg.Done()
			txHash, err := p.rpcClient.GetTransactionHash(srcTx.Hash)
			if err != nil {
				// Return the same transaction if an error occurs
				out <- srcTx
				logger.Error("GetTransactionsByBlockChan", err, logger.Params{"hash": srcTx.Hash})
			}
			out <- txHash
		}(srcTx, outChan, &wg)
	}

	wg.Wait()
	close(out)

	for r := range out {
		txs = append(txs, r)
	}
	return txs, nil
}

// Converts multiple transactions
func NormalizeTxs(srcTxs []DexTx, adress, token string) (txs []blockatlas.Tx) {
	for _, srcTx := range srcTxs {
		tx, ok := NormalizeTx(srcTx, adress, token)
		if !ok {
			continue
		}
		txs = append(txs, tx...)
	}
	return
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx DexTx, address, token string) (blockatlas.TxPage, bool) {
	tx := blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.BNB,
		From:   srcTx.FromAddr,
		To:     srcTx.ToAddr,
		Fee:    blockatlas.Amount(srcTx.getDexFee()),
		Date:   srcTx.Timestamp / 1000,
		Block:  srcTx.BlockHeight,
		Status: blockatlas.StatusCompleted,
		Memo:   srcTx.Memo,
	}

	switch srcTx.Type {
	case TxTransfer:
		return normalizeTransfer(tx, srcTx, token, address)
	}
	return blockatlas.TxPage{tx}, false
}

// Converts a Binance transaction into the generic model
func normalizeTransfer(tx blockatlas.Tx, srcTx DexTx, token, address string) (blockatlas.TxPage, bool) {
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

// Extract BEP2 token symbol from asset name e.g: TWT-8C2 => TWT
func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}
