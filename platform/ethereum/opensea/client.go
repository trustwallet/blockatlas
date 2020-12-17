package opensea

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/golibs/client"
)

type Client struct {
	client.Request
}

func InitClient(api string, apiKey string) *Client {
	c := Client{internal.InitClient(api)}
	c.Headers["X-API-KEY"] = apiKey
	return &c
}

func (c Client) GetCollectionsByOwner(owner string) (page []Collection, err error) {
	query := url.Values{
		"asset_owner": {owner},
		"limit":       {"1000"},
	}
	err = c.Get(&page, "api/v1/collections", query)
	return
}

func (c Client) GetCollectiblesByCollectionId(owner string, collectionId string) ([]Collectible, error) {
	query := url.Values{
		"owner":      {owner},
		"collection": {collectionId},
		"limit":      {strconv.Itoa(300)},
	}

	var page CollectiblePage
	err := c.Get(&page, "api/v1/assets", query)
	return page.Collectibles, err
}
