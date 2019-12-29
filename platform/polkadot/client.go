package polkadot

import (
	"encoding/json"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTransfersOfAddress(address string) ([]Transfer, error) {
	var res SubscanResponse
	err := c.Post(&res, "/scan/transfers", TransferRequest{Address: address, Row: blockatlas.TxPerPage})
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(res.Data)
	if err != nil {
		return nil, err
	}
	var data SubscanResponseData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Transfers, nil
}

func (c *Client) GetExtrinsicsOfAddress(address string) ([]Extrinsic, error) {
	var res SubscanResponse
	err := c.Post(&res, "/scan/extrinsics", TransferRequest{Address: address, Row: blockatlas.TxPerPage})
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(res.Data)
	if err != nil {
		return nil, err
	}
	var data SubscanResponseData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data.Extrinsics, nil
}

func (c *Client) GetCurrentBlock() (int64, error) {
	var res SubscanResponse
	err := c.Post(&res, "/scan/metadata", nil)
	if err != nil {
		return 0, err
	}
	bytes, err := json.Marshal(res.Data)
	if err != nil {
		return 0, err
	}
	var data Metadata
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return 0, err
	}
	block, err := strconv.ParseInt(data.BlockNum, 10, 64)
	if err != nil {
		return 0, err
	}
	return block, nil
}
