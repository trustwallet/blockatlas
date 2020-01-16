package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"net/url"
	"strconv"
)

type CollectionsClient struct {
	blockatlas.Request
}

func (c CollectionsClient) GetCollections(owner string) (page []Collection, err error) {
	query := url.Values{
		"asset_owner": {owner},
		"limit":       {"1000"},
	}
	err = c.Get(&page, "api/v1/collections", query)
	return
}

func (c CollectionsClient) GetCollectibles(owner string, collectibleID string) (*Collection, []Collectible, error) {
	collections, err := c.GetCollections(owner)
	if err != nil {
		return nil, nil, err
	}
	id := getCollectionId(collectibleID)
	collection := searchCollection(collections, id)
	if collection == nil {
		return nil, nil, errors.E("collectible not found", errors.TypePlatformClient,
			errors.Params{"collectibleID": collectibleID}).PushToSentry()
	}

	query := url.Values{
		"owner": {owner},
		"limit": {strconv.Itoa(300)},
	}

	for _, i := range collection.Contracts {
		if _, ok := slugTokens[i.Type]; ok {
			query.Set("collection", collection.Slug)
			break
		}
		query.Add("asset_contract_addresses", i.Address)
	}

	var page CollectiblePage
	err = c.Get(&page, "api/v1/assets", query)
	return collection, page.Collectibles, err
}
