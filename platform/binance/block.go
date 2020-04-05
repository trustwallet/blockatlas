package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	// No native function to get height in explorer API
	// Workaround: Request list of blocks
	// and return number of the newest one
	list, err := p.client.GetBlockList(1)
	if err != nil {
		return 0, err
	}
	if len(list.BlockArray) == 0 {
		return 0, errors.E("no block descriptor found", errors.TypePlatformApi)
	}
	return list.BlockArray[0].BlockHeight, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	var txs blockatlas.TxPage
	childTxs, err := p.getTxChildChan(srcTxs.Txs)
	if err == nil {
		txs = NormalizeTxs(childTxs, "", "")
	} else {
		txs = NormalizeTxs(srcTxs.Txs, "", "")
	}
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
