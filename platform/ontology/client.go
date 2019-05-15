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

const TxPerPage = 20

func (c *Client) GetTxsOfAddress(address string, assetName string, page uint) (*TxPage, error) {
	uri := fmt.Sprintf("%s/address/%s/%s/%d/%d",
		c.BaseURL,
		address,
		assetName,
		TxPerPage,
		page,
	)

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Errorf("Ontology: Failed to get transactions for address %s for asset %s", address, assetName)
	}
	defer res.Body.Close()

	print(res)

	txPage := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(txPage)
	return txPage, err
}