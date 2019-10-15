package blockatlas

import "github.com/trustwallet/blockatlas/coin"

type ValidatorPage []Validator
type DelegationsPage []Delegation
type DelegationsBatchPage []DelegationsBatch

type DocsResponse struct {
	Docs interface{} `json:"docs"`
}

type DelegationStatus string
type DelegationType string

type ValidatorMap map[string]StakeValidator

const (
	DelegationStatusActive  DelegationStatus = "active"
	DelegationStatusPending DelegationStatus = "pending"

	DelegationTypeAuto     DelegationType = "auto"
	DelegationTypeDelegate DelegationType = "delegate"
)

const ValidatorsPerPage = 100

type StakingReward struct {
	Annual float64 `json:"annual"`
}

type StakingDetails struct {
	Reward        StakingReward  `json:"reward"`
	LockTime      int            `json:"locktime"`
	MinimumAmount Amount         `json:"minimum_amount"`
	Type          DelegationType `json:"type"`
}

type Validator struct {
	ID      string         `json:"id"`
	Status  bool           `json:"status"`
	Details StakingDetails `json:"details"`
}

type Delegation struct {
	Delegator StakeValidator   `json:"delegator"`
	Value     string           `json:"value"`
	Status    DelegationStatus `json:"status"`
	Metadata  interface{}      `json:"metadata,omitempty"`
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
	ID      string             `json:"id"`
	Status  bool               `json:"status,omitempty"`
	Info    StakeValidatorInfo `json:"info,omitempty"`
	Details StakingDetails     `json:"details,omitempty"`
}

type DelegationsBatch struct {
	Address     string             `json:"address"`
	Coin        *coin.ExternalCoin `json:"coin"`
	Details     StakingDetails     `json:"details"`
	Delegations DelegationsPage    `json:"delegations,omitempty"`
	Balance     string             `json:"balance"`
	Error       string             `json:"error,omitempty"`
}
