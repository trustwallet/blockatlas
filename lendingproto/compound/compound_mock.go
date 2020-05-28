package compound

// Simulates Compound API; see https://compound.finance/docs/api

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

	CMCToken struct {
		TokenAddress     string `json:"token_address"`
		TotalSupply      string `json:"total_supply"`
		ExchangeRate     string `json:"exchange_rate"`
		SupplyRate       string `json:"supply_rate"`
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		UnderlyingSymbol string `json:"underlying_symbol"`
		UnderlyingName   string `json:"underlying_name"`
	}

	CMCTokenResponse struct {
		CToken []CMCToken `json:"cToken"`
	}
)

var (
	tokenAddressDAI  = "0x6aabbcc000001"
	tokenAddressUSDC = "0x6aabbcc000002"
	tokenAddressETH  = "0x6aabbcc000003"
	tokenAddressWBTC = "0x6aabbcc000004"

	sampleTokenInfo = []CMCToken{
		{tokenAddressDAI, "200000", "1.5678", "0.0132", "cDAI", "Compound DAI", "DAI", "DAI"},
		{tokenAddressUSDC, "300000", "1.4259", "0.0167", "cUSDC", "Compound USD Coin", "USDC", "Circle USD Coin"},
		{tokenAddressETH, "800000", "1.2657", "0.0085", "cETH", "Compound Ether", "ETH", "Ether"},
		{tokenAddressWBTC, "60000", "1.0456", "0.0220", "cWBTC", "Compound Wrapped Bitcoin", "WBTC", "Wrpaped Bitcoin"},
	}

	sampleContractsInternal = []contractInternalType{
		{"0x12340000", CMAccountCToken{tokenAddressUSDC, "USDC", 200.45, 0.45}},
		{"0x12340000", CMAccountCToken{tokenAddressDAI, "DAI", 300.85, 0.85}},
		{"0x12360000", CMAccountCToken{tokenAddressUSDC, "USDC", 1001.25, 2.25}},
	}
)

type contractInternalType struct {
	userAddress string
	contract    CMAccountCToken
}

func CMockAccount(request CMAccountRequest) (CMAccountResponse, error) {
	resp := CMAccountResponse{}
	for _, addr := range request.Addresses {
		cmAccount := CMAccount{addr, []CMAccountCToken{}}
		for _, sc := range sampleContractsInternal {
			if sc.userAddress == addr {
				cmAccount.Tokens = append(cmAccount.Tokens, sc.contract)
			}
		}
		resp.Account = append(resp.Account, cmAccount)
	}
	return resp, nil
}

// See "https://api.compound.finance/api/v2/ctoken"
func CMockCToken(tokenAddresses []string) (CMCTokenResponse, error) {
	var res CMCTokenResponse
	for _, ct := range sampleTokenInfo {
		if matchAddress(ct.TokenAddress, tokenAddresses) {
			res.CToken = append(res.CToken, ct)
		}
	}
	return res, nil
}

func matchAddress(address string, addresses []string) bool {
	if len(addresses) == 0 {
		return true
	}
	for _, a := range addresses {
		if address == a {
			return true
		}
	}
	return false
}
