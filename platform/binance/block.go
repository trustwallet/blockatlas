package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"time"
)

func (p Platform) CurrentBlockNumber() (int64, error) {
	info, err := p.rpcClient.fetchNodeInfo()
	if err != nil {
		return 0, err
	}

	return info.SyncInfo.LatestBlockHeight, nil
}

func (p Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	blockTransactions, err := p.rpcClient.fetchBlockTransactions(num)
	if err != nil {
		return nil, err
	}

	explorerTransactions := make([]ExplorerTxs, 0, len(blockTransactions))
	for _, tx := range blockTransactions {
		explorerTransactions = append(explorerTransactions, normalizeTxsForExplorer(tx))
	}

	var normalizedTxs []blockatlas.Tx
	for _, tx := range explorerTransactions {
		normalizedTx := normalizeTx(tx, "", "")
		if normalizedTx == nil {
			continue
		}
		normalizedTxs = append(normalizedTxs, normalizedTx...)
	}

	return &blockatlas.Block{Number: num, Txs: normalizedTxs}, nil
}

func normalizeTxsForExplorer(txV2 TxV2) ExplorerTxs {
	tx := ExplorerTxs{
		TxAsset:     txV2.Asset,
		Code:        txV2.Code,
		FromAddr:    txV2.FromAddr,
		TxHash:      txV2.TxHash,
		Memo:        txV2.Memo,
		ToAddr:      txV2.ToAddr,
		TxType:      txV2.Type,
		BlockHeight: txV2.BlockHeight,
	}

	if value, err := numbers.StringNumberToFloat64(txV2.Value); err == nil {
		tx.Value = value
	}
	if txV2.Fee == "" && len(txV2.SubTransactions) > 1 {
		txV2.Fee = txV2.SubTransactions[0].Fee
	}
	if fee, err := numbers.StringNumberToFloat64(txV2.Fee); err == nil {
		tx.TxFee = fee
	}
	if len(txV2.SubTransactions) > 0 {
		tx.HasChildren = 1
	}
	if t, err := time.Parse(time.RFC3339, txV2.Timestamp); err == nil {
		tx.Timestamp = t.Unix()
	}

	mts := make([]MultiTransfer, len(txV2.SubTransactions))
	for i, st := range txV2.SubTransactions {
		mts[i] = MultiTransfer{Amount: st.Value, Asset: st.Asset, From: st.FromAddr, To: st.ToAddr}
	}
	tx.MultisendTransfers = mts
	return tx
}
