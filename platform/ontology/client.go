package ontology

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Errorf("Ontology: Failed to get transactions for address %s for asset %s", address, assetName)
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
