package zilliqa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	APIKey     string
}

func (c *Client) newRequest(method, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	req, error := http.NewRequest(method, url, nil)
	req.Header.Set("X-APIKEY", c.APIKey)
	return req, error
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	path := fmt.Sprintf("/addresses/%s/txs", address)
	req, _ := c.newRequest("GET", path)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Zilliqa: Failed to get transactions for ", address)
		return nil, err
	}
	defer res.Body.Close()
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.WithError(err).Error("Zilliqa: Error read response body")
		return nil, err
	}

	if bytes.HasPrefix(body, []byte(`{"message":"Invalid API key specified"`)) {
		return nil, fmt.Errorf("invalid Zilliqa API key")
	}
	
	txs := make([]Tx, 0)
	err = json.Unmarshal(body, &txs)
	if err != nil {
		logrus.WithError(err).Error("Zilliqa: Error decode json transaction response")
		return nil, err
	}

	return txs, nil
}
