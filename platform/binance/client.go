package binance

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas/models"
	"net/http"
	"net/url"
)

// TODO Headers + rate limiting

type Client struct {
	HTTPClient         *http.Client
	ExplorerBaseURL    string
	RPCBaseURL         string
}

func (c *Client) GetTxsOfAddress(address string, token string) (*TxPage, error) {
	uri := fmt.Sprintf("%s/txs?%s",
		c.ExplorerBaseURL,
		url.Values{
			"address": {address},
			"rows":    {"100"},
			"page":    {"1"},
		}.Encode())
		println(uri)
	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get transactions")
		return nil, models.ErrSourceConn
	}

	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusNotFound:
		return nil, getHTTPError(res, "get transactions")
	case http.StatusOK:
		break
	default:
		return nil, fmt.Errorf("%s", res.Status)
	}

	stx := new(TxPage)
	err = json.NewDecoder(res.Body).Decode(stx)
	return stx, nil
}

func (c *Client) getTransactionReceipt(hash string) (*Receipt, error) {
	url := fmt.Sprintf("%s/tx/%s?format=json", c.RPCBaseURL, hash)
	println(url)
	res, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	recp := new(Receipt)
	err = json.NewDecoder(res.Body).Decode(recp)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to decode transaction receipt API response")
		return nil, models.ErrSourceConn
	}

	return recp, nil

}
 
func getHTTPError(res *http.Response, desc string) error {
	var sErr Error
	err := json.NewDecoder(res.Body).Decode(&sErr)
	if err != nil {
		logrus.WithError(err).Error("Binance: Failed to get error")
		return models.ErrSourceConn
	}

	switch sErr.Message {
	case "address is not valid":
		return models.ErrInvalidAddr
	}

	logrus.WithFields(logrus.Fields {
		"status":  res.StatusCode,
		"code":    sErr.Code,
		"message": sErr.Message,
	}).Error("Binance: Failed to " + desc)
	return models.ErrSourceConn
}
