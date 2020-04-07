package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

var (
	supportedTypes = map[string]bool{"ERC721": true, "ERC1155": true}
)

func (p *Platform) GetCollections(owner string) (blockatlas.CollectionPage, error) {
	collections, err := p.collectionsClient.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	return NormalizeCollections(collections, p.CoinIndex, owner), nil
}

func (p *Platform) GetCollectibles(owner, collectibleID string) (blockatlas.CollectiblePage, error) {
	items, err := p.collectionsClient.GetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	return NormalizeCollectiblePage(items, p.CoinIndex), nil
}

func NormalizeCollections(collections []Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
	for _, collection := range collections {
		item := NormalizeCollection(collection, coinIndex, owner)
		page = append(page, item)
	}
	return page
}

func NormalizeCollection(c Collection, coinIndex uint, owner string) blockatlas.Collection {
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

func NormalizeCollectiblePage(collectibles []Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
	for _, collectible := range collectibles {
		item := NormalizeCollectible(collectible, coinIndex)
		if _, ok := supportedTypes[item.Type]; ok {
			page = append(page, item)
		}
	}
	return page
}

func NormalizeCollectible(a Collectible, coinIndex uint) blockatlas.Collectible {
	id := strings.Join([]string{a.AssetContract.Address, a.TokenId}, "-")
	return blockatlas.Collectible{
		ID:              id,
		CollectionID:    a.Collection.Slug,
		TokenID:         a.TokenId,
		ContractAddress: a.AssetContract.Address,
		Name:            a.Name,
		Category:        a.Collection.Name,
		ImageUrl:        a.ImagePreviewUrl,
		ProviderLink:    a.Permalink,
		ExternalLink:    a.Collection.ExternalLink,
		Type:            a.AssetContract.Type,
		Description:     a.Description,
		Coin:            coinIndex,
		Version:         a.AssetContract.Version,
	}
}
