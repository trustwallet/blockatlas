package zilliqa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type Client struct {
	client  *http.Client
	apiKey  string
	baseURL string
}

func NewClient() *Client {
	apiKey := os.Getenv("ATLAS_ZILLIQA_KEY")
	return &Client{
		client: http.DefaultClient,
		apiKey: apiKey,
	}
}

func (c *Client) newRequest(method string, path string) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	req, error := http.NewRequest(method, url, nil)
	req.Header.Set("X-APIKEY", c.apiKey)
	return req, error
}

func (c *Client) GetTxsOfAddress(address string) ([]Tx, error) {
	path := fmt.Sprintf("/addresses/%s/txs", address)
	req, _ := c.newRequest("GET", path)
	res, err := c.client.Do(req)
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
	
	txs := make([]Tx, 0)
	err = json.Unmarshal(body, &txs)
	if err != nil {
		logrus.WithError(err).Error("Zilliqa: Error decode json transaction response")
		return nil, err
	}

	return txs, nil
}
