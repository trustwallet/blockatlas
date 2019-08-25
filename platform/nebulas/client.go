package nebulas

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/client"
	"net/http"
	"net/url"
	"strconv"
)

const TxType = "binary"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	RPCURL     string
	Request    blockatlas.Request
	URL        string
}

func InitClient(BaseURL string, RPCURL string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient: http.DefaultClient,
			ErrorHandler: func(res *http.Response, uri string) error {
				return nil
			},
		},
		BaseURL: BaseURL,
		RPCURL:  RPCURL,
	}
}

func (c *Client) GetTxs(address string, page int) ([]Transaction, error) {
	var response Response
	values := url.Values{
		"a": {address},
		"p": {strconv.Itoa(page)},
	}
	var path = ""
	err := client.Request(c.HTTPClient, c.BaseURL, path, values, &response)
	if err != nil {
		return nil, err
	}

	var result []Transaction
	for _, tx := range response.Data.TxnList {
		if tx.Type == TxType {
			result = append(result, tx)
		}
	}

	return result, nil
}

func (c *Client) GetLatestIrreversibleBlock() (int64, error) {
	path := fmt.Sprintf("v1/user/lib")
	var response NasResponse

	err := client.Request(c.HTTPClient, c.RPCURL, path, nil, &response)
	if err != nil {
		logrus.Error("Error loading latest block height")
		return 0, err
	}

	return int64(response.Result.Height), nil
}

func (c *Client) GetBlockByNumber(num int64) (NasBlock, error) {
	path := fmt.Sprintf("v1/user/getBlockByHeight")
	var response NasResponse
	m := make(map[string]string)
	m["height"] = strconv.FormatInt(int64(num), 10)
	m["full_fill_transaction"] = "true"

	err := client.RequestPost(c.HTTPClient, c.RPCURL, path, "application/json", m, &response)
	if err != nil {
		logrus.Error("Error loading current nebulas block")
		return NasBlock{}, err
	}

	return response.Result, nil
}
