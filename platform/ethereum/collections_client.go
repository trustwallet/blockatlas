package ethereum

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

func (c CollectionsClient) GetCollectibles(owner string, collectibleID string) (*Collection, []Collectible, error) {
	collections, err := c.GetCollections(owner)
	if err != nil {
		return nil, nil, err
	}
	collection := searchCollection(&collections, collectibleID)
	if collection == nil {
		return nil, nil, errors.New(fmt.Sprintf("%s not found", collectibleID))
	}

	uriValues := url.Values{
		"owner": {owner},
		"limit": {strconv.Itoa(300)},
	}
	for _, i := range collection.Contracts {
		uriValues.Add("asset_contract_addresses", i.Address)
	}
	if len(collection.Contracts) == 0 {
		uriValues.Set("collection", collection.Slug)
	}
	uri := fmt.Sprintf("%s/api/v1/assets/?%s",
		c.CollectionsURL,
		uriValues.Encode())

	req, _ := http.NewRequest("GET", uri, nil)
	req.Header.Set("X-API-KEY", c.CollectionsApiKey)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	var page CollectiblePage
	err = json.NewDecoder(res.Body).Decode(&page)
	return collection, page.Collectibles, err
}

func searchCollection(collections *[]Collection, collectibleID string) *Collection {
	for _, i := range *collections {
		if strings.EqualFold(i.Slug, collectibleID) {
			return &i
		}
		for _, contract := range i.Contracts {
			if strings.EqualFold(contract.Address, collectibleID) {
				return &i
			}
		}
	}
	return nil
}
