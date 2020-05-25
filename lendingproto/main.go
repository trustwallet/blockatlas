package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func printContract(contract LendingContract) {
	daysAgo := float64(contract.CurrentTime-contract.StartTime) / 86400.0
	startAmnt, _ := strconv.ParseFloat(contract.StartAmount, 64)
	currentAmnt, _ := strconv.ParseFloat(contract.CurrentAmount, 64)
	fmt.Printf("You have lent the amount of %v %v %v days ago, yield so far: %v %v, %v%% APR \n",
		strconv.FormatFloat(startAmnt, 'f', 2, 64), contract.Asset,
		strconv.FormatFloat(daysAgo, 'f', 1, 64),
		strconv.FormatFloat(currentAmnt-startAmnt, 'f', 2, 64), contract.Asset,
		strconv.FormatFloat(contract.CurrentAPR, 'f', 2, 64))
}

func printContracts(address string, asset string) {
	fmt.Printf("GetAccountLendingContracts for address %v and ", address)
	assetList := []string{}
	if len(asset) > 0 {
		fmt.Printf("asset %v ", asset)
		assetList = []string{asset}
	} else {
		fmt.Printf("all assets ")
	}
	fmt.Printf(":\n")
	contracts, _ := GetAccountLendingContracts(address, assetList)
	for _, c := range contracts.Contracts {
		printContract(c)
	}
	b, _ := json.MarshalIndent(contracts, "    ", "    ")
	fmt.Println(string(b))
}

func main() {
	fmt.Println("Lending proto")

	providerInfo, _ := GetProviderInfo()
	fmt.Println("ProviderInfo:")
	b, _ := json.MarshalIndent(providerInfo, "    ", "    ")
	fmt.Println(string(b))

	asset := "WBTC"
	rate, _ := GetCurrentLendingRates([]string{asset})
	fmt.Println("CurrentLendingRates for " + asset + ":")
	b, _ = json.MarshalIndent(rate, "    ", "    ")
	fmt.Println(string(b))

	rates, _ := GetCurrentLendingRates([]string{})
	fmt.Println("CurrentLendingRates for all:")
	b, _ = json.MarshalIndent(rates, "    ", "    ")
	fmt.Println(string(b))

	printContracts("0x12340000", "")
	printContracts("0x12560000", "")
	printContracts("0x12340000", "USDC")
}
