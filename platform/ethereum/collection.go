package ethereum

import (
	"fmt"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

var (
	supportedTypes = map[string]bool{"ERC721": true, "ERC1155": true}
	slugTokens     = map[string]bool{"ERC1155": true}
)

func (p *Platform) GetCollections(owner string) (blockatlas.CollectionPage, error) {
	collections, err := p.collectionsClient.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectionPage(collections, p.CoinIndex, owner)
	return page, nil
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func (p *Platform) OldGetCollections(owner string) (blockatlas.CollectionPage, error) {
	collections, err := p.collectionsClient.GetCollections(owner)
	if err != nil {
		return nil, err
	}
	page := OldNormalizeCollectionPage(collections, p.CoinIndex, owner)
	return page, nil
}

func (p *Platform) GetCollectibles(owner, collectibleID string) (blockatlas.CollectiblePage, error) {
	collection, items, err := p.collectionsClient.GetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	page := NormalizeCollectiblePage(collection, items, p.CoinIndex)
	return page, nil
}

func NormalizeCollectionPage(collections []Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
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

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func (p *Platform) OldGetCollectibles(owner, collectibleID string) (blockatlas.CollectiblePage, error) {
	collection, items, err := p.collectionsClient.OldGetCollectibles(owner, collectibleID)
	if err != nil {
		return nil, err
	}
	page := OldNormalizeCollectiblePage(collection, items, p.CoinIndex)
	return page, nil
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func OldNormalizeCollectionPage(collections []Collection, coinIndex uint, owner string) (page blockatlas.CollectionPage) {
	for _, collection := range collections {
		if len(collection.Contracts) == 0 {
			continue
		}
		item := OldNormalizeCollection(collection, coinIndex, owner)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func OldNormalizeCollection(c Collection, coinIndex uint, owner string) blockatlas.Collection {
	if len(c.Contracts) == 0 {
		return blockatlas.Collection{}
	}

	description := blockatlas.GetValidParameter(c.Description, c.Contracts[0].Description)
	symbol := blockatlas.GetValidParameter(c.Contracts[0].Symbol, "")
	collectionId := blockatlas.GetValidParameter(c.Contracts[0].Address, "")
	version := blockatlas.GetValidParameter(c.Contracts[0].NftVersion, "")
	collectionType := blockatlas.GetValidParameter(c.Contracts[0].Type, "")
	if _, ok := slugTokens[collectionType]; ok {
		collectionId = createCollectionId(collectionId, c.Slug)
	}

	return blockatlas.Collection{
		Name:            c.Name,
		Symbol:          symbol,
		Slug:            c.Slug,
		ImageUrl:        c.ImageUrl,
		Description:     description,
		ExternalLink:    c.ExternalUrl,
		Total:           int(c.Total.Int64()),
		Id:              collectionId,
		CategoryAddress: collectionId,
		Address:         owner,
		Version:         version,
		Coin:            coinIndex,
		Type:            collectionType,
	}
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func OldNormalizeCollectible(c *Collection, a Collectible, coinIndex uint) blockatlas.Collectible {
	address := blockatlas.GetValidParameter(c.Contracts[0].Address, "")
	collectionType := blockatlas.GetValidParameter(c.Contracts[0].Type, "")
	collectionID := address
	if _, ok := slugTokens[collectionType]; ok {
		collectionID = createCollectionId(address, c.Slug)
	}
	externalLink := blockatlas.GetValidParameter(a.ExternalLink, a.AssetContract.ExternalLink)
	id := strings.Join([]string{a.AssetContract.Address, a.TokenId}, "-")
	return blockatlas.Collectible{
		ID:               id,
		CollectionID:     collectionID,
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

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func createCollectionId(address, slug string) string {
	return fmt.Sprintf("%s---%s", address, slug)
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func getCollectionId(collectionId string) string {
	s := strings.Split(collectionId, "---")
	if len(s) != 2 {
		return collectionId
	}
	return s[1]
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func OldNormalizeCollectiblePage(c *Collection, srcPage []Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
	if len(c.Contracts) == 0 {
		return
	}
	for _, src := range srcPage {
		item := OldNormalizeCollectible(c, src, coinIndex)
		if _, ok := supportedTypes[item.Type]; !ok {
			continue
		}
		page = append(page, item)
	}
	return
}

func NormalizeCollection(c Collection, coinIndex uint, owner string) blockatlas.Collection {
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

func normalizeSupportedContracts(c *Collection) {
	supportedContracts := make([]PrimaryAssetContract, 0)
	for _, contract := range c.Contracts {
		if _, ok := supportedTypes[contract.Type]; !ok {
			continue
		}
		supportedContracts = append(supportedContracts, contract)
	}
	c.Contracts = supportedContracts
}

func NormalizeCollectiblePage(c *Collection, srcPage []Collectible, coinIndex uint) (page blockatlas.CollectiblePage) {
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

func NormalizeCollectible(c *Collection, a Collectible, coinIndex uint) blockatlas.Collectible {
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

func searchCollection(collections []Collection, collectibleID string) *Collection {
	for _, i := range collections {
		if strings.EqualFold(i.Slug, collectibleID) {
			return &i
		}
	}
	return nil
}

//TODO: remove once most of the clients will be updated (deadline: March 17th)
func oldSearchCollection(collections []Collection, collectibleID string) *Collection {
	for _, i := range collections {
		if strings.EqualFold(i.Slug, collectibleID) {
			return &i
		}
		for _, contract := range i.Contracts {
			if strings.EqualFold(contract.Address, collectibleID) {
				return &i
			}
		}
	}
	return nil
}
