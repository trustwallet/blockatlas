package tokenindexer

import (
	"time"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"github.com/trustwallet/golibs/types"
)

type Instance struct {
	database *db.Instance
}

func Init(database *db.Instance) Instance {
	return Instance{database: database}
}

func (i Instance) GetNewTokensRequest(r Request) (blockatlas.ResultsResponse, error) {
	from := time.Unix(r.From, 0)
	result, err := i.database.GetAssetsFrom(from)
	if err != nil {
		return blockatlas.ResultsResponse{}, err
	}
	return normalize(result), nil
}

func (i Instance) GetTokensByAddress(r GetTokensByAddressRequest) (GetTokensByAddressResponse, error) {
	list := make([]string, 0)

	for coin, coins := range r.AddressesByCoin {
		for _, address := range coins {
			list = append(list, types.GetAddressID(coin, address))
		}
	}
	from := time.Unix(int64(r.From), 0)
	associations, err := i.database.GetSubscriptionsByAddressIDs(list, from)
	if err != nil {
		return GetTokensByAddressResponse{}, err
	}

	assetIds := make([]GetTokensAsset, 0)

	for _, association := range associations {
		assetIds = append(assetIds, GetTokensAsset{
			AssetId:   association.Asset.Asset,
			CreatedAt: association.CreatedAt.Unix(),
			UpdatedAt: association.UpdatedAt.Unix(),
		})
	}

	return assetIds, nil
}

func normalize(dbAssets []models.Asset) blockatlas.ResultsResponse {
	result := make([]types.Asset, 0)
	for _, a := range dbAssets {
		asset := types.Asset{
			Id:       a.Asset,
			Name:     a.Name,
			Symbol:   a.Symbol,
			Type:     types.TokenType(a.Type),
			Decimals: a.Decimals,
		}
		result = append(result, asset)
	}
	return blockatlas.ResultsResponse{Results: result}
}
