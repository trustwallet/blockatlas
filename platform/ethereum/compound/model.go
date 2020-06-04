package compound

type (
	Account struct {
		Address string          `json:"address"`
		Tokens  []AccountCToken `json:"tokens"`
	}

	AccountCToken struct {
		// Token address
		Address string `json:"address"`
		Symbol  string `json:"symbol"`
		// The cToken balance converted to underlying tokens; cTokens held x exchange rate
		SupplyBalanceUnderlying Precise `json:"supply_balance_underlying"`
		SupplyInterest          Precise `json:"lifetime_supply_interest_accrued"`
	}

	AccountResponse struct {
		Error    string    `json:"error"`
		Accounts []Account `json:"accounts"`
	}

	CTokenResponse struct {
		CToken []CToken `json:"cToken"`
	}

	CToken struct {
		TokenAddress     string  `json:"token_address"`
		TotalSupply      Precise `json:"total_supply"`
		ExchangeRate     Precise `json:"exchange_rate"`
		SupplyRate       Precise `json:"supply_rate"`
		Symbol           string  `json:"symbol"`
		Name             string  `json:"name"`
		UnderlyingSymbol string  `json:"underlying_symbol"`
		UnderlyingName   string  `json:"underlying_name"`
	}

	Precise struct {
		Value string
	}
)

const (
	Chain string = "ETH"
)
