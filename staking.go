package blockatlas

type ValidatorPage []Validator

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

const ValidatorsPerPage = 100

type StakingReward struct {
	Annual float64 `json:"annual"`
}

type Validator struct {
	ID     string        `json:"id"`
	Status bool          `json:"status"`
	Reward StakingReward `json:"reward"`
}

type StakeValidatorInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Website     string `json:"website"`
}

type StakeValidator struct {
	ID     string             `json:"id"`
	Status bool               `json:"status"`
	Info   StakeValidatorInfo `json:"info"`
	Reward StakingReward      `json:"reward"`
}
