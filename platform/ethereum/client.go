package ethereum

import (
	"github.com/trustwallet/golibs/txtype"
)

type EthereumClient interface {
	GetTransactions(address string, coinIndex uint) (txtype.TxPage, error)
	GetTokenTxs(address, token string, coinIndex uint) (txtype.TxPage, error)
	GetTokenList(address string, coinIndex uint) (txtype.TokenPage, error)
	GetCurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64, coinIndex uint) (*txtype.Block, error)
}

func (p *Platform) GetTokenListByAddress(address string) (txtype.TokenPage, error) {
	return p.client.GetTokenList(address, p.CoinIndex)
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*txtype.Block, error) {
	return p.client.GetBlockByNumber(num, p.CoinIndex)
}
