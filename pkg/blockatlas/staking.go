package blockatlas

import "github.com/trustwallet/blockatlas/coin"

type ValidatorPage []Validator
type DelegationsPage []Delegation
type DelegationsBatchPage []DelegationsBatch

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

type DelegationStatus string

type ValidatorMap map[string]StakeValidator

const (
	DelegationStatusActive  DelegationStatus = "active"
	DelegationStatusPending DelegationStatus = "pending"
)

const ValidatorsPerPage = 100

type StakingReward struct {
	Annual float64 `json:"annual"`
}

type Validator struct {
	ID            string        `json:"id"`
	Status        bool          `json:"status"`
	Reward        StakingReward `json:"reward"`
	LockTime      int           `json:"locktime"`
	MinimumAmount Amount        `json:"minimum_amount"`
}

type Delegation struct {
	Delegator StakeValidator `json:"delegator"`

	Value    string           `json:"value"`
	Status   DelegationStatus `json:"status"`
	Metadata interface{}      `json:"metadata,omitempty"`
}

type DelegationMetaDataPending struct {
	AvailableDate uint `json:"available_date"`
}

type StakeValidatorInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Website     string `json:"website"`
}

type StakeValidator struct {
	ID            string             `json:"id"`
	Status        bool               `json:"status,omitempty"`
	Info          StakeValidatorInfo `json:"info,omitempty"`
	Reward        StakingReward      `json:"reward,omitempty"`
	LockTime      int                `json:"locktime,omitempty"`
	MinimumAmount Amount             `json:"minimum_amount,omitempty"`
}

type DelegationsBatch struct {
	Address     string             `json:"address"`
	Coin        *coin.ExternalCoin `json:"coin"`
	Delegations DelegationsPage    `json:"delegations,omitempty"`
	Error       string             `json:"error,omitempty"`
}
