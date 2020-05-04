package binance

import (
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

func (p *Platform) getTxChildChan(srcTxs []DexTx) ([]DexTx, error) {
	var (
		wg      sync.WaitGroup
		hashChan = make(chan TxHashRPC)
	)

	for i, srcT := range srcTxs {
		if srcT.HasChildren == 1 {
			wg.Add(1)
			go func(srcTx DexTx, out chan TxHashRPC, wg *sync.WaitGroup) {
				defer wg.Done()
				defer close(out)

				txHash, err := p.rpcClient.GetTransactionHash(srcTx.TxHash)
				if err == nil {
					out <- *txHash
					return
				}
				logger.Error("GetTransactionHash", err, logger.Params{"hash": srcTx.TxHash})
			}(srcT, hashChan, &wg)

			for hash := range hashChan {
				if len(hash.Tx.Value.Msg) > 0 {
					srcTxs[i].MultisendTransfers = extractMultiTransfers(hash.Tx.Value)
				}
			}
		}
	}
	wg.Wait()

	return srcTxs, nil
}

// NormalizeTx converts a Binance transaction into the generic model
func NormalizeTx(srcTx DexTx, address, token string) ([]blockatlas.Tx, bool) {
	switch srcTx.TxType {
	case TxTransfer:
		switch srcTx.QuantityTransferType() {
		case SingleTransfer:
			return normalizeSingleTransfer(srcTx, address, token)
		case MultiTransfer:
			return normalizeMultiTransfer(srcTx, address, token)
		default:
			return blockatlas.TxPage{}, false
		}
	default:
		return blockatlas.TxPage{}, false
	}
}

func normalizeSingleTransfer(srcTx DexTx, address, token string) (blockatlas.TxPage, bool) {
	tx := getBase(&srcTx)
	tx.Direction = srcTx.Direction(address)
	bnbCoin := coin.Coins[coin.BNB]

	// Native coin transfer condition e.g: BNB
	if srcTx.TxAsset == bnbCoin.Symbol {
		tx.Type = blockatlas.TxTransfer
		tx.Meta = blockatlas.Transfer{
			Decimals: bnbCoin.Decimals,
			Symbol:   bnbCoin.Symbol,
			Value:    srcTx.getDexValue(),
		}
		return blockatlas.TxPage{tx}, true
	}

	// Native token transfer condition e.g: TWT-8C2
	if srcTx.TxAsset == token {
		tx.Type = blockatlas.TxNativeTokenTransfer
		tx.Meta = blockatlas.NativeTokenTransfer{
			Decimals: bnbCoin.Decimals,
			From:     srcTx.FromAddr,
			Name:     "", // TODO add name
			Symbol:   TokenSymbol(srcTx.TxAsset),
			To:       srcTx.ToAddr,
			TokenID:  srcTx.TxAsset,
			Value:    srcTx.getDexValue(),
		}
		return blockatlas.TxPage{tx}, true
	}

	return blockatlas.TxPage{}, false
}

func normalizeMultiTransfer(srcTx DexTx, address, token string) ([]blockatlas.Tx, bool) {
	var txs blockatlas.TxPage
	for _, m := range srcTx.MultisendTransfers {
		if m.From == address || m.To == address {
			srcTx.FromAddr = m.From
			srcTx.ToAddr = m.To
			srcTx.TxAsset = m.Asset
			value, err := strconv.ParseFloat(numbers.ToDecimal(m.Amount, 8), 64)
			if err != nil {
				value = 0
			}
			srcTx.Value = value
			single, ok := normalizeSingleTransfer(srcTx, address, token)
			if !ok {
				continue
			}
			txs = append(txs, single...)
		}
	}
	return txs, true
}

func getBase(srcTx *DexTx) blockatlas.Tx {
	return blockatlas.Tx{
		ID:       srcTx.TxHash,
		Coin:     coin.BNB,
		From:     srcTx.FromAddr,
		Fee:      srcTx.getDexFee(),
		Date:     srcTx.Timestamp / 1000,
		Block:    srcTx.BlockHeight,
		Status:   srcTx.getStatus(),
		Memo:     srcTx.Memo,
		Error:    "",
		Sequence: 0,
		To:       srcTx.ToAddr,
	}
}

// Extract BEP2 token symbol from asset name e.g: TWT-8C2 => TWT
func TokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}
