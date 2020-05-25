package main

import ()

type (
	CMAccountRequest struct {
		Addresses []string `json:"addresses"`
	}

	CMAccountResponse struct {
		Account []CMAccount `json:"accounts"`
	}

	CMAccount struct {
		Address string            `json:"address"`
		Tokens  []CMAccountCToken `json:"tokens"`
	}

	CMAccountCToken struct {
		// Token address
		Address string `json:"address"`
		Symbol  string `json:"symbol"`
		// The cToken balance converted to underlying tokens; cTokens held x exchange rate
		SupplyBalanceUnderlying float64 `json:"supply_balance_underlying"`
		SupplyInterest          float64 `json:"lifetime_supply_interest_accrued"`
	}
)

var tokenAddressDAI = "0x6aabbcc000001"
var tokenAddressUSDC = "0x6aabbcc000002"

type contractInternalType struct {
	userAddress string
	contract    CMAccountCToken
}

var sampleContractsInternal []contractInternalType = []contractInternalType{
	{"0x12340000", CMAccountCToken{tokenAddressUSDC, "USDC", 200.45, 0.45}},
	{"0x12340000", CMAccountCToken{tokenAddressDAI, "DAI", 300.85, 0.85}},
	{"0x12560000", CMAccountCToken{tokenAddressUSDC, "USDC", 1001.25, 2.25}},
}

func CompoundMock_GetContracts(request CMAccountRequest) (CMAccountResponse, error) {
	var resp CMAccountResponse
	for _, addr := range request.Addresses {
		cmAccount := CMAccount{addr, []CMAccountCToken{}}
		for _, sc := range sampleContractsInternal {
			if sc.userAddress != addr {
				continue
			}
			cmAccount.Tokens = append(cmAccount.Tokens, sc.contract)
		}
		resp.Account = append(resp.Account, cmAccount)
	}
	return resp, nil
}
