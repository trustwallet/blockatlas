package binance

import (
	"encoding/json"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
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

	var normTxs []blockatlas.Tx
	for _, srcTx := range txs {
		tx, ok := NormalizeTx(srcTx, address, token)
		if !ok {
			continue
		}
		normTxs = append(normTxs, tx...)
	}

	return normTxs, nil
}

// getTxChildChan get transaction hash for multisend transfer
func (p *Platform) getTxChildChan(srcTxs []DexTx) ([]DexTx, error) {
	var (
		wg      sync.WaitGroup
		outChan = make(chan TxHashRPC)
	)

	for i, srcT := range srcTxs {
		if srcT.HasChildren == 1 {
			wg.Add(1)
			go func(srcTx DexTx, out chan TxHashRPC, wg *sync.WaitGroup) {
				defer wg.Done()
				defer close(out)

				txHash, err := p.rpcClient.GetTransactionHash(srcTx.Hash)
				if err == nil {
					out <- *txHash
					return
					logger.Error("GetTransactionHash", err, logger.Params{"hash": srcTx.Hash})
				}
			}(srcT, outChan, &wg)

			select {
			case hash := <-outChan:
				if len(hash.Tx.V.Messages) > 0 {
					srcTxs[i].MultisendTransfers = hash.Tx.V.Messages[0]
				}
			}
		}
	}
	wg.Wait()

	return srcTxs, nil
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx DexTx, address, token string) (blockatlas.TxPage, bool) {
	txBase := blockatlas.Tx{
		ID:     srcTx.Hash,
		Coin:   coin.BNB,
		Fee:    blockatlas.Amount(srcTx.getDexFee()),
		Date:   srcTx.Timestamp / 1000,
		Block:  srcTx.BlockHeight,
		Status: blockatlas.StatusCompleted, // TODO status
		Memo:   srcTx.Memo,
	}

	switch srcTx.Type {
	case TxTransfer:
		switch srcTx.QuantityTransferType() {
		case SingleTransfer:
			return normalizeSingleTransfer(txBase, srcTx, address, token)
		case MultiTransfer:
			return normalizeMultiTransfer(txBase, srcTx, address, token)
		default:
			return blockatlas.TxPage{}, false
		}
	default:
		return blockatlas.TxPage{}, false
	}

	return blockatlas.TxPage{}, false
}

func normalizeSingleTransfer(tx blockatlas.Tx, srcTx DexTx, address, token string) (blockatlas.TxPage, bool) {
	tx.From = srcTx.FromAddr
	tx.To = srcTx.ToAddr

	bnbCoin := coin.Coins[coin.BNB]
	value := numbers.DecimalExp(string(srcTx.Value), 8)
	if srcTx.Asset == bnbCoin.Symbol {
		tx.Direction = srcTx.Direction(address)
		// Native coin transfer condition e.g: BNB
		tx.Meta = blockatlas.Transfer{
			Decimals: bnbCoin.Decimals,
			Symbol:   bnbCoin.Symbol,
			Value:    blockatlas.Amount(value),
		}
		return blockatlas.TxPage{tx}, true
	}

	// Native token transfer condition e.g: TWT-8C2
	if srcTx.Asset == token {
		tx.Meta = blockatlas.NativeTokenTransfer{
			Decimals: bnbCoin.Decimals,
			From:     srcTx.FromAddr,
			Name:     "", // TODO add name
			Symbol:   TokenSymbol(srcTx.Asset),
			To:       srcTx.ToAddr,
			TokenID:  srcTx.Asset,
			Value:    blockatlas.Amount(value),
		}
		return blockatlas.TxPage{tx}, true
	}

	return blockatlas.TxPage{}, false
}

func normalizeMultiTransfer(tx blockatlas.Tx, srcTx DexTx, address, token string) (page []blockatlas.Tx, ok bool) {
	var multisends = srcTx.extractMultiTransfers(address)

	for _, m := range multisends {
		srcTx.FromAddr = m.From
		srcTx.ToAddr = m.To
		srcTx.Asset = m.Asset
		val, err := strconv.ParseFloat(m.Amount, 64)
		if err != nil {
			val = 0
		}
		srcTx.Value = val
		single, ok := normalizeSingleTransfer(tx, srcTx, address, token)
		if !ok {
			continue
		}
		page = append(page, single)
	}

	return
}

// Extract BEP2 token symbol from asset name e.g: TWT-8C2 => TWT
func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}
