package iotex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxsOfAddress(address string) (*Response, error) {
	uri := fmt.Sprintf("%s/actions/addr/%s?%s",
		c.BaseURL,
		address,
		url.Values{
			"start": {"0"},
			"count": {strconv.FormatInt(models.TxPerPage, 10)},
		}.Encode(),
	)

	res, err := c.HTTPClient.Get(uri)
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		logrus.WithError(err).Errorf("IOTEX: Failed to get transactions for address %s", address)
		return nil, models.ErrSourceConn
	}

	if res.StatusCode != http.StatusOK {
		logrus.WithError(err).Error(res.Status)
		return nil, fmt.Errorf("%s", res.Status)
	}

	var act Response
	if err := json.NewDecoder(res.Body).Decode(&act); err != nil {
		return nil, models.ErrNotFound
	}

	return &act, nil
}
