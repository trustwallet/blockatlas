package ethereum

import (
	"fmt"
	"strings"
)

func getValidParameter(first, second string) string {
	if len(first) > 0 {
		return first
	}
	return second
}

func createCollectionId(address, slug string) string {
	return fmt.Sprintf("%s---%s", address, slug)
}

func getCollectionId(collectionId string) string {
	s := strings.Split(collectionId, "---")
	if len(s) != 2 {
		return collectionId
	}
	return s[1]
}

func searchCollection(collections []Collection, collectibleID string) *Collection {
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
