package assets

type AssetValidator struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Website     string          `json:"website"`
	Payout      ValidatorPayout `json:"payout,omitempty"`
	Status      ValidatorStatus `json:"status,omitempty"`
}

type ValidatorPayout struct {
	Commission float64 `json:"commission"`
}

type ValidatorStatus struct {
	Disabled bool `json:"disabled"`
}

type AssetValidators []AssetValidator
type AssetValidatorMap map[string]AssetValidator

func (av AssetValidators) toMap() AssetValidatorMap {
	validators := make(AssetValidatorMap)
	for _, v := range av {
		validators[v.ID] = v
	}
	return validators
}
