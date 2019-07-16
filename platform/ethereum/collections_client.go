package ethereum

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CollectionsClient struct {
	HTTPClient        *http.Client
	CollectionsURL    string
	CollectionsApiKey string
}

func (c CollectionsClient) GetCollections(owner string) ([]Collection, error) {
	uri := fmt.Sprintf("%s/api/v1/collections?%s",
		c.CollectionsURL,
		url.Values{
			"asset_owner": {owner},
			"limit":       {strconv.Itoa(1000)},
		}.Encode())

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("X-API-KEY", c.CollectionsApiKey)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var page []Collection
	err = json.NewDecoder(res.Body).Decode(&page)
	return page, err
}

func (c CollectionsClient) GetCollectibles(owner string, collectibleID string) ([]Collectible, error) {
	uri := fmt.Sprintf("%s/api/v1/assets/?%s",
		c.CollectionsURL,
		url.Values{
			"owner":                  {owner},
			"asset_contract_address": {collectibleID},
			"limit":                  {strconv.Itoa(1000)},
		}.Encode())

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("X-API-KEY", c.CollectionsApiKey)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var page CollectiblePage
	err = json.NewDecoder(res.Body).Decode(&page)
	return page.Collectibles, err
}
