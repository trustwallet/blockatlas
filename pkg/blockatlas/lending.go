package blockatlas

type ProviderType string

const (
	ProviderTypeLending ProviderType = "lending"
	ProviderTypeStaking ProviderType = "staking"
)

type (
	// LendingProvider static info about the lending provider, such as name and asset classes supported.
	LendingProvider struct {
		ID     string              `json:"id"`
		Info   LendingProviderInfo `json:"info"`
		Type   ProviderType        `json:"type"`
		Assets []AssetInfo         `json:"assets"`
	}

	// LendingProviderInfo basic information about a lending provider.
	LendingProviderInfo struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Image       string `json:"image"`
		Website     string `json:"website"`
	}

	// AssetInfo Info about an asset that can be lent
	AssetInfo struct {
		Symbol         string        `json:"symbol"`
		Description    string        `json:"description"`
		APY            float64       `json:"apy"`
		YieldPeriod    int64         `json:"yield_period"` // the period of validity of current APY, 0 for variable APY
		YieldFrequency int64         `json:"yield_freq"`   // the period of yield computation, in seconds, e.g. 86400 for daily yield writeoff.
		TotalSupply    string        `json:"total_supply"`
		MinimumAmount  string        `json:"minimum_amount"`
		MetaInfo       AssetMetaInfo `json:"meta_info,omitempty"`
	}

	AssetMetaInfo struct {
		DefiInfo *DefiAssetInfo `json:"defi_info,omitempty"` // pointer for omit to work
	}

	DefiAssetInfo struct {
		AssetToken     DefiTokenInfo `json:"asset_token,omitempty"`
		TechnicalToken DefiTokenInfo `json:"technical_token,omitempty"`
	}

	DefiTokenInfo struct {
		Symbol          string `json:"symbol"`
		Chain           string `json:"chain"`
		ContractAddress string `json:"contract_address,omitempty"`
	}

	// AccountRequest Account API request
	AccountRequest struct {
		Addresses []string `json:"addresses"`
		Assets    []string `json:"assets"`
	}

	// LendingContract Describes a lending contract, of a user, of an asset.
	LendingContract struct {
		Asset         AssetInfo `json:"asset"`
		CurrentAmount string    `json:"current_amount"`
	}

	// AccountLendingContracts Contracts of an address
	AccountLendingContracts struct {
		Address   string            `json:"address"`
		Contracts []LendingContract `json:"contracts"`
	}
)
