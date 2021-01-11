package near

import (
	"github.com/trustwallet/golibs/txtype"
)

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	normalized := make([]txtype.Tx, 0)
	return normalized, nil
}
