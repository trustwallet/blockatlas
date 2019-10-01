package nebulas

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"net/url"
	"strconv"
)

const TxTypeBinary = "binary"

type Client struct {
	blockatlas.Request
}

func (c *Client) GetTxs(address string, page int) ([]Transaction, error) {
	values := url.Values{
		"a": {address},
		"p": {strconv.Itoa(page)},
	}

	return c.GetTransactions(values)
}

func (c *Client) GetLatestBlock() (int64, error) {
	values := url.Values{
		"type": {"newblock"},
	}
	var response NewBlockResponse

	err := c.Get(&response, "block", values)
	if err != nil || len(response.Data) == 0 {
		return 0, err
	}

	return response.Data[0].Height, nil
}

func (c *Client) GetBlockByNumber(num int64) ([]Transaction, error) {
	values := url.Values{
		"block": {strconv.Itoa(int(num))},
	}
	return c.GetTransactions(values)
}

func (c *Client) GetTransactions(values url.Values) ([]Transaction, error) {
	var response Response
	err := c.Get(&response, "tx", values)
	if err != nil {
		return nil, err
	}

	result := make([]Transaction, 0)
	for _, tx := range response.Data.Transactions {
		if tx.Type == TxTypeBinary {
			result = append(result, tx)
		}
	}

	return result, nil
}
