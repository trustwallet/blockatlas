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

func associationsToAdd(oldAssociations map[string][]string, newAssociations map[string][]string) map[string][]string {
	result := make(map[string][]string)
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

func newAssociationsForAddress(oldAssociations []string, newAssociations []string) []string {
	var result []string
	oldM := make(map[string]bool)
	for _, o := range oldAssociations {
		oldM[o] = true
	}
	for _, n := range newAssociations {
		if ok := oldM[n]; !ok {
			result = append(result, n)
		}
	}
	return result
}
