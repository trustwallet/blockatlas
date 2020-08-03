package binance

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.FetchLatestBlockNumber()
	if err != nil {
		return 0, err
	}

	return block, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return nil, nil
}

//
//func normalizeTxsToExplorer(txV2 TxV2) ExplorerTxs {
//	tx := ExplorerTxs{
//		TxAsset:     txV2.Asset,
//		Code:        txV2.Code,
//		FromAddr:    txV2.FromAddr,
//		TxHash:      txV2.TxHash,
//		Memo:        txV2.Memo,
//		ToAddr:      txV2.ToAddr,
//		TxType:      txV2.Type,
//		BlockHeight: txV2.BlockHeight,
//	}
//
//	if value, err := numbers.StringNumberToFloat64(txV2.Value); err == nil {
//		tx.Value = value
//	}
//	if txV2.Fee == "" && len(txV2.SubTransactions) > 1 {
//		txV2.Fee = txV2.SubTransactions[0].Fee
//	}
//	if fee, err := numbers.StringNumberToFloat64(txV2.Fee); err == nil {
//		tx.TxFee = fee
//	}
//	if len(txV2.SubTransactions) > 0 {
//		tx.HasChildren = 1
//	}
//	if t, err := time.Parse(time.RFC3339, txV2.Timestamp); err == nil {
//		tx.Timestamp = t.Unix() * 1000
//	}
//
//	mts := make([]MultiTransfer, len(txV2.SubTransactions))
//	for i, st := range txV2.SubTransactions {
//		mts[i] = MultiTransfer{Amount: st.Value, Asset: st.Asset, From: st.FromAddr, To: st.ToAddr}
//	}
//	tx.MultisendTransfers = mts
//	return tx
//}
