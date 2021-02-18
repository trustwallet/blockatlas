package near

import "github.com/trustwallet/golibs/types"

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	normalized := make(types.Txs, 0)
	return normalized, nil
}
