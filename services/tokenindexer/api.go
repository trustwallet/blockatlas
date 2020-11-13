package tokenindexer

import (
	"context"
	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
	"time"
)

type Instance struct {
	database *db.Instance
}

func Init(database *db.Instance) Instance {
	return Instance{database: database}
}

func (i Instance) HandleNewTokensRequest(r Request, ctx context.Context) (Response, error) {
	from := time.Unix(r.From, 0)
	result, err := i.database.GetAssetsFrom(from, r.Coin, ctx)
	if err != nil {
		return Response{}, err
	}
	return normalize(result), nil
}

func normalize(dbAssets []models.Asset) Response {
	var result []Asset
	for _, a := range dbAssets {
		asset := Asset{
			Asset:    a.Asset,
			Name:     a.Name,
			Symbol:   a.Symbol,
			Type:     a.Type,
			Decimals: a.Decimals,
		}
		result = append(result, asset)
	}
	return Response{Assets: result}
}
