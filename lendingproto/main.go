package main

import (
	"fmt"

	"github.com/trustwallet/blockatlas/lendingproto/lending"
)

func main() {
	fmt.Println("Lending proto")

	lending.Init(":8080")

	// curl "http://localhost:8080/v1/lending/providers"
	// curl -d '{}' "http://localhost:8080/v1/lending/rates/compound"
	// curl -d '{"assets":["USDC"]}' "http://localhost:8080/v1/lending/rates/compound"
	// curl -d '{"address":"0x12340000"}' "http://localhost:8080/v1/lending/account/compound"
	// curl -d '{"address":"0x12360000"}' "http://localhost:8080/v1/lending/account/compound"
	// curl -d '{"address":"0x12340000","assets":["USDC"]}' "http://localhost:8080/v1/lending/account/compound"
}
