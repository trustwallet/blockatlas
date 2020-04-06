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
	var transferTypes = []TxType{TxTransfer}
	var wg sync.WaitGroup
	out := make(chan []Tx, len(transferTypes))
	wg.Add(len(transferTypes))
	for _, t := range transferTypes {
		go func(txType TxType, address, token string, wg *sync.WaitGroup) {
			defer wg.Done()
			txs, err := p.client.GetAddressAssetTrx(address, token, string(txType))
			if err != nil {
				log.Error("GetAddressAssetTrx : ", err, logger.Params{"txType": txType, "address": address, "token": token})
				return
			}
			out <- txs
		}(t, address, token, &wg)
	}
	wg.Wait()
	close(out)

	srcTx := make([]Tx, 0)
	for r := range out {
		srcTx = append(srcTx, r...)
	}

	return NormalizeTxs(srcTx, address, token), nil
}

// NormalizeTxs converts multiple Binance transactions
func NormalizeTxs(srcTxs []Tx, address, token string) (txs []blockatlas.Tx) {
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
func NormalizeTx(t Tx, address, token string) (blockatlas.TxPage, bool) {
	tBase := blockatlas.Tx{
		ID:        t.TxHash,
		Coin:      coin.BNB,
		From:      t.FromAddr,
		To:        t.ToAddr,
		Fee:       blockatlas.Amount(t.getFee()),
		Date:      t.BlockTimestamp(),
		Block:     t.BlockHeight,
		Status:    blockatlas.StatusCompleted, // FIX
		Error:     "",
		Sequence:  t.Sequence,
		Direction: t.Direction(address),
		Memo:      t.Memo,
	}

	switch t.Type {
	case TxTransfer:
		normalized, ok := normalizeTransfer(tBase, t)
		if !ok {
			return blockatlas.TxPage{}, false
		}
		return blockatlas.TxPage{normalized}, true
	}

	return blockatlas.TxPage{}, false
}

func normalizeTransfer(t blockatlas.Tx, srcTx Tx) (blockatlas.Tx, bool) {
	t.Type = blockatlas.TxTransfer
	bnbCoin := coin.Coins[coin.BNB]
	value := blockatlas.Amount(numbers.DecimalExp(srcTx.Value, 8))

	// Condition for native transfer (BNB)
	if srcTx.Asset == bnbCoin.Symbol {
		t.Meta = blockatlas.Transfer{
			Decimals: bnbCoin.Decimals,
			Symbol:   bnbCoin.Symbol,
			Value:    value,
		}
		return t, true
	}

	// Condition for BEP2 token transfer e.g. TWT-8C2
	t.Meta = blockatlas.NativeTokenTransfer{
		Decimals: bnbCoin.Decimals,
		From:     srcTx.FromAddr,
		Symbol:   TokenSymbol(srcTx.Asset),
		To:       srcTx.ToAddr,
		TokenID:  srcTx.Asset,
		Value:    value,
		Name:     "",
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
