package solana

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	results := make(blockatlas.TxPage, 0)
	return results, nil
}
