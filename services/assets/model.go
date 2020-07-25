package assets

type (
	AssetValidators []AssetValidator

	AssetValidatorMap map[string]AssetValidator

	AssetValidator struct {
		ID          string          `json:"id"`
		Name        string          `json:"name"`
		Description string          `json:"description"`
		Website     string          `json:"website"`
		Payout      ValidatorPayout `json:"payout,omitempty"`
	}

	ValidatorPayout struct {
		Commission float64 `json:"commission"`
	}

	ValidatorStatus struct {
		Disabled bool `json:"disabled"`
	}

	StakingInfo struct {
		MinDelegation float64 `json:"minDelegation"`
	}
)

func (av AssetValidators) toMap() AssetValidatorMap {
	validators := make(AssetValidatorMap)
	for _, v := range av {
		validators[v.ID] = v
	}
	return validators
}
