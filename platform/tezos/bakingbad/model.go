package bakingbad

type Baker struct {
	Address       string  `json:"address"`
	Fee           float64 `json:"fee"`
	PayoutDelay   int     `json:"payoutDelay"`
	PayoutPeriod  int     `json:"payoutPeriod"`
	MinDelegation int     `json:"minDelegation"`
	FreeSpace     float64 `json:"freeSpace"`
	EstimatedRoi  float64 `json:"estimatedRoi"`
}
