package binance

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

func (p *Platform) GetTxsByAddress(address string) (blockatlas.TxPage, error) {
	txsFromClient, err := p.client.FetchTransactionsByAddressAndTokenID(address, coin.Binance().Symbol)
	if err != nil {
		return nil, err
	}
	return normalizeTransactions(txsFromClient), nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (blockatlas.TxPage, error) {
	txsFromClient, err := p.client.FetchTransactionsByAddressAndTokenID(address, token)
	if err != nil {
		return nil, err
	}
	return normalizeTransactions(txsFromClient), nil
}
