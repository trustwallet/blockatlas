package main

import (
	"fmt"

	"github.com/trustwallet/blockatlas/lendingproto/lending"
)

func main() {
	fmt.Println("Lending proto")

	if err := lending.Init(":8080"); err != nil {
		panic(err.Error())
	}

	// curl "http://localhost:8080/v1/lending/providers"
	// curl -d '{}' "http://localhost:8080/v1/lending/rates/compound"
	// curl -d '{"assets":["USDC"]}' "http://localhost:8080/v1/lending/rates/compound"
	// curl -d '{"addresses":["0x12340000"]}' "http://localhost:8080/v1/lending/account/compound"
	// curl -d '{"addresses":["0x12360000"]}' "http://localhost:8080/v1/lending/account/compound"
	// curl -d '{"addresses":["0x12340000"],"assets":["USDC"]}' "http://localhost:8080/v1/lending/account/compound"
	// curl -d '{"addresses":["0x12340000", "0x12360000"],"assets":[]}' "http://localhost:8080/v1/lending/account/compound"
	// curl -d '{"addresses":[]}' "http://localhost:8080/v1/lending/account/compound"
}
