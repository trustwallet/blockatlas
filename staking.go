package blockatlas

type ValidatorPage []StakeValidator

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

const ValidatorsPerPage = 100

type StakeValidatorInfo struct {
	Name     string `json:"website"`
	Description      string `json:"image"`
}

type StakeValidator struct {
	//Name     string `json:"name"`
	//Description      string `json:"description"`
	//Status    string   `json:"status"`
	//Uptime int `json:"uptime"`
	//Info        StakeValidatorInfo `json:"info"`
	Address        string `json:"address"`
	PublicKey        string `json:"pubkey"`
}
