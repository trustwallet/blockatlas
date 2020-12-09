package tokenindexer

import (
	"time"

	"github.com/trustwallet/blockatlas/db"
	"github.com/trustwallet/blockatlas/db/models"
)

type Instance struct {
	database *db.Instance
}

func Init(database *db.Instance) Instance {
	return Instance{database: database}
}

func (i Instance) HandleNewTokensRequest(r Request) (Response, error) {
	from := time.Unix(r.From, 0)
	result, err := i.database.GetAssetsFrom(from, r.Coin)
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
