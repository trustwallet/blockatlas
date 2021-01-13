package elrond

import "github.com/trustwallet/golibs/types"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	return p.client.GetBlockByNumber(num)
}
