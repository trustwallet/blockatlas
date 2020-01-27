package ontology

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

// Explorer API max returned transactions per page
const (
	TxPerPage                    = 20
	requestOnlyLatestBlockAmount = 1
)

func (c *Client) GetTxsOfAddress(address, assetName string) (txPage *TxPage, err error) {
	url := fmt.Sprintf("address/%s/%s/%d/1", address, assetName, TxPerPage)
	err = c.Get(&txPage, url, nil)
	return
}

func (c *Client) CurrentBlockNumber() (response *BlockResults, err error) {
	url := fmt.Sprintf("blocklist/%d", requestOnlyLatestBlockAmount)
	err = c.Get(&response, url, nil)
	return
}

func (c *Client) GetBlockByNumber(num int64) (block *BlockResult, err error) {
	url := fmt.Sprintf("block/%d", num)
	err = c.Get(&block, url, nil)
	return
}
