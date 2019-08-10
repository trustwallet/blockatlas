package assets

import (
	"github.com/trustwallet/blockatlas"
	"time"

	"github.com/trustwallet/blockatlas/coin"
	"net/http"
	"net/url"
)

var client = http.Client{
	Timeout: time.Second * 5,
}

const (
	AssetsURL = "https://raw.githubusercontent.com/trustwallet/assets/master/blockchains/"
)

func GetValidators(coin coin.Coin) ([]AssetValidator, error) {
	var results []AssetValidator
	err := blockatlas.Request(&client, AssetsURL+coin.Handle, "/validators/list.json", url.Values{}, &results)
	return results, err
}

func NormalizeValidators(validators []blockatlas.Validator, assets []AssetValidator) []blockatlas.StakeValidator {
	var results []blockatlas.StakeValidator

	for _, v := range validators {
		for _, v2 := range assets {
			if v.ID == v2.ID {
				results = append(results, NormalizeValidator(v, v2))
			}
		}
	}

	return results
}

func NormalizeValidator(plainValidator blockatlas.Validator, validator AssetValidator) blockatlas.StakeValidator {
	return blockatlas.StakeValidator{
		ID:     validator.ID,
		Status: plainValidator.Status,
		Info: blockatlas.StakeValidatorInfo{
			Name:        validator.Name,
			Description: validator.Description,
			Image:       GetImage(plainValidator.Coin, plainValidator.ID),
			Website:     validator.Website,
		},
	}
}

func GetImage(c coin.Coin, ID string) string {
	return AssetsURL + c.Handle + "/validators/assets/" + ID + "/logo.png"
}
