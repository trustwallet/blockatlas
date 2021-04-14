package bounce

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/types"
)

var (
	nftVersion = "3.0" // opensea nft_version compatible
)

func (c *Client) GetCollections(owner string, coinIndex uint) (types.CollectionPage, error) {
	collections, err := c.getCollections(owner)
	if err != nil {
		return nil, err
	}
	return c.processCollections(collections, coinIndex, owner)

}

func (c *Client) GetCollectibles(owner, collectionID string, coinIndex uint) (types.CollectiblePage, error) {
	collectibles, err := c.getCollectibles(owner, collectionID)
	if err != nil {
		return nil, err
	}
	return c.processCollectibles(collectibles, coinIndex)
}

func (c *Client) processCollections(collections []Collection, coinIndex uint, owner string) (types.CollectionPage, error) {
	page := make(types.CollectionPage, 0)
	categories := map[string]*types.Collection{}

	for _, cl := range collections {

		// skip invalid balance
		total, err := strconv.Atoi(cl.Balance)
		if err != nil {
			continue
		}

		// udpate existed balance
		existed, ok := categories[cl.ContractAddr]
		if ok {
			existed.Total = existed.Total + total
			continue
		}

		// skip empty info token
		if len(cl.TokenURI) == 0 {
			continue
		}

		// fetch token info
		info, err := fetchTokenURI(cl.TokenURI)
		if err != nil {
			return nil, err
		}

		// skip empty name/image
		if len(info.Name) == 0 || len(info.Image) == 0 {
			continue
		}

		contractName := cl.ContractName
		if len(contractName) == 0 {
			contractName = info.Name
		}

		categories[cl.ContractAddr] = &types.Collection{
			Id:           cl.ContractAddr,
			Name:         contractName,
			ImageUrl:     normalizeUrl(info.Image),
			Description:  info.Description,
			ExternalLink: normalizeUrl(cl.TokenURI),
			Total:        total,
			Address:      owner,
			Coin:         coinIndex,
			Type:         "ERC" + cl.TokenType,
		}
	}

	for _, c := range categories {
		page = append(page, *c)
	}
	return page, nil
}

func (c *Client) processCollectibles(collectibles []Collectible, coinIndex uint) (types.CollectiblePage, error) {
	if len(collectibles) == 0 {
		return types.CollectiblePage{}, nil
	}
	page := make(types.CollectiblePage, 0)
	for _, c := range collectibles {
		info, err := fetchTokenURI(c.TokenURI)
		if err != nil {
			return nil, err
		}
		normalized := normalizeCollectible(c, coinIndex, info)
		page = append(page, normalized)
	}
	return page, nil
}

func normalizeCollectible(c Collectible, coinIndex uint, info CollectionInfo) types.Collectible {
	category := c.ContractName
	if len(category) == 0 {
		category = info.Name
	}
	return types.Collectible{
		ID:              blockatlas.GenCollectibleId(c.ContractAddr, c.TokenID),
		CollectionID:    c.ContractAddr,
		TokenID:         c.TokenID,
		ContractAddress: c.ContractAddr,
		Category:        category,
		ImageUrl:        normalizeUrl(info.Image),
		ExternalLink:    normalizeUrl(c.TokenURI),
		Type:            string(types.ERC721),
		Description:     info.Description,
		Coin:            coinIndex,
		Name:            info.Name,
		Version:         nftVersion,
	}
}
