package oasis

import (
	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var blk int64

	err := c.Post(&blk, "block/tip", nil)
	if err != nil {
		return 0, err
	}

	return blk, nil
}

func (c *Client) GetBlockByNumber(num int64) (*[]Transaction, error) {
	var txs []Transaction

	err := c.Post(&txs, "transactions/block", BlockRequest{BlockIdentifier: num})
	if err != nil {
		return nil, err
	}

	return &txs, nil
}

func (c *Client) GetTrxOfAddress(address string) (*[]Transaction, error) {
	var txs []Transaction

	err := c.Post(&txs, "transactions/address", TransactionsByAddressRequest{Address: address})
	if err != nil {
		return nil, err
	}

	return &txs, nil
}
