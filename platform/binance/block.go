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
	childTxs := make([]DexTx, 0, len(blockTxs))
	for _, bTrx := range blockTxs {
		childTxs = append(childTxs, normalizeBlockSubTx(bTrx))
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
func normalizeBlockSubTx(txV2 TxV2) DexTx {
	tx := DexTx{
		TxAsset:     txV2.Asset,
		Code:        txV2.Code,
		FromAddr:    txV2.FromAddr,
		TxHash:      txV2.TxHash,
		Memo:        txV2.Memo,
		ToAddr:      txV2.ToAddr,
		TxType:      txV2.Type,
		BlockHeight: txV2.BlockHeight,
	}

	value, err := numbers.StringNumberToFloat64(txV2.Value)
	if err != nil {
		tx.Value = 0
	} else {
		tx.Value = value
	}

	var feeStr string
	if txV2.Fee == "" && len(txV2.SubTransactions) > 1 {
		feeStr = txV2.SubTransactions[0].Fee
	} else {
		feeStr = txV2.Fee
	}

	feeFloat, err := numbers.StringNumberToFloat64(feeStr)
	if err != nil {
		tx.TxFee = 0
	}
	tx.TxFee = feeFloat

	if len(txV2.SubTransactions) > 0 {
		tx.HasChildren = 1
	} else {
		tx.HasChildren = 0
	}

	time, err := time.Parse(time.RFC3339, txV2.Timestamp)
	if err != nil {
		tx.Timestamp = time.Unix()
	}

	multisend := make([]multiTransfer, 0, len(txV2.SubTransactions))
	for _, st := range txV2.SubTransactions {
		amount, _ := numbers.StringNumberToFloat64(st.Value)
		m := multiTransfer{
			Amount: numbers.Float64toString(amount),
			Asset:  st.Asset,
			From:   st.FromAddr,
			To:     st.ToAddr,
		}
		multisend = append(multisend, m)
	}
	tx.MultisendTransfers = multisend

	return tx
}
