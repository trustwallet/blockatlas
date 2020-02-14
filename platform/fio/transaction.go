package fio

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (page blockatlas.TxPage, err error) {
	return page, err
}
