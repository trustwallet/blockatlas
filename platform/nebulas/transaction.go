package nebulas

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	txs, err := p.client.GetTxs(address, 1)
	if err != nil {
		return nil, err
	}

	return NormalizeTxs(txs), nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	txs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	return &types.Block{
		Number: num,
		Txs:    NormalizeTxs(txs),
	}, nil
}

func NormalizeTxs(txs []Transaction) types.Txs {
	normalizeTxs := make(types.Txs, 0)
	for _, srcTx := range txs {
		normalizeTxs = append(normalizeTxs, NormalizeTx(srcTx))
	}
	return normalizeTxs
}

func NormalizeTx(srcTx Transaction) types.Tx {
	var status = types.StatusCompleted
	if srcTx.Status == 0 {
		status = types.StatusError
	}
	return types.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.NEBULAS,
		From:     srcTx.From.Hash,
		To:       srcTx.To.Hash,
		Fee:      types.Amount(srcTx.TxFee),
		Date:     int64(srcTx.Timestamp) / 1000,
		Block:    srcTx.Block.Height,
		Status:   status,
		Sequence: srcTx.Nonce,
		Meta: types.Transfer{
			Value:    types.Amount(srcTx.Value),
			Symbol:   coin.Nebulas().Symbol,
			Decimals: coin.Nebulas().Decimals,
		},
	}
}
