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
		Status      ValidatorStatus `json:"status,omitempty"`
		Staking     StakingInfo     `json:"staking,omitempty"`
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

func (av AssetValidators) activeValidators() AssetValidators {
	activeAssets := make(AssetValidators, 0)
	for _, a := range av {
		if !a.Status.Disabled {
			activeAssets = append(activeAssets, a)
		}
	}
	return activeAssets
}
