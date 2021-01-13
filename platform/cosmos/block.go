package cosmos

import "github.com/trustwallet/golibs/types"

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	srcTxs, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}

	txs := p.NormalizeTxs(srcTxs.Txs)
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}
