package polkadot

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func (c *Client) GetTransfersOfAddress(address string) ([]Transfer, error) {
	var res SubscanResponse
	err := c.Post(&res, "scan/transfers", TransfersRequest{Address: address, Row: blockatlas.TxPerPage})
	if err != nil {
		return nil, err
	}
	return res.Data.Transfers, nil
}

func (c *Client) GetExtrinsicsOfAddress(address string) ([]Extrinsic, error) {
	var res SubscanResponse
	err := c.Post(&res, "scan/extrinsics", TransfersRequest{Address: address, Row: blockatlas.TxPerPage})
	if err != nil {
		return nil, err
	}
	return res.Data.Extrinsics, nil
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var res SubscanResponse
	err := c.Post(&res, "scan/metadata", nil)
	if err != nil {
		return 0, err
	}
	block, err := strconv.ParseInt(res.Data.BlockNumber, 10, 64)
	if err != nil {
		return 0, err
	}
	return block, nil
}

func (c *Client) GetBlockByNumber(number int64) ([]Extrinsic, error) {
	var res SubscanResponse
	err := c.Post(&res, "scan/block", BlockRequest{BlockNumber: number})
	if err != nil {
		return nil, err
	}
	return res.Data.Extrinsics, nil
}
