package collection

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c Client) GetCollections(owner string) (page []Collection, err error) {
	query := url.Values{
		"asset_owner": {owner},
		"limit":       {"1000"},
	}
	err = c.Get(&page, "api/v1/collections", query)
	return
}

func (c Client) GetCollectibles(owner string, collectibleID string) ([]Collectible, error) {
	query := url.Values{
		"owner":      {owner},
		"collection": {collectibleID},
		"limit":      {strconv.Itoa(300)},
	}

	var page CollectiblePage
	err := c.Get(&page, "api/v1/assets", query)
	return page.Collectibles, err
}
