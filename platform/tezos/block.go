package tezos

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlock()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	txTypes := []string{TxTypeTransaction, TxTypeDelegation}
	srcTxs, err := p.client.GetBlockByNumber(num, txTypes)
	if err != nil {
		return nil, err
	}
	txs := NormalizeTxs(srcTxs, "")
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
