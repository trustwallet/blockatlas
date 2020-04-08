package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/platform/ethereum/collection"
	"strings"
)

func (p *Platform) GetCollectionsV3(owner string) (blockatlas.CollectionPageV3, error) {
	collections, err := p.collectible.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectionPageV3(collections, p.CoinIndex, owner)
	return page, nil
}

func (p *Platform) GetCollectiblesV3(owner, collectibleID string) (blockatlas.CollectiblePageV3, error) {
	collection, items, err := p.collectible.GetCollectiblesV3(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectiblePageV3(collection, items, p.CoinIndex)
	return page, nil
}

func NormalizeCollectionPageV3(collections []collection.Collection, coinIndex uint, owner string) (page blockatlas.CollectionPageV3) {
	for _, collection := range collections {
		if len(collection.Contracts) == 0 {
			continue
		}
		item := NormalizeCollectionV3(collection, coinIndex, owner)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollectionV3(c collection.Collection, coinIndex uint, owner string) blockatlas.CollectionV3 {
	normalizeSupportedContracts(&c)
	if len(c.Contracts) == 0 {
		return blockatlas.CollectionV3{}
	}

	description := blockatlas.GetValidParameter(c.Description, c.Contracts[0].Description)
	symbol := blockatlas.GetValidParameter(c.Contracts[0].Symbol, "")
	version := blockatlas.GetValidParameter(c.Contracts[0].NftVersion, "")
	collectionType := blockatlas.GetValidParameter(c.Contracts[0].Type, "")

	return blockatlas.CollectionV3{
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

func NormalizeCollectiblePageV3(c *collection.Collection, srcPage []collection.Collectible, coinIndex uint) (page blockatlas.CollectiblePageV3) {
	normalizeSupportedContracts(c)
	if len(c.Contracts) == 0 {
		return
	}
	for _, src := range srcPage {
		item := NormalizeCollectibleV3(c, src, coinIndex)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollectibleV3(c *collection.Collection, a collection.Collectible, coinIndex uint) blockatlas.CollectibleV3 {
	address := blockatlas.GetValidParameter(c.Contracts[0].Address, "")
	collectionType := blockatlas.GetValidParameter(c.Contracts[0].Type, "")
	externalLink := blockatlas.GetValidParameter(a.ExternalLink, a.AssetContract.ExternalLink)
	id := strings.Join([]string{a.AssetContract.Address, a.TokenId}, "-")
	return blockatlas.CollectibleV3{
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
