package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	srcTxs, err := p.client.GetBlockByNumber(num, txTypes)
	if err != nil {
		logger.Error("GetAddrTxs", err, logger.Params{"txType": txTypes, "num": num})
		return nil, err
	}
	txs := NormalizeTxs(srcTxs, "")
	return &blockatlas.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
