package collection

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
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

func (c Client) GetCollectiblesV3(owner string, collectibleID string) (*Collection, []Collectible, error) {
	collections, err := c.GetCollections(owner)
	if err != nil {
		return nil, nil, err
	}
	collection := SearchCollection(collections, collectibleID)
	if collection == nil {
		return nil, nil, errors.E("collectible not found", errors.TypePlatformClient,
			errors.Params{"collectibleID": collectibleID})
	}

	query := url.Values{
		"owner": {owner},
		"limit": {strconv.Itoa(300)},
	}

	query.Set("collection", collection.Slug)

	var page CollectiblePage
	err = c.Get(&page, "api/v1/assets", query)
	return collection, page.Collectibles, err
}

func SearchCollection(collections []Collection, collectibleID string) *Collection {
	for _, i := range collections {
		if strings.EqualFold(i.Slug, collectibleID) {
			return &i
		}
	}
	return nil
}
