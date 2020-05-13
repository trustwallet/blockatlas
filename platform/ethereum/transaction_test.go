package ethereum

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"testing"
)

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

func getTxClientMock() EthereumClient {
	return &c
}

var tx = blockatlas.Tx{
	ID:        "1",
	Coin:      60,
	From:      "A",
	To:        "B",
	Fee:       "",
	Date:      0,
	Block:     0,
	Status:    "",
	Error:     "",
	Sequence:  0,
	Type:      "",
	Inputs:    nil,
	Outputs:   nil,
	Direction: "",
	Memo:      "",
	Meta:      nil,
}

var page = blockatlas.TxPage([]blockatlas.Tx{tx})

type Client string

var c Client

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
