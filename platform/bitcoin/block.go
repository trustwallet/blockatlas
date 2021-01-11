package bitcoin

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	block, err := p.client.GetAllTransactionsByBlockNumber(num)
	if err != nil {
		return nil, err
	}
	var normalized []txtype.Tx
	for _, tx := range block {
		normalized = append(normalized, normalizeTransaction(tx, p.CoinIndex))
	}
	return &txtype.Block{
		Number: num,
		Txs:    normalized,
	}, nil

}
