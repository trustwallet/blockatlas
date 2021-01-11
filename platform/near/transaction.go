package near

import "github.com/trustwallet/golibs/types"

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	normalized := make([]types.Tx, 0)
	return normalized, nil
}
