package zilliqa

import (
	"strconv"

	"github.com/trustwallet/blockatlas/platform/zilliqa/viewblock"
	"github.com/trustwallet/golibs/types"
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

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	var normalized types.Txs
	var txs []viewblock.Tx

	header, rpcTxs, err := p.rpcClient.GetTxInBlock(num)
	if err != nil {
		return nil, err
	}

	for _, tx := range rpcTxs {
		if tx := TxFromRpc(tx, header); tx != nil {
			txs = append(txs, *tx)
		}
	}

	for _, srcTx := range txs {
		tx := Normalize(&srcTx)
		normalized = append(normalized, tx)
	}

	block := types.Block{
		Number: num,
		Txs:    normalized,
	}
	return &block, nil
}
