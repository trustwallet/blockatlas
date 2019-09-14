package ethereum

import (
	"encoding/json"
	"errors"
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

func (c CollectionsClient) GetCollectibles(owner string, collectibleID string) (*Collection, []Collectible, error) {
	collections, err := c.GetCollections(owner)
	if err != nil {
		return nil, nil, err
	}
	id := getCollectionId(collectibleID)
	collection := searchCollection(collections, id)
	if collection == nil {
		return nil, nil, errors.New(fmt.Sprintf("%s not found", collectibleID))
	}

	uriValues := url.Values{
		"owner": {owner},
		"limit": {strconv.Itoa(300)},
	}

	for _, i := range collection.Contracts {
		if _, ok := slugTokens[i.Type]; ok {
			uriValues.Set("collection", collection.Slug)
			break
		}
		uriValues.Add("asset_contract_addresses", i.Address)
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
