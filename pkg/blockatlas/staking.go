package blockatlas

import (
	"github.com/trustwallet/golibs/coin"
	"github.com/trustwallet/golibs/types"
)

const (
	DelegationStatusActive  DelegationStatus = "active"
	DelegationStatusPending DelegationStatus = "pending"

	DelegationTypeAuto     DelegationType = "auto"
	DelegationTypeDelegate DelegationType = "delegate"

	DefaultAnnualReward = 0.0
)

type (
	ValidatorPage        []Validator
	DelegationsPage      []Delegation
	DelegationsBatchPage []DelegationResponse
	StakingBatchPage     []StakingResponse
	StakeValidators      []StakeValidator

	DelegationStatus string
	DelegationType   string

	ValidatorMap map[string]StakeValidator

	StakingReward struct {
		Annual float64 `json:"annual"`
	}

	StakingDetails struct {
		Reward        StakingReward  `json:"reward"`
		LockTime      int            `json:"locktime"`
		MinimumAmount types.Amount   `json:"minimum_amount"`
		Type          DelegationType `json:"type"`
	}

	Validator struct {
		ID      string         `json:"id"`
		Status  bool           `json:"status"`
		Details StakingDetails `json:"details"`
	}

	Delegation struct {
		Delegator StakeValidator   `json:"delegator"`
		Value     string           `json:"value"`
		Status    DelegationStatus `json:"status"`
		Metadata  interface{}      `json:"metadata,omitempty"`
	}

	DelegationMetaDataPending struct {
		AvailableDate uint `json:"available_date"`
	}

	StakeValidatorInfo struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Website     string `json:"website"`
	}

	StakeValidator struct {
		ID      string             `json:"id"`
		Status  bool               `json:"status"`
		Info    StakeValidatorInfo `json:"info,omitempty"`
		Details StakingDetails     `json:"details,omitempty"`
	}

	DelegationResponse struct {
		Delegations DelegationsPage `json:"delegations"`
		Balance     string          `json:"balance"`
		Address     string          `json:"address"`
		StakingResponse
	}

	StakingResponse struct {
		Coin    *coin.ExternalCoin `json:"coin"`
		Details StakingDetails     `json:"details"`
	}
)

func (s StakeValidators) ToMap() ValidatorMap {
	validators := make(ValidatorMap)
	for _, v := range s {
		validators[v.ID] = v
	}
	return validators
}

func FindHightestAPR(validators []Validator) float64 {
	var apr = 0.0
	for _, v := range validators {
		if apr < v.Details.Reward.Annual {
			apr = v.Details.Reward.Annual
		}
	}
	return apr
}
