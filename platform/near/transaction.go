package near

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	normalized := make([]blockatlas.Tx, 0)
	return normalized, nil
}
