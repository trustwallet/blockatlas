package ethereum

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client string

var (
	tx = blockatlas.Tx{
		ID:   "1",
		Coin: 60,
		From: "A",
		To:   "B",
	}
	page = blockatlas.TxPage([]blockatlas.Tx{tx})
	c    Client
)

func getTxClientMock() EthereumClient {
	return &c
}
func TestPlatform_GetTokenTxsByAddress(t *testing.T) {
	p := Platform{
		client: getTxClientMock(),
	}

	resp, err := p.GetTxsByAddress("A")
	assert.Nil(t, err)
	assert.Equal(t, page, resp)
}

func TestPlatform_GetTxsByAddress(t *testing.T) {
	p := Platform{
		client: getTxClientMock(),
	}

	resp, err := p.GetTokenTxsByAddress("A", "")
	assert.Nil(t, err)
	assert.Equal(t, page, resp)
}

func (c Client) GetTransactions(address string, coinIndex uint) (blockatlas.TxPage, error) {
	txs := make([]blockatlas.Tx, 0)
	txs = append(txs, tx)
	return txs, nil
}

func (c Client) GetTokenTxs(address, token string, coinIndex uint) (blockatlas.TxPage, error) {
	txs := make([]blockatlas.Tx, 0)
	txs = append(txs, tx)
	return txs, nil
}

func (c Client) GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error) {
	return blockatlas.TokenPage{}, nil
}

func (c Client) GetCurrentBlockNumber() (int64, error) {
	return 0, nil
}

func (c Client) GetBlockByNumber(num int64, coinIndex uint) (*blockatlas.Block, error) {
	return nil, nil
}
