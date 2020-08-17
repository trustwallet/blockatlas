package tokensearcher

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/watchmarket/pkg/watchmarket"
	"strconv"
	"strings"
)

func assetsMap(txs blockatlas.Txs) map[string][]string {
	if len(txs) == 0 {
		return nil
	}
	coin := txs[0].Coin
	result := make(map[string][]string)
	prefix := strconv.Itoa(int(coin)) + "_"
	for _, tx := range txs {
		addresses := tx.GetAddresses()
		tokenID, ok := tx.TokenID()
		if !ok {
			continue
		}
		assetID := watchmarket.BuildID(tx.Coin, tokenID)
		for _, a := range addresses {
			assetIDs := result[prefix+a]
			result[prefix+a] = append(assetIDs, assetID)
		}
	}
	return result
}

func associationsToAdd(associations map[string][]string, assetIDsMap map[string][]string) map[string][]string {
	result := make(map[string][]string)
	for addressFromAssociation, currentAssets := range associations {
		for addressFromTransactions, newAssets := range assetIDsMap {
			if strings.EqualFold(addressFromAssociation, addressFromTransactions) {
				m := result[addressFromTransactions]
				result[addressFromTransactions] = append(m, newAssociationsForAddress(currentAssets, newAssets)...)
			}
		}
	}
	return result
}

func newAssociationsForAddress(oldAssociations []string, assetIDs []string) []string {
	var result []string
	for _, o := range oldAssociations {
		for _, n := range assetIDs {
			if strings.EqualFold(o, n) {
				continue
			}
			result = append(result, n)
		}
	}
	return result
}
