package ethereum

import (
	"strings"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum/collection"
)

var (
	supportedTypes = map[string]bool{"ERC721": true, "ERC1155": true}
)

func (p *Platform) GetCollections(owner string) (blockatlas.CollectionPage, error) {
	collections, err := p.collectible.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	return NormalizeCollections(collections, p.CoinIndex, owner), nil
}

func (p *Platform) GetCollectibles(owner, collectibleID string) (blockatlas.CollectiblePage, error) {
	items, err := p.collectible.GetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	return NormalizeCollectiblePage(items, p.CoinIndex), nil
}

func NormalizeCollections(collections []collection.Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
	for _, collection := range collections {
		item := NormalizeCollection(collection, coinIndex, owner)
		page = append(page, item)
	}
	return page
}

func NormalizeCollection(c collection.Collection, coinIndex uint, owner string) blockatlas.Collection {
	return blockatlas.Collection{
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

func NormalizeCollectiblePage(collectibles []collection.Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
	for _, collectible := range collectibles {
		item := NormalizeCollectible(collectible, coinIndex)
		if _, ok := supportedTypes[item.Type]; ok {
			page = append(page, item)
		}
	}
	return page
}

func NormalizeCollectible(c collection.Collectible, coinIndex uint) blockatlas.Collectible {
	id := strings.Join([]string{c.AssetContract.Address, c.TokenId}, "-")
	return blockatlas.Collectible{
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
