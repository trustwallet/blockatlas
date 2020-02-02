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

func (c *Client) GetTxsOfAddress(address string, assetName AssetType) (txPage *TxPage, err error) {
	url := fmt.Sprintf("api/v1/explorer/address/%s/%s/%d/1", address, assetName, TxPerPage)
	err = c.Get(&txPage, url, nil)
	return
}

func (c *Client) CurrentBlockNumber() (response *BlockResults, err error) {
	url := fmt.Sprintf("api/v1/explorer/blocklist/%d", requestOnlyLatestBlockAmount)
	err = c.Get(&response, url, nil)
	return
}

func (c *Client) GetBlockByNumber(num int64) (block *BlockResult, err error) {
	url := fmt.Sprintf("api/v1/explorer/block/%d", num)
	err = c.Get(&block, url, nil)
	return
}

func (c *Client) GetTxDetailsByHash(hash string) (*TxV2, error) {
	url := fmt.Sprintf("v2/transactions/%s", hash)
	var response TxResponse
	err := c.Get(&response, url, nil)
	if err != nil || response.Msg != "SUCCESS" {
		return nil, errors.E(err, "explorer client GetTxDetailsByHash", errors.Params{"platform": "ONT"})
	}
	var ontTxV2 TxV2
	if response.Result.EventType == 3 {
		ontTxV2 = response.Result
	}
	return &ontTxV2, nil
}
