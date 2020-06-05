package blockatlas

type ProviderType string

const (
	ProviderTypeLending ProviderType = "lending"
	ProviderTypeStaking ProviderType = "staking"
)

type (
	LendingProvider struct {
		ID     string             `json:"id"`
		Info   StakeValidatorInfo `json:"info"`
		Type   ProviderType       `json:"type"`
		Assets []AssetInfo        `json:"assets"`
	}

	AssetInfo struct {
		Symbol         string        `json:"symbol"`
		Description    string        `json:"description"`
		APY            float64       `json:"apy"`
		YieldPeriod    int64         `json:"yield_period"` // the period of validity of current APY, 0 for variable APY
		YieldFrequency int64         `json:"yield_freq"`   // the period of yield computation, in seconds, e.g. 86400 for daily yield writeoff.
		TotalSupply    Amount        `json:"total_supply"`
		MinimumAmount  Amount        `json:"minimum_amount"`
		LockTime       int64         `json:"lock_time"`
		MetaInfo       AssetMetaInfo `json:"meta_info,omitempty"`
	}

	AssetMetaInfo struct {
		DefiInfo *DefiAssetInfo `json:"defi_info,omitempty"` // pointer type for omit to work
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

	AccountRequest struct {
		Addresses []string `json:"addresses"`
		Assets    []string `json:"assets"`
	}

	LendingContract struct {
		Asset         AssetInfo `json:"asset"`
		CurrentAmount string    `json:"current_amount"`
	}

	AccountLendingContracts struct {
		Address   string            `json:"address"`
		Contracts []LendingContract `json:"contracts"`
	}
)
