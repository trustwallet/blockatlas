package assets

import (
	"sort"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/numbers"
)

func GetValidatorsMap(api blockatlas.StakeAPI) (blockatlas.ValidatorMap, error) {
	assets, validators, err := getValidators(api)
	if err != nil {
		return nil, err
	}
	results := normalizeValidators(assets, validators, api.Coin())
	return results.ToMap(), nil
}

func getValidators(api blockatlas.StakeAPI) (AssetValidators, blockatlas.ValidatorPage, error) {
	assetsValidators, err := fetchValidatorsInfo(api.Coin())
	if err != nil {
		return nil, nil, err
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
