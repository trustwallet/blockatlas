package oasis

import (
	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var blk int64

	err := c.Get(&blk, "block", nil)
	if err != nil {
		return 0, err
	}

	return blk, nil
}

func (c *Client) GetBlockByNumber(num int64) (*[]Transaction, error) {
	var trxs []Transaction

	err := c.Post(&trxs, "transactions/block", BlockRequest{BlockIdentifier: num})
	if err != nil {
		return nil, err
	}

	return &trxs, nil
}

func (c *Client) GetTrxOfAddress(address string) (*[]Transaction, error) {
	var trxs []Transaction

	err := c.Post(&trxs, "transactions/address", TransactionsByAddressRequest{Address: address})
	if err != nil {
		return nil, err
	}

	return &trxs, nil
}
