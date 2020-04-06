package ontology

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/url"
)

type Client struct {
	blockatlas.Request
}

func (c *Client) GetBalances(address string) (balances BalancesResult, err error) {
	path := fmt.Sprintf("v2/addresses/%s/native/balances", address)
	err = c.Get(&balances, path, nil)
	if err != nil || balances.Msg != MsgSuccess {
		return balances, errors.E(err, "explorer client GetBalances", errors.Params{"platform": "ONT"})
	}
	return
}

func (c *Client) GetTxsOfAddress(address string) (txPage TxsResult, err error) {
	query := url.Values{"page_size": {"20"}, "page_number": {"1"}}
	path := fmt.Sprintf("v2/addresses/%s/transactions", address)
	err = c.Get(&txPage, path, query)
	if err != nil || txPage.Msg != MsgSuccess {
		return txPage, errors.E(err, "explorer client GetTxsOfAddress", errors.Params{"platform": "ONT"})
	}
	return
}

func (c *Client) CurrentBlockNumber() (blocks BlockResult, err error) {
	query := url.Values{"page_size": {"1"}, "page_number": {"1"}}
	path := "v2/blocks"
	err = c.Get(&blocks, path, query)
	if err != nil || blocks.Msg != MsgSuccess {
		return blocks, errors.E(err, "explorer client CurrentBlockNumber", errors.Params{"platform": "ONT"})
	}
	return
}

func (c *Client) GetBlockByNumber(num int64) (block BlockResults, err error) {
	path := fmt.Sprintf("v2/blocks/%d", num)
	err = c.Get(&block, path, nil)
	if err != nil || block.Msg != MsgSuccess {
		return block, errors.E(err, "explorer client GetBlockByNumber", errors.Params{"platform": "ONT"})
	}
	return
}

func (c *Client) GetTxDetailsByHash(hash string) (Tx, error) {
	path := fmt.Sprintf("v2/transactions/%s", hash)
	var response TxResult
	err := c.Get(&response, path, nil)
	if err != nil || response.Msg != MsgSuccess {
		return Tx{}, errors.E(err, "explorer client GetTxDetailsByHash", errors.Params{"platform": "ONT"})
	}
	var ontTxV2 Tx
	if response.Result.EventType == 3 {
		ontTxV2 = response.Result
	}
	return ontTxV2, nil
}
