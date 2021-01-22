package blockatlas

import "github.com/trustwallet/golibs/types"

type (
	ResultsResponse struct {
		Results interface{} `json:"docs"`
	}
)

func GetAssetsIds(assets types.TokenPage) []string {
	assetIds := make([]string, 0)
	for _, asset := range assets {
		assetIds = append(assetIds, asset.AssetId())
	}
	return assetIds
}
