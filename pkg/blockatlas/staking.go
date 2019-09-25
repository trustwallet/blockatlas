package blockatlas

type ValidatorPage []Validator
type DelegationsPage []Validator

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

type DelegationStatus string

const (
	DelegationStatusActive  DelegationStatus = "active"
	DelegationStatusPending DelegationStatus = "pending"
)

const ValidatorsPerPage = 100

type StakingReward struct {
	Annual float64 `json:"annual"`
}

type Validator struct {
	ID     string        `json:"id"`
	Status bool          `json:"status"`
	Reward StakingReward `json:"reward"`
}

type Delegation struct {
	Delegator string           `json:"delegator"`
	Value     string           `json:"value"`
	Symbol    string           `json:"symbol"`
	Decimals  uint64           `json:"decimals"`
	Status    DelegationStatus `json:"status"`
	Metadata  interface{}      `json:"metadata"`
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
