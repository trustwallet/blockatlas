package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	return p.client.GetTransactions(address, p.CoinIndex)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	return p.client.GetTokenTxs(address, token, p.CoinIndex)
}
