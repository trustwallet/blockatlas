package blockatlas

type ProviderType string

const (
	ProviderTypeLending ProviderType = "lending"
	ProviderTypeStaking ProviderType = "staking"
)

type (
	// LendingProvider static info about the lending provider, such as name and asset classes supported.
	LendingProvider struct {
		ID     string              `json:"id"`
		Info   LendingProviderInfo `json:"info"`
		Type   ProviderType        `json:"type"`
		Assets []AssetClass        `json:"assets"`
	}

	// LendingProviderInfo basic information about a lending provider.
	LendingProviderInfo struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Website     string `json:"website"`
	}

	// AssetClass Info about an asset that can be lent
	AssetClass struct {
		Symbol      string `json:"symbol"`
		Chain       string `json:"chain"`
		Description string `json:"description"`
		// YieldFrequency the period of yield computation in seconds, e.g. 86400 for daily.
		YieldFrequency int64 `json:"yield_freq"`
		// Terms Predefined lending term periods, like [7, 30.5, 180].
		Terms []Term `json:"terms"`
	}

	// Term length of a predefined term, in days
	Term float64

	// RatesRequest Rates API request
	RatesRequest struct {
		Assets []string `json:"assets"`
	}

	// LendingTermAPR Asset yield APR, for an asset for a term.  E.g. {30, 1.45}
	LendingTermAPR struct {
		Term          `json:"term"`
		APR           float64 `json:"apr"`
		MinimumAmount string  `json:"minimum_amount"`
	}

	// LendingAssetRates Asset yield rates, for an asset for one or more periods.  E.g. [{7, 0.9}, {30, 1.45}]
	LendingAssetRates struct {
		Asset     string           `json:"asset"`
		TermRates []LendingTermAPR `json:"term_rates"`
		// MaxAPR the rate of the term with the highest rate
		MaxAPR float64 `json:"max_apr"`
	}

	// AccountRequest Account API request
	AccountRequest struct {
		Addresses []string `json:"addresses"`
		Assets    []string `json:"assets"`
	}

	// Time Second-granular UNIX time
	Time int32

	// LendingContract Describes a lending contract, of a user, of an asset.
	LendingContract struct {
		Asset             string  `json:"asset"`
		Term              Term    `json:"term"`
		StartAmount       string  `json:"start_amount"`
		CurrentAmount     string  `json:"current_amount"`
		EndAmountEstimate string  `json:"end_amount_estimate"`
		CurrentAPR        float64 `json:"current_apr"`
		StartTime         Time    `json:"start_time"`
		CurrentTime       Time    `json:"current_time"`
		EndTime           Time    `json:"end_time"`
	}

	// AccountLendingContracts Contracts of an address
	AccountLendingContracts struct {
		Address   string            `json:"address"`
		Contracts []LendingContract `json:"contracts"`
	}
)
