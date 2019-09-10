package nebulas

import (
	"fmt"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"net/http"
	"net/url"
	"strconv"
)

const TxTypeBinary = "binary"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Request    blockatlas.Request
	URL        string
}

func InitClient(BaseURL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		BaseURL: BaseURL,
	}
}

func (c *Client) GetTxs(address string, page int) ([]Transaction, error) {
	values := url.Values{
		"a": {address},
		"p": {strconv.Itoa(page)},
	}

	return c.GetTransactions(values)
}

func (c *Client) GetLatestBlock() (int64, error) {
	path := fmt.Sprintf("/block")
	values := url.Values{
		"type": {"newblock"},
	}
	var response NewBlockResponse

	err := c.Request.Get(&response, c.BaseURL, path, values)
	if err != nil || len(response.Data) == 0 {
		logger.Error("Nebulas: Error loading latest block height")
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
	err := c.Request.Get(&response, c.BaseURL, "tx", values)
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
