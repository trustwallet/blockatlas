package ontology

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
)

type Client struct {
	blockatlas.Request
}

// Explorer API max returned transactions per page
const (
	TxPerPage                    = 20
	requestOnlyLatestBlockAmount = 1
)

func (c *Client) GetTxsOfAddress(address, assetName string) (*TxPage, error) {
	url := fmt.Sprintf("api/v1/explorer/address/%s/%s/%d/1", address, assetName, TxPerPage)
	var txPage TxPage
	err := c.Get(&txPage, url, nil)
	if err != nil {
		return nil, err
	}
	return &txPage, nil
}

func (c *Client) CurrentBlockNumber() (*BlockResults, error) {
	url := fmt.Sprintf("api/v1/explorer/blocklist/%d", requestOnlyLatestBlockAmount)
	var response BlockResults
	err := c.Get(&response, url, nil)
	if err != nil {
		return nil, err
	}
	if response.Error != 0 {
		return nil, errors.E("explorer client CurrentBlockNumber", errors.Params{"platform": "ONT"})
	}
	return &response, nil
}

func (c *Client) GetBlockByNumber(num int64) (*BlockResult, error) {
	url := fmt.Sprintf("api/v1/explorer/block/%d", num)
	var block BlockResult
	err := c.Get(&block, url, nil)
	if err != nil {
		return nil, err
	}
	if block.Error != 0 {
		return nil, errors.E("explorer client GetBlockByNumber", errors.Params{"platform": "ONT"})
	}
	return &block, nil
}

func (c *Client) GetTxDetailsByHash(hash string) (*TxV2, error) {
	url := fmt.Sprintf("v2/transactions/%s", hash)
	var response TxResponse
	err := c.Get(&response, url, nil)
	if err != nil {
		return nil, err
	}
	if response.Msg != "SUCCESS" {
		return nil, errors.E("explorer client GetTxDetailsByHash", errors.Params{"platform": "ONT"})
	}
	var ontTxV2 TxV2
	if response.Result.EventType == 3 {
		ontTxV2 = response.Result
	}
	return &ontTxV2, nil
}
