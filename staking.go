package blockatlas

import "github.com/trustwallet/blockatlas/coin"

type ValidatorPage []StakeValidator

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

const ValidatorsPerPage = 100

type StakeValidatorInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Website     string `json:"website"`
}

type PlainStakeValidator struct {
	Coin   coin.Coin
	ID     string `json:"id"`
	Status bool   `json:"status"`
}

type StakeValidator struct {
	ID     string             `json:"id"`
	Status bool               `json:"status"`
	Info   StakeValidatorInfo `json:"info"`
}
