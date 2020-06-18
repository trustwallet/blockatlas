package elrond

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.CurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return p.client.GetBlockByNumber(num)
}
