package ethereum

import (
	"github.com/trustwallet/golibs/types"
)

func (p *Platform) GetTxsByAddress(address string) (types.Txs, error) {
	return p.client.GetTransactions(address, p.CoinIndex)
}

func (p *Platform) GetTokenTxsByAddress(address string, token string) (types.Txs, error) {
	return p.client.GetTokenTxs(address, token, p.CoinIndex)
}

func (p *Platform) GetTokenListByAddress(address string) ([]types.Token, error) {
	return p.client.GetTokenList(address, p.CoinIndex)
}

func (p *Platform) GetTokenListIdsByAddress(address string) ([]string, error) {
	assets, err := p.GetTokenListByAddress(address)
	if err != nil {
		return []string{}, err
	}
	return types.GetAssetsIds(assets), nil
}
