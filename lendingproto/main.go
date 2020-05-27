package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/trustwallet/blockatlas/lendingproto/compound"
	"github.com/trustwallet/blockatlas/lendingproto/lending"
	"github.com/trustwallet/blockatlas/lendingproto/model"
)

func printContract(contract model.LendingContract) {
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
	contracts, _ := compound.GetAccountLendingContracts(address, assetList)
	for _, c := range contracts.Contracts {
		printContract(c)
	}
	b, _ := json.MarshalIndent(contracts, "    ", "    ")
	fmt.Println(string(b))
}

func main() {
	fmt.Println("Lending proto")

	lending.Init(":8080")

	// curl "http://localhost:8080/v1/lending/providers"

	printContracts("0x12340000", "")
	printContracts("0x12560000", "")
	printContracts("0x12340000", "USDC")
}
