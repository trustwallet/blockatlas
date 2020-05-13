package assets

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/numbers"
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
		return nil, errors.E(err, errors.Params{"coin": coin.Handle})
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

// Get validators from assets repository and RPC
func GetValidators(api blockatlas.StakeAPI) (AssetValidators, blockatlas.ValidatorPage, error) {
	assetsValidators, err := requestValidatorsInfo(api.Coin())
	if err != nil {
		return nil, nil, errors.E(err, "unable to fetch validators list from the registry")
	}

	validators, err := api.GetValidators()
	if err != nil {
		return nil, nil, err
	}
	return assetsValidators, validators, nil
}

func normalizeValidators(assetsValidators AssetValidators, rpcValidators []blockatlas.Validator, coin coin.Coin) blockatlas.StakeValidators {
	results := make(blockatlas.StakeValidators, 0)
	assetsMap := assetsValidators.toMap()
	for _, v := range rpcValidators {
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

func normalizeValidator(rpcValidator blockatlas.Validator, assetValidator AssetValidator, coin coin.Coin) blockatlas.StakeValidator {
	details := rpcValidator.Details
	details.MinimumAmount = blockatlas.Amount(numbers.Float64toString(assetValidator.Staking.MinDelegation))
	details.Reward.Annual = calculateAnnual(details.Reward.Annual, assetValidator.Payout.Commission)

	return blockatlas.StakeValidator{
		ID:     assetValidator.ID,
		Status: rpcValidator.Status && !assetValidator.Status.Disabled,
		Info: blockatlas.StakeValidatorInfo{
			Name:        assetValidator.Name,
			Description: assetValidator.Description,
			Image:       getImage(coin, rpcValidator.ID),
			Website:     assetValidator.Website,
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
