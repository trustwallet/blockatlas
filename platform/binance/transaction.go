package binance

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	txsFromClient, err := p.client.FetchTransactionsByAddressAndTokenID(address, coin.Binance().Symbol)
	if err != nil {
		return nil, err
	}
	return normalizeTransactions(txsFromClient), nil
}

func (p *Platform) GetTokenTxsByAddress(address, token string) (types.Txs, error) {
	txsFromClient, err := p.client.FetchTransactionsByAddressAndTokenID(address, token)
	if err != nil {
		return nil, err
	}
	return normalizeTransactions(txsFromClient), nil
}
