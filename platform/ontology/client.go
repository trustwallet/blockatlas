package ontology

import (
	"encoding/json"
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/errors"
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
		return nil, errors.E(err, errors.TypePlatformRequest, errors.Params{"url": uri,"platform": "ontology"})
	}
	defer res.Body.Close()

	txPage := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(txPage)
	if err != nil {
		return nil, errors.E(err, errors.TypePlatformUnmarshal, errors.Params{"url": uri,"platform": "ontology"})
	}

	return txPage, nil
}
