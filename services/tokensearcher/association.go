package tokensearcher

import (
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

func assetsMap(txs blockatlas.Txs, coinID string) map[string][]models.Asset {
	result := make(map[string][]models.Asset)
	prefix := coinID + "_"
	for _, tx := range txs {
		addresses := tx.GetAddresses()
		asset, ok := tx.AssetModel()
		if !ok {
			continue
		}
		for _, a := range addresses {
			assetIDs := result[prefix+a]
			result[prefix+a] = append(assetIDs, asset)
		}
	}
	return result
}

func associationsToAdd(oldAssociations map[string][]models.Asset, newAssociations map[string][]models.Asset) map[string][]models.Asset {
	result := make(map[string][]models.Asset)
	for oldAddresses, oldAssets := range oldAssociations {
		for newAddresses, newAssets := range newAssociations {
			if strings.EqualFold(oldAddresses, newAddresses) {
				m := result[newAddresses]
				result[newAddresses] = append(m, newAssociationsForAddress(oldAssets, newAssets)...)
			}
		}
	}
	return result
}

func newAssociationsForAddress(oldAssociations []models.Asset, newAssociations []models.Asset) []models.Asset {
	var result []models.Asset
	oldM := make(map[string]bool)
	for _, o := range oldAssociations {
		oldM[o.Asset] = true
	}
	for _, n := range newAssociations {
		if ok := oldM[n.Asset]; !ok {
			result = append(result, n)
		}
	}
	return result
}

func fromModelToAssociation(associations []models.AddressToAssetAssociation) map[string][]models.Asset {
	result := make(map[string][]models.Asset)
	for _, a := range associations {
		m := result[a.Address.Address]
		result[a.Address.Address] = append(m, a.Asset)
	}
	return result
}
