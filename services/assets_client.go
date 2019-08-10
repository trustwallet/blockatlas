package services

import (
	"github.com/trustwallet/blockatlas"
	"log"
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

type Validator struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
}

func GetValidators(coin coin.Coin) ([]Validator, error) {
	var results []Validator
	err := blockatlas.Request(&client, AssetsURL+coin.Handle, "/validators/list.json", url.Values{}, &results)
	log.Print(err)
	return results, err
}

func NormalizeValidators(stakersValidators []blockatlas.PlainStakeValidator, validators []Validator) ([]blockatlas.StakeValidator, error) {
	var results []blockatlas.StakeValidator

	for _, v := range stakersValidators {
		for _, v2 := range validators {
			if v.ID == v2.ID {
				results = append(results, NormalizeValidator(v, v2))
			}
		}
	}

	return results, nil
}

func NormalizeValidator(plainValidator blockatlas.PlainStakeValidator, validator Validator) blockatlas.StakeValidator {
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
