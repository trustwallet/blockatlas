package ethereum

import "github.com/trustwallet/golibs/txtype"

func (p *Platform) GetTxsByAddress(address string) (txtype.TxPage, error) {
	return p.client.GetTransactions(address, p.CoinIndex)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (txtype.TxPage, error) {
	return p.client.GetTokenTxs(address, token, p.CoinIndex)
}
