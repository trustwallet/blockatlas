package binance

import "github.com/trustwallet/golibs/txtype"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	block, err := p.client.FetchLatestBlockNumber()
	if err != nil {
		return 0, err
	}
	return block, nil
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	transactionInBlockResponse, err := p.client.FetchTransactionsInBlock(num)
	if err != nil {
		return nil, err
	}
	block := normalizeBlock(transactionInBlockResponse)
	return &block, nil
}
