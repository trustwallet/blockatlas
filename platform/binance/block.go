package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/numbers"
	"time"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	info, err := p.rpcClient.fetchNodeInfo()
	if err != nil {
		return 0, err
	}

	return info.SyncInfo.LatestBlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	blockTxs, err := p.rpcClient.GetBlockTransactions(num)
	if err != nil {
		return nil, err
	}
	childTxs := make([]DexTx, 0)
	for _, bTrx := range blockTxs {
		childTxs = append(childTxs, normalizeBlockSubTx(&bTrx))
	}

	var normTxs []blockatlas.Tx
	print(len(childTxs))
	for _, srcTx := range childTxs {
		normT, ok := NormalizeTx(srcTx, "", "")
		if !ok {
			continue
		}
		normTxs = append(normTxs, normT...)
	}

	return &blockatlas.Block{Number: num, Txs: normTxs}, nil
}

// Normalize block sub transaction from RPC to explorer transaction
func normalizeBlockSubTx(t *TxV2) DexTx {
	tx := DexTx{
		TxAsset:     t.Asset,
		Code:        t.Code,
		FromAddr:    t.FromAddr,
		TxHash:      t.TxHash,
		Memo:        t.Memo,
		ToAddr:      t.ToAddr,
		TxType:      t.Type,
		BlockHeight: t.BlockHeight,
	}

	tx.Value = numbers.StringNumberToFloat64(t.Value)

	if t.Fee == "" && len(t.SubTransactions) > 1 {
		tx.TxFee = numbers.StringNumberToFloat64(t.SubTransactions[0].Fee)
	} else {
		tx.TxFee = numbers.StringNumberToFloat64(t.Fee)
	}

	if len(t.SubTransactions) > 0 {
		tx.HasChildren = 1
	} else {
		tx.HasChildren = 0
	}

	time, err := time.Parse(time.RFC3339, t.Timestamp)
	if err != nil {
		tx.Timestamp = time.Unix()
	}

	var multisend []multiTransfer
	for _, st := range t.SubTransactions {
		v := multiTransfer{
			Amount: numbers.ToDecimal(st.Value, 8),
			Asset:  st.Asset,
			From:   st.FromAddr,
			To:     st.ToAddr,
		}
		multisend = append(multisend, v)
	}
	tx.MultisendTransfers = multisend

	return tx
}
