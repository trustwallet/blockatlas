package cosmos

import (
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

func (c *Client) GetTxsOfAddress(address string) (*Tx, error) {
	uri := fmt.Sprintf("%s/address/txList?%s",
		c.BaseURL,
		url.Values{
			"sender": {address},
			"page":   {strconv.FormatInt(1, 10)},
			"limit":  {strconv.FormatInt(models.TxPerPage, 10)},
		}.Encode())

	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Cosmos: Failed to get transactions")
		return nil, models.ErrSourceConn
	}

	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		logrus.WithError(err).Error("Cosmos: Bad request or 404 error")
		return nil, models.ErrSourceConn
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("%s", res.Status)
	}

	return res, nil
}
