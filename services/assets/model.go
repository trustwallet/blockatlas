package assets

type AssetValidator struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Website     string          `json:"website"`
	Payout      ValidatorPayout `json:"payout,omitempty"`
}

type ValidatorPayout struct {
	Commission float64 `json:"commission"`
}
