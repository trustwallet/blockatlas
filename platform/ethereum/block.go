package ethereum

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	return p.client.GetBlockByNumber(num, p.CoinIndex)
}
