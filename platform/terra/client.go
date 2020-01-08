package terra

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

// Client - the HTTP client
type Client struct {
	blockatlas.Request
}

// GetAddrTxs - get all LUNA transactions for a given address
func (c *Client) GetAddrTxs(address string) (txs TxPage, err error) {
	query := url.Values{
		"account": {address},
		"page":    {"1"},
		"limit":   {"25"},
	}
	err = c.Get(&txs, "v1/txs", query)
	if err != nil {
		return TxPage{}, err
	}
	return
}

// GetBlockByNumber return txs with block number
func (c *Client) GetBlockByNumber(num int64) (txs TxPage, err error) {
	err = c.Get(&txs, "txs", url.Values{"tx.height": {strconv.FormatInt(num, 10)}})
	return
}

// CurrentBlockNumber return current block height
func (c *Client) CurrentBlockNumber() (num int64, err error) {
	var block Block
	err = c.Get(&block, "blocks/latest", nil)

	if err != nil {
		return num, err
	}

	num, err = strconv.ParseInt(block.Meta.Header.Height, 10, 64)
	if err != nil {
		return num, errors.E("error to ParseInt", errors.TypePlatformUnmarshal).PushToSentry()
	}

	return
}

// GetAccount loads current account information from the chain
func (c *Client) GetAccount(address string) (result AuthAccount, err error) {
	path := fmt.Sprintf("auth/accounts/%s", address)
	err = c.Get(&result, path, nil)
	return
}
