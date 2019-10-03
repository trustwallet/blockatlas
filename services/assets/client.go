package assets

import (
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"strings"
)

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func GetValidators(api blockatlas.StakeAPI) ([]blockatlas.StakeValidator, error) {
	assetsValidators, err := GetValidatorsInfo(api.Coin())
	if err != nil {
		return nil, errors.E(err, "unable to fetch validators list from the registry")
	}

	validators, err := api.GetValidators()
	if err != nil {
		return nil, err
	}
	results := NormalizeValidators(validators, assetsValidators, api.Coin())
	return results, nil
}

func GetValidatorsInfo(coin coin.Coin) ([]AssetValidator, error) {
	var results []AssetValidator
	request := blockatlas.Request{
		BaseUrl:      AssetsURL + coin.Handle,
		HttpClient:   blockatlas.DefaultClient,
		ErrorHandler: blockatlas.DefaultErrorHandler,
	}
	err := request.Get(&results, "validators/list.json", nil)
	if err != nil {
		return nil, errors.E(err, errors.Params{"coin": coin.Handle})
	}
	return results, nil
}

func NormalizeValidators(validators []blockatlas.Validator, assets []AssetValidator, coin coin.Coin) []blockatlas.StakeValidator {
	results := make([]blockatlas.StakeValidator, 0)
	for _, v := range validators {
		for _, v2 := range assets {
			if v.ID == v2.ID {
				results = append(results, NormalizeValidator(v, v2, coin))
			}
		}
	}
	return results
}

func NormalizeValidator(plainValidator blockatlas.Validator, validator AssetValidator, coin coin.Coin) blockatlas.StakeValidator {
	return blockatlas.StakeValidator{
		ID:     validator.ID,
		Status: plainValidator.Status,
		Info: blockatlas.StakeValidatorInfo{
			Name:        validator.Name,
			Description: validator.Description,
			Image:       GetImage(coin, plainValidator.ID),
			Website:     validator.Website,
		},
		Reward:        plainValidator.Reward,
		LockTime:      plainValidator.LockTime,
		MinimumAmount: plainValidator.MinimumAmount,
	}
}

func GetImage(c coin.Coin, ID string) string {
	return AssetsURL + c.Handle + "/validators/assets/" + strings.ToLower(ID) + "/logo.png"
}
