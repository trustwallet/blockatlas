package main

import (
	"strconv"
	"time"
)

// Lending API, as realized by Compound.
// As currently only Compund is planned, API is not made entirely generic, but prepared for later generalization.
// Compound does not use fixed terms, only open-ended, but structure is done to support different predefined terms.

// GetProviderInfo return static info about the lending provider, such as name and asset classes supported.
func GetProviderInfo() (LendingProvider, error) {
	// Note: should be cached
	return LendingProvider{
		"compound",
		LendingProviderInfo{
			"compound",
			"Compound Decentralized Finance Protocol",
			"https://compound.finance/images/compound-logo.svg",
			"https://compound.finance",
		},
		getTokensNormalized(),
	}, nil
}

// GetCurrentLendingRates return current estimated yield rates for assets.  Rates are annualized.  Rates vary over time.
// assets: List asset IDs to consider, or empty for all
// Note: can use the CTokenRequest compound API
func GetCurrentLendingRates(assets []string) (LendingRates, error) {
	res := LendingRates{}
	if len(assets) == 0 {
		// empty filter, get all available assets
		tokens := getTokens()
		for t := range tokens {
			assets = append(assets, t)
		}
	}
	for _, asset := range assets {
		if rates, err := getCurrentLendingRatesForAsset(asset); err == nil {
			res = append(res, rates)
		}
	}
	return res, nil
}

// GetAccountLendingContracts return current contract details for a given address.
// assets: List asset IDs to consider, or empty for all
func GetAccountLendingContracts(address string, assets []string) (AccountLendingContracts, error) {
	now := Time(time.Now().Unix())
	res := AccountLendingContracts{address, []LendingContract{}}
	contracts, _ := CMockAccount(CMAccountRequest{[]string{address}})
	for _, sc := range contracts.Account {
		for _, t := range sc.Tokens {
			asset := t.Symbol
			// APR: no info, take general current APR
			var apr float64 = 0
			if assetInfo, err := getCurrentLendingRatesForAsset(asset); err == nil {
				apr = assetInfo.MaxAPR
			}
			res.Contracts = append(res.Contracts, LendingContract{
				t.Symbol,
				0, // term
				// startAmount: not available in API, derive as currentAmount - interest earn
				strconv.FormatFloat(t.SupplyBalanceUnderlying-t.SupplyInterest, 'f', 10, 64),
				strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
				strconv.FormatFloat(t.SupplyBalanceUnderlying, 'f', 10, 64),
				apr,
				// startTime: no info, use current time
				now,
				now,
				now,
			})
		}
	}
	return res, nil
}
