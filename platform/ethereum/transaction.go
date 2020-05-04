package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	// Endpoint supports queries without token query parameter
	return p.GetTokenTxsByAddress(address, p.Coin().Symbol)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (blockatlas.TxPage, error) {
	return p.client.GetTransactions(address, p.CoinIndex)
}
