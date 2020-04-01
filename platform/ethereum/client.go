package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type EthereumClient interface {
	GetTransactions(address string, coinIndex uint) (blockatlas.TxPage, error)
	GetTokenTxs(address, token string, coinIndex uint) (blockatlas.TxPage, error)
	GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error)
	GetCurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64, coinIndex uint) (*blockatlas.Block, error)
}

func (p *Platform) GetTokenListByAddress(address string) (blockatlas.TokenPage, error) {
	return p.client.GetTokenList(address, p.CoinIndex)
}

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return p.client.GetBlockByNumber(num, p.CoinIndex)
}
