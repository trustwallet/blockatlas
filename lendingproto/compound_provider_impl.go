package main

import (
	"errors"
	//"fmt"
	//"strconv"
	//"time"
)

var sampleCurrentRates LendingRates = LendingRates{
	LendingAssetRates{"ETH", []LendingTermAPR{LendingTermAPR{0.00017, 0.01}}, 0},
	LendingAssetRates{"DAI", []LendingTermAPR{LendingTermAPR{0.00017, 0.73}}, 0},
	LendingAssetRates{"USDC", []LendingTermAPR{LendingTermAPR{0.00017, 1.67}}, 0},
	LendingAssetRates{"WBTC", []LendingTermAPR{LendingTermAPR{0.00017, 0.15}}, 0},
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

func matchAsset(asset string, assets []string) bool {
	if len(assets) == 0 {
		return true
	}
	for _, a := range assets {
		if asset == a {
			return true
		}
	}
	// no match
	return false
}

func matchAddress(address string, addresses []string) bool {
	for _, a := range addresses {
		if address == a {
			return true
		}
	}
	// no match
	return false
}

func getAssets() []string {
	res := make([]string, len(sampleCurrentRates))
	for i := range sampleCurrentRates {
		res[i] = sampleCurrentRates[i].Asset
	}
	return res
}

func getCurrentLendingRatesForAsset(asset string) (LendingAssetRates, error) {
	for i := range sampleCurrentRates {
		if sampleCurrentRates[i].Asset == asset {
			r := &sampleCurrentRates[i]
			enrichAssetRatesWithMax(r)
			return *r, nil
		}
	}
	return LendingAssetRates{}, errors.New("Asset not found")
}
