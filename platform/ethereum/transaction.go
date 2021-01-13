package ethereum

import "github.com/trustwallet/golibs/types"

func (p *Platform) GetTxsByAddress(address string) (types.TxPage, error) {
	return p.client.GetTransactions(address, p.CoinIndex)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (types.TxPage, error) {
	return p.client.GetTokenTxs(address, token, p.CoinIndex)
}

func (p *Platform) GetTokenListByAddress(address string) ([]string, error) {
	return p.client.GetTokenList(address, p.CoinIndex)
}
