package elrond

import "github.com/trustwallet/golibs/txtype"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	return p.client.GetBlockByNumber(num)
}
