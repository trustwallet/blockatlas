package source

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

// TODO Headers + rate limiting

type Client struct {
	HttpClient *http.Client
	RpcUrl     string
}

func (c *Client) GetTxsOfAddress(address string) (*TxPage, error) {
	uri := fmt.Sprintf("%s/transactions?%s",
		c.RpcUrl,
		url.Values{"address": {address}}.Encode())
	res, err := c.HttpClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get transactions")
		return nil, ErrSourceConn
	}

	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		return nil, getHttpError(res, "get transactions")
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
