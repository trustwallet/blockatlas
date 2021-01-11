package nebulas

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	txs, err := p.client.GetTxs(address, 1)
	if err != nil {
		return nil, err
	}

	return NormalizeTxs(txs), nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetLatestBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	txs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	return &txtype.Block{
		Number: num,
		Txs:    NormalizeTxs(txs),
	}, nil
}

func NormalizeTxs(txs []Transaction) []txtype.Tx {
	normalizeTxs := make([]txtype.Tx, 0)
	for _, srcTx := range txs {
		normalizeTxs = append(normalizeTxs, NormalizeTx(srcTx))
	}
	return normalizeTxs
}

func NormalizeTx(srcTx Transaction) txtype.Tx {
	var status = txtype.StatusCompleted
	if srcTx.Status == 0 {
		status = txtype.StatusError
	}
	return txtype.Tx{
		ID:       srcTx.Hash,
		Coin:     coin.NAS,
		From:     srcTx.From.Hash,
		To:       srcTx.To.Hash,
		Fee:      txtype.Amount(srcTx.TxFee),
		Date:     int64(srcTx.Timestamp) / 1000,
		Block:    srcTx.Block.Height,
		Status:   status,
		Sequence: srcTx.Nonce,
		Meta: txtype.Transfer{
			Value:    txtype.Amount(srcTx.Value),
			Symbol:   coin.Coins[coin.NAS].Symbol,
			Decimals: coin.Coins[coin.NAS].Decimals,
		},
	}
}
