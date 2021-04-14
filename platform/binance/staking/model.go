package staking

type Validators struct {
	Total      int         `json:"total"`
	Validators []Validator `json:"validators"`
}

type Validator struct {
	Validator string  `json:"validator"`
	Status    int     `json:"status"`
	APR       float64 `json:"apr"`
}
