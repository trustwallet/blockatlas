package zilliqa

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	info, err := p.rpcClient.GetBlockchainInfo()
	if err != nil {
		return 0, err
	}
	block, err := strconv.ParseInt(info.NumTxBlocks, 10, 64)
	if err != nil {
		return 0, err
	}

	return block, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	var normalized []blockatlas.Tx
	txs, err := p.rpcClient.GetTxInBlock(num)
	if err != nil {
		return nil, err
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		normalized = append(normalized, tx)
	}
	block := blockatlas.Block{
		Number: num,
		Txs:    normalized,
	}

	return &block, nil
}
