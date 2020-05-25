package main

import (
	//"errors"
	//"fmt"
	"strconv"
	"time"
)

// Lending API, as realized by Compound.
// As currently only Compund is planned, API is not made entirely generic, but prepared for later generalization.
// Compound does not use fixed terms, only open-ended, but structure is done to support different predefined terms.

// GetProviderInfo return static info about the lending provider, such as name and asset classes supported.
func GetProviderInfo() (LendingProvider, error) {
	return LendingProvider{
		"compound",
		LendingProviderInfo{
			"compound",
			"Compound Decentralized Finance Protocol",
			"https://compound.finance/images/compound-logo.svg",
			"https://compound.finance",
		},
		// In compound all assets are updated with each ETH block, about each 15 seconds.  There are no predefined terms.
		AssetClasses{
			AssetClass{"ETH", "ETH", "Ethereum native coin", 15, AssetTerms{}},
			AssetClass{"USDC", "ETH", "USDC stablecoin token", 15, AssetTerms{}},
			AssetClass{"DAI", "ETH", "DAI stablecoin token", 15, AssetTerms{}},
			AssetClass{"WBTC", "ETH", "Wrapped Bitcoin token", 15, AssetTerms{}},
		},
	}, nil
}

// GetCurrentLendingRates return current estimated yield rates for assets.  Rates are annualized.  Rates vary over time.
// assets: List asset IDs to consider, or empty for all
// Note: can use the CTokenRequest compound API
func GetCurrentLendingRates(assets []string) (LendingRates, error) {
	var res LendingRates
	if len(assets) == 0 {
		assets = getAssets()
	}
	for _, asset := range assets {
		rates, err := getCurrentLendingRatesForAsset(asset)
		if err == nil {
			res = append(res, rates)
		}
	}
	return res, nil
}

// GetAccountLendingContracts return current contract details for a given address.
// assets: List asset IDs to consider, or empty for all
func GetAccountLendingContracts(address string, assets []string) (AccountLendingContracts, error) {
	var now int32 = int32(time.Now().Unix())
	res := AccountLendingContracts{
		address,
		LendingContracts{},
	}
	contracts, _ := CompoundMockGetContracts(CMAccountRequest{[]string{address}})
	for _, sc := range contracts.Account {
		for _, t := range sc.Tokens {
			asset := t.Symbol
			// APR: no info, take general current APR
			var apr float64 = 0
			assetInfo, err := getCurrentLendingRatesForAsset(asset)
			if err == nil {
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
