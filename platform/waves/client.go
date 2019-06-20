package waves

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

func (c *Client) GetTxs(address string, limit int, after string) ([]Transaction, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions/address/%s/limit/%d?after=%s",
		c.BaseURL,
		address,
		limit,
		after))
}

func (c *Client) getTxs(uri string) ([]Transaction, error) {
	req, _ := http.NewRequest("GET", uri, nil)

	res, err := c.HTTPClient.Do((req))
	if err != nil {
		logrus.WithError(err).Error("Waves: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	txs := new([][]Transaction)
	err = json.NewDecoder(res.Body).Decode(txs)
	if err != nil {
		return nil, err
	}
	txsObj := *txs

	return txsObj[0], nil
}

func (c *Client) GetTokenInfo(tokenId string) (*TokenInfo, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/assets/details/%s",
		c.BaseURL,
		tokenId), nil)

	res, err := c.HTTPClient.Do((req))
	if err != nil {
		logrus.WithError(err).Error("Waves: Failed to get token info")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s", res.Status)
	}

	tokenInfo := new(TokenInfo)
	err = json.NewDecoder(res.Body).Decode(tokenInfo)
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
