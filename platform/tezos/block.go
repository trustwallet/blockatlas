package tezos

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	srcTxs, err := p.client.GetBlockByNumber(num, txTypes)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcTxs, "")
	return &txtype.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
