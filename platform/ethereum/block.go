package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strconv"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	if srcPage, err := p.client.GetBlockByNumber(num); err == nil {
		var txs []blockatlas.Tx
		for _, srcTx := range srcPage {
			txs = AppendTxs(txs, &srcTx, p.CoinIndex)
		}
		return &blockatlas.Block{
			Number: num,
			ID:     strconv.FormatInt(num, 10),
			Txs:    txs,
		}, nil
	} else {
		return nil, err
	}
}
