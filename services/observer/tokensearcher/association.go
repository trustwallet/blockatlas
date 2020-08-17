package tokensearcher

import (
	"context"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"strings"
)

func allAssociationsToAdd(database *db.Instance, associations []models.AddressToTokenAssociation, txs map[string]blockatlas.Txs) []models.AddressToTokenAssociation {
	var result []models.AddressToTokenAssociation
	for _, a := range associations {
		for address, tx := range txs {
			if strings.EqualFold(a.Address, address) {
				result = append(result, addAssociationsForAddress(database, a, tx)...)
			}
		}
	}
	return result
}

func addAssociationsForAddress(database *db.Instance, association models.AddressToTokenAssociation, txs blockatlas.Txs) []models.AddressToTokenAssociation {
	var result []models.AddressToTokenAssociation
	for _, tx := range txs {
		if isTokenAlreadyAssociated(association, tx) == 2 {
			t, _, _, ok := getInfoByMeta(tx)
			if !ok {
				continue
			}
			a, err := addAssociationForAddress(database, association.Address, t, uint(tx.Block))
			if err != nil {
				continue
			}
			result = append(result, a)
		}
	}
	return result
}

func isTokenAlreadyAssociated(association models.AddressToTokenAssociation, tx blockatlas.Tx) int {
	token, _, _, ok := getInfoByMeta(tx)
	if !ok {
		return 0
	}
	if strings.EqualFold(token.TokenID, association.Token.TokenID) {
		return 1
	}
	return 2
}

func addAssociationForAddress(database *db.Instance, address string, token models.Token, blockNumber uint) (models.AddressToTokenAssociation, error) {
	token, err := database.GetTokenByTokenID(token.TokenID, context.TODO())
	if err != nil {
		err := database.AddToken(token, context.TODO())
		if err != nil {
			return models.AddressToTokenAssociation{}, err
		}
		token, err = database.GetTokenByTokenID(token.TokenID, context.TODO())
		if err != nil {
			return models.AddressToTokenAssociation{}, err
		}
	}
	association := models.AddressToTokenAssociation{
		Address:          address,
		TokenID:          token.ID,
		LastUpdatedBlock: blockNumber,
	}
	return association, nil
}
