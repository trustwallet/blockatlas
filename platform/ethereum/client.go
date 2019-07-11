package ethereum

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/trustwallet/blockatlas"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	HTTPClient        *http.Client
	BaseURL           string
	CollectionsURL    string
	CollectionsApiKey string
}

func (c *Client) GetTxs(address string, build string) (*Page, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions?%s",
		c.BaseURL,
		url.Values{
			"address": {address},
		}.Encode()), build)
}

func (c *Client) GetTxsWithContract(address, contract string, build string) (*Page, error) {
	return c.getTxs(fmt.Sprintf("%s/transactions?%s",
		c.BaseURL,
		url.Values{
			"address":  {address},
			"contract": {contract},
		}.Encode()), build)
}

func (c *Client) getTxs(uri string, build string) (*Page, error) {
	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("client-build", build)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		logrus.WithError(err).Error("Ethereum/Trust Ray: Failed to get transactions")
		return nil, blockatlas.ErrSourceConn
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http %s (%s)", res.Status, uri)
	}

	txs := new(Page)
	err = json.NewDecoder(res.Body).Decode(txs)
	return txs, nil
}

func (c Client) GetCollections(owner string) ([]Collection, error) {
	uri := fmt.Sprintf("%s/api/v1/collections?%s",
		c.CollectionsURL,
		url.Values{
			"asset_owner": {owner},
			"limit":       {strconv.Itoa(1000)},
		}.Encode())
	req, _ := http.NewRequest("GET", uri, nil)
	//req.Header.Set("X-API-KEY", c.CollectionsApiKey)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var page []Collection
	err = json.NewDecoder(res.Body).Decode(&page)
	return page, err
}

func (c Client) GetCollectibles(owner string, contract string) ([]Collectible, error) {
	uri := fmt.Sprintf("%s/api/v1/assets/?%s",
		c.CollectionsURL,
		url.Values{
			"owner":                  {owner},
			"asset_contract_address": {contract},
			"limit":                  {strconv.Itoa(1000)},
		}.Encode())
	res, err := c.HTTPClient.Get(uri)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var page CollectiblePage
	err = json.NewDecoder(res.Body).Decode(&page)
	return page.Collectibles, err
}
