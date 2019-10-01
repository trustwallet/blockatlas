package ontology

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func InitClient(baseUrl string) Client {
	return Client{
		Request: blockatlas.Request{
			HttpClient:   blockatlas.DefaultClient,
			ErrorHandler: blockatlas.DefaultErrorHandler,
			BaseUrl:      baseUrl,
		},
	}
}

// Explorer API max returned transactions per page
const TxPerPage = 20

func (c *Client) GetTxsOfAddress(address, assetName string) (txPage *TxPage, err error) {
	url := fmt.Sprintf("/address/%s/%s/%d/1", address, assetName, TxPerPage)
	err = c.Get(&txPage, url, nil)
	return
}
