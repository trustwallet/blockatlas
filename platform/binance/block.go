package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	info, err := p.client.fetchNodeInfo()
	if err != nil {
		return 0, err
	}

	return info.SyncInfo.LatestBlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	blockTxs, err := p.client.GetBlockTransactions(num)
	if err != nil {
		return nil, err
	}

	childTxs := make([]Tx, 0)
	for _, t := range blockTxs {
		if len(t.SubTransactions) > 0 {
			for _, tSub := range t.SubTransactions {
				childTxs = append(childTxs, normalizeBlockSubTx(&t, &tSub))
			}
		}
	}

	txs := NormalizeTxs(childTxs, "")

	return &blockatlas.Block{Number: num, Txs: txs}, nil
}

func normalizeBlockSubTx(t *TxV2, tSub *SubTx) Tx {
	return Tx{
		Asset:       tSub.Asset,
		BlockHeight: t.BlockHeight,
		Code:        t.Code,
		Data:        t.Data,
		Fee:         tSub.Fee,
		FromAddr:    tSub.FromAddr,
		Memo:        t.Memo,
		OrderID:     t.OrderID,
		Sequence:    t.Sequence,
		Source:      t.Source,
		Timestamp:   t.Timestamp,
		ToAddr:      tSub.ToAddr,
		TxHash:      t.TxHash,
		Type:        tSub.Type,
		Value:       tSub.Value,
	}
}

func normalizeBlockSubTx(t *TxV2, tSub *SubTx) Tx {
	return Tx{
		Asset:       tSub.Asset,
		BlockHeight: t.BlockHeight,
		Code:        t.Code,
		Data:        t.Data,
		Fee:         tSub.Fee,
		FromAddr:    tSub.FromAddr,
		Memo:        t.Memo,
		OrderID:     t.OrderID,
		Sequence:    t.Sequence,
		Source:      t.Source,
		Timestamp:   t.Timestamp,
		ToAddr:      tSub.ToAddr,
		TxHash:      t.TxHash,
		Type:        tSub.Type,
		Value:       tSub.Value,
	}
}
