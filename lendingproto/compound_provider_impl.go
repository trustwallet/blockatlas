package main

import (
	"fmt"
	"strconv"
)

type tokenInfo struct {
	address string
	name    string
}

func getTokensNormalized() []AssetClass {
	// In compound all assets are updated with each ETH block, about each 15 seconds.  There are no predefined terms.
	tokens := getTokens()
	res := []AssetClass{}
	for s, t := range tokens {
		res = append(res, AssetClass{s, "ETH", t.name, 15, []Term{}})
	}
	return res
}

// Returns a info on tokens, map by symbol
// Note: this should be cached
func getTokens() map[string]tokenInfo {
	tokens := CMockCToken([]string{})
	res := make(map[string]tokenInfo)
	for _, t := range tokens.CToken {
		res[t.UnderlyingSymbol] = tokenInfo{t.TokenAddress, t.Name}
	}
	return res
}

// Note: should work from cached data
func addressOfToken(symbol string) (string, bool) {
	tokens := getTokens()
	tokenInfo, ok := tokens[symbol]
	if !ok {
		return "", false
	}
	return tokenInfo.address, true
}

func getCurrentLendingRatesForAsset(asset string) (LendingAssetRates, error) {
	res := LendingAssetRates{asset, []LendingTermAPR{}, 0}
	address, ok := addressOfToken(asset)
	if !ok {
		return res, fmt.Errorf("Token not found %v", asset)
	}
	tokens := CMockCToken([]string{address})
	for _, t := range tokens.CToken {
		apr, err := strconv.ParseFloat(t.SupplyRate, 64)
		if err != nil {
			apr = 0
		} else {
			apr = 100.0 * apr
		}
		res.TermRates = append(res.TermRates, LendingTermAPR{0.00017, apr})
	}
	enrichAssetRatesWithMax(&res)
	return res, nil
}

func enrichAssetRatesWithMax(rates *LendingAssetRates) {
	var max float64 = 0
	for _, r := range rates.TermRates {
		if r.APR > max {
			max = r.APR
		}
	}
	rates.MaxAPR = max
}
