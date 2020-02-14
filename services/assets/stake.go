package assets

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"sort"
	"time"
)

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func requestValidatorsInfo(coin coin.Coin) (AssetValidators, error) {
	var results AssetValidators
	request := blockatlas.InitClient(AssetsURL + coin.Handle)
	err := request.GetWithCache(&results, "validators/list.json", nil, time.Hour*1)
	if err != nil {
		return nil, errors.E(err, errors.Params{"coin": coin.Handle}).PushToSentry()
	}
	return results, nil
}

func GetValidatorsMap(api blockatlas.StakeAPI) (blockatlas.ValidatorMap, error) {
	assets, validators, err := GetValidators(api)
	if err != nil {
		return nil, err
	}
	results := normalizeValidators(assets, validators, api.Coin())
	return results.ToMap(), nil
}

func GetActiveValidators(api blockatlas.StakeAPI) (blockatlas.StakeValidators, error) {
	assets, validators, err := GetValidators(api)
	if err != nil {
		return nil, err
	}
	results := normalizeValidators(assets.activeValidators(), validators, api.Coin())
	return results, nil
}

func GetValidators(api blockatlas.StakeAPI) (AssetValidators, blockatlas.ValidatorPage, error) {
	assetsValidators, err := requestValidatorsInfo(api.Coin())
	if err != nil {
		return nil, nil, errors.E(err, "unable to fetch validators list from the registry").PushToSentry()
	}

	validators, err := api.GetValidators()
	if err != nil {
		return nil, nil, err
	}
	return assetsValidators, validators, nil
}

func normalizeValidators(assets AssetValidators, validators []blockatlas.Validator, coin coin.Coin) blockatlas.StakeValidators {
	results := make(blockatlas.StakeValidators, 0)
	assetsMap := assets.toMap()
	for _, v := range validators {
		asset, ok := assetsMap[v.ID]
		if !ok {
			continue
		}
		results = append(results, normalizeValidator(v, asset, coin))
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Details.Reward.Annual > results[j].Details.Reward.Annual
	})
	return results
}

func normalizeValidator(plainValidator blockatlas.Validator, validator AssetValidator, coin coin.Coin) blockatlas.StakeValidator {
	details := plainValidator.Details
	details.Reward.Annual = calculateAnnual(details.Reward.Annual, validator.Payout.Commission)
	return blockatlas.StakeValidator{
		ID:     validator.ID,
		Status: plainValidator.Status && !validator.Status.Disabled,
		Info: blockatlas.StakeValidatorInfo{
			Name:        validator.Name,
			Description: validator.Description,
			Image:       getImage(coin, plainValidator.ID),
			Website:     validator.Website,
		},
		Details: details,
	}
}
func calculateAnnual(annual float64, commission float64) float64 {
	return (annual * (100 - commission)) / 100
}

func getImage(c coin.Coin, ID string) string {
	return AssetsURL + c.Handle + "/validators/assets/" + ID + "/logo.png"
}
