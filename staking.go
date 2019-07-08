package blockatlas

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

type StakeValidator struct {
	Status    bool               `json:"status"`
	Info      StakeValidatorInfo `json:"info"`
	Address   string             `json:"address"`
	PublicKey string             `json:"pubkey"`
}
