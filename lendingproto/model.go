package main

import ()

type (
	// LendingProvider static info about the lending provider, such as name and asset classes supported.
	LendingProvider struct {
		ID           string              `json:"id"`
		Info         LendingProviderInfo `json:"info"`
		AssetClasses AssetClasses        `json:"assets"`
	}

	LendingProviderInfo struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Website     string `json:"website"`
	}

	AssetClass struct {
		Symbol      string `json:"symbol"`
		Chain       string `json:"chain"`
		Description string `json:"description"`
		// YieldFrequency the period of yield computation in seconds, e.g. 86400 for daily.
		YieldFrequency int64 `json:"yield_freq"`
		// Terms Predefined lending term periods, like [7, 30.5, 180].
		Terms AssetTerms `json:"terms"`
	}

	AssetClasses []AssetClass

	// AssetTerm length of a predefined term for the asset, in days
	AssetTerm struct {
		Asset string  `json:"asset"`
		Term  float64 `json:"term"`
	}

	AssetTerms []AssetTerm

	// LendingAssetAPR Asset yield APR, for an asset for a period.  E.g. {30, 1.45}
	LendingTermAPR struct {
		Term float64 `json:"term"`
		APR  float64 `json:"apr"`
	}

	// LendingAssetRates Asset yield rates, for an asset for one or more periods.  E.g. [{7, 0.9}, {30, 1.45}]
	LendingAssetRates struct {
		Asset     string           `json:"asset"`
		TermRates []LendingTermAPR `json:"term_rates"`
		// MaxAPR the rate of the term with the highest rate
		MaxAPR float64 `json:"max_apr"`
	}

	LendingRates []LendingAssetRates

	LendingContract struct {
		Asset             string  `json:"asset"`
		Term              float64 `json:"term"`
		StartAmount       string  `json:"start_amount"`
		CurrentAmount     string  `json:"current_amount"`
		EndAmountEstimate string  `json:"end_amount_estimate"`
		CurrentAPR        float64 `json:"current_apr"`
		StartTime         int32   `json:"start_time"`
		CurrentTime       int32   `json:"current_time"`
		EndTime           int32   `json:"end_time"`
	}

	LendingContracts []LendingContract

	AccountLendingContracts struct {
		Address   string           `json:"address"`
		Contracts LendingContracts `json:"contracts"`
	}
)
