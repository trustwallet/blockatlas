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
