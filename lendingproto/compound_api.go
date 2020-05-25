package main

import (
	//"errors"
	//"fmt"
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
	var res LendingRates = LendingRates{}
	for i := range sampleCurrentRates {
		if !matchAsset(sampleCurrentRates[i].Asset, assets) {
			continue
		}
		r := &sampleCurrentRates[i]
		enrichAssetRatesWithMax(r)
		res = append(res, *r)
	}
	return res, nil
}

// GetAccountLendingContracts return current contract details for a given address.
// assets: List asset IDs to consider, or empty for all
func GetAccountLendingContracts(address string, assets []string) (AccountLendingContracts, error) {
	res := AccountLendingContracts{
		address,
		LendingContracts{},
	}
	for _, sc := range sampleContracts {
		if sc.address != address {
			continue
		}
		if !matchAsset(sc.asset, assets) {
			continue
		}
		res.Contracts = append(res.Contracts, getCurrentContractValues(sc))
	}
	return res, nil
}
