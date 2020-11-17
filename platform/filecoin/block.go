package filecoin

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	response, err := p.client.getBlockHeight()
	if err != nil {
		return 0, err
	}
	return int64(response.Height), nil
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return nil, nil
}
