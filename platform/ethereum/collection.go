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
	page := NormalizeCollectionPage(collections, p.CoinIndex, owner)
	return page, nil
}

func (p *Platform) GetCollectibles(owner, collectibleID string) (blockatlas.CollectiblePage, error) {
	collection, items, err := p.collectible.GetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectiblePage(collection, items, p.CoinIndex)
	return page, nil
}

func NormalizeCollectionPage(collections []collection.Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
	for _, collection := range collections {
		if len(collection.Contracts) == 0 {
			continue
		}
		item := NormalizeCollection(collection, coinIndex, owner)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollection(c collection.Collection, coinIndex uint, owner string) blockatlas.Collection {
	normalizeSupportedContracts(&c)
	if len(c.Contracts) == 0 {
		return blockatlas.Collection{}
	}

	description := blockatlas.GetValidParameter(c.Description, c.Contracts[0].Description)
	symbol := blockatlas.GetValidParameter(c.Contracts[0].Symbol, "")
	version := blockatlas.GetValidParameter(c.Contracts[0].NftVersion, "")
	collectionType := blockatlas.GetValidParameter(c.Contracts[0].Type, "")

	return blockatlas.Collection{
		Name:            c.Name,
		Symbol:          symbol,
		Slug:            c.Slug,
		ImageUrl:        c.ImageUrl,
		Description:     description,
		ExternalLink:    c.ExternalUrl,
		Total:           int(c.Total.Int64()),
		Id:              c.Slug,
		CategoryAddress: c.Slug,
		Address:         owner,
		Version:         version,
		Coin:            coinIndex,
		Type:            collectionType,
	}
}

func normalizeSupportedContracts(c *collection.Collection) {
	supportedContracts := make([]collection.PrimaryAssetContract, 0)
	for _, contract := range c.Contracts {
		if _, ok := supportedTypes[contract.Type]; !ok {
			continue
		}
		supportedContracts = append(supportedContracts, contract)
	}
	c.Contracts = supportedContracts
}

func NormalizeCollectiblePage(c *collection.Collection, srcPage []collection.Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
	normalizeSupportedContracts(c)
	if len(c.Contracts) == 0 {
		return
	}
	for _, src := range srcPage {
		item := NormalizeCollectible(c, src, coinIndex)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollectible(c *collection.Collection, a collection.Collectible, coinIndex uint) blockatlas.Collectible {
	address := blockatlas.GetValidParameter(c.Contracts[0].Address, "")
	collectionType := blockatlas.GetValidParameter(c.Contracts[0].Type, "")
	externalLink := blockatlas.GetValidParameter(a.ExternalLink, a.AssetContract.ExternalLink)
	id := strings.Join([]string{a.AssetContract.Address, a.TokenId}, "-")
	return blockatlas.Collectible{
		ID:               id,
		CollectionID:     c.Slug,
		ContractAddress:  address,
		TokenID:          a.TokenId,
		CategoryContract: a.AssetContract.Address,
		Name:             a.Name,
		Category:         c.Name,
		ImageUrl:         a.ImagePreviewUrl,
		ProviderLink:     a.Permalink,
		ExternalLink:     externalLink,
		Type:             collectionType,
		Description:      a.Description,
		Coin:             coinIndex,
	}
}
