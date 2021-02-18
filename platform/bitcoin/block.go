package bitcoin

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	block, err := p.client.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		return nil, err
	}
	var normalized types.Txs
	for _, tx := range block {
		normalized = append(normalized, normalizeTransaction(tx, p.CoinIndex))
	}
	return &types.Block{
		Number: num,
		Txs:    normalized,
	}, nil

}
