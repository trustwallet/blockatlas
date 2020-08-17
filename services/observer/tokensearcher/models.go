package tokensearcher

import "github.com/trustwallet/blockatlas/db/models"

func fromModelToAssociation(associations []models.AddressToTokenAssociation) map[string][]string {
	result := make(map[string][]string)
	for _, a := range associations {
		m := result[a.Address.Address]
		result[a.Address.Address] = append(m, a.Asset.AssetID)
	}
	return result
}
