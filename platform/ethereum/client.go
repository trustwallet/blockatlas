package ethereum

import "github.com/trustwallet/golibs/types"

type EthereumClient interface {
	GetTransactions(address string, coinIndex uint) (types.TxPage, error)
	GetTokenTxs(address, token string, coinIndex uint) (types.TxPage, error)
	GetTokenList(address string, coinIndex uint) (types.TokenPage, error)
	GetCurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64, coinIndex uint) (*types.Block, error)
}

func (p *Platform) GetTokenListByAddress(address string) (types.TokenPage, error) {
	return p.client.GetTokenList(address, p.CoinIndex)
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*types.Block, error) {
	return p.client.GetBlockByNumber(num, p.CoinIndex)
}
