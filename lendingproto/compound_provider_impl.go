package main

import (
	//"fmt"
	"strconv"
	"time"
)

var sampleCurrentRates LendingRates = LendingRates{
	LendingAssetRates{"ETH", []LendingTermAPR{LendingTermAPR{0.00017, 0.01}}, 0},
	LendingAssetRates{"DAI", []LendingTermAPR{LendingTermAPR{0.00017, 0.73}}, 0},
	LendingAssetRates{"USDC", []LendingTermAPR{LendingTermAPR{0.00017, 1.67}}, 0},
	LendingAssetRates{"WBTC", []LendingTermAPR{LendingTermAPR{0.00017, 0.15}}, 0},
}

type contractInfoType struct {
	address     string
	asset       string
	startAmount float64
	startTime   int32
	realRate    float64
}

var sampleContracts []contractInfoType = []contractInfoType{
	contractInfoType{"0x12340000", "USDC", 200, 1589533690, 2.67},
	contractInfoType{"0x12340000", "DAI", 300, 1587002000, 1.73},
	contractInfoType{"0x12560000", "USDC", 1000, 1589210000, 2.67},
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

func getAssets() []string {
	res := make([]string, len(sampleCurrentRates))
	for i := range sampleCurrentRates {
		res[i] = sampleCurrentRates[i].Asset
	}
	return res
}

func getCurrentContractValues(contract contractInfoType) LendingContract {
	var now int32 = int32(time.Now().Unix())
	elapsedSecs := now - contract.startTime
	var elapsedDays float64 = 1.0 / 86400.0 * float64(elapsedSecs)
	currentAmount := contract.startAmount * (1.0 + 0.01*contract.realRate/365*elapsedDays)
	return LendingContract{
		contract.asset,
		0,
		strconv.FormatFloat(contract.startAmount, 'f', 10, 64),
		strconv.FormatFloat(currentAmount, 'f', 10, 64),
		strconv.FormatFloat(currentAmount, 'f', 10, 64), // no term, no end amount, use current
		contract.realRate,
		contract.startTime,
		now,
		now, // no term, no end time, use current
	}
}
