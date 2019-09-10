package ontology

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Client - the HTTP client
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// Explorer API max returned transactions per page
const TxPerPage = 20

func (c *Client) GetTxsOfAddress(address, assetName string) (*TxPage, error) {
	uri := fmt.Sprintf("%s/address/%s/%s/%d/1",
		c.BaseURL,
		address,
		assetName,
		TxPerPage,
	)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	txPage := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(txPage)
	if err != nil {
		return nil, err
	}

	return txPage, nil
}
