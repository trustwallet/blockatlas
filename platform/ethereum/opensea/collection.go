package opensea

import (
	"strings"

	"github.com/trustwallet/golibs/types"
)

var (
	supportedTypes = map[string]bool{"ERC721": true, "ERC1155": true}
)

func (c *Client) GetCollections(owner string, coinIndex uint) (types.CollectionPage, error) {
	collections, err := c.GetCollectionsByOwner(owner)
	if err != nil {
		return nil, err
	}
	return NormalizeCollections(collections, coinIndex, owner), nil

}

func (c *Client) GetCollectibles(owner, collectionId string, coinIndex uint) (types.CollectiblePage, error) {
	items, err := c.GetCollectiblesByCollectionId(owner, collectionId)
	if err != nil {
		return nil, err
	}
	return NormalizeCollectiblePage(items, coinIndex), nil
}

func NormalizeCollections(collections []Collection, coinIndex uint, owner string) (page types.CollectionPage) {
	for _, collection := range collections {
		item := NormalizeCollection(collection, coinIndex, owner)
		page = append(page, item)
	}
	return page
}

func NormalizeCollection(c Collection, coinIndex uint, owner string) types.Collection {
	return types.Collection{
		Name:         c.Name,
		ImageUrl:     c.ImageUrl,
		Description:  c.Description,
		ExternalLink: c.ExternalUrl,
		Total:        int(c.Total.Int64()),
		Id:           c.Slug,
		Address:      owner,
		Coin:         coinIndex,
	}
}

func NormalizeCollectiblePage(collectibles []Collectible, coinIndex uint) (page types.CollectiblePage) {
	for _, collectible := range collectibles {
		item := NormalizeCollectible(collectible, coinIndex)
		if _, ok := supportedTypes[item.Type]; ok {
			page = append(page, item)
		}
	}
	return page
}

func NormalizeCollectible(c Collectible, coinIndex uint) types.Collectible {
	id := strings.Join([]string{c.AssetContract.Address, c.TokenId}, "-")
	return types.Collectible{
		ID:              id,
		CollectionID:    c.Collection.Slug,
		TokenID:         c.TokenId,
		ContractAddress: c.AssetContract.Address,
		Name:            c.Name,
		Category:        c.Collection.Name,
		ImageUrl:        c.ImagePreviewUrl,
		ProviderLink:    c.Permalink,
		ExternalLink:    c.Collection.ExternalLink,
		Type:            c.AssetContract.Type,
		Description:     c.Description,
		Coin:            coinIndex,
		Version:         c.AssetContract.Version,
	}
}
