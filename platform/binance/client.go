package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"net/url"
	"strconv"
)

// TODO Headers + rate limiting

type Client struct {
	HTTPClient *http.Client
	RpcURL     string
}

func (c *Client) GetTxsOfAddress(address string) (*TxPage, error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.RpcURL,
		url.Values{
			"address": {address},
			"rows":    {strconv.Itoa(models.TxPerPage)},
			"page":    {"1"},
		}.Encode())
	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get transactions")
		return nil, ErrSourceConn
	}

	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		return nil, getHttpError(res, "get transactions")
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("%s", res.Status)
	}

	stx := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(stx)
	return stx, nil
}

func getHttpError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get error")
		return ErrSourceConn
	} else {
		switch sErr.Message {
		case "address is not valid":
			return ErrInvalidAddr
		}

		logrus.WithFields(logrus.Fields {
			"status":  res.StatusCode,
			"code":    sErr.Code,
			"message": sErr.Message,
		}).Error("Binance: Failed to " + desc)
		return ErrSourceConn
	}
}
