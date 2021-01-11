package ontology

import (
	"errors"

	"github.com/trustwallet/golibs/types"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.CurrentBlockNumber()
	if err != nil {
		return 0, err
	}
	if len(block.Result.Records) == 0 {
		return 0, errors.New("invalid block height result")
	}
	return block.Result.Records[0].Height, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	blockOnt, err := p.client.GetBlockByNumber(num)
	if err != nil {
		return nil, err
	}
	txsRaw, err := p.getTxDetails(blockOnt.Result.Txs)
	if err != nil {
		return nil, err
	}
	txs := normalizeTxs(txsRaw, AssetAll)
	return &types.Block{
		Number: num,
		Txs:    txs,
	}, nil
}
