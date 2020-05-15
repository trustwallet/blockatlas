package binance

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"strconv"
	"strings"
	"sync"
)

const emptyToken = ""

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	explorerResponse, err := p.GetTokenTxsByAddress(address, emptyToken)
	if err != nil {
		return nil, err
	}
	return filterTxsByType(explorerResponse, blockatlas.TxTransfer), nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	explorerResponse, err := p.explorerClient.getTxsOfAddress(address, token)
	if err != nil {
		return nil, err
	}

	explorerTxs, err := p.addTxDetails(explorerResponse.Txs)
	if err != nil {
		return nil, err
	}

	return normalizeTxs(explorerTxs, address), nil
}

func normalizeTxs(explorerTxs []ExplorerTxs, address string) []blockatlas.Tx {
	var txs []blockatlas.Tx
	for _, tx := range explorerTxs {
		normalizedTxs := normalizeTx(tx, address)
		if normalizedTxs == nil {
			continue
		}
		txs = append(txs, normalizedTxs...)
	}
	return txs
}

func filterTxsByType(txs []blockatlas.Tx, txType blockatlas.TransactionType) []blockatlas.Tx {
	var result = make([]blockatlas.Tx, 0, len(txs))
	for _, tx := range txs {
		if tx.Type == txType {
			result = append(result, tx)
		}
	}
	return result
}

func normalizeTx(srcTx ExplorerTxs, address string) []blockatlas.Tx {
	explorerTxType := srcTx.getTransactionType()
	switch explorerTxType {
	case SingleTransferOperation:
		return normalizeSingleTransfer(srcTx, address)
	case MultiTransferOperation:
		return normalizeMultiTransfer(srcTx, address)
	default:
		return nil
	}
}

func normalizeSingleTransfer(srcTx ExplorerTxs, address string) blockatlas.TxPage {
	if srcTx.TxType != TxTransfer {
		return nil
	}
	tx := getBase(srcTx)
	tx.Direction = srcTx.getDirection(address)
	bnbCoin := coin.Coins[coin.BNB]

	if srcTx.TxAsset == bnbCoin.Symbol {
		tx.Type = blockatlas.TxTransfer
		tx.Meta = blockatlas.Transfer{
			Decimals: bnbCoin.Decimals,
			Symbol:   bnbCoin.Symbol,
			Value:    srcTx.getDexValue(),
		}
		return blockatlas.TxPage{tx}
	}

	if srcTx.TxAsset != "" {
		tx.Type = blockatlas.TxNativeTokenTransfer
		tx.Meta = blockatlas.NativeTokenTransfer{
			Decimals: bnbCoin.Decimals,
			From:     srcTx.FromAddr,
			Symbol:   tokenSymbol(srcTx.TxAsset),
			To:       srcTx.ToAddr,
			TokenID:  srcTx.TxAsset,
			Value:    srcTx.getDexValue(),
		}
		return blockatlas.TxPage{tx}
	}
	return nil
}

func normalizeMultiTransfer(srcTx ExplorerTxs, address string) []blockatlas.Tx {
	var txs blockatlas.TxPage
	for _, t := range srcTx.MultisendTransfers {
		if t.From == address || t.To == address {
			srcTx.FromAddr = t.From
			srcTx.ToAddr = t.To
			srcTx.TxAsset = t.Asset

			if value, err := strconv.ParseFloat(numbers.ToDecimal(t.Amount, 8), 64); err == nil {
				srcTx.Value = value
			}

			if single := normalizeSingleTransfer(srcTx, address); single != nil {
				txs = append(txs, single...)
			}
		}
	}
	return txs
}

func (p *Platform) addTxDetails(txs []ExplorerTxs) ([]ExplorerTxs, error) {
	var (
		wg              sync.WaitGroup
		txsWithDetails  = make([]ExplorerTxs, 0, len(txs))
		txHashChan      = make(chan TxHashRPC, len(txs))
		multiSendTxsMap = make(map[string]ExplorerTxs, len(txs))
	)

	for _, tx := range txs {
		multiSendTxsMap[tx.TxHash] = tx

		if tx.HasChildren != 1 {
			continue
		}

		wg.Add(1)

		go func(srcTx ExplorerTxs, txHashChan chan TxHashRPC, wg *sync.WaitGroup) {
			defer wg.Done()
			txHash, err := p.rpcClient.fetchTransactionHash(srcTx.TxHash)
			if err != nil {
				return
			}
			txHashChan <- *txHash
		}(tx, txHashChan, &wg)
	}

	wg.Wait()
	close(txHashChan)

	for res := range txHashChan {
		if len(res.Tx.Value.Msg) > 0 {
			a, ok := multiSendTxsMap[res.Hash]
			if !ok {
				continue
			}
			a.MultisendTransfers = extractMultiTransfers(res.Tx.Value)
			multiSendTxsMap[res.Hash] = a
		}
	}

	for _, tx := range multiSendTxsMap {
		txsWithDetails = append(txsWithDetails, tx)
	}

	return txsWithDetails, nil
}

// Construct base Tx out of explorer transfer using common fields for all type of Blockatlas transfers
func getBase(srcTx ExplorerTxs) blockatlas.Tx {
	base := blockatlas.Tx{
		ID:    srcTx.TxHash,
		Coin:  coin.BNB,
		From:  srcTx.FromAddr,
		Fee:   srcTx.getDexFee(),
		Date:  srcTx.Timestamp / 1000,
		Block: srcTx.BlockHeight,
		Memo:  srcTx.Memo,
		To:    srcTx.ToAddr,
	}

	status := srcTx.getStatus()
	base.Status = status
	if status == blockatlas.StatusError {
		base.Error = srcTx.getError()
	}

	return base
}

// Extract BEP2 token symbol from asset name e.g: TWT-8C2 => TWT
func tokenSymbol(asset string) string {
	s := strings.Split(asset, "-")
	if len(s) > 1 {
		return s[0]
	}
	return asset
}
