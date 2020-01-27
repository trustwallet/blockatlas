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

func GetCoinInfo(coinId int) (info blockatlas.CoinInfo, err error) {
	c, ok := coin.Coins[uint(coinId)]
	if !ok {
		return info, errors.E("coin not found")
	}
	request := blockatlas.InitClient(AssetsURL + c.Handle)
	err = request.GetWithCache(&info, "info/info.json", nil, time.Hour*1)
	return
}

func GetValidatorsMap(api blockatlas.StakeAPI) (blockatlas.ValidatorMap, error) {
	validatorList, err := GetValidators(api)
	if err != nil {
		return nil, err
	}
	validators := make(blockatlas.ValidatorMap)
	for _, v := range validatorList {
		validators[v.ID] = v
	}
	return validators, nil
}

func GetValidators(api blockatlas.StakeAPI) ([]blockatlas.StakeValidator, error) {
	assetsValidators, err := getValidatorsInfo(api.Coin())
	if err != nil {
		return nil, errors.E(err, "unable to fetch validators list from the registry").PushToSentry()
	}

	validators, err := api.GetValidators()
	if err != nil {
		return nil, err
	}
	results := normalizeValidators(validators, assetsValidators, api.Coin())
	sort.Slice(results, func(i, j int) bool {
		if results[i].Details.FreeSpace < 1 && results[j].Details.FreeSpace > 1 {
			return false
		}
		if results[j].Details.FreeSpace < 1 && results[i].Details.FreeSpace > 1 {
			return true
		}
		return results[i].Details.Reward.Annual > results[j].Details.Reward.Annual
	})
	return results, nil
}

func getValidatorsInfo(coin coin.Coin) ([]AssetValidator, error) {
	var results []AssetValidator
	request := blockatlas.InitClient(AssetsURL + coin.Handle)
	err := request.GetWithCache(&results, "validators/list.json", nil, time.Hour*1)
	if err != nil {
		return nil, errors.E(err, errors.Params{"coin": coin.Handle}).PushToSentry()
	}
	return results, nil
}

func normalizeValidators(validators []blockatlas.Validator, assets []AssetValidator, coin coin.Coin) []blockatlas.StakeValidator {
	results := make([]blockatlas.StakeValidator, 0)
	for _, v := range validators {
		for _, v2 := range assets {
			if v.ID == v2.ID && !v2.Status.Disabled {
				results = append(results, normalizeValidator(v, v2, coin))
			}
		}
	}
	return results
}

func normalizeValidator(plainValidator blockatlas.Validator, validator AssetValidator, coin coin.Coin) blockatlas.StakeValidator {
	details := plainValidator.Details
	details.Reward.Annual = calculateAnnual(details.Reward.Annual, validator.Payout.Commission)
	return blockatlas.StakeValidator{
		ID:     validator.ID,
		Status: plainValidator.Status,
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
